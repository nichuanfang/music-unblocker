package config

// Config 配置文件
type Config struct {
	Log        *LogConfig     `json:"log"`         // 日志配置
	Music      []*MusicConfig `json:"music"`       // 音乐源配置
	FfmpegPath string         `json:"ffmpeg_path"` // ffmepg路径
	TargetPath string         `json:"target_path"` // 破解后的整理目录
}

// LogConfig 日志配置
type LogConfig struct {
	Mode  int    `json:"mode"`  // 日志模式：1标准输出，2日志文件，3标准输出和日志文件
	Level int    `json:"level"` // 日志等级，0-4分别是：debug，info，warning，error，fatal
	File  string `json:"file"`  // 日志文件路径
}

// MusicConfig 音乐配置
type MusicConfig struct {
	Name         string `json:"name"`          // 音乐源
	NameZh       string `json:"name_zh"`       // 音乐源(中文名称)
	DownloadPath string `json:"download_path"` // 下载路径 会监听该路径并破解
}
