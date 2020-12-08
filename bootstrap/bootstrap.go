package bootstrap

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dotamixer/doom/pkg/di"
	"github.com/dotamixer/doom/pkg/log"
	"github.com/dotamixer/doom/pkg/register"
	"github.com/dotamixer/doom/pkg/register/consul"
	"github.com/dotamixer/doom/pkg/store/mongo"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"google.golang.org/grpc"
)

type Server struct {
	// grpc server
	grpcSrv *grpc.Server
	// registry is a service registry
	registry register.Register
	// Container is a dependence inject container
	Container *dig.Container
	// opts is server options
	opts *Options
	// logConfig is log configuration
	logConfig *LogConfig
	// srvConfig is server configuration
	srvConfig *ServerConfig
	// redisConfig is redis configuration
	redisConfig *RedisConfig
	// mongoConfig is redis configuration
	mongoConfig *MongoConfig
}

func NewServer(opts ...Option) (s *Server) {

	options := NewOptions(opts...)

	s = &Server{
		opts:      options,
		Container: dig.New(),
	}

	s.loadConfig()

	return s
}

func (s *Server) Init(opts ...Option) {

	for _, o := range opts {
		o(s.opts)
	}

	s.initLog()

	s.initDI()

	s.initGrpcServer()
}

func (s *Server) initLog() {
	opt := &log.Options{
		LogGrpc:            s.logConfig.LogGrpc,
		OutputPaths:        s.logConfig.OutputPaths,
		RotationMaxSize:    s.logConfig.RotationMaxSize,
		RotationMaxBackups: s.logConfig.RotationMaxBackups,
		RotationMaxAge:     s.logConfig.RotationMaxAge,
	}

	logrus.Info("begin to init log")
	err := log.Configure(opt)
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info("Init log success")
}

func (s *Server) initDI() {

	if s.opts.withRedis {
		di.MustContainerProvide(s.Container, s.NewRedisOptions)
		di.MustContainerProvide(s.Container, redis.NewClient)
	}

	if s.opts.withMongoDB {
		di.MustContainerProvide(s.Container, s.NewMongoOptions)
		di.MustContainerProvide(s.Container, mongo.NewClient)
	}

	logrus.Info("Init DI success")
}

func (s *Server) initGrpcServer() {
	var (
		err error
	)

	if s.opts.withRegister {
		s.registry, err = consul.NewRegister(&register.Options{
			RegistryAddr: s.srvConfig.RegistryAddr,
			Name:         s.opts.serviceName,
			Port:         s.srvConfig.Port,
		})
		if err != nil {
			logrus.Fatalf("Failed to new register. err:[%v]", err)
			return
		}
	}

	s.grpcSrv = grpc.NewServer()

	s.opts.registerGrpcCallback(s.grpcSrv)

	logrus.Info("Init grpc server success")
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", s.srvConfig.Port))
	if err != nil {
		logrus.Fatalf("Failed to listen. err:[%v]", err)
	}

	go func() {
		logrus.Infof("grpc Server is serving(addr:%s)...", lis.Addr())

		err = s.grpcSrv.Serve(lis)
		if err != nil {
			logrus.Fatalf("Failed to serve grpc server. err:[%v]", err)
		}
	}()

	if s.opts.withRegister {
		// 注册服务
		err = s.registry.Register()
		if err != nil {
			logrus.Fatalf("Failed to register service. err:[%v]", err)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-c
		logrus.Infof("capture a signal. signal:[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:

			s.Stop()

			if s.opts.withRegister {
				_ = s.registry.Deregister()
			}

			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func (s *Server) Stop() {
	s.grpcSrv.GracefulStop()
}
