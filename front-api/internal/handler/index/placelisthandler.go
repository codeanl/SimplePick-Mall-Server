package index

import (
	"net/http"

	"SimplePick-Mall-Server/front-api/internal/logic/index"
	"SimplePick-Mall-Server/front-api/internal/svc"
	"SimplePick-Mall-Server/front-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PlaceListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListPlaceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := index.NewPlaceListLogic(r.Context(), svcCtx)
		resp, err := l.PlaceList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
