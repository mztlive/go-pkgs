package client

import (
	"sync"

	"google.golang.org/grpc"
)

var (

	// clientMap 客户端连接池, key是服务名, value是客户端连接
	clientMap = make(map[string]*grpc.ClientConn)
	lock      sync.RWMutex
)

// GetClient 获取客户端连接
func GetClient(serviceName string) *grpc.ClientConn {
	lock.RLock()
	defer lock.RUnlock()

	if client, ok := clientMap[serviceName]; ok {
		return client
	}

	return nil
}

// AddClient 新建客户端连接
func AddClient(serviceName, address string) (*grpc.ClientConn, error) {
	lock.Lock()
	defer lock.Unlock()

	if client, ok := clientMap[serviceName]; ok {
		return client, nil
	}

	client, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	clientMap[serviceName] = client
	return client, nil
}

// CloseAll 关闭所有客户端连接
// 一般在程序退出时调用
func CloseAll() {
	lock.Lock()
	defer lock.Unlock()

	for _, client := range clientMap {
		client.Close()
	}
}
