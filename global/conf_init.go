package global

import "github.com/sllt/tp-tcp-plugin/conf"

func init() {
	conf, err := conf.LoadConfig("default")
	if err != nil {
		panic(err)
	}
	Config = conf
}
