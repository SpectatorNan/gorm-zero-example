package nocache

import (
	"context"

	"gorm-zero-example/services/api/internal/svc"
	"gorm-zero-example/services/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NocacheDelUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNocacheDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NocacheDelUserLogic {
	return &NocacheDelUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NocacheDelUserLogic) NocacheDelUser(req *types.DeleteUserReq) error {
	err := l.svcCtx.UserNoCacheModel.Delete(l.ctx, nil, req.Id)
	return err
}
