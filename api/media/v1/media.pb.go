// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: api/media/v1/media.proto

package mediav1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 视频上传请求
type VideoUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoData []byte `protobuf:"bytes,1,opt,name=videoData,proto3" json:"videoData,omitempty"`
	Filename  string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"` // 其他字段...
}

func (x *VideoUploadRequest) Reset() {
	*x = VideoUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_media_v1_media_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoUploadRequest) ProtoMessage() {}

func (x *VideoUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_media_v1_media_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoUploadRequest.ProtoReflect.Descriptor instead.
func (*VideoUploadRequest) Descriptor() ([]byte, []int) {
	return file_api_media_v1_media_proto_rawDescGZIP(), []int{0}
}

func (x *VideoUploadRequest) GetVideoData() []byte {
	if x != nil {
		return x.VideoData
	}
	return nil
}

func (x *VideoUploadRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

// 视频上传响应
type VideoUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId string `protobuf:"bytes,1,opt,name=videoId,proto3" json:"videoId,omitempty"` // 其他字段...
}

func (x *VideoUploadResponse) Reset() {
	*x = VideoUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_media_v1_media_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoUploadResponse) ProtoMessage() {}

func (x *VideoUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_media_v1_media_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoUploadResponse.ProtoReflect.Descriptor instead.
func (*VideoUploadResponse) Descriptor() ([]byte, []int) {
	return file_api_media_v1_media_proto_rawDescGZIP(), []int{1}
}

func (x *VideoUploadResponse) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

// 视频转码请求
type VideoTranscodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId       string   `protobuf:"bytes,1,opt,name=videoId,proto3" json:"videoId,omitempty"`
	OutputFormats []string `protobuf:"bytes,2,rep,name=outputFormats,proto3" json:"outputFormats,omitempty"` // 其他字段...
}

func (x *VideoTranscodeRequest) Reset() {
	*x = VideoTranscodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_media_v1_media_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoTranscodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoTranscodeRequest) ProtoMessage() {}

func (x *VideoTranscodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_media_v1_media_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoTranscodeRequest.ProtoReflect.Descriptor instead.
func (*VideoTranscodeRequest) Descriptor() ([]byte, []int) {
	return file_api_media_v1_media_proto_rawDescGZIP(), []int{2}
}

func (x *VideoTranscodeRequest) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

func (x *VideoTranscodeRequest) GetOutputFormats() []string {
	if x != nil {
		return x.OutputFormats
	}
	return nil
}

// 视频转码响应
type VideoTranscodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TranscodedVideoIds []string `protobuf:"bytes,1,rep,name=transcodedVideoIds,proto3" json:"transcodedVideoIds,omitempty"` // 其他字段...
}

func (x *VideoTranscodeResponse) Reset() {
	*x = VideoTranscodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_media_v1_media_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoTranscodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoTranscodeResponse) ProtoMessage() {}

func (x *VideoTranscodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_media_v1_media_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoTranscodeResponse.ProtoReflect.Descriptor instead.
func (*VideoTranscodeResponse) Descriptor() ([]byte, []int) {
	return file_api_media_v1_media_proto_rawDescGZIP(), []int{3}
}

func (x *VideoTranscodeResponse) GetTranscodedVideoIds() []string {
	if x != nil {
		return x.TranscodedVideoIds
	}
	return nil
}

// 视频播放请求
type VideoPlaybackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId string `protobuf:"bytes,1,opt,name=videoId,proto3" json:"videoId,omitempty"` // 其他字段...
}

func (x *VideoPlaybackRequest) Reset() {
	*x = VideoPlaybackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_media_v1_media_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoPlaybackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoPlaybackRequest) ProtoMessage() {}

func (x *VideoPlaybackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_media_v1_media_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoPlaybackRequest.ProtoReflect.Descriptor instead.
func (*VideoPlaybackRequest) Descriptor() ([]byte, []int) {
	return file_api_media_v1_media_proto_rawDescGZIP(), []int{4}
}

func (x *VideoPlaybackRequest) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

// 视频播放响应
type VideoPlaybackResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlaybackUrl string `protobuf:"bytes,1,opt,name=playbackUrl,proto3" json:"playbackUrl,omitempty"` // 其他字段...
}

func (x *VideoPlaybackResponse) Reset() {
	*x = VideoPlaybackResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_media_v1_media_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoPlaybackResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoPlaybackResponse) ProtoMessage() {}

func (x *VideoPlaybackResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_media_v1_media_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoPlaybackResponse.ProtoReflect.Descriptor instead.
func (*VideoPlaybackResponse) Descriptor() ([]byte, []int) {
	return file_api_media_v1_media_proto_rawDescGZIP(), []int{5}
}

func (x *VideoPlaybackResponse) GetPlaybackUrl() string {
	if x != nil {
		return x.PlaybackUrl
	}
	return ""
}

// 视频截图请求
type VideoScreenshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId   string  `protobuf:"bytes,1,opt,name=videoId,proto3" json:"videoId,omitempty"`
	Timestamp float64 `protobuf:"fixed64,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // 其他字段...
}

func (x *VideoScreenshotRequest) Reset() {
	*x = VideoScreenshotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_media_v1_media_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoScreenshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoScreenshotRequest) ProtoMessage() {}

func (x *VideoScreenshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_media_v1_media_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoScreenshotRequest.ProtoReflect.Descriptor instead.
func (*VideoScreenshotRequest) Descriptor() ([]byte, []int) {
	return file_api_media_v1_media_proto_rawDescGZIP(), []int{6}
}

func (x *VideoScreenshotRequest) GetVideoId() string {
	if x != nil {
		return x.VideoId
	}
	return ""
}

func (x *VideoScreenshotRequest) GetTimestamp() float64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// 视频截图响应
type VideoScreenshotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScreenshotUrl string `protobuf:"bytes,1,opt,name=screenshotUrl,proto3" json:"screenshotUrl,omitempty"` // 其他字段...
}

func (x *VideoScreenshotResponse) Reset() {
	*x = VideoScreenshotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_media_v1_media_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoScreenshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoScreenshotResponse) ProtoMessage() {}

func (x *VideoScreenshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_media_v1_media_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoScreenshotResponse.ProtoReflect.Descriptor instead.
func (*VideoScreenshotResponse) Descriptor() ([]byte, []int) {
	return file_api_media_v1_media_proto_rawDescGZIP(), []int{7}
}

func (x *VideoScreenshotResponse) GetScreenshotUrl() string {
	if x != nil {
		return x.ScreenshotUrl
	}
	return ""
}

var File_api_media_v1_media_proto protoreflect.FileDescriptor

var file_api_media_v1_media_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x76, 0x31, 0x22, 0x4e, 0x0a, 0x12, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x13, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x15, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x73, 0x22, 0x48, 0x0a,
	0x16, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x63, 0x6f, 0x64, 0x65, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x73, 0x22, 0x30, 0x0a, 0x14, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x50, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x15, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63,
	0x6b, 0x55, 0x72, 0x6c, 0x22, 0x50, 0x0a, 0x16, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x63, 0x72,
	0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x3f, 0x0a, 0x17, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53,
	0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74, 0x55,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x73, 0x68, 0x6f, 0x74, 0x55, 0x72, 0x6c, 0x32, 0xcc, 0x02, 0x0a, 0x0c, 0x4d, 0x65, 0x64, 0x69,
	0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x76,
	0x31, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x76, 0x31, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x76, 0x31, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x76, 0x31, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x79, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x12, 0x1d, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x50, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x50, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x53, 0x0a, 0x0e, 0x54, 0x61, 0x6b, 0x65, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73,
	0x68, 0x6f, 0x74, 0x12, 0x1f, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x76, 0x31, 0x2e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x76, 0x31, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x64, 0x73, 0x73, 0x78, 0x2f, 0x6d, 0x6f, 0x72, 0x70, 0x68,
	0x69, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2f, 0x76, 0x31, 0x3b,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_media_v1_media_proto_rawDescOnce sync.Once
	file_api_media_v1_media_proto_rawDescData = file_api_media_v1_media_proto_rawDesc
)

func file_api_media_v1_media_proto_rawDescGZIP() []byte {
	file_api_media_v1_media_proto_rawDescOnce.Do(func() {
		file_api_media_v1_media_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_media_v1_media_proto_rawDescData)
	})
	return file_api_media_v1_media_proto_rawDescData
}

var file_api_media_v1_media_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_media_v1_media_proto_goTypes = []interface{}{
	(*VideoUploadRequest)(nil),      // 0: mediav1.VideoUploadRequest
	(*VideoUploadResponse)(nil),     // 1: mediav1.VideoUploadResponse
	(*VideoTranscodeRequest)(nil),   // 2: mediav1.VideoTranscodeRequest
	(*VideoTranscodeResponse)(nil),  // 3: mediav1.VideoTranscodeResponse
	(*VideoPlaybackRequest)(nil),    // 4: mediav1.VideoPlaybackRequest
	(*VideoPlaybackResponse)(nil),   // 5: mediav1.VideoPlaybackResponse
	(*VideoScreenshotRequest)(nil),  // 6: mediav1.VideoScreenshotRequest
	(*VideoScreenshotResponse)(nil), // 7: mediav1.VideoScreenshotResponse
}
var file_api_media_v1_media_proto_depIdxs = []int32{
	0, // 0: mediav1.MediaService.UploadVideo:input_type -> mediav1.VideoUploadRequest
	2, // 1: mediav1.MediaService.TranscodeVideo:input_type -> mediav1.VideoTranscodeRequest
	4, // 2: mediav1.MediaService.PlayVideo:input_type -> mediav1.VideoPlaybackRequest
	6, // 3: mediav1.MediaService.TakeScreenshot:input_type -> mediav1.VideoScreenshotRequest
	1, // 4: mediav1.MediaService.UploadVideo:output_type -> mediav1.VideoUploadResponse
	3, // 5: mediav1.MediaService.TranscodeVideo:output_type -> mediav1.VideoTranscodeResponse
	5, // 6: mediav1.MediaService.PlayVideo:output_type -> mediav1.VideoPlaybackResponse
	7, // 7: mediav1.MediaService.TakeScreenshot:output_type -> mediav1.VideoScreenshotResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_media_v1_media_proto_init() }
func file_api_media_v1_media_proto_init() {
	if File_api_media_v1_media_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_media_v1_media_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoUploadRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_media_v1_media_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoUploadResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_media_v1_media_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoTranscodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_media_v1_media_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoTranscodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_media_v1_media_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoPlaybackRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_media_v1_media_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoPlaybackResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_media_v1_media_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoScreenshotRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_media_v1_media_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoScreenshotResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_media_v1_media_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_media_v1_media_proto_goTypes,
		DependencyIndexes: file_api_media_v1_media_proto_depIdxs,
		MessageInfos:      file_api_media_v1_media_proto_msgTypes,
	}.Build()
	File_api_media_v1_media_proto = out.File
	file_api_media_v1_media_proto_rawDesc = nil
	file_api_media_v1_media_proto_goTypes = nil
	file_api_media_v1_media_proto_depIdxs = nil
}
