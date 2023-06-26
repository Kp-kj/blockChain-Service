package user

import (
	"net/http"

	"block/internal/logic/user"
	"block/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetHelpCategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetHelpCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.GetHelpCategoryList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
