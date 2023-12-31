// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models_spots.proto

package spots

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// QueryType - how do we search - in circle or in square
type QueryType int32

const (
	QueryType_QUERY_TYPE_INVALID QueryType = 0
	QueryType_QUERY_TYPE_CIRCLE  QueryType = 1
	QueryType_QUERY_TYPE_SQUARE  QueryType = 2
)

var QueryType_name = map[int32]string{
	0: "QUERY_TYPE_INVALID",
	1: "QUERY_TYPE_CIRCLE",
	2: "QUERY_TYPE_SQUARE",
}

var QueryType_value = map[string]int32{
	"QUERY_TYPE_INVALID": 0,
	"QUERY_TYPE_CIRCLE":  1,
	"QUERY_TYPE_SQUARE":  2,
}

func (x QueryType) String() string {
	return proto.EnumName(QueryType_name, int32(x))
}

func (QueryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad90d2f3a7059ae1, []int{0}
}

// ArtworkTokenIdRequest is request model
// for ListSpotsNearPoint
type ListSpotsNearPointRequest struct {
	Longitude            float64   `protobuf:"fixed64,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude             float64   `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Length               float32   `protobuf:"fixed32,3,opt,name=length,proto3" json:"length,omitempty"`
	SearchType           QueryType `protobuf:"varint,4,opt,name=search_type,json=searchType,proto3,enum=spots.models.QueryType" json:"search_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListSpotsNearPointRequest) Reset()         { *m = ListSpotsNearPointRequest{} }
func (m *ListSpotsNearPointRequest) String() string { return proto.CompactTextString(m) }
func (*ListSpotsNearPointRequest) ProtoMessage()    {}
func (*ListSpotsNearPointRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad90d2f3a7059ae1, []int{0}
}

func (m *ListSpotsNearPointRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSpotsNearPointRequest.Unmarshal(m, b)
}
func (m *ListSpotsNearPointRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSpotsNearPointRequest.Marshal(b, m, deterministic)
}
func (m *ListSpotsNearPointRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSpotsNearPointRequest.Merge(m, src)
}
func (m *ListSpotsNearPointRequest) XXX_Size() int {
	return xxx_messageInfo_ListSpotsNearPointRequest.Size(m)
}
func (m *ListSpotsNearPointRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSpotsNearPointRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSpotsNearPointRequest proto.InternalMessageInfo

func (m *ListSpotsNearPointRequest) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *ListSpotsNearPointRequest) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *ListSpotsNearPointRequest) GetLength() float32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *ListSpotsNearPointRequest) GetSearchType() QueryType {
	if m != nil {
		return m.SearchType
	}
	return QueryType_QUERY_TYPE_INVALID
}

// ChangeArtworkTokenIdRequest is request model
// for ListSpotsNearPoint
type ListSpotsNearPointResponse struct {
	Spots                []*Spot  `protobuf:"bytes,1,rep,name=spots,proto3" json:"spots,omitempty"`
	Total                int64    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSpotsNearPointResponse) Reset()         { *m = ListSpotsNearPointResponse{} }
func (m *ListSpotsNearPointResponse) String() string { return proto.CompactTextString(m) }
func (*ListSpotsNearPointResponse) ProtoMessage()    {}
func (*ListSpotsNearPointResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad90d2f3a7059ae1, []int{1}
}

func (m *ListSpotsNearPointResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSpotsNearPointResponse.Unmarshal(m, b)
}
func (m *ListSpotsNearPointResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSpotsNearPointResponse.Marshal(b, m, deterministic)
}
func (m *ListSpotsNearPointResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSpotsNearPointResponse.Merge(m, src)
}
func (m *ListSpotsNearPointResponse) XXX_Size() int {
	return xxx_messageInfo_ListSpotsNearPointResponse.Size(m)
}
func (m *ListSpotsNearPointResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSpotsNearPointResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSpotsNearPointResponse proto.InternalMessageInfo

func (m *ListSpotsNearPointResponse) GetSpots() []*Spot {
	if m != nil {
		return m.Spots
	}
	return nil
}

func (m *ListSpotsNearPointResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

// Spot model
type Spot struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Rating               float32  `protobuf:"fixed32,2,opt,name=rating,proto3" json:"rating,omitempty"`
	Website              string   `protobuf:"bytes,3,opt,name=website,proto3" json:"website,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Coordinate           string   `protobuf:"bytes,5,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
	Uuid                 string   `protobuf:"bytes,6,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Spot) Reset()         { *m = Spot{} }
func (m *Spot) String() string { return proto.CompactTextString(m) }
func (*Spot) ProtoMessage()    {}
func (*Spot) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad90d2f3a7059ae1, []int{2}
}

func (m *Spot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Spot.Unmarshal(m, b)
}
func (m *Spot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Spot.Marshal(b, m, deterministic)
}
func (m *Spot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spot.Merge(m, src)
}
func (m *Spot) XXX_Size() int {
	return xxx_messageInfo_Spot.Size(m)
}
func (m *Spot) XXX_DiscardUnknown() {
	xxx_messageInfo_Spot.DiscardUnknown(m)
}

var xxx_messageInfo_Spot proto.InternalMessageInfo

func (m *Spot) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Spot) GetRating() float32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *Spot) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Spot) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Spot) GetCoordinate() string {
	if m != nil {
		return m.Coordinate
	}
	return ""
}

func (m *Spot) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

// Empty is just empty message
// that could return error
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad90d2f3a7059ae1, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("spots.models.QueryType", QueryType_name, QueryType_value)
	proto.RegisterType((*ListSpotsNearPointRequest)(nil), "spots.models.ListSpotsNearPointRequest")
	proto.RegisterType((*ListSpotsNearPointResponse)(nil), "spots.models.ListSpotsNearPointResponse")
	proto.RegisterType((*Spot)(nil), "spots.models.Spot")
	proto.RegisterType((*Empty)(nil), "spots.models.Empty")
}

func init() {
	proto.RegisterFile("models_spots.proto", fileDescriptor_ad90d2f3a7059ae1)
}

var fileDescriptor_ad90d2f3a7059ae1 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x8b, 0xd3, 0x40,
	0x14, 0xc6, 0x9d, 0xb4, 0xe9, 0x9a, 0x57, 0xd1, 0xee, 0x43, 0xd7, 0xb8, 0x88, 0x84, 0x9c, 0x82,
	0x87, 0x0a, 0xeb, 0xc5, 0xeb, 0xba, 0xe6, 0x50, 0x28, 0xcb, 0x66, 0x76, 0x57, 0x58, 0x11, 0x42,
	0x9a, 0x0c, 0xed, 0x40, 0x3a, 0x13, 0x33, 0x2f, 0x48, 0xfe, 0x1d, 0x2f, 0xfe, 0x9b, 0x92, 0x99,
	0x5a, 0xa3, 0x78, 0x7b, 0xdf, 0xef, 0x7d, 0x33, 0x7c, 0xdf, 0x30, 0x80, 0x7b, 0x5d, 0x89, 0xda,
	0xe4, 0xa6, 0xd1, 0x64, 0x96, 0x4d, 0xab, 0x49, 0xe3, 0x13, 0x27, 0xdc, 0x26, 0xfe, 0xc9, 0xe0,
	0xd5, 0x5a, 0x1a, 0xba, 0x1d, 0xe0, 0xb5, 0x28, 0xda, 0x1b, 0x2d, 0x15, 0x71, 0xf1, 0xad, 0x13,
	0x86, 0xf0, 0x35, 0x04, 0xb5, 0x56, 0x5b, 0x49, 0x5d, 0x25, 0x42, 0x16, 0xb1, 0x84, 0xf1, 0x3f,
	0x00, 0xcf, 0xe1, 0x71, 0x5d, 0x90, 0x5b, 0x7a, 0x76, 0x79, 0xd4, 0x78, 0x06, 0xb3, 0x5a, 0xa8,
	0x2d, 0xed, 0xc2, 0x49, 0xc4, 0x12, 0x8f, 0x1f, 0x14, 0x7e, 0x80, 0xb9, 0x11, 0x45, 0x5b, 0xee,
	0x72, 0xea, 0x1b, 0x11, 0x4e, 0x23, 0x96, 0x3c, 0xbd, 0x78, 0xb9, 0x1c, 0x67, 0x5a, 0x66, 0x9d,
	0x68, 0xfb, 0xbb, 0xbe, 0x11, 0x1c, 0x9c, 0x77, 0x98, 0xe3, 0xaf, 0x70, 0xfe, 0xbf, 0xa0, 0xa6,
	0xd1, 0xca, 0x08, 0x4c, 0xc0, 0xb7, 0x77, 0x84, 0x2c, 0x9a, 0x24, 0xf3, 0x0b, 0xfc, 0xfb, 0xc6,
	0xe1, 0x10, 0x77, 0x06, 0x7c, 0x0e, 0x3e, 0x69, 0x2a, 0x6a, 0x1b, 0x79, 0xc2, 0x9d, 0x88, 0x7f,
	0x30, 0x98, 0x0e, 0x2e, 0x44, 0x98, 0xaa, 0x62, 0xef, 0xda, 0x06, 0xdc, 0xce, 0x43, 0x99, 0xb6,
	0x20, 0xa9, 0xb6, 0xf6, 0x8c, 0xc7, 0x0f, 0x0a, 0x43, 0x38, 0xf9, 0x2e, 0x36, 0x46, 0x92, 0xb0,
	0x2d, 0x03, 0xfe, 0x5b, 0x62, 0x04, 0xf3, 0x4a, 0x98, 0xb2, 0x95, 0x0d, 0x49, 0xad, 0x6c, 0xcd,
	0x80, 0x8f, 0x11, 0xbe, 0x01, 0x28, 0xb5, 0x6e, 0x2b, 0xa9, 0x0a, 0x12, 0xa1, 0x6f, 0x0d, 0x23,
	0x32, 0xe4, 0xe8, 0x3a, 0x59, 0x85, 0x33, 0x97, 0x63, 0x98, 0xe3, 0x13, 0xf0, 0xd3, 0x7d, 0x43,
	0xfd, 0xdb, 0x0c, 0x82, 0xe3, 0x23, 0xe1, 0x19, 0x60, 0x76, 0x9f, 0xf2, 0x87, 0xfc, 0xee, 0xe1,
	0x26, 0xcd, 0x57, 0xd7, 0x9f, 0x2f, 0xd7, 0xab, 0x4f, 0x8b, 0x47, 0xf8, 0x02, 0x4e, 0x47, 0xfc,
	0x6a, 0xc5, 0xaf, 0xd6, 0xe9, 0x82, 0xfd, 0x83, 0x6f, 0xb3, 0xfb, 0x4b, 0x9e, 0x2e, 0xbc, 0x8f,
	0xa7, 0x5f, 0x9e, 0xd9, 0xff, 0x51, 0xea, 0xda, 0xbc, 0xb3, 0x2f, 0xb5, 0x99, 0x59, 0xf0, 0xfe,
	0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x74, 0x3b, 0x58, 0x46, 0x02, 0x00, 0x00,
}
