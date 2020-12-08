package log

type Options struct {
	LogGrpc            bool     //记录GRPC 框架日志
	OutputPaths        []string //日志输出路径
	RotationMaxSize    int      // 切割日志大小
	RotationMaxBackups int      // 备份多少个
	RotationMaxAge     int      //最大时间
	LogLevel           string
}

func DefaultOptions() *Options {
	return &Options{
		LogGrpc:            false,
		OutputPaths:        []string{"stdout"},
		RotationMaxSize:    10,
		RotationMaxBackups: 10,
		RotationMaxAge:     10,
		LogLevel:           "debug",
	}
}
