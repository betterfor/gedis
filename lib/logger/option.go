package logger

type Option func(*option)

type option struct {
	path     string
	name     string
	ext      string
	logLevel logLevel
	//timeFormat string

	// todo 日志轮转
}

// WithLogPath 设置日志路径
func WithLogPath(path string) Option {
	return func(o *option) {
		o.path = path
	}
}

// WithLogName 设置日志文件名
func WithLogName(name string) Option {
	return func(o *option) {
		o.name = name
	}
}

// WithLogExt 设置日志后缀
func WithLogExt(ext string) Option {
	return func(o *option) {
		o.ext = ext
	}
}

func WithLogLevel(level logLevel) Option {
	return func(o *option) {
		o.logLevel = level
	}
}
