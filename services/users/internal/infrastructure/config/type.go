package config

type Config struct {
	Port int `mapstructure:"port"`

	Db struct {
		Port     int    `mapstructure:"port"`
		Host     string `mapstructure:"host"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DbName   string `mapstructure:"db_name"`
	} `mapstructure:"DB"`
}
