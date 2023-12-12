package biz

import (
	"context"

	aiartv1 "github.com/ydssx/morphix/api/aiart/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)
// Transaction is a transaction interface.
type Transaction interface {
	// InTx executes f in a transaction.
	// If f returns a non-nil error, the transaction will be rolled back.
	// If f returns nil, the transaction will be committed.
	InTx(context.Context, func(ctx context.Context) error) error
}

type AiartUseCase struct{}

func NewAiartUseCase() *AiartUseCase {
	return &AiartUseCase{}
}

// 生成图像
func (uc *AiartUseCase) GenerateImage(ctx context.Context, req *aiartv1.GenerateImageRequest) (res *aiartv1.GenerateImageResponse, err error) {
	res = new(aiartv1.GenerateImageResponse)

	// TODO:ADD logic here and delete this line.

	return
}

// 获取生成任务状态
func (uc *AiartUseCase) GetGenerateStatus(ctx context.Context, req *aiartv1.GetGenerateStatusRequest) (res *aiartv1.GenerateStatusResponse, err error) {
	res = new(aiartv1.GenerateStatusResponse)

	// TODO:ADD logic here and delete this line.

	return
}

// 获取已生成的图像
func (uc *AiartUseCase) GetGeneratedImage(ctx context.Context, req *aiartv1.GetGeneratedImageRequest) (res *aiartv1.GetGeneratedImageResponse, err error) {
	res = new(aiartv1.GetGeneratedImageResponse)

	// TODO:ADD logic here and delete this line.

	return
}

// 获取模型信息
func (uc *AiartUseCase) GetModelInfo(ctx context.Context, req *emptypb.Empty) (res *aiartv1.GetModelInfoResponse, err error) {
	res = new(aiartv1.GetModelInfoResponse)

	// TODO:ADD logic here and delete this line.

	return
}
