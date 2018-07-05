// @Time : 2018/6/12 17:15
// @Author : minigeek
package client

import (
	"log"
	"mini_crawler/engine"

	"mini_crawler_distributed/config"
	"mini_crawler_distributed/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item, error) {
	out := make(chan engine.Item)
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	go func() {
		var itemCount int
		for {
			item := <-out
			log.Printf("ItemSaver got item %d,%v", itemCount, item)
			itemCount++
			// call rpc
			result := ""
			err = client.Call(config.ItemSaverMethodSave, item, &result)
			if err != nil {
				log.Printf("item: %v,save faild:%v", item, err)
			}
		}
	}()
	return out, nil
}
