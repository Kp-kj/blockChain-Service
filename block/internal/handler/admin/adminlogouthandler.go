package admin

import (
	"net/http"

	"block/internal/logic/admin"
	"block/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AdminLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewAdminLogoutLogic(r.Context(), svcCtx)
		resp, err := l.AdminLogout()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
