package logic

import (
	"context"

	"gorm-zero-example/services/api/internal/svc"
	"gorm-zero-example/services/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) DelUserLogic {
	return DelUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelUserLogic) DelUser(req types.DeleteUserReq) error {
	// todo: add your logic here and delete this line

	err := l.svcCtx.UserModel.Delete(l.ctx, nil, req.Id)

	return err
}
