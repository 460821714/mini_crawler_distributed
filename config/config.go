// @Time : 2018/6/12 17:30
// @Author : minigeek
package config

const (
	//ItemSaverRpcServicePort = 8888
	ElasticIndex        = "crawler"
	ItemSaverMethodSave = "ItemSaverService.Save"

	//WorkerRpcServicePort0 = 9000
	WorkerMethodProcess = "CrawService.Process"
)

// parser names
const (
	ParseCityList = "ParseCityList"
	ParseCity     = "ParseCity"
	ParseProfile  = "ParseProfile"
	ParseNil      = "ParseNil"
)

// rate limiter
const QPS = 20
