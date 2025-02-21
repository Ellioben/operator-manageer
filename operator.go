package main

import (
	"context"
	"flag"
	"fmt"
	"operator-manager/internal/model"

	"operator-manager/internal/config"
	"operator-manager/internal/handler"
	"operator-manager/internal/svc"

	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

var configFile = flag.String("f", "etc/operator.yaml", "the config file")

type Operator struct {
	svcCtx *svc.ServiceContext
	ctx    context.Context
}

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)

	op := &Operator{
		svcCtx: ctx,
		ctx:    context.Background(),
	}

	op.Start()
}

func (op *Operator) Start() {
	if err := op.initClusterConfigs(); err != nil {
		logx.Errorf("集群配置初始化失败: %v", err)
		return
	}

	server := rest.MustNewServer(op.svcCtx.Config.RestConf)
	defer server.Stop()

	server.Use(op.svcCtx.EarlyCheckMiddleware)
	handler.RegisterHandlers(server, op.svcCtx)

	fmt.Printf("Starting server at %s:%d...\n",
		op.svcCtx.Config.Host,
		op.svcCtx.Config.Port)
	server.Start()
}

func (op *Operator) initClusterConfigs() error {
	clusters, err := op.svcCtx.GetAllClusters()
	if err != nil {
		return fmt.Errorf("获取集群配置失败: %w", err)
	}

	for _, cluster := range clusters {
		if err := op.svcCtx.StoreClusterConfig(cluster); err != nil {
			logx.Errorf("存储集群配置失败 [%s]: %v", cluster.Name, err)
			continue
		}

		if err := op.initClusterInformer(cluster); err != nil {
			logx.Errorf("初始化informer失败 [%s]: %v", cluster.Name, err)
			continue
		}
	}
	return nil
}

func (op *Operator) initClusterInformer(cluster *model.ClusterConfig) error {
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(cluster.KubeConfig))
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	informerFactory := informers.NewSharedInformerFactory(clientset, time.Minute)

	namespaceInformer := informerFactory.Core().V1().Namespaces()
	namespaceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			ns := obj.(*v1.Namespace)
			logx.Infof("检测到新命名空间 [%s] 在集群 [%s]", ns.Name, cluster.Name)
		},
	})

	informerFactory.Start(op.ctx.Done())
	return nil
}
