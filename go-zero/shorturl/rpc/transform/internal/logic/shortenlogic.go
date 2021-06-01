package logic

import (
	"context"

	"shorturl/rpc/transform"
	"shorturl/rpc/transform/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ShortenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ShortenLogic) Shorten(in *__.ShortenReq) (*__.ShortenResp, error) {
	// todo: add your logic here and delete this line

	return &__.ShortenResp{}, nil
}
