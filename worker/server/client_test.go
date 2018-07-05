package main

import (
	"fmt"
	"mini_crawler_distributed/config"
	"mini_crawler_distributed/rpcsupport"
	"mini_crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawService(t *testing.T) {
	go rpcsupport.ServeRpc(":9000", &worker.CrawService{})
	time.Sleep(2 * time.Second)

	client, err := rpcsupport.NewClient(":9000")
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/109782247",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "MeiZi",
		},
	}

	var result worker.ParseResult

	err = client.Call(config.WorkerMethodProcess, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("items:")
		for _, item := range result.Items {
			fmt.Println(item)
		}
		fmt.Println("requests:")
		for _, req := range result.Requests {
			fmt.Println(req)
		}
	}
}
