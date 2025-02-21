package model

type ClusterConfig struct {
	Name       string `db:"cluster_name"`
	KubeConfig string `db:"kube_config"`
	IsActive   bool   `db:"is_active"`
}
