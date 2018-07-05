// @Time : 2018/5/23 15:07
// @Author : minigeek
package main

import (
	"flag"
	"fmt"
	"log"
	"mini_crawler/engine"
	"mini_crawler/scheduler"
	"mini_crawler/zhenai/parser"
	itemsaver "mini_crawler_distributed/persist/client"
	"mini_crawler_distributed/rpcsupport"
	worker "mini_crawler_distributed/worker/client"
	"net/rpc"
	"strings"
)

var (
	itemsaver_host = flag.String("itemsaver_host", "", "itemsaver_host")
	worker_hosts   = flag.String("worker_hosts", "", "worker_hosts(comma separate)")
)

// start url for fetch.
const startUrl = "http://www.zhenai.com/zhenghun"

func main() {
	flag.Parse()
	if *itemsaver_host == "" {
		fmt.Println("must apply itemsaver_host")
		return
	}
	if *worker_hosts == "" {
		fmt.Println("must apply worker_hosts")
		return
	}
	itemChan, err := itemsaver.ItemSaver(*itemsaver_host)
	if err != nil {
		panic(err)
	}
	log.Println("start fetch...")
	clientChan := createClientPool(strings.Split(*worker_hosts, ","))
	processor, err := worker.CreateProcessor(clientChan)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    startUrl,
		Parser: engine.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err != nil {
			log.Printf("connect to %s fail:%v", h, err)
		} else {
			clients = append(clients, client)
		}
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}
