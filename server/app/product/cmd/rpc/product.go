package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"server/app/product/cmd/rpc/internal/config"
	"server/app/product/cmd/rpc/internal/server"
	"server/app/product/cmd/rpc/internal/svc"
	"server/app/product/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// var configFile = flag.String("f", "etc/product.yaml", "the config file")
var configFile = flag.String("f", "app/product/cmd/rpc/etc/product.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterProductServer(grpcServer, server.NewProductServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	logx.DisableStat()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
