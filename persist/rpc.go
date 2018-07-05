// @Time : 2018/6/12 14:23
// @Author : minigeek
package persist

import (
	"mini_crawler/engine"
	"mini_crawler/persist"

	"gopkg.in/olivere/elastic.v5"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	if err != nil {
		return err
	}
	*result = "ok"
	return nil
}
