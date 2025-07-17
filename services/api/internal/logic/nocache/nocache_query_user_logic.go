package nocache

import (
	"context"
	"errors"
	"gorm-zero-example/app/errorx"
	"gorm-zero-example/services/model"

	"gorm-zero-example/services/api/internal/svc"
	"gorm-zero-example/services/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NocacheQueryUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNocacheQueryUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NocacheQueryUserLogic {
	return &NocacheQueryUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NocacheQueryUserLogic) NocacheQueryUser(req *types.QueryUserReq) (resp *types.QueryUserResp, err error) {
	user, err := l.svcCtx.UserNoCacheModel.FindOne(l.ctx, req.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errorx.NewDefaultError("用户不存在")
		}
		return nil, err
	}
	resp = &types.QueryUserResp{
		Id:         user.Id,
		Account:    user.Account.String,
		NickName:   user.NickName.String,
		CreateTime: user.CreatedAt.Time.Format("2006-01-02 15:04:05"),
	}
	return
}
