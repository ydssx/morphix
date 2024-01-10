package service

import (
	"context"

	aiartv1 "github.com/ydssx/morphix/api/aiart/v1"
	"github.com/ydssx/morphix/app/aiart/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArtService struct {
	uc *biz.AiartUseCase

	aiartv1.UnimplementedArtServiceServer
}

func NewArtService(uc *biz.AiartUseCase) *ArtService {
	return &ArtService{uc: uc}
}

// 生成图像
func (s *ArtService) GenerateImage(ctx context.Context, req *aiartv1.GenerateImageRequest) (res *aiartv1.GenerateImageResponse, err error) {
	return s.uc.GenerateImage(ctx, req)
}

// 获取生成任务状态
func (s *ArtService) GetGenerateStatus(ctx context.Context, req *aiartv1.GetGenerateStatusRequest) (res *aiartv1.GenerateStatusResponse, err error) {
	return s.uc.GetGenerateStatus(ctx, req)
}

// 获取已生成的图像
func (s *ArtService) GetGeneratedImage(ctx context.Context, req *aiartv1.GetGeneratedImageRequest) (res *aiartv1.GetGeneratedImageResponse, err error) {
	return s.uc.GetGeneratedImage(ctx, req)
}

// 获取模型信息
func (s *ArtService) GetModelInfo(ctx context.Context, req *emptypb.Empty) (res *aiartv1.GetModelInfoResponse, err error) {
	return s.uc.GetModelInfo(ctx, req)
}

func (s *ArtService) ImageToImage(ctx context.Context, req *aiartv1.ImageToImageRequest) (res *aiartv1.ImageToImageResponse, err error) {
	return s.uc.ImageToImage(ctx, req)
}
