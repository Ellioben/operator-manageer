package svc

import (
	"operator-manager/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	ClusterMap map[string]string
}

func NewServiceContext(c config.Config) *ServiceContext {
	//TODO 将多个 cluster 集群信息加入到
	// id - clientset
	m := make(map[string]string)

	return &ServiceContext{
		Config:     c,
		ClusterMap: m,
	}
}
