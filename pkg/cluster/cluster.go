package cluster

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"operator-manager/internal/config"
)

type Cluster struct {
	ID        string
	ClientSet *kubernetes.Clientset
}

func NewCluster(c config.Config, id string) *Cluster {
	// TODO 通过id获取对应集群的kubefile(通过配置文件id - kubeconfig生成对应的clientset)
	var kubeFile = c.ClientSets[id]
	var clientset *kubernetes.Clientset
	var err error

	if kubeFile == "" {
		var clusterConfig *rest.Config
		clusterConfig, err = rest.InClusterConfig()
		if err != nil {
			log.Fatalf("Error creating in-cluster config: %v", err)
		}
		clientset, err = kubernetes.NewForConfig(clusterConfig)
		if err != nil {
			log.Fatalf("Error creating clientset: %v", err)
		}
	} else {
		var config *rest.Config
		config, err = clientcmd.RESTConfigFromKubeConfig([]byte(kubeFile))
		if err != nil {
			log.Fatalf("Error creating config from kubeconfig: %v", err)
		}
		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatalf("Error creating clientset: %v", err)
		}
	}

	return &Cluster{
		ID:        id,
		ClientSet: clientset,
	}
}
