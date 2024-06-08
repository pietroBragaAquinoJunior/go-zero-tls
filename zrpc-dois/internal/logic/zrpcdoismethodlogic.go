package logic

import (
	"context"

	"go-zero-tls/common/pb"
	"go-zero-tls/zrpc-dois/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ZrpcDoisMethodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewZrpcDoisMethodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ZrpcDoisMethodLogic {
	return &ZrpcDoisMethodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ZrpcDoisMethodLogic) ZrpcDoisMethod(in *__.ZrpcDoisMethodRequest) (*__.ZrpcDoisMethodResponse, error) {
	// todo: add your logic here and delete this line

	frase := "Eu sou o zrpc dois ! Olá, " + in.Nome

	return &__.ZrpcDoisMethodResponse{Frase: frase}, nil
}
