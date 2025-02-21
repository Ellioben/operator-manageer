package middleware

import (
	"context"
	"net/http"
	"operator-manager/internal/config"
	"operator-manager/pkg/cluster"
)

type EarlyCheckMiddleware struct {
	config config.Config
}

func NewEarlyCheckMiddleware(c config.Config) *EarlyCheckMiddleware {
	return &EarlyCheckMiddleware{config: c}
}

func (m *EarlyCheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 使用配置中的AuthToken进行验证
		if r.Header.Get("X-Auth") != m.config.AuthToken {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func (m *EarlyCheckMiddleware) SetContext(r *http.Request) *http.Request {
	// 获取当前集群信息
	clientSet := cluster.NewCluster(m.config, GetHeaderField(r, "clusterid")).ClientSet
	// 把获取到的clientset放在上下文里
	ctx := context.WithValue(r.Context(), "clientset", clientSet)
	r = r.WithContext(ctx)
	return r
}

func GetHeaderField(r *http.Request, headerField string) string {
	return r.Header.Get(headerField)
}
