package configs

type Configuration struct {
	Owner   OwnerConfig `mapstructure:"owner"`
	Project Project     `mapstructure:"project"`
	Http    Http        `mapstructure:"http"`
}

type OwnerConfig struct {
	Name string `mapstructure:"name"`
}

type Project struct {
	Name    string `mapstructure:"name"`
	Dob     string `mapstructure:"dob"`
	Version string `mapstructure:"version"`
	Log     int    `mapstructure:"log"`
	Status  string `mapstructure:"status"`
}

type Http struct {
	Port     string `mapstructure:"port"`
	Address  string `mapstructure:"address"`
	Waittime int    `mapstructure:"waittime"`
}
