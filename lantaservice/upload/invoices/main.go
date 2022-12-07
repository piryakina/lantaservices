package main

import (
	"log"
	"nninvest/cmd"
)

func main() {
	//var service = flag.StringP("service", "s", "rpc", "kind of service")
	//flag.Parse()
	//switch *service {
	//case "rpc":
	//	rpc.Server()
	//case "web":
	//	web.Server()
	//}

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
