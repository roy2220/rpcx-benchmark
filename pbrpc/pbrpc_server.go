package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"

	"github.com/let-z-go/pbrpc/channel"
	"github.com/rpcx-ecosystem/rpcx-benchmark/proto"
)

func SayHello(rpc *channel.RPC) {
	// , ctx context.Context, args *proto.BenchmarkMessage, reply *proto.BenchmarkMessage
	args := rpc.Request.(*proto.BenchmarkMessage)
	reply := *args
	reply.Field1 = "OK"
	reply.Field2 = 100
	rpc.Response = &reply
	if *delay > 0 {
		time.Sleep(*delay)
	} else {
		runtime.Gosched()
	}
	return
}

var (
	host       = flag.String("s", "127.0.0.1:8972", "listened ip and port")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
	delay      = flag.Duration("delay", 0, "delay to mock business processing")
	debugAddr  = flag.String("d", "127.0.0.1:9981", "server ip and port")
)

func main() {
	flag.Parse()
	go func() {
		log.Println(http.ListenAndServe(*debugAddr, nil))
	}()
	l, err := net.Listen("tcp", *host)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			opts := channel.Options{}
			opts.SetMethod("Hello", "Say").
				SetRequestFactory(func() channel.Message { return new(proto.BenchmarkMessage) }).
				SetResponseFactory(func() channel.Message { return new(proto.BenchmarkMessage) }).
				SetIncomingRPCHandler(SayHello)
			cn := new(channel.Channel).Init(&opts)
			log.Println(cn.AcceptAndServe(context.Background(), c))
		}()
	}
}
