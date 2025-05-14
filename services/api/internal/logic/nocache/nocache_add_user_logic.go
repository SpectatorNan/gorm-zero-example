package nocache

import (
	"context"
	"database/sql"
	"gorm-zero-example/services/model_noCache"

	"gorm-zero-example/services/api/internal/svc"
	"gorm-zero-example/services/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NocacheAddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNocacheAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NocacheAddUserLogic {
	return &NocacheAddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NocacheAddUserLogic) NocacheAddUser(req *types.AddUserReq) (resp *types.AddUserResp, err error) {

	u := model_noCache.Users{
		Account: sql.NullString{
			String: req.Account,
			Valid:  true,
		},
		NickName: sql.NullString{
			String: req.NickName,
			Valid:  true,
		},
		Password: sql.NullString{
			String: req.Password,
			Valid:  true,
		},
	}

	err = l.svcCtx.UserNoCacheModel.Insert(l.ctx, nil, &u)
	if err != nil {
		return nil, err
	}
	resp = &types.AddUserResp{Id: u.Id}
	return
}
