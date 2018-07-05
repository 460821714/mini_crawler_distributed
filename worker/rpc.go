// @Time : 2018/7/3 19:30
// @Author : minigeek
package worker

import "mini_crawler/engine"

type CrawService struct{}

func (c *CrawService) Process(req Request, result *ParseResult) error {
	enginReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	enginParseResult, err := engine.Worker(enginReq)
	if err != nil {
		return err
	}
	*result = SerializeParseResult(enginParseResult)
	return nil
}
