package clusterInfo

import (
	"context"

	"operator-manager/internal/svc"
	"operator-manager/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClusterNamespaceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClusterNamespaceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClusterNamespaceLogic {
	return &GetClusterNamespaceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClusterNamespaceLogic) GetClusterNamespace(req *types.GetClusterNamespaceReq) (resp *types.GetClusterNamespaceResp, err error) {
	// 获取上下文的clientset来执行集群操作
	return
}
