package main

import (
	"log"
	"net"
	"time"

	"github.com/bookun/sandbox/go/grpc/pb"
	"github.com/bookun/sandbox/go/grpc/service"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

func main() {
	opts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}

	zapLogger, _ := zap.NewProduction()
	grpc_zap.ReplaceGrpcLogger(zapLogger)

	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(zapLogger, opts...),
		),
	)
	catService := &service.MyCatService{}
	helloService := &service.HelloService{}

	pb.RegisterCatServer(server, catService)
	pb.RegisterHelloServiceServer(server, helloService)

	server.Serve(listenPort)

}
