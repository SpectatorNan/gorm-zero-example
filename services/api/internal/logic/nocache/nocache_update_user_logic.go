package nocache

import (
	"context"

	"gorm-zero-example/services/api/internal/svc"
	"gorm-zero-example/services/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NocacheUpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNocacheUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NocacheUpdateUserLogic {
	return &NocacheUpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NocacheUpdateUserLogic) NocacheUpdateUser(req *types.UpdateUserReq) error {
	// todo: add your logic here and delete this line

	return nil
}
