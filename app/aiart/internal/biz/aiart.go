package biz

import (
	"context"

	pb "github.com/ydssx/morphix/api/aiart/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AiartUseCase struct {
}

func NewAiartUseCase() *AiartUseCase {
	return &AiartUseCase{}
}

// 生成图像
func (*AiartUseCase) GenerateImage(ctx context.Context, req *pb.GenerateImageRequest) (res *pb.GenerateImageResponse, err error) {
	res = new(pb.GenerateImageResponse)

	// TODO:ADD logic here and delete this line.

	return
}

// 获取生成任务状态
func (*AiartUseCase) GetGenerateStatus(ctx context.Context, req *pb.GetGenerateStatusRequest) (res *pb.GenerateStatusResponse, err error) {
	res = new(pb.GenerateStatusResponse)

	// TODO:ADD logic here and delete this line.

	return
}

// 获取已生成的图像
func (*AiartUseCase) GetGeneratedImage(ctx context.Context, req *pb.GetGeneratedImageRequest) (res *pb.GetGeneratedImageResponse, err error) {
	res = new(pb.GetGeneratedImageResponse)

	// TODO:ADD logic here and delete this line.

	return
}

// 获取模型信息
func (*AiartUseCase) GetModelInfo(ctx context.Context, req *emptypb.Empty) (res *pb.GetModelInfoResponse, err error) {
	res = new(pb.GetModelInfoResponse)

	// TODO:ADD logic here and delete this line.

	return
}
