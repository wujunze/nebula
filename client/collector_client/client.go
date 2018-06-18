package collector_client

import (
	"context"
	"fmt"
	"time"

	"github.com/robfig/cron"
	"github.com/samoslab/nebula/provider/node"
	pb "github.com/samoslab/nebula/tracker/collector/client/pb"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var NodePtr *node.Node

func Collect(al *pb.ActionLog) {
	l := len(queue)
	if l < cap(queue)*9/10 {
		queue <- al
	} else {
		log.Warnf("queue will be full, abandon action log, ticket: %s", al.Ticket)
	}
	if l > send_immediate_min {
		go send()
	}
}

const batch_max = 500
const send_immediate_min = 20

var queue = make(chan *pb.ActionLog, 2000)
var cronRunner *cron.Cron
var sendLock = make(chan bool, 1)
var conn *grpc.ClientConn

func sendLockOff() {
	sendLock <- false
}
func Start(collectServer string) {
	sendLockOff()
	var err error
	conn, err = grpc.Dial(collectServer, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial collector failed: %s", err)
	}

	cronRunner = cron.New()
	cronRunner.AddFunc("4,19,34,49 * * * * *", send)
	cronRunner.Start()
}

func Stop() {
	cronRunner.Stop()
	conn.Close()
}

func send() {
	select {
	case _ = <-sendLock:
		defer sendLockOff()
		if err := doSend(); err != nil {
			log.Warnf("send action log to collector error: %s", err)
		}
	default:
	}
}

func doSend() error {
	pcsc := pb.NewClientCollectorServiceClient(conn)
	stream, err := pcsc.Collect(context.Background())
	if err != nil {
		fmt.Printf("RPC Collect failed: %s", err.Error())
		return err
	}
	for {
		if len(queue) == 0 {
			break
		}
		size := len(queue)
		if size > batch_max {
			size = batch_max
		}
		if err = stream.Send(buildReq(size)); err != nil {
			return err
		}
	}
	_, err = stream.CloseAndRecv()
	return err
}

func buildReq(size int) *pb.CollectReq {
	bs := make([]*pb.ActionLog, 0, size)
	for i := 0; i < size; i++ {
		bs = append(bs, <-queue)
	}
	//no := node.LoadFormConfig()
	req := &pb.CollectReq{NodeId: NodePtr.NodeId,
		Timestamp: uint64(time.Now().UnixNano()),
		ActionLog: bs}
	req.SignReq(NodePtr.PriKey)
	return req
}
