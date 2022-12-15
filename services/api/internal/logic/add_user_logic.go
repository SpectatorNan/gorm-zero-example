package logic

import (
	"context"
	"database/sql"
	"gorm-zero-example/services/api/internal/svc"
	"gorm-zero-example/services/api/internal/types"
	"gorm-zero-example/services/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddUserLogic {
	return AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req types.AddUserReq) (resp *types.AddUserResp, err error) {
	// todo: add your logic here and delete this line

	u := model.Users{
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

	err = l.svcCtx.UserModel.Insert(l.ctx, nil, &u)
	if err != nil {
		return nil, err
	}
	resp = &types.AddUserResp{Id: u.Id}
	return
}
