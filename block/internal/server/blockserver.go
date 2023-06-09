// Code generated by goctl. DO NOT EDIT.
// Source: block.proto

package server

import (
	"context"

	"block/block"
	"block/internal/logic"
	"block/internal/svc"
)

type BlockServer struct {
	svcCtx *svc.ServiceContext
	block.UnimplementedBlockServer
}

func NewBlockServer(svcCtx *svc.ServiceContext) *BlockServer {
	return &BlockServer{
		svcCtx: svcCtx,
	}
}

// 管理后台接口
func (s *BlockServer) Ping(ctx context.Context, in *block.Request) (*block.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *BlockServer) CreateCryptominer(ctx context.Context, in *block.CreateCryptominerRequest) (*block.IsSuccessResponse, error) {
	l := logic.NewCreateCryptominerLogic(ctx, s.svcCtx)
	return l.CreateCryptominer(in)
}

func (s *BlockServer) CreateProp(ctx context.Context, in *block.CreatePropRequest) (*block.IsSuccessResponse, error) {
	l := logic.NewCreatePropLogic(ctx, s.svcCtx)
	return l.CreateProp(in)
}

func (s *BlockServer) AdminGoodList(ctx context.Context, in *block.AdminGoodListRequest) (*block.AdminGoodListResponse, error) {
	l := logic.NewAdminGoodListLogic(ctx, s.svcCtx)
	return l.AdminGoodList(in)
}

func (s *BlockServer) StartGood(ctx context.Context, in *block.StartGoodRequest) (*block.IsSuccessResponse, error) {
	l := logic.NewStartGoodLogic(ctx, s.svcCtx)
	return l.StartGood(in)
}

func (s *BlockServer) CreateActivity(ctx context.Context, in *block.CreateActivityRequest) (*block.IsSuccessResponse, error) {
	l := logic.NewCreateActivityLogic(ctx, s.svcCtx)
	return l.CreateActivity(in)
}

func (s *BlockServer) AdminActivityList(ctx context.Context, in *block.AdminActivityListRequest) (*block.AdminActivityListResponse, error) {
	l := logic.NewAdminActivityListLogic(ctx, s.svcCtx)
	return l.AdminActivityList(in)
}

func (s *BlockServer) StartActivity(ctx context.Context, in *block.StartActivityRequest) (*block.IsSuccessResponse, error) {
	l := logic.NewStartActivityLogic(ctx, s.svcCtx)
	return l.StartActivity(in)
}

// PC接口
func (s *BlockServer) GetGoodsList(ctx context.Context, in *block.GetGoodsListRequest) (*block.GetGoodsListResponse, error) {
	l := logic.NewGetGoodsListLogic(ctx, s.svcCtx)
	return l.GetGoodsList(in)
}

func (s *BlockServer) GetPurchaseRecord(ctx context.Context, in *block.GetPurchaseRecordRequest) (*block.GetPurchaseRecordResponse, error) {
	l := logic.NewGetPurchaseRecordLogic(ctx, s.svcCtx)
	return l.GetPurchaseRecord(in)
}

func (s *BlockServer) JudgeBargain(ctx context.Context, in *block.JudgeBargainRequest) (*block.JudgeBargainResponse, error) {
	l := logic.NewJudgeBargainLogic(ctx, s.svcCtx)
	return l.JudgeBargain(in)
}

func (s *BlockServer) CryptominerFullPurchase(ctx context.Context, in *block.CryptominerPurchaseRequest) (*block.IsSuccessResponse, error) {
	l := logic.NewCryptominerFullPurchaseLogic(ctx, s.svcCtx)
	return l.CryptominerFullPurchase(in)
}

func (s *BlockServer) PropPurchase(ctx context.Context, in *block.PropPurchaseRequest) (*block.IsSuccessResponse, error) {
	l := logic.NewPropPurchaseLogic(ctx, s.svcCtx)
	return l.PropPurchase(in)
}

func (s *BlockServer) CryptominerBargainPurchase(ctx context.Context, in *block.CryptominerBargainRequest) (*block.CryptominerBargainResponse, error) {
	l := logic.NewCryptominerBargainPurchaseLogic(ctx, s.svcCtx)
	return l.CryptominerBargainPurchase(in)
}

func (s *BlockServer) AssistorBargain(ctx context.Context, in *block.AssistorBargainRequest) (*block.AssistorBargainResponse, error) {
	l := logic.NewAssistorBargainLogic(ctx, s.svcCtx)
	return l.AssistorBargain(in)
}

func (s *BlockServer) GetBargainRecord(ctx context.Context, in *block.GetBargainRecordRequest) (*block.GetBargainRecordResponse, error) {
	l := logic.NewGetBargainRecordLogic(ctx, s.svcCtx)
	return l.GetBargainRecord(in)
}

func (s *BlockServer) BargainPay(ctx context.Context, in *block.BargainPayRequest) (*block.IsSuccessResponse, error) {
	l := logic.NewBargainPayLogic(ctx, s.svcCtx)
	return l.BargainPay(in)
}

// 外部rpc接口
func (s *BlockServer) JudgeGoodsPurchase(ctx context.Context, in *block.JudgeGoodsPurchaseRequest) (*block.JudgeGoodsPurchaseResponse, error) {
	l := logic.NewJudgeGoodsPurchaseLogic(ctx, s.svcCtx)
	return l.JudgeGoodsPurchase(in)
}

func (s *BlockServer) GetUserCryptominerStatus(ctx context.Context, in *block.GetUserCryptominerStatusRequest) (*block.GetUserCryptominerStatusResponse, error) {
	l := logic.NewGetUserCryptominerStatusLogic(ctx, s.svcCtx)
	return l.GetUserCryptominerStatus(in)
}

func (s *BlockServer) GetUserPropStatus(ctx context.Context, in *block.GetUserPropStatusRequest) (*block.GetUserPropStatusResponse, error) {
	l := logic.NewGetUserPropStatusLogic(ctx, s.svcCtx)
	return l.GetUserPropStatus(in)
}

func (s *BlockServer) ActivateCryptominer(ctx context.Context, in *block.ActivateCryptominerRequest) (*block.ActivateCryptominerResponse, error) {
	l := logic.NewActivateCryptominerLogic(ctx, s.svcCtx)
	return l.ActivateCryptominer(in)
}
