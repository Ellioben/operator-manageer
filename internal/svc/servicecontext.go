package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"operator-manager/internal/config"
	"operator-manager/internal/middleware"
)

type ServiceContext struct {
	Config                  config.Config
	NewEarlyCheckMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	c.ClientSets = initCluster()
	return &ServiceContext{
		Config:                  c,
		NewEarlyCheckMiddleware: middleware.NewEarlyCheckMiddleware(c).Handle,
	}
}

func initCluster() map[string]string {
	//TODO 将多个 cluster 集群信息加入到
	// id - clientset

	m := make(map[string]string)
	// 把所有集群kubeconfig加入到 map
	for i := 0; i < 10; i++ {
		m["cluster-"+string(rune(i))] = "kubeconfig" + string(rune(i))
	}
	return m
}
