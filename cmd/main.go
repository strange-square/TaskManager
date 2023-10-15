package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"
	"main.go/internal/db"
	"main.go/internal/grpc_api"
	pb "main.go/internal/proto"
	"main.go/internal/repository/postgresql"
)

const (
	service     = "api"
	environment = "development"
	tracerURL   = "http://localhost:14268/api/traces"
	serviceAddr = ":9091"
)

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service),
			attribute.String("environment", environment),
		)),
	)
	return tp, nil
}

func main() {
	tp, err := tracerProvider(tracerURL)
	if err != nil {
		log.Fatal(err)
	}

	otel.SetTracerProvider(tp)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func(ctx context.Context) {
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	lsn, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.NewDB(ctx)
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	defer database.GetPool(ctx).Close()
	Tasks := postgresql.NewTasks(database)
	Projects := postgresql.NewProjects(database)

	go http.ListenAndServe(serviceAddr, promhttp.Handler())

	serv := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))

	pb.RegisterTaskManagerServiceServer(serv, grpc_api.NewGRPCServer(ctx, Tasks, Projects))

	log.Printf("starting server on %s", lsn.Addr().String())
	if err := serv.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
