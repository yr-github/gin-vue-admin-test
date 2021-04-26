package config

type Rabbit struct {
	User     string    	`mapstructure:"user" json:"user" yaml:"user"`
	Password string 	`mapstructure:"password" json:"password" yaml:"password"`
	Addr     string 	`mapstructure:"addr" json:"addr" yaml:"addr"`
}
