package user

import (
	"SimplePick-Mall-Server/service/sys/rpc/sysclient"
	"context"
	"encoding/json"

	"SimplePick-Mall-Server/api/internal/svc"
	"SimplePick-Mall-Server/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (*types.UserInfoResp, error) {
	id, _ := l.ctx.Value("id").(json.Number).Int64()
	resp, _ := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.InfoReq{ID: id})
	UserInfo := types.UserInfo{
		ID:       resp.UserInfo.ID,
		Username: resp.UserInfo.Username,
		NickName: resp.UserInfo.Nickname,
		Phone:    resp.UserInfo.Phone,
		Gander:   resp.UserInfo.Gender,
		Avatar:   resp.UserInfo.Avatar,
		Email:    resp.UserInfo.Email,
		Status:   resp.UserInfo.Status,
		CreateAt: resp.UserInfo.CreateAt,
		UpdateAt: resp.UserInfo.UpdateAt,
		CreateBy: resp.UserInfo.CreateBy,
		UpdateBy: resp.UserInfo.UpdateBy,
	}
	data := types.Data{
		UserInfo: UserInfo,
		Routes:   resp.Routes,
		Roles:    resp.Roles,
		Buttons:  resp.Buttons,
	}
	return &types.UserInfoResp{
		Code:    200,
		Data:    data,
		Message: "success",
	}, nil
}
