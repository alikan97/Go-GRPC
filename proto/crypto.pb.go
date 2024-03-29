// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/crypto.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type Asset struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dda53b2096cbd59, []int{0}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetAssetReq struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAssetReq) Reset()         { *m = GetAssetReq{} }
func (m *GetAssetReq) String() string { return proto.CompactTextString(m) }
func (*GetAssetReq) ProtoMessage()    {}
func (*GetAssetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dda53b2096cbd59, []int{1}
}

func (m *GetAssetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAssetReq.Unmarshal(m, b)
}
func (m *GetAssetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAssetReq.Marshal(b, m, deterministic)
}
func (m *GetAssetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAssetReq.Merge(m, src)
}
func (m *GetAssetReq) XXX_Size() int {
	return xxx_messageInfo_GetAssetReq.Size(m)
}
func (m *GetAssetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAssetReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetAssetReq proto.InternalMessageInfo

func (m *GetAssetReq) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type GetAssetResp struct {
	Response             *Asset   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAssetResp) Reset()         { *m = GetAssetResp{} }
func (m *GetAssetResp) String() string { return proto.CompactTextString(m) }
func (*GetAssetResp) ProtoMessage()    {}
func (*GetAssetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dda53b2096cbd59, []int{2}
}

func (m *GetAssetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAssetResp.Unmarshal(m, b)
}
func (m *GetAssetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAssetResp.Marshal(b, m, deterministic)
}
func (m *GetAssetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAssetResp.Merge(m, src)
}
func (m *GetAssetResp) XXX_Size() int {
	return xxx_messageInfo_GetAssetResp.Size(m)
}
func (m *GetAssetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAssetResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetAssetResp proto.InternalMessageInfo

func (m *GetAssetResp) GetResponse() *Asset {
	if m != nil {
		return m.Response
	}
	return nil
}

type GetAllAssetResp struct {
	Response             []*Asset `protobuf:"bytes,1,rep,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllAssetResp) Reset()         { *m = GetAllAssetResp{} }
func (m *GetAllAssetResp) String() string { return proto.CompactTextString(m) }
func (*GetAllAssetResp) ProtoMessage()    {}
func (*GetAllAssetResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dda53b2096cbd59, []int{3}
}

func (m *GetAllAssetResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllAssetResp.Unmarshal(m, b)
}
func (m *GetAllAssetResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllAssetResp.Marshal(b, m, deterministic)
}
func (m *GetAllAssetResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllAssetResp.Merge(m, src)
}
func (m *GetAllAssetResp) XXX_Size() int {
	return xxx_messageInfo_GetAllAssetResp.Size(m)
}
func (m *GetAllAssetResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllAssetResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllAssetResp proto.InternalMessageInfo

func (m *GetAllAssetResp) GetResponse() []*Asset {
	if m != nil {
		return m.Response
	}
	return nil
}

type RecentTradesReq struct {
	Since                uint64   `protobuf:"varint,1,opt,name=since,proto3" json:"since,omitempty"`
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecentTradesReq) Reset()         { *m = RecentTradesReq{} }
func (m *RecentTradesReq) String() string { return proto.CompactTextString(m) }
func (*RecentTradesReq) ProtoMessage()    {}
func (*RecentTradesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dda53b2096cbd59, []int{4}
}

func (m *RecentTradesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecentTradesReq.Unmarshal(m, b)
}
func (m *RecentTradesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecentTradesReq.Marshal(b, m, deterministic)
}
func (m *RecentTradesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecentTradesReq.Merge(m, src)
}
func (m *RecentTradesReq) XXX_Size() int {
	return xxx_messageInfo_RecentTradesReq.Size(m)
}
func (m *RecentTradesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RecentTradesReq.DiscardUnknown(m)
}

var xxx_messageInfo_RecentTradesReq proto.InternalMessageInfo

func (m *RecentTradesReq) GetSince() uint64 {
	if m != nil {
		return m.Since
	}
	return 0
}

func (m *RecentTradesReq) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type RecentTradesResp struct {
	SymbolName           string   `protobuf:"bytes,1,opt,name=symbolName,proto3" json:"symbolName,omitempty"`
	Quantity             float64  `protobuf:"fixed64,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Time                 uint64   `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecentTradesResp) Reset()         { *m = RecentTradesResp{} }
func (m *RecentTradesResp) String() string { return proto.CompactTextString(m) }
func (*RecentTradesResp) ProtoMessage()    {}
func (*RecentTradesResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dda53b2096cbd59, []int{5}
}

func (m *RecentTradesResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecentTradesResp.Unmarshal(m, b)
}
func (m *RecentTradesResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecentTradesResp.Marshal(b, m, deterministic)
}
func (m *RecentTradesResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecentTradesResp.Merge(m, src)
}
func (m *RecentTradesResp) XXX_Size() int {
	return xxx_messageInfo_RecentTradesResp.Size(m)
}
func (m *RecentTradesResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RecentTradesResp.DiscardUnknown(m)
}

var xxx_messageInfo_RecentTradesResp proto.InternalMessageInfo

func (m *RecentTradesResp) GetSymbolName() string {
	if m != nil {
		return m.SymbolName
	}
	return ""
}

func (m *RecentTradesResp) GetQuantity() float64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *RecentTradesResp) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type RecentTradesMultiple struct {
	ListResponse         []*RecentTradesResp `protobuf:"bytes,1,rep,name=listResponse,proto3" json:"listResponse,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RecentTradesMultiple) Reset()         { *m = RecentTradesMultiple{} }
func (m *RecentTradesMultiple) String() string { return proto.CompactTextString(m) }
func (*RecentTradesMultiple) ProtoMessage()    {}
func (*RecentTradesMultiple) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dda53b2096cbd59, []int{6}
}

func (m *RecentTradesMultiple) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecentTradesMultiple.Unmarshal(m, b)
}
func (m *RecentTradesMultiple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecentTradesMultiple.Marshal(b, m, deterministic)
}
func (m *RecentTradesMultiple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecentTradesMultiple.Merge(m, src)
}
func (m *RecentTradesMultiple) XXX_Size() int {
	return xxx_messageInfo_RecentTradesMultiple.Size(m)
}
func (m *RecentTradesMultiple) XXX_DiscardUnknown() {
	xxx_messageInfo_RecentTradesMultiple.DiscardUnknown(m)
}

var xxx_messageInfo_RecentTradesMultiple proto.InternalMessageInfo

func (m *RecentTradesMultiple) GetListResponse() []*RecentTradesResp {
	if m != nil {
		return m.ListResponse
	}
	return nil
}

type AddRecentTradeReq struct {
	AssetName            string   `protobuf:"bytes,1,opt,name=assetName,proto3" json:"assetName,omitempty"`
	AssetCode            string   `protobuf:"bytes,2,opt,name=assetCode,proto3" json:"assetCode,omitempty"`
	Price                uint32   `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Quantity             float64  `protobuf:"fixed64,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Time                 uint64   `protobuf:"varint,5,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRecentTradeReq) Reset()         { *m = AddRecentTradeReq{} }
func (m *AddRecentTradeReq) String() string { return proto.CompactTextString(m) }
func (*AddRecentTradeReq) ProtoMessage()    {}
func (*AddRecentTradeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dda53b2096cbd59, []int{7}
}

func (m *AddRecentTradeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRecentTradeReq.Unmarshal(m, b)
}
func (m *AddRecentTradeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRecentTradeReq.Marshal(b, m, deterministic)
}
func (m *AddRecentTradeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRecentTradeReq.Merge(m, src)
}
func (m *AddRecentTradeReq) XXX_Size() int {
	return xxx_messageInfo_AddRecentTradeReq.Size(m)
}
func (m *AddRecentTradeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRecentTradeReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddRecentTradeReq proto.InternalMessageInfo

func (m *AddRecentTradeReq) GetAssetName() string {
	if m != nil {
		return m.AssetName
	}
	return ""
}

func (m *AddRecentTradeReq) GetAssetCode() string {
	if m != nil {
		return m.AssetCode
	}
	return ""
}

func (m *AddRecentTradeReq) GetPrice() uint32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *AddRecentTradeReq) GetQuantity() float64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *AddRecentTradeReq) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type AddRecentTradeReqMulti struct {
	RecentTrade          []*AddRecentTradeReq `protobuf:"bytes,1,rep,name=recentTrade,proto3" json:"recentTrade,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AddRecentTradeReqMulti) Reset()         { *m = AddRecentTradeReqMulti{} }
func (m *AddRecentTradeReqMulti) String() string { return proto.CompactTextString(m) }
func (*AddRecentTradeReqMulti) ProtoMessage()    {}
func (*AddRecentTradeReqMulti) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dda53b2096cbd59, []int{8}
}

func (m *AddRecentTradeReqMulti) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRecentTradeReqMulti.Unmarshal(m, b)
}
func (m *AddRecentTradeReqMulti) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRecentTradeReqMulti.Marshal(b, m, deterministic)
}
func (m *AddRecentTradeReqMulti) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRecentTradeReqMulti.Merge(m, src)
}
func (m *AddRecentTradeReqMulti) XXX_Size() int {
	return xxx_messageInfo_AddRecentTradeReqMulti.Size(m)
}
func (m *AddRecentTradeReqMulti) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRecentTradeReqMulti.DiscardUnknown(m)
}

var xxx_messageInfo_AddRecentTradeReqMulti proto.InternalMessageInfo

func (m *AddRecentTradeReqMulti) GetRecentTrade() []*AddRecentTradeReq {
	if m != nil {
		return m.RecentTrade
	}
	return nil
}

type GetQuoteReq struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetQuoteReq) Reset()         { *m = GetQuoteReq{} }
func (m *GetQuoteReq) String() string { return proto.CompactTextString(m) }
func (*GetQuoteReq) ProtoMessage()    {}
func (*GetQuoteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dda53b2096cbd59, []int{9}
}

func (m *GetQuoteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQuoteReq.Unmarshal(m, b)
}
func (m *GetQuoteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQuoteReq.Marshal(b, m, deterministic)
}
func (m *GetQuoteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQuoteReq.Merge(m, src)
}
func (m *GetQuoteReq) XXX_Size() int {
	return xxx_messageInfo_GetQuoteReq.Size(m)
}
func (m *GetQuoteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQuoteReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetQuoteReq proto.InternalMessageInfo

func (m *GetQuoteReq) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type QuoteResp struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Price                float64  `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	LastUpdated          string   `protobuf:"bytes,3,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuoteResp) Reset()         { *m = QuoteResp{} }
func (m *QuoteResp) String() string { return proto.CompactTextString(m) }
func (*QuoteResp) ProtoMessage()    {}
func (*QuoteResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dda53b2096cbd59, []int{10}
}

func (m *QuoteResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuoteResp.Unmarshal(m, b)
}
func (m *QuoteResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuoteResp.Marshal(b, m, deterministic)
}
func (m *QuoteResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuoteResp.Merge(m, src)
}
func (m *QuoteResp) XXX_Size() int {
	return xxx_messageInfo_QuoteResp.Size(m)
}
func (m *QuoteResp) XXX_DiscardUnknown() {
	xxx_messageInfo_QuoteResp.DiscardUnknown(m)
}

var xxx_messageInfo_QuoteResp proto.InternalMessageInfo

func (m *QuoteResp) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *QuoteResp) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *QuoteResp) GetLastUpdated() string {
	if m != nil {
		return m.LastUpdated
	}
	return ""
}

type UpdateQuoteReq struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	UpdatedPrice         float64  `protobuf:"fixed64,2,opt,name=updatedPrice,proto3" json:"updatedPrice,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateQuoteReq) Reset()         { *m = UpdateQuoteReq{} }
func (m *UpdateQuoteReq) String() string { return proto.CompactTextString(m) }
func (*UpdateQuoteReq) ProtoMessage()    {}
func (*UpdateQuoteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dda53b2096cbd59, []int{11}
}

func (m *UpdateQuoteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateQuoteReq.Unmarshal(m, b)
}
func (m *UpdateQuoteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateQuoteReq.Marshal(b, m, deterministic)
}
func (m *UpdateQuoteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateQuoteReq.Merge(m, src)
}
func (m *UpdateQuoteReq) XXX_Size() int {
	return xxx_messageInfo_UpdateQuoteReq.Size(m)
}
func (m *UpdateQuoteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateQuoteReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateQuoteReq proto.InternalMessageInfo

func (m *UpdateQuoteReq) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *UpdateQuoteReq) GetUpdatedPrice() float64 {
	if m != nil {
		return m.UpdatedPrice
	}
	return 0
}

func init() {
	proto.RegisterType((*Asset)(nil), "com.server.protos.asset")
	proto.RegisterType((*GetAssetReq)(nil), "com.server.protos.getAssetReq")
	proto.RegisterType((*GetAssetResp)(nil), "com.server.protos.getAssetResp")
	proto.RegisterType((*GetAllAssetResp)(nil), "com.server.protos.getAllAssetResp")
	proto.RegisterType((*RecentTradesReq)(nil), "com.server.protos.recentTradesReq")
	proto.RegisterType((*RecentTradesResp)(nil), "com.server.protos.recentTradesResp")
	proto.RegisterType((*RecentTradesMultiple)(nil), "com.server.protos.recentTradesMultiple")
	proto.RegisterType((*AddRecentTradeReq)(nil), "com.server.protos.addRecentTradeReq")
	proto.RegisterType((*AddRecentTradeReqMulti)(nil), "com.server.protos.addRecentTradeReqMulti")
	proto.RegisterType((*GetQuoteReq)(nil), "com.server.protos.getQuoteReq")
	proto.RegisterType((*QuoteResp)(nil), "com.server.protos.quoteResp")
	proto.RegisterType((*UpdateQuoteReq)(nil), "com.server.protos.updateQuoteReq")
}

func init() { proto.RegisterFile("proto/crypto.proto", fileDescriptor_4dda53b2096cbd59) }

var fileDescriptor_4dda53b2096cbd59 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x6d, 0x6f, 0xd3, 0x30,
	0x10, 0x56, 0x5f, 0xd5, 0x5e, 0xcb, 0xc6, 0xac, 0xa9, 0xaa, 0xc2, 0xb4, 0x75, 0x19, 0x88, 0xf1,
	0x81, 0x54, 0xda, 0x90, 0x10, 0x9f, 0x10, 0x14, 0xa8, 0x26, 0xde, 0x86, 0x81, 0x2f, 0x20, 0x01,
	0x69, 0x63, 0x4a, 0x44, 0x12, 0xbb, 0xb1, 0x83, 0xd4, 0x1f, 0xc2, 0x0f, 0xe5, 0x1f, 0xa0, 0x9c,
	0xd3, 0xd6, 0xe9, 0x92, 0x75, 0x9f, 0x92, 0xbb, 0x7b, 0xce, 0xbe, 0xe7, 0xee, 0x1e, 0x03, 0x11,
	0x31, 0x57, 0x7c, 0x38, 0x8d, 0x17, 0x42, 0x71, 0x07, 0x0d, 0xb2, 0x37, 0xe5, 0xa1, 0x23, 0x59,
	0xfc, 0x87, 0xc5, 0xda, 0x23, 0xad, 0x3b, 0x33, 0xce, 0x67, 0x01, 0x1b, 0xa2, 0x39, 0x49, 0x7e,
	0x0e, 0x59, 0x28, 0xd4, 0x42, 0x47, 0xad, 0xa3, 0xcd, 0xa0, 0xf2, 0x43, 0x26, 0x95, 0x1b, 0x0a,
	0x0d, 0xb0, 0xcf, 0xa1, 0xe1, 0x4a, 0xc9, 0x14, 0xe9, 0x41, 0x53, 0x2e, 0xc2, 0x09, 0x0f, 0xfa,
	0x95, 0x41, 0xe5, 0xb4, 0x4d, 0x33, 0x8b, 0x10, 0xa8, 0x47, 0x6e, 0xc8, 0xfa, 0x55, 0xf4, 0xe2,
	0xbf, 0x7d, 0x0f, 0x3a, 0x33, 0xa6, 0x9e, 0xa5, 0x79, 0x94, 0xcd, 0xcb, 0x52, 0xed, 0x17, 0xd0,
	0x5d, 0xc3, 0xa4, 0x20, 0x8f, 0xa0, 0x15, 0x33, 0x29, 0x78, 0x24, 0x19, 0x22, 0x3b, 0x67, 0x7d,
	0xe7, 0x0a, 0x1f, 0x07, 0xcb, 0xa1, 0x2b, 0xa4, 0x3d, 0x86, 0xdd, 0xf4, 0x94, 0x20, 0x28, 0x3b,
	0xa8, 0x76, 0xc3, 0x83, 0x9e, 0xc2, 0x6e, 0xcc, 0xa6, 0x2c, 0x52, 0x9f, 0x62, 0xd7, 0x63, 0x32,
	0xad, 0x7c, 0x1f, 0x1a, 0xd2, 0x8f, 0xa6, 0xba, 0x9c, 0x3a, 0xd5, 0x86, 0xc1, 0xa7, 0x9a, 0xe3,
	0x33, 0x81, 0xdb, 0xf9, 0x03, 0xa4, 0x20, 0x87, 0x00, 0x3a, 0xfa, 0x2e, 0x6d, 0x92, 0xe6, 0x6f,
	0x78, 0x88, 0x05, 0xad, 0x79, 0xe2, 0x46, 0xca, 0x57, 0x8b, 0x7e, 0x6d, 0x50, 0x39, 0xad, 0xd0,
	0x95, 0x9d, 0xb6, 0x36, 0x1d, 0x47, 0xbf, 0x8e, 0x97, 0xe3, 0xbf, 0xfd, 0x1d, 0xf6, 0xcd, 0x3b,
	0xde, 0x26, 0x81, 0xf2, 0x45, 0xc0, 0xc8, 0x18, 0xba, 0x81, 0x2f, 0x91, 0xbe, 0x41, 0xfb, 0xa4,
	0x80, 0xf6, 0x66, 0x89, 0x34, 0x97, 0x68, 0xff, 0xad, 0xc0, 0x9e, 0xeb, 0x79, 0x74, 0x8d, 0x4a,
	0x1b, 0x71, 0x00, 0x6d, 0x6c, 0x97, 0xc1, 0x62, 0xed, 0x58, 0x45, 0x47, 0xdc, 0x5b, 0x2e, 0xc2,
	0xda, 0x91, 0x36, 0x51, 0xc4, 0xfe, 0x94, 0x21, 0xbf, 0x5b, 0x54, 0x1b, 0x39, 0xe2, 0xf5, 0x12,
	0xe2, 0x0d, 0x83, 0xf8, 0x0f, 0xe8, 0x5d, 0x29, 0x0b, 0xd9, 0x93, 0x57, 0xd0, 0x31, 0x38, 0x65,
	0xcc, 0xef, 0x16, 0x0d, 0x7c, 0x33, 0x9f, 0x9a, 0x89, 0xd9, 0xd6, 0x7e, 0x48, 0xb8, 0x62, 0xd7,
	0x6d, 0xed, 0x57, 0x68, 0xcf, 0x35, 0x46, 0x8a, 0x52, 0x55, 0xac, 0x38, 0x57, 0x91, 0x5a, 0xc6,
	0x79, 0x00, 0x9d, 0xc0, 0x95, 0xea, 0xb3, 0xf0, 0x5c, 0xc5, 0x3c, 0xec, 0x47, 0x9b, 0x9a, 0x2e,
	0xfb, 0x0d, 0xec, 0x24, 0xf8, 0xbb, 0xad, 0x0c, 0x62, 0x43, 0x57, 0x23, 0xbd, 0x4b, 0xe3, 0xa2,
	0x9c, 0xef, 0xec, 0x5f, 0x0d, 0x9a, 0x23, 0x7c, 0x1e, 0xc8, 0x85, 0x96, 0x64, 0xa6, 0x12, 0xd2,
	0x73, 0xb4, 0xf0, 0x9d, 0xa5, 0xf0, 0x9d, 0x97, 0xe9, 0xab, 0x60, 0xd9, 0x05, 0x6d, 0xdb, 0x54,
	0xd7, 0x6b, 0x68, 0x2d, 0x65, 0x4b, 0x0e, 0x4b, 0xf0, 0x99, 0xf4, 0xad, 0xa3, 0x6b, 0xe3, 0x52,
	0x90, 0x6f, 0xa8, 0x5e, 0x63, 0x2c, 0x92, 0xd8, 0x5b, 0x97, 0x76, 0x6e, 0xdd, 0xdf, 0x82, 0x59,
	0xe9, 0xe2, 0x23, 0xec, 0xe4, 0xc7, 0x4e, 0x1e, 0xdc, 0x64, 0x33, 0x30, 0xdf, 0x2a, 0xe9, 0x12,
	0x79, 0x8f, 0x45, 0x8f, 0x92, 0x38, 0x66, 0x91, 0x5e, 0x98, 0xb2, 0x46, 0x2c, 0xc7, 0x68, 0x1d,
	0x14, 0xc4, 0xd7, 0x6b, 0x74, 0xb1, 0x1c, 0x26, 0xe2, 0x25, 0x39, 0x2e, 0x40, 0xe7, 0xf7, 0xa2,
	0xac, 0xb6, 0xe7, 0x27, 0x5f, 0x8e, 0x67, 0xbe, 0xfa, 0x95, 0x4c, 0xd2, 0x23, 0x86, 0x6e, 0xe0,
	0xff, 0x76, 0xa3, 0x27, 0x8f, 0x87, 0x63, 0xfe, 0x70, 0x4c, 0x2f, 0x47, 0xd9, 0x43, 0xdf, 0xc4,
	0xcf, 0xf9, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37, 0x0d, 0xaa, 0xd2, 0x43, 0x06, 0x00, 0x00,
}
