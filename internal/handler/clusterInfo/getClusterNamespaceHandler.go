package clusterInfo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"operator-manager/internal/logic/clusterInfo"
	"operator-manager/internal/svc"
	"operator-manager/internal/types"
)

func GetClusterNamespaceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetClusterNamespaceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := clusterInfo.NewGetClusterNamespaceLogic(r.Context(), svcCtx)
		resp, err := l.GetClusterNamespace(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
