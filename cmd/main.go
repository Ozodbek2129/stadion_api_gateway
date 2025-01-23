package main

import (
	"apigateway/api"
	"apigateway/config"
)

func main() {

	r := api.NewRouter()
	r.Run(config.Load().API_GATEWAY)

}