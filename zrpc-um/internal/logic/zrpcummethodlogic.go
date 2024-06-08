package logic

import (
	"context"

	"go-zero-tls/common/pb"
	"go-zero-tls/zrpc-um/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ZrpcUmMethodLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewZrpcUmMethodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ZrpcUmMethodLogic {
	return &ZrpcUmMethodLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ZrpcUmMethodLogic) ZrpcUmMethod(in *__.ZrpcUmMethodRequest) (*__.ZrpcUmMethodResponse, error) {
	// todo: add your logic here and delete this line

	resposta, err := l.svcCtx.ZrpcDoisClient.ZrpcDoisMethod(l.ctx, &__.ZrpcDoisMethodRequest{Nome: in.GetNome()})
	if err != nil {
		return nil, err
	}

	return &__.ZrpcUmMethodResponse{Frase: resposta.Frase}, nil
}
