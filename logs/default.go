package logs

// LogConfig 日志配置结构体
type LogConfig struct {
	Path       string // 文件路径，默认 logs/
	Level      string // 日志等级，默认 Debug
	MaxSize    int    // 日志容量
	MaxAge     int    // 日志保存时间
	MaxBackups int    // 日志备份数量
	Compress   bool   // 是否压缩
}

// defaultOption 默认的日志配置选项
func defaultOption() *LogConfig {
	return &LogConfig{
		Path:       "logs/",
		Level:      "debug",
		MaxSize:    20,
		MaxAge:     30,
		MaxBackups: 1000,
		Compress:   true,
	}
}
