package coupon

import (
	"net/http"

	"SimplePick-Mall-Server/api/internal/logic/sms/coupon"
	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CouponAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCouponReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := coupon.NewCouponAddLogic(r.Context(), svcCtx)
		resp, err := l.CouponAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
