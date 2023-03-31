package common

import "flag"

func SetCmdParams() {
	var vip string
	flag.StringVar(&vip, "virtual ip", "11.11.11.11", "cluster mgt vip")
	flag.Parse()
}
