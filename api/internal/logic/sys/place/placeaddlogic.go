package place

import (
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceAddLogic {
	return &PlaceAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceAddLogic) PlaceAdd(req *types.AddPlaceReq) (resp *types.AddPlaceResp, err error) {
	_, err = l.svcCtx.Sys.PlaceAdd(l.ctx, &sysclient.PlaceAddReq{
		Name:      req.Name,
		Place:     req.Place,
		Status:    req.Status,
		Pic:       req.Pic,
		Phone:     req.Phone,
		Principal: req.Principal,
		CreateBy:  l.ctx.Value("username").(string),
	})
	if err != nil {
		return nil, err
	}
	return &types.AddPlaceResp{
		Code:    200,
		Message: "添加成功",
	}, nil
}
