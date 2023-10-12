package service

import (
	"context"

	pb "github.com/ydssx/morphix/api/aiart/v1"
	"github.com/ydssx/morphix/app/aiart/internal/biz"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ArtService struct {
	uc *biz.AiartUseCase

	pb.UnimplementedArtServiceServer
}

func NewArtService(uc *biz.AiartUseCase) *ArtService {
	return &ArtService{uc: uc}
}

// 生成图像
func (s *ArtService) GenerateImage(ctx context.Context, req *pb.GenerateImageRequest) (res *pb.GenerateImageResponse, err error) {
	return s.uc.GenerateImage(ctx, req)
}

// 获取生成任务状态
func (s *ArtService) GetGenerateStatus(ctx context.Context, req *pb.GetGenerateStatusRequest) (res *pb.GenerateStatusResponse, err error) {
	return s.uc.GetGenerateStatus(ctx, req)
}

// 获取已生成的图像
func (s *ArtService) GetGeneratedImage(ctx context.Context, req *pb.GetGeneratedImageRequest) (res *pb.GetGeneratedImageResponse, err error) {
	return s.uc.GetGeneratedImage(ctx, req)
}

// 获取模型信息
func (s *ArtService) GetModelInfo(ctx context.Context, req *emptypb.Empty) (res *pb.GetModelInfoResponse, err error) {
	return s.uc.GetModelInfo(ctx, req)
}
