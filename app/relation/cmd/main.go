package main

import (
	"dzug/app/relation/service"
	"dzug/discovery"
	pb "dzug/idl/relation"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	endpoints := []string{"localhost:2379"}
	// lease 应该是租约时间，这里是5秒
	etcdRegister, err := discovery.NewServiceRegister(endpoints, "relation", "localhost:9001", 5)
	if err != nil {
		log.Fatal(err)
	}
	defer etcdRegister.Close()
	go etcdRegister.ListenLeaseRespChan() // 启用协程，监听续租响应通道

	// 创建grpc服务器并监听9000端口
	server := grpc.NewServer()
	defer server.Stop()
	pb.RegisterDouyinRelationActionServiceServer(server, &service.RelationSrv{})

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listening ")
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}