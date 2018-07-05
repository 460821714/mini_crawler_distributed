// @Time : 2018/7/5 14:33
// @Author : minigeek
package client

import (
	"mini_crawler/engine"
	"mini_crawler_distributed/config"
	"mini_crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor, error) {
	return func(request engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(request)
		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.WorkerMethodProcess, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, nil
		}
		return worker.DeserializeParseResult(sResult), nil
	}, nil
}
