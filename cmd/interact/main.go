package main

import (
	"demo/tiktok/cmd/interact/dal"
	"demo/tiktok/cmd/interact/rpc"
	interactdemo "demo/tiktok/kitex_gen/interactdemo/interactservice"
	"demo/tiktok/pkg/constants"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	dal.Init()
	rpc.Init()
}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
	if err != nil {
		panic(err)
	}

	svr := interactdemo.NewServer(new(InteractServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.InteractServiceName}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
