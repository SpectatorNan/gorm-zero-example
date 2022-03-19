package logic

import (
	"context"
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
		Account:  req.Account,
		NickName: req.NickName,
		Password: req.Password,
	}

	err = l.svcCtx.UserModel.Insert(&u)
	if err != nil {
		return nil, err
	}
	resp = &types.AddUserResp{Id: int64(u.ID)}
	return
}
