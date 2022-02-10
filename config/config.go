package config

type Configuration struct {
	App        App      `mapstructure:"app" json:"app" yaml:"app"`
	Log        Log      `mapstructure:"log" json:"log" yaml:"log"`
	Datasource Database `mapstructure:"datasource" json:"datasource" yaml:"datasource"`
}

type App struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	Name string `mapstructure:"name" json:"name" yaml:"name"`
	Url  string `mapstructure:"url" json:"url" yaml:"url"`
}
