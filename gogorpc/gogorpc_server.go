package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"

	"github.com/let-z-go/gogorpc/channel"
	"github.com/let-z-go/gogorpc/server"
	"github.com/rpcx-ecosystem/rpcx-benchmark/gogorpc/gogorpc"
	"github.com/rpcx-ecosystem/rpcx-benchmark/proto"
)

var (
	host       = flag.String("s", "127.0.0.1:8972", "listened ip and port")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	delay      = flag.Duration("delay", 0, "delay to mock business processing")
	debugAddr  = flag.String("d", "127.0.0.1:9981", "server ip and port")
)

type serviceHandler struct{}

func (serviceHandler) Say(ctx context.Context, args *proto.BenchmarkMessage) (*proto.BenchmarkMessage, error) {
	reply := *args
	args.Field1 = "OK"
	args.Field2 = 100
	if *delay > 0 {
		time.Sleep(*delay)
	} else {
		runtime.Gosched()
	}
	return &reply, nil
}

func main() {
	flag.Parse()
	go func() {
		log.Println(http.ListenAndServe(*debugAddr, nil))
	}()
	opts := server.Options{
		Channel: (&channel.Options{}).Do(gogorpc.RegisterHelloHandler(serviceHandler{})),
	}
	svr := new(server.Server).Init(&opts, "tcp://"+*host)
	defer svr.Close()
	svr.Run()
}
