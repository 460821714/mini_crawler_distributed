// @Time : 2018/6/12 14:28
// @Author : minigeek
package main

import (
	"flag"
	"log"
	"mini_crawler_distributed/config"
	"mini_crawler_distributed/persist"
	"mini_crawler_distributed/rpcsupport"

	"fmt"

	"gopkg.in/olivere/elastic.v5"
)

var port = flag.Int("port", 0, "the port for itemsaver lisetening on")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Println("must supply a port")
		return
	}
	err := ServeRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex)
	if err != nil {
		panic(err)
	}
}

func ServeRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
