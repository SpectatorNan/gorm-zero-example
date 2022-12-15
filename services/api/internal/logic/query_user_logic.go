package logic

import (
	"context"
	"gorm-zero-example/app/errorx"
	"gorm-zero-example/services/model"
	"time"

	"gorm-zero-example/services/api/internal/svc"
	"gorm-zero-example/services/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryUserLogic {
	return QueryUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserLogic) QueryUser(req types.QueryUserReq) (resp *types.QueryUserResp, err error) {

	user, err := l.svcCtx.UserModel.FindOneWithExpire(l.ctx, req.Id, time.Minute*3)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError("用户不存在")
		}
		return nil, err
	}
	resp = &types.QueryUserResp{
		Id:         user.Id,
		Account:    user.Account.String,
		NickName:   user.NickName.String,
		CreateTime: user.CreateAt.Time.Format("2006-01-02 15:04:05"),
	}
	return
}
