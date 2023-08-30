package models

type Config struct {
	Server Server `yaml:"Server"`
	DB     DB     `yaml:"DB"`
}

type Server struct {
	Host     string `yaml:"Host,omitempty"`
	Port     uint   `yaml:"Port,omitempty"`
	RTimeout int64  `yaml:"RTimeout,omitempty"`
	WTimeout int64  `yaml:"WTimeout,omitempty"`
}

type DB struct {
	DriverName  string `yaml:"DriverName"`
	ConnString  string `yaml:"ConnString"`
	MaxConns    int    `yaml:"MaxConns"`
	ConnTimeout int64  `yaml:"ConnTimeout"`
}
