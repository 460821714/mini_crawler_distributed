package main

import (
	"flag"
	"fmt"
	"log"
	"mini_crawler_distributed/rpcsupport"
	"mini_crawler_distributed/worker"
)

var port = flag.Int("port", 0, "the port for worker lisetening on")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Println("must supply a port")
		return
	}
	err := rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), &worker.CrawService{})
	if err != nil {
		panic(err)
	}
}
