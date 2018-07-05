// @Time : 2018/6/12 14:10
// @Author : minigeek
package rpcsupport

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) error {
	//注册一个rpc服务
	rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	fmt.Println("listening on " + host)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("one client connect.")
		go jsonrpc.ServeConn(conn)
	}
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
