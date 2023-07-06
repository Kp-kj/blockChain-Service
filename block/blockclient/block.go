// Code generated by goctl. DO NOT EDIT.
// Source: block.proto

package blockclient

import (
	"context"

	"block/block"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BargainAmount                 = block.BargainAmount
	BargainRecord                 = block.BargainRecord
	CreateActivityRequest         = block.CreateActivityRequest
	CreateCryptominerRequest      = block.CreateCryptominerRequest
	CreatePropRequest             = block.CreatePropRequest
	Cryptominer                   = block.Cryptominer
	CryptominerBargainRequest     = block.CryptominerBargainRequest
	CryptominerBargainResponse    = block.CryptominerBargainResponse
	CryptominerPurchaseRequest    = block.CryptominerPurchaseRequest
	GetBargainCryptominerRequest  = block.GetBargainCryptominerRequest
	GetBargainCryptominerResponse = block.GetBargainCryptominerResponse
	GetBargainProgressRequest     = block.GetBargainProgressRequest
	GetBargainProgressResponse    = block.GetBargainProgressResponse
	GetBargainRuleRequest         = block.GetBargainRuleRequest
	GetBargainRuleResponse        = block.GetBargainRuleResponse
	GetGoodsListRequest           = block.GetGoodsListRequest
	GetGoodsListResponse          = block.GetGoodsListResponse
	GetPurchaseRecordRequest      = block.GetPurchaseRecordRequest
	GetPurchaseRecordResponse     = block.GetPurchaseRecordResponse
	IsSuccessResponse             = block.IsSuccessResponse
	JudgeBargainRequest           = block.JudgeBargainRequest
	JudgeBargainResponse          = block.JudgeBargainResponse
	JudgeGoodsPurchaseRequest     = block.JudgeGoodsPurchaseRequest
	JudgeGoodsPurchaseResponse    = block.JudgeGoodsPurchaseResponse
	OneBargain                    = block.OneBargain
	Profile                       = block.Profile
	Prop                          = block.Prop
	PropPurchaseRequest           = block.PropPurchaseRequest
	PurchaseRecord                = block.PurchaseRecord
	Request                       = block.Request
	Response                      = block.Response

	Block interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		CreateCryptominer(ctx context.Context, in *CreateCryptominerRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error)
		CreateProp(ctx context.Context, in *CreatePropRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error)
		CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error)
		GetGoodsList(ctx context.Context, in *GetGoodsListRequest, opts ...grpc.CallOption) (*GetGoodsListResponse, error)
		JudgeBargain(ctx context.Context, in *JudgeBargainRequest, opts ...grpc.CallOption) (*JudgeBargainResponse, error)
		CryptominerFullPurchase(ctx context.Context, in *CryptominerPurchaseRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error)
		CryptominerBargainPurchase(ctx context.Context, in *CryptominerBargainRequest, opts ...grpc.CallOption) (*CryptominerBargainResponse, error)
		PropPurchase(ctx context.Context, in *PropPurchaseRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error)
		GetBargainRule(ctx context.Context, in *GetBargainRuleRequest, opts ...grpc.CallOption) (*GetBargainRuleResponse, error)
		GetBargainCryptominer(ctx context.Context, in *GetBargainCryptominerRequest, opts ...grpc.CallOption) (*GetBargainCryptominerResponse, error)
		GetBargainProgress(ctx context.Context, in *GetBargainProgressRequest, opts ...grpc.CallOption) (*GetBargainProgressResponse, error)
		GetPurchaseRecord(ctx context.Context, in *GetPurchaseRecordRequest, opts ...grpc.CallOption) (*GetPurchaseRecordResponse, error)
		JudgeGoodsPurchase(ctx context.Context, in *JudgeGoodsPurchaseRequest, opts ...grpc.CallOption) (*JudgeGoodsPurchaseResponse, error)
	}

	defaultBlock struct {
		cli zrpc.Client
	}
)

func NewBlock(cli zrpc.Client) Block {
	return &defaultBlock{
		cli: cli,
	}
}

func (m *defaultBlock) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultBlock) CreateCryptominer(ctx context.Context, in *CreateCryptominerRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.CreateCryptominer(ctx, in, opts...)
}

func (m *defaultBlock) CreateProp(ctx context.Context, in *CreatePropRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.CreateProp(ctx, in, opts...)
}

func (m *defaultBlock) CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.CreateActivity(ctx, in, opts...)
}

func (m *defaultBlock) GetGoodsList(ctx context.Context, in *GetGoodsListRequest, opts ...grpc.CallOption) (*GetGoodsListResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.GetGoodsList(ctx, in, opts...)
}

func (m *defaultBlock) JudgeBargain(ctx context.Context, in *JudgeBargainRequest, opts ...grpc.CallOption) (*JudgeBargainResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.JudgeBargain(ctx, in, opts...)
}

func (m *defaultBlock) CryptominerFullPurchase(ctx context.Context, in *CryptominerPurchaseRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.CryptominerFullPurchase(ctx, in, opts...)
}

func (m *defaultBlock) CryptominerBargainPurchase(ctx context.Context, in *CryptominerBargainRequest, opts ...grpc.CallOption) (*CryptominerBargainResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.CryptominerBargainPurchase(ctx, in, opts...)
}

func (m *defaultBlock) PropPurchase(ctx context.Context, in *PropPurchaseRequest, opts ...grpc.CallOption) (*IsSuccessResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.PropPurchase(ctx, in, opts...)
}

func (m *defaultBlock) GetBargainRule(ctx context.Context, in *GetBargainRuleRequest, opts ...grpc.CallOption) (*GetBargainRuleResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.GetBargainRule(ctx, in, opts...)
}

func (m *defaultBlock) GetBargainCryptominer(ctx context.Context, in *GetBargainCryptominerRequest, opts ...grpc.CallOption) (*GetBargainCryptominerResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.GetBargainCryptominer(ctx, in, opts...)
}

func (m *defaultBlock) GetBargainProgress(ctx context.Context, in *GetBargainProgressRequest, opts ...grpc.CallOption) (*GetBargainProgressResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.GetBargainProgress(ctx, in, opts...)
}

func (m *defaultBlock) GetPurchaseRecord(ctx context.Context, in *GetPurchaseRecordRequest, opts ...grpc.CallOption) (*GetPurchaseRecordResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.GetPurchaseRecord(ctx, in, opts...)
}

func (m *defaultBlock) JudgeGoodsPurchase(ctx context.Context, in *JudgeGoodsPurchaseRequest, opts ...grpc.CallOption) (*JudgeGoodsPurchaseResponse, error) {
	client := block.NewBlockClient(m.cli.Conn())
	return client.JudgeGoodsPurchase(ctx, in, opts...)
}
