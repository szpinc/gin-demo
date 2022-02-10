package config

type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	RootDir    string `mapstructure:"root-dir" json:"root-dir" yaml:"root-dir"`
	FileName   string `mapstructure:"file-name" json:"file-name" yaml:"file-name"`
	Format     string `mapstructure:"format" json:"format" yaml:"format"`
	ShowLine   bool   `mapstructure:"show-line" json:"show-line" yaml:"show-line"`
	MaxBackups int    `mapstructure:"max-backups" json:"max-backups" yaml:"max-backups"`
	MaxSize    int    `mapstructure:"max-size" json:"max-size" yaml:"max-size"`
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}
