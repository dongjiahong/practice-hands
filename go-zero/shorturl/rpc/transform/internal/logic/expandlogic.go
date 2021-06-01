package logic

import (
	"context"

	"shorturl/rpc/transform"
	"shorturl/rpc/transform/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ExpandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExpandLogic) Expand(in *__.ExpandReq) (*__.ExpandResp, error) {
	// todo: add your logic here and delete this line

	return &__.ExpandResp{}, nil
}
