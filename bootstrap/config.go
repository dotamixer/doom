package bootstrap

import (
	"net/url"
	"os"
	"path/filepath"

	"github.com/dotamixer/doom/pkg/lion"
	"github.com/dotamixer/doom/pkg/lion/source/file"
	"github.com/dotamixer/doom/pkg/store/mongo"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
)

type ServerConfig struct {
	RegistryAddr string `yaml:"registryAddr"`
	Port int `yaml:"port"`
}

type LogConfig struct {
	LogGrpc            bool   `yaml:"logGrpc"`            //记录GRPC 框架日志
	OutputPaths        []string `yaml:"outputPaths"`         //日志输出路径
	RotationMaxSize    int    `yaml:"rotationMaxSize"`    // 切割日志大小
	RotationMaxBackups int    `yaml:"rotationMaxBackups"` // 备份多少个
	RotationMaxAge     int    `yaml:"rotationMaxAge"`     //最大时间
	LogLevel           string `yaml:"logLevel"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

type MongoConfig struct {
	Hosts       []string `yaml:"hosts"`
	Username    string   `yaml:"username"`
	Password    string   `yaml:"password"`
	AuthSource  string   `yaml:"authSource"`
	MaxPoolSize uint64   `yaml:"maxPoolSize"`
	ReplicaSet  string   `yaml:"replicaSet"`
}

func (s *Server) loadConfig() {

	rawUrl := os.Getenv("DOOM_SERVICE_CONFIG")
	logrus.Infof("ralUrl: %s", rawUrl)

	ret, err := url.Parse(rawUrl)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("ret:%+v", *ret)
	if ret.Scheme == "file" {
		err = lion.Load(file.NewSource(file.WithPath(filepath.Join(".", ret.Path))))
		if err != nil {
			logrus.Fatal(err)
		}
	} else if ret.Scheme == "apollo" {
		//TODO:
	}

	s.loadLogConfig()

	s.loadServerConfig()

	if s.opts.withRedis {
		s.loadRedisConfig()
	}

	if s.opts.withMongoDB {
		s.loadMongoConfig()
	}
}

func (s *Server) loadLogConfig() {
	logrus.Info("load log config")
	s.logConfig = &LogConfig{}

	err := lion.Get("log").Scan(s.logConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("load log config success. config:[%+v]", *s.logConfig)
}

func (s *Server) loadServerConfig() {
	s.srvConfig = &ServerConfig{}

	err := lion.Get("server").Scan(s.srvConfig)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("load server config success. config:[%+v]", s.srvConfig)
}

func (s *Server) loadRedisConfig() {
	s.redisConfig = &RedisConfig{}

	err := lion.Get("redis").Scan(s.redisConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("load redis config success. config:[%+v]", *s.redisConfig)

}

func (s *Server) loadMongoConfig() {
	s.mongoConfig = &MongoConfig{}

	err := lion.Get("mongodb").Scan(s.mongoConfig)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("load mongo config success, config:[%+v]", *s.mongoConfig)

}

func (s *Server) NewRedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     s.redisConfig.Addr,
		Password: s.redisConfig.Password,
	}
}

func (s *Server) NewMongoOptions() *mongo.Options {
	return &mongo.Options{
		Hosts:       s.mongoConfig.Hosts,
		Username:    s.mongoConfig.Username,
		Password:    s.mongoConfig.Password,
		AuthSource:  s.mongoConfig.AuthSource,
		MaxPoolSize: s.mongoConfig.MaxPoolSize,
		ReplicaSet:  s.mongoConfig.ReplicaSet,
	}
}
