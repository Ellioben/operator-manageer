package svc

import (
	"context"
	"fmt"
	"operator-manager/internal/config"
	"operator-manager/internal/middleware"
	"operator-manager/internal/model"

	"github.com/redis/go-redis/v9"

	"github.com/jmoiron/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config               config.Config
	EarlyCheckMiddleware rest.Middleware
	Redis                *redis.Client
	MySQL                *sqlx.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化MySQL连接
	mysqlConn := sqlx.MustConnect("mysql", c.MySQL.DSN)

	// 初始化Redis连接
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})

	c.ClientSets = initCluster()
	return &ServiceContext{
		Config:               c,
		EarlyCheckMiddleware: middleware.NewEarlyCheckMiddleware(c).Handle,
		Redis:                rdb,
		MySQL:                mysqlConn,
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

func (s *ServiceContext) GetAllClusters() ([]*model.ClusterConfig, error) {
	query := `SELECT cluster_name, kube_config, is_active FROM cluster_configs WHERE is_active = 1`
	var clusters []*model.ClusterConfig
	err := s.MySQL.Select(&clusters, query)
	return clusters, err
}

func (s *ServiceContext) StoreClusterConfig(cluster *model.ClusterConfig) error {
	return s.Redis.Set(
		context.Background(),
		fmt.Sprintf("cluster:%s:config", cluster.Name),
		cluster.KubeConfig,
		0,
	).Err()
}
