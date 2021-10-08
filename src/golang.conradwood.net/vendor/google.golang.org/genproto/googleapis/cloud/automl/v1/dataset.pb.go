// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1/dataset.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A workspace for solving a single, particular machine learning (ML) problem.
// A workspace contains examples that may be annotated.
type Dataset struct {
	// Required.
	// The dataset metadata that is specific to the problem type.
	//
	// Types that are valid to be assigned to DatasetMetadata:
	//	*Dataset_TranslationDatasetMetadata
	//	*Dataset_ImageClassificationDatasetMetadata
	//	*Dataset_TextClassificationDatasetMetadata
	//	*Dataset_ImageObjectDetectionDatasetMetadata
	//	*Dataset_TextExtractionDatasetMetadata
	//	*Dataset_TextSentimentDatasetMetadata
	DatasetMetadata isDataset_DatasetMetadata `protobuf_oneof:"dataset_metadata"`
	// Output only. The resource name of the dataset.
	// Form: `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of the dataset to show in the interface. The name can be
	// up to 32 characters long and can consist only of ASCII Latin letters A-Z
	// and a-z, underscores
	// (_), and ASCII digits 0-9.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// User-provided description of the dataset. The description can be up to
	// 25000 characters long.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. The number of examples in the dataset.
	ExampleCount int32 `protobuf:"varint,21,opt,name=example_count,json=exampleCount,proto3" json:"example_count,omitempty"`
	// Output only. Timestamp when this dataset was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,14,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Used to perform consistent read-modify-write updates. If not set, a blind
	// "overwrite" update happens.
	Etag string `protobuf:"bytes,17,opt,name=etag,proto3" json:"etag,omitempty"`
	// Optional. The labels with user-defined metadata to organize your dataset.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	// Label values are optional. Label keys must start with a letter.
	//
	// See https://goo.gl/xmQnxf for more information on and examples of labels.
	Labels               map[string]string `protobuf:"bytes,39,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Dataset) Reset()         { *m = Dataset{} }
func (m *Dataset) String() string { return proto.CompactTextString(m) }
func (*Dataset) ProtoMessage()    {}
func (*Dataset) Descriptor() ([]byte, []int) {
	return fileDescriptor_78541757d26dc96c, []int{0}
}

func (m *Dataset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dataset.Unmarshal(m, b)
}
func (m *Dataset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dataset.Marshal(b, m, deterministic)
}
func (m *Dataset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dataset.Merge(m, src)
}
func (m *Dataset) XXX_Size() int {
	return xxx_messageInfo_Dataset.Size(m)
}
func (m *Dataset) XXX_DiscardUnknown() {
	xxx_messageInfo_Dataset.DiscardUnknown(m)
}

var xxx_messageInfo_Dataset proto.InternalMessageInfo

type isDataset_DatasetMetadata interface {
	isDataset_DatasetMetadata()
}

type Dataset_TranslationDatasetMetadata struct {
	TranslationDatasetMetadata *TranslationDatasetMetadata `protobuf:"bytes,23,opt,name=translation_dataset_metadata,json=translationDatasetMetadata,proto3,oneof"`
}

type Dataset_ImageClassificationDatasetMetadata struct {
	ImageClassificationDatasetMetadata *ImageClassificationDatasetMetadata `protobuf:"bytes,24,opt,name=image_classification_dataset_metadata,json=imageClassificationDatasetMetadata,proto3,oneof"`
}

type Dataset_TextClassificationDatasetMetadata struct {
	TextClassificationDatasetMetadata *TextClassificationDatasetMetadata `protobuf:"bytes,25,opt,name=text_classification_dataset_metadata,json=textClassificationDatasetMetadata,proto3,oneof"`
}

type Dataset_ImageObjectDetectionDatasetMetadata struct {
	ImageObjectDetectionDatasetMetadata *ImageObjectDetectionDatasetMetadata `protobuf:"bytes,26,opt,name=image_object_detection_dataset_metadata,json=imageObjectDetectionDatasetMetadata,proto3,oneof"`
}

type Dataset_TextExtractionDatasetMetadata struct {
	TextExtractionDatasetMetadata *TextExtractionDatasetMetadata `protobuf:"bytes,28,opt,name=text_extraction_dataset_metadata,json=textExtractionDatasetMetadata,proto3,oneof"`
}

type Dataset_TextSentimentDatasetMetadata struct {
	TextSentimentDatasetMetadata *TextSentimentDatasetMetadata `protobuf:"bytes,30,opt,name=text_sentiment_dataset_metadata,json=textSentimentDatasetMetadata,proto3,oneof"`
}

func (*Dataset_TranslationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_ImageClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_ImageObjectDetectionDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextExtractionDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextSentimentDatasetMetadata) isDataset_DatasetMetadata() {}

func (m *Dataset) GetDatasetMetadata() isDataset_DatasetMetadata {
	if m != nil {
		return m.DatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTranslationDatasetMetadata() *TranslationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TranslationDatasetMetadata); ok {
		return x.TranslationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetImageClassificationDatasetMetadata() *ImageClassificationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_ImageClassificationDatasetMetadata); ok {
		return x.ImageClassificationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTextClassificationDatasetMetadata() *TextClassificationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TextClassificationDatasetMetadata); ok {
		return x.TextClassificationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetImageObjectDetectionDatasetMetadata() *ImageObjectDetectionDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_ImageObjectDetectionDatasetMetadata); ok {
		return x.ImageObjectDetectionDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTextExtractionDatasetMetadata() *TextExtractionDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TextExtractionDatasetMetadata); ok {
		return x.TextExtractionDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTextSentimentDatasetMetadata() *TextSentimentDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TextSentimentDatasetMetadata); ok {
		return x.TextSentimentDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Dataset) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Dataset) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Dataset) GetExampleCount() int32 {
	if m != nil {
		return m.ExampleCount
	}
	return 0
}

func (m *Dataset) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Dataset) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

func (m *Dataset) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Dataset) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Dataset_TranslationDatasetMetadata)(nil),
		(*Dataset_ImageClassificationDatasetMetadata)(nil),
		(*Dataset_TextClassificationDatasetMetadata)(nil),
		(*Dataset_ImageObjectDetectionDatasetMetadata)(nil),
		(*Dataset_TextExtractionDatasetMetadata)(nil),
		(*Dataset_TextSentimentDatasetMetadata)(nil),
	}
}

func init() {
	proto.RegisterType((*Dataset)(nil), "google.cloud.automl.v1.Dataset")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1.Dataset.LabelsEntry")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1/dataset.proto", fileDescriptor_78541757d26dc96c)
}

var fileDescriptor_78541757d26dc96c = []byte{
	// 639 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0xcd, 0xd6, 0x56, 0x3a, 0x5b, 0xa5, 0x0e, 0x5a, 0xd3, 0xb0, 0xda, 0xf4, 0x8f, 0x74,
	0x41, 0x48, 0xd8, 0xaa, 0x60, 0x53, 0x6f, 0xec, 0xb6, 0xa8, 0xd0, 0xaa, 0xac, 0xa5, 0x17, 0x52,
	0x08, 0xb3, 0xd9, 0x69, 0x88, 0x4e, 0x32, 0x21, 0x73, 0xb2, 0x6c, 0x2f, 0x04, 0x1f, 0x40, 0xf0,
	0xc2, 0x6b, 0x5f, 0xa0, 0x8f, 0xe2, 0xa3, 0xf8, 0x14, 0x32, 0x7f, 0x62, 0x6b, 0xb7, 0xd9, 0xbd,
	0x9b, 0x3d, 0xdf, 0x77, 0xce, 0xfc, 0xce, 0x97, 0x65, 0xd0, 0x46, 0xcc, 0x79, 0xcc, 0xa8, 0x1f,
	0x31, 0x5e, 0x0e, 0x7c, 0x52, 0x02, 0x4f, 0x99, 0x3f, 0xec, 0xf8, 0x03, 0x02, 0x44, 0x50, 0xf0,
	0xf2, 0x82, 0x03, 0xc7, 0x4b, 0xda, 0xe5, 0x29, 0x97, 0xa7, 0x5d, 0xde, 0xb0, 0xe3, 0x2c, 0x9b,
	0x6e, 0x92, 0x27, 0x7e, 0x41, 0x05, 0x2f, 0x8b, 0x88, 0xea, 0x16, 0x67, 0xad, 0x66, 0x70, 0x92,
	0x92, 0xb8, 0xf2, 0xac, 0xd6, 0x78, 0x80, 0x8e, 0xcc, 0xcd, 0x4e, 0xbb, 0xce, 0x52, 0x90, 0x4c,
	0x30, 0x02, 0x09, 0xcf, 0x8c, 0x73, 0xc5, 0x38, 0xd5, 0xaf, 0x7e, 0x79, 0xea, 0x43, 0x92, 0x52,
	0x01, 0x24, 0xcd, 0x8d, 0xa1, 0x75, 0x09, 0x96, 0x64, 0x19, 0x07, 0xd5, 0x2d, 0xb4, 0xba, 0xf6,
	0x6b, 0x1e, 0xdd, 0xda, 0xd3, 0x4b, 0xe3, 0x21, 0x6a, 0x5d, 0x9a, 0x1f, 0x9a, 0x2c, 0xc2, 0x94,
	0x02, 0x91, 0x67, 0xfb, 0x81, 0x6b, 0xb5, 0x9b, 0x5b, 0x5b, 0xde, 0xf5, 0xa9, 0x78, 0x47, 0x17,
	0xbd, 0x66, 0xe2, 0xa1, 0xe9, 0x7c, 0x73, 0xa3, 0xe7, 0x40, 0xad, 0x8a, 0x7f, 0x58, 0xe8, 0xb1,
	0xca, 0x27, 0x8c, 0x18, 0x11, 0x22, 0x39, 0x4d, 0xa2, 0x1a, 0x02, 0x5b, 0x11, 0x04, 0x75, 0x04,
	0x6f, 0xe5, 0x90, 0xee, 0x7f, 0x33, 0xc6, 0x49, 0xd6, 0x92, 0xa9, 0x2e, 0xfc, 0xdd, 0x42, 0x1b,
	0xf2, 0x6b, 0x4c, 0x05, 0x5a, 0x56, 0x40, 0xdb, 0xb5, 0x91, 0xd0, 0x11, 0x4c, 0xe3, 0x59, 0x85,
	0x69, 0x26, 0xfc, 0xd3, 0x42, 0x9b, 0x3a, 0x20, 0xde, 0xff, 0x4c, 0x23, 0x08, 0x07, 0x14, 0x68,
	0x74, 0x3d, 0x91, 0xa3, 0x88, 0x76, 0x26, 0x46, 0xf4, 0x5e, 0x4d, 0xd9, 0xab, 0x86, 0x8c, 0x33,
	0xad, 0x27, 0xd3, 0x6d, 0xf8, 0x9b, 0x85, 0x5c, 0x15, 0x12, 0x1d, 0x41, 0x41, 0x6a, 0x70, 0x5a,
	0x0a, 0xe7, 0xf9, 0xa4, 0x80, 0xf6, 0xff, 0xb5, 0x8f, 0x83, 0x3c, 0x84, 0x49, 0x06, 0xfc, 0x15,
	0xad, 0x28, 0x02, 0x41, 0x33, 0xf9, 0xb7, 0xcf, 0x60, 0x1c, 0xe0, 0x91, 0x02, 0x78, 0x36, 0x09,
	0xe0, 0x63, 0xd5, 0x3d, 0x7e, 0x7f, 0x0b, 0x26, 0xe8, 0x18, 0xa3, 0x9b, 0x19, 0x49, 0xa9, 0x6d,
	0xb9, 0x56, 0x7b, 0xbe, 0xa7, 0xce, 0x78, 0x15, 0x2d, 0x0c, 0x12, 0x91, 0x33, 0x72, 0x16, 0x2a,
	0xad, 0xa1, 0xb4, 0xa6, 0xa9, 0xbd, 0x93, 0x16, 0x17, 0x35, 0x07, 0x54, 0x44, 0x45, 0x92, 0xcb,
	0x9d, 0xec, 0x19, 0xe3, 0xb8, 0x28, 0xe1, 0x75, 0x74, 0x9b, 0x8e, 0x48, 0x9a, 0x33, 0x1a, 0x46,
	0xbc, 0xcc, 0xc0, 0xbe, 0xef, 0x5a, 0xed, 0xd9, 0xde, 0x82, 0x29, 0x76, 0x65, 0x0d, 0xef, 0xa0,
	0x66, 0x54, 0x50, 0x02, 0x34, 0x94, 0x74, 0xf6, 0x1d, 0xb5, 0xa8, 0x53, 0x2d, 0x5a, 0xbd, 0x07,
	0xde, 0x51, 0xf5, 0x1e, 0xf4, 0x90, 0xb6, 0xcb, 0x82, 0x44, 0xa7, 0x40, 0x62, 0xfb, 0xae, 0x46,
	0x97, 0x67, 0xdc, 0x45, 0x73, 0x8c, 0xf4, 0x29, 0x13, 0xf6, 0xa6, 0x3b, 0xd3, 0x6e, 0x6e, 0x3d,
	0xa9, 0x0b, 0xcd, 0xe4, 0xe0, 0x1d, 0x28, 0xf7, 0x7e, 0x06, 0xc5, 0x59, 0xcf, 0xb4, 0x3a, 0xdb,
	0xa8, 0x79, 0xa9, 0x8c, 0x17, 0xd1, 0xcc, 0x17, 0x7a, 0x66, 0x12, 0x92, 0x47, 0x7c, 0x0f, 0xcd,
	0x0e, 0x09, 0x2b, 0xab, 0x64, 0xf4, 0x8f, 0xa0, 0xf1, 0xc2, 0xda, 0xc5, 0x68, 0xf1, 0xea, 0xe7,
	0xdb, 0x3d, 0xb7, 0x90, 0x13, 0xf1, 0xb4, 0x86, 0xe4, 0x83, 0xf5, 0xe9, 0xa5, 0x51, 0x62, 0xce,
	0x48, 0x16, 0x7b, 0xbc, 0x88, 0xfd, 0x98, 0x66, 0x6a, 0x7b, 0x5f, 0x4b, 0x24, 0x4f, 0xc4, 0xd5,
	0x87, 0x74, 0x47, 0x9f, 0xce, 0x1b, 0x4b, 0xaf, 0x75, 0x7b, 0x57, 0x0d, 0x7e, 0x55, 0x02, 0x3f,
	0x3c, 0xf0, 0x8e, 0x3b, 0xbf, 0x2b, 0xe1, 0x44, 0x09, 0x27, 0x4a, 0x60, 0x27, 0xc7, 0x9d, 0x3f,
	0x8d, 0x65, 0x2d, 0x04, 0x81, 0x52, 0x82, 0x40, 0xf7, 0x04, 0xc1, 0x71, 0xa7, 0x3f, 0xa7, 0xae,
	0x7d, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0x82, 0x5c, 0x35, 0xa8, 0x5e, 0x06, 0x00, 0x00,
}
