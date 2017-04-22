// Code generated by protoc-gen-go.
// source: ns.proto
// DO NOT EDIT!

/*
Package ns is a generated protocol buffer package.

It is generated from these files:
	ns.proto

It has these top-level messages:
	CreateNodeSessionRequest
	CreateNodeSessionResponse
	GetNodeSessionRequest
	GetNodeSessionResponse
	UpdateNodeSessionRequest
	UpdateNodeSessionResponse
	DeleteNodeSessionRequest
	DeleteNodeSessionResponse
	GetRandomDevAddrRequest
	GetRandomDevAddrResponse
	EnqueueDataDownMACCommandRequest
	EnqueueDataDownMACCommandResponse
	PushDataDownRequest
	PushDataDownResponse
*/
package ns

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RXWindow int32

const (
	// Receive window 1
	RXWindow_RX1 RXWindow = 0
	// Receive window 2
	RXWindow_RX2 RXWindow = 1
)

var RXWindow_name = map[int32]string{
	0: "RX1",
	1: "RX2",
}
var RXWindow_value = map[string]int32{
	"RX1": 0,
	"RX2": 1,
}

func (x RXWindow) String() string {
	return proto.EnumName(RXWindow_name, int32(x))
}
func (RXWindow) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateNodeSessionRequest struct {
	// The address of the device (4 bytes).
	DevAddr []byte `protobuf:"bytes,1,opt,name=devAddr,proto3" json:"devAddr,omitempty"`
	// The application EUI (8 bytes).
	AppEUI []byte `protobuf:"bytes,2,opt,name=appEUI,proto3" json:"appEUI,omitempty"`
	// The device EUI (8 bytes).
	DevEUI []byte `protobuf:"bytes,3,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	// The network-session key (16 bytes).
	NwkSKey []byte `protobuf:"bytes,4,opt,name=nwkSKey,proto3" json:"nwkSKey,omitempty"`
	// The next expected uplink frame-counter.
	FCntUp uint32 `protobuf:"varint,5,opt,name=fCntUp" json:"fCntUp,omitempty"`
	// The frame-counter used for the next downlink frame.
	FCntDown uint32 `protobuf:"varint,6,opt,name=fCntDown" json:"fCntDown,omitempty"`
	// the RX delay value (0 = 1 sec, 1 = 1 sec, 2 = 2 sec ...).
	RxDelay uint32 `protobuf:"varint,7,opt,name=rxDelay" json:"rxDelay,omitempty"`
	// The data-rate offset used for RX1 (see LoRaWAN specs for valid values).
	Rx1DROffset uint32 `protobuf:"varint,8,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	// The custom channel frequency list (this is not supported for every
	// ISM band, see the LoRaWAN specs for more information).
	CFList []uint32 `protobuf:"varint,9,rep,packed,name=cFList" json:"cFList,omitempty"`
	// The RX window to use for downlink transmissions.
	RxWindow RXWindow `protobuf:"varint,10,opt,name=rxWindow,enum=ns.RXWindow" json:"rxWindow,omitempty"`
	// The data-rate to use for RX2 transmissions.
	Rx2DR uint32 `protobuf:"varint,11,opt,name=rx2DR" json:"rx2DR,omitempty"`
	// Use relax frame-counter mode for ABP devices (this is insecure!).
	RelaxFCnt bool `protobuf:"varint,12,opt,name=relaxFCnt" json:"relaxFCnt,omitempty"`
	// The interval (based on frame-counter) on which to calculate the ideal
	// data-rate and tx-power of the node and if needed, request an adaption.
	AdrInterval uint32 `protobuf:"varint,13,opt,name=adrInterval" json:"adrInterval,omitempty"`
	// The installation margin to take into account when calculating the ideal
	// data-rate and tx-power. The default recommended value is 5dB.
	InstallationMargin float64 `protobuf:"fixed64,14,opt,name=installationMargin" json:"installationMargin,omitempty"`
}

func (m *CreateNodeSessionRequest) Reset()                    { *m = CreateNodeSessionRequest{} }
func (m *CreateNodeSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateNodeSessionRequest) ProtoMessage()               {}
func (*CreateNodeSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateNodeSessionRequest) GetDevAddr() []byte {
	if m != nil {
		return m.DevAddr
	}
	return nil
}

func (m *CreateNodeSessionRequest) GetAppEUI() []byte {
	if m != nil {
		return m.AppEUI
	}
	return nil
}

func (m *CreateNodeSessionRequest) GetDevEUI() []byte {
	if m != nil {
		return m.DevEUI
	}
	return nil
}

func (m *CreateNodeSessionRequest) GetNwkSKey() []byte {
	if m != nil {
		return m.NwkSKey
	}
	return nil
}

func (m *CreateNodeSessionRequest) GetFCntUp() uint32 {
	if m != nil {
		return m.FCntUp
	}
	return 0
}

func (m *CreateNodeSessionRequest) GetFCntDown() uint32 {
	if m != nil {
		return m.FCntDown
	}
	return 0
}

func (m *CreateNodeSessionRequest) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *CreateNodeSessionRequest) GetRx1DROffset() uint32 {
	if m != nil {
		return m.Rx1DROffset
	}
	return 0
}

func (m *CreateNodeSessionRequest) GetCFList() []uint32 {
	if m != nil {
		return m.CFList
	}
	return nil
}

func (m *CreateNodeSessionRequest) GetRxWindow() RXWindow {
	if m != nil {
		return m.RxWindow
	}
	return RXWindow_RX1
}

func (m *CreateNodeSessionRequest) GetRx2DR() uint32 {
	if m != nil {
		return m.Rx2DR
	}
	return 0
}

func (m *CreateNodeSessionRequest) GetRelaxFCnt() bool {
	if m != nil {
		return m.RelaxFCnt
	}
	return false
}

func (m *CreateNodeSessionRequest) GetAdrInterval() uint32 {
	if m != nil {
		return m.AdrInterval
	}
	return 0
}

func (m *CreateNodeSessionRequest) GetInstallationMargin() float64 {
	if m != nil {
		return m.InstallationMargin
	}
	return 0
}

type CreateNodeSessionResponse struct {
}

func (m *CreateNodeSessionResponse) Reset()                    { *m = CreateNodeSessionResponse{} }
func (m *CreateNodeSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateNodeSessionResponse) ProtoMessage()               {}
func (*CreateNodeSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetNodeSessionRequest struct {
	DevEUI []byte `protobuf:"bytes,1,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
}

func (m *GetNodeSessionRequest) Reset()                    { *m = GetNodeSessionRequest{} }
func (m *GetNodeSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNodeSessionRequest) ProtoMessage()               {}
func (*GetNodeSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetNodeSessionRequest) GetDevEUI() []byte {
	if m != nil {
		return m.DevEUI
	}
	return nil
}

type GetNodeSessionResponse struct {
	// The address of the device (4 bytes).
	DevAddr []byte `protobuf:"bytes,1,opt,name=devAddr,proto3" json:"devAddr,omitempty"`
	// The application EUI (8 bytes).
	AppEUI []byte `protobuf:"bytes,2,opt,name=appEUI,proto3" json:"appEUI,omitempty"`
	// The device EUI (8 bytes).
	DevEUI []byte `protobuf:"bytes,3,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	// The network-session key (16 bytes).
	NwkSKey []byte `protobuf:"bytes,4,opt,name=nwkSKey,proto3" json:"nwkSKey,omitempty"`
	// The next expected uplink frame-counter.
	FCntUp uint32 `protobuf:"varint,5,opt,name=fCntUp" json:"fCntUp,omitempty"`
	// The frame-counter used for the next downlink frame.
	FCntDown uint32 `protobuf:"varint,6,opt,name=fCntDown" json:"fCntDown,omitempty"`
	// the RX delay value (0 = 1 sec, 1 = 1 sec, 2 = 2 sec ...).
	RxDelay uint32 `protobuf:"varint,7,opt,name=rxDelay" json:"rxDelay,omitempty"`
	// The data-rate offset used for RX1 (see LoRaWAN specs for valid values).
	Rx1DROffset uint32 `protobuf:"varint,8,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	// The custom channel frequency list (this is not supported for every
	// ISM band, see the LoRaWAN specs for more information).
	CFList []uint32 `protobuf:"varint,9,rep,packed,name=cFList" json:"cFList,omitempty"`
	// The RX window to use for downlink transmissions.
	RxWindow RXWindow `protobuf:"varint,10,opt,name=rxWindow,enum=ns.RXWindow" json:"rxWindow,omitempty"`
	// The data-rate to use for RX2 transmissions.
	Rx2DR uint32 `protobuf:"varint,11,opt,name=rx2DR" json:"rx2DR,omitempty"`
	// Use relax frame-counter mode for ABP devices (this is insecure!).
	RelaxFCnt bool `protobuf:"varint,12,opt,name=relaxFCnt" json:"relaxFCnt,omitempty"`
	// The interval (based on frame-counter) on which to calculate the ideal
	// data-rate and tx-power of the node and if needed, request an adaption.
	AdrInterval uint32 `protobuf:"varint,13,opt,name=adrInterval" json:"adrInterval,omitempty"`
	// The installation margin to take into account when calculating the ideal
	// data-rate and tx-power. The default recommended value is 5dB.
	InstallationMargin float64 `protobuf:"fixed64,14,opt,name=installationMargin" json:"installationMargin,omitempty"`
	// The number of times the node should re-transmit an uplink frame.
	// This is controlled by the ADR engine.
	NbTrans uint32 `protobuf:"varint,15,opt,name=nbTrans" json:"nbTrans,omitempty"`
	// The TX power of the node. This is controlled by the ADR engine.
	TxPower uint32 `protobuf:"varint,16,opt,name=txPower" json:"txPower,omitempty"`
}

func (m *GetNodeSessionResponse) Reset()                    { *m = GetNodeSessionResponse{} }
func (m *GetNodeSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNodeSessionResponse) ProtoMessage()               {}
func (*GetNodeSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetNodeSessionResponse) GetDevAddr() []byte {
	if m != nil {
		return m.DevAddr
	}
	return nil
}

func (m *GetNodeSessionResponse) GetAppEUI() []byte {
	if m != nil {
		return m.AppEUI
	}
	return nil
}

func (m *GetNodeSessionResponse) GetDevEUI() []byte {
	if m != nil {
		return m.DevEUI
	}
	return nil
}

func (m *GetNodeSessionResponse) GetNwkSKey() []byte {
	if m != nil {
		return m.NwkSKey
	}
	return nil
}

func (m *GetNodeSessionResponse) GetFCntUp() uint32 {
	if m != nil {
		return m.FCntUp
	}
	return 0
}

func (m *GetNodeSessionResponse) GetFCntDown() uint32 {
	if m != nil {
		return m.FCntDown
	}
	return 0
}

func (m *GetNodeSessionResponse) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *GetNodeSessionResponse) GetRx1DROffset() uint32 {
	if m != nil {
		return m.Rx1DROffset
	}
	return 0
}

func (m *GetNodeSessionResponse) GetCFList() []uint32 {
	if m != nil {
		return m.CFList
	}
	return nil
}

func (m *GetNodeSessionResponse) GetRxWindow() RXWindow {
	if m != nil {
		return m.RxWindow
	}
	return RXWindow_RX1
}

func (m *GetNodeSessionResponse) GetRx2DR() uint32 {
	if m != nil {
		return m.Rx2DR
	}
	return 0
}

func (m *GetNodeSessionResponse) GetRelaxFCnt() bool {
	if m != nil {
		return m.RelaxFCnt
	}
	return false
}

func (m *GetNodeSessionResponse) GetAdrInterval() uint32 {
	if m != nil {
		return m.AdrInterval
	}
	return 0
}

func (m *GetNodeSessionResponse) GetInstallationMargin() float64 {
	if m != nil {
		return m.InstallationMargin
	}
	return 0
}

func (m *GetNodeSessionResponse) GetNbTrans() uint32 {
	if m != nil {
		return m.NbTrans
	}
	return 0
}

func (m *GetNodeSessionResponse) GetTxPower() uint32 {
	if m != nil {
		return m.TxPower
	}
	return 0
}

type UpdateNodeSessionRequest struct {
	// The address of the device (4 bytes).
	DevAddr []byte `protobuf:"bytes,1,opt,name=devAddr,proto3" json:"devAddr,omitempty"`
	// The application EUI (8 bytes).
	AppEUI []byte `protobuf:"bytes,2,opt,name=appEUI,proto3" json:"appEUI,omitempty"`
	// The device EUI (8 bytes).
	DevEUI []byte `protobuf:"bytes,3,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	// The network-session key (16 bytes).
	NwkSKey []byte `protobuf:"bytes,4,opt,name=nwkSKey,proto3" json:"nwkSKey,omitempty"`
	// The next expected uplink frame-counter.
	FCntUp uint32 `protobuf:"varint,5,opt,name=fCntUp" json:"fCntUp,omitempty"`
	// The frame-counter used for the next downlink frame.
	FCntDown uint32 `protobuf:"varint,6,opt,name=fCntDown" json:"fCntDown,omitempty"`
	// the RX delay value (0 = 1 sec, 1 = 1 sec, 2 = 2 sec ...).
	RxDelay uint32 `protobuf:"varint,7,opt,name=rxDelay" json:"rxDelay,omitempty"`
	// The data-rate offset used for RX1 (see LoRaWAN specs for valid values).
	Rx1DROffset uint32 `protobuf:"varint,8,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	// The custom channel frequency list (this is not supported for every
	// ISM band, see the LoRaWAN specs for more information).
	CFList []uint32 `protobuf:"varint,9,rep,packed,name=cFList" json:"cFList,omitempty"`
	// The RX window to use for downlink transmissions.
	RxWindow RXWindow `protobuf:"varint,10,opt,name=rxWindow,enum=ns.RXWindow" json:"rxWindow,omitempty"`
	// The data-rate to use for RX2 transmissions.
	Rx2DR uint32 `protobuf:"varint,11,opt,name=rx2DR" json:"rx2DR,omitempty"`
	// Use relax frame-counter mode for ABP devices (this is insecure!).
	RelaxFCnt bool `protobuf:"varint,12,opt,name=relaxFCnt" json:"relaxFCnt,omitempty"`
	// The interval (based on frame-counter) on which to calculate the ideal
	// data-rate and tx-power of the node and if needed, request an adaption.
	AdrInterval uint32 `protobuf:"varint,13,opt,name=adrInterval" json:"adrInterval,omitempty"`
	// The installation margin to take into account when calculating the ideal
	// data-rate and tx-power. The default recommended value is 5dB.
	InstallationMargin float64 `protobuf:"fixed64,14,opt,name=installationMargin" json:"installationMargin,omitempty"`
}

func (m *UpdateNodeSessionRequest) Reset()                    { *m = UpdateNodeSessionRequest{} }
func (m *UpdateNodeSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateNodeSessionRequest) ProtoMessage()               {}
func (*UpdateNodeSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateNodeSessionRequest) GetDevAddr() []byte {
	if m != nil {
		return m.DevAddr
	}
	return nil
}

func (m *UpdateNodeSessionRequest) GetAppEUI() []byte {
	if m != nil {
		return m.AppEUI
	}
	return nil
}

func (m *UpdateNodeSessionRequest) GetDevEUI() []byte {
	if m != nil {
		return m.DevEUI
	}
	return nil
}

func (m *UpdateNodeSessionRequest) GetNwkSKey() []byte {
	if m != nil {
		return m.NwkSKey
	}
	return nil
}

func (m *UpdateNodeSessionRequest) GetFCntUp() uint32 {
	if m != nil {
		return m.FCntUp
	}
	return 0
}

func (m *UpdateNodeSessionRequest) GetFCntDown() uint32 {
	if m != nil {
		return m.FCntDown
	}
	return 0
}

func (m *UpdateNodeSessionRequest) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *UpdateNodeSessionRequest) GetRx1DROffset() uint32 {
	if m != nil {
		return m.Rx1DROffset
	}
	return 0
}

func (m *UpdateNodeSessionRequest) GetCFList() []uint32 {
	if m != nil {
		return m.CFList
	}
	return nil
}

func (m *UpdateNodeSessionRequest) GetRxWindow() RXWindow {
	if m != nil {
		return m.RxWindow
	}
	return RXWindow_RX1
}

func (m *UpdateNodeSessionRequest) GetRx2DR() uint32 {
	if m != nil {
		return m.Rx2DR
	}
	return 0
}

func (m *UpdateNodeSessionRequest) GetRelaxFCnt() bool {
	if m != nil {
		return m.RelaxFCnt
	}
	return false
}

func (m *UpdateNodeSessionRequest) GetAdrInterval() uint32 {
	if m != nil {
		return m.AdrInterval
	}
	return 0
}

func (m *UpdateNodeSessionRequest) GetInstallationMargin() float64 {
	if m != nil {
		return m.InstallationMargin
	}
	return 0
}

type UpdateNodeSessionResponse struct {
}

func (m *UpdateNodeSessionResponse) Reset()                    { *m = UpdateNodeSessionResponse{} }
func (m *UpdateNodeSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateNodeSessionResponse) ProtoMessage()               {}
func (*UpdateNodeSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type DeleteNodeSessionRequest struct {
	DevEUI []byte `protobuf:"bytes,1,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
}

func (m *DeleteNodeSessionRequest) Reset()                    { *m = DeleteNodeSessionRequest{} }
func (m *DeleteNodeSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeSessionRequest) ProtoMessage()               {}
func (*DeleteNodeSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DeleteNodeSessionRequest) GetDevEUI() []byte {
	if m != nil {
		return m.DevEUI
	}
	return nil
}

type DeleteNodeSessionResponse struct {
}

func (m *DeleteNodeSessionResponse) Reset()                    { *m = DeleteNodeSessionResponse{} }
func (m *DeleteNodeSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeSessionResponse) ProtoMessage()               {}
func (*DeleteNodeSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type GetRandomDevAddrRequest struct {
}

func (m *GetRandomDevAddrRequest) Reset()                    { *m = GetRandomDevAddrRequest{} }
func (m *GetRandomDevAddrRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRandomDevAddrRequest) ProtoMessage()               {}
func (*GetRandomDevAddrRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type GetRandomDevAddrResponse struct {
	DevAddr []byte `protobuf:"bytes,1,opt,name=devAddr,proto3" json:"devAddr,omitempty"`
}

func (m *GetRandomDevAddrResponse) Reset()                    { *m = GetRandomDevAddrResponse{} }
func (m *GetRandomDevAddrResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRandomDevAddrResponse) ProtoMessage()               {}
func (*GetRandomDevAddrResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetRandomDevAddrResponse) GetDevAddr() []byte {
	if m != nil {
		return m.DevAddr
	}
	return nil
}

type EnqueueDataDownMACCommandRequest struct {
	DevEUI     []byte `protobuf:"bytes,1,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	FrmPayload bool   `protobuf:"varint,2,opt,name=frmPayload" json:"frmPayload,omitempty"`
	Data       []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *EnqueueDataDownMACCommandRequest) Reset()         { *m = EnqueueDataDownMACCommandRequest{} }
func (m *EnqueueDataDownMACCommandRequest) String() string { return proto.CompactTextString(m) }
func (*EnqueueDataDownMACCommandRequest) ProtoMessage()    {}
func (*EnqueueDataDownMACCommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10}
}

func (m *EnqueueDataDownMACCommandRequest) GetDevEUI() []byte {
	if m != nil {
		return m.DevEUI
	}
	return nil
}

func (m *EnqueueDataDownMACCommandRequest) GetFrmPayload() bool {
	if m != nil {
		return m.FrmPayload
	}
	return false
}

func (m *EnqueueDataDownMACCommandRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type EnqueueDataDownMACCommandResponse struct {
}

func (m *EnqueueDataDownMACCommandResponse) Reset()         { *m = EnqueueDataDownMACCommandResponse{} }
func (m *EnqueueDataDownMACCommandResponse) String() string { return proto.CompactTextString(m) }
func (*EnqueueDataDownMACCommandResponse) ProtoMessage()    {}
func (*EnqueueDataDownMACCommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{11}
}

type PushDataDownRequest struct {
	// DevEUI of the node to which to push the data.
	DevEUI []byte `protobuf:"bytes,1,opt,name=devEUI,proto3" json:"devEUI,omitempty"`
	// Data (encrypted with the AppSKey) to push to the node.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// Payload must be acknowledged by the node.
	Confirmed bool `protobuf:"varint,3,opt,name=confirmed" json:"confirmed,omitempty"`
	// FPort to use for transmitting the payload.
	FPort uint32 `protobuf:"varint,4,opt,name=fPort" json:"fPort,omitempty"`
	// FCnt used for encrypting the data. When this does not match the FCntDown
	// of the network-server, an error is returned.
	FCnt uint32 `protobuf:"varint,5,opt,name=fCnt" json:"fCnt,omitempty"`
}

func (m *PushDataDownRequest) Reset()                    { *m = PushDataDownRequest{} }
func (m *PushDataDownRequest) String() string            { return proto.CompactTextString(m) }
func (*PushDataDownRequest) ProtoMessage()               {}
func (*PushDataDownRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *PushDataDownRequest) GetDevEUI() []byte {
	if m != nil {
		return m.DevEUI
	}
	return nil
}

func (m *PushDataDownRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PushDataDownRequest) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func (m *PushDataDownRequest) GetFPort() uint32 {
	if m != nil {
		return m.FPort
	}
	return 0
}

func (m *PushDataDownRequest) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

type PushDataDownResponse struct {
}

func (m *PushDataDownResponse) Reset()                    { *m = PushDataDownResponse{} }
func (m *PushDataDownResponse) String() string            { return proto.CompactTextString(m) }
func (*PushDataDownResponse) ProtoMessage()               {}
func (*PushDataDownResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*CreateNodeSessionRequest)(nil), "ns.CreateNodeSessionRequest")
	proto.RegisterType((*CreateNodeSessionResponse)(nil), "ns.CreateNodeSessionResponse")
	proto.RegisterType((*GetNodeSessionRequest)(nil), "ns.GetNodeSessionRequest")
	proto.RegisterType((*GetNodeSessionResponse)(nil), "ns.GetNodeSessionResponse")
	proto.RegisterType((*UpdateNodeSessionRequest)(nil), "ns.UpdateNodeSessionRequest")
	proto.RegisterType((*UpdateNodeSessionResponse)(nil), "ns.UpdateNodeSessionResponse")
	proto.RegisterType((*DeleteNodeSessionRequest)(nil), "ns.DeleteNodeSessionRequest")
	proto.RegisterType((*DeleteNodeSessionResponse)(nil), "ns.DeleteNodeSessionResponse")
	proto.RegisterType((*GetRandomDevAddrRequest)(nil), "ns.GetRandomDevAddrRequest")
	proto.RegisterType((*GetRandomDevAddrResponse)(nil), "ns.GetRandomDevAddrResponse")
	proto.RegisterType((*EnqueueDataDownMACCommandRequest)(nil), "ns.EnqueueDataDownMACCommandRequest")
	proto.RegisterType((*EnqueueDataDownMACCommandResponse)(nil), "ns.EnqueueDataDownMACCommandResponse")
	proto.RegisterType((*PushDataDownRequest)(nil), "ns.PushDataDownRequest")
	proto.RegisterType((*PushDataDownResponse)(nil), "ns.PushDataDownResponse")
	proto.RegisterEnum("ns.RXWindow", RXWindow_name, RXWindow_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NetworkServer service

type NetworkServerClient interface {
	// CreateNodeSession creates the given node-session.
	// The DevAddr must contain the same NwkID as the configured NetID.
	// Node-sessions will expire automatically after the configured TTL.
	CreateNodeSession(ctx context.Context, in *CreateNodeSessionRequest, opts ...grpc.CallOption) (*CreateNodeSessionResponse, error)
	// GetNodeSession returns the node-session matching the given DevEUI.
	GetNodeSession(ctx context.Context, in *GetNodeSessionRequest, opts ...grpc.CallOption) (*GetNodeSessionResponse, error)
	// UpdateNodeSession updates the given node-session.
	UpdateNodeSession(ctx context.Context, in *UpdateNodeSessionRequest, opts ...grpc.CallOption) (*UpdateNodeSessionResponse, error)
	// DeleteNodeSession deletes the node-session matching the given DevAddr.
	DeleteNodeSession(ctx context.Context, in *DeleteNodeSessionRequest, opts ...grpc.CallOption) (*DeleteNodeSessionResponse, error)
	// GetRandomDevAddr returns a random DevAddr taking the NwkID prefix into account.
	GetRandomDevAddr(ctx context.Context, in *GetRandomDevAddrRequest, opts ...grpc.CallOption) (*GetRandomDevAddrResponse, error)
	// EnqueueDataDownMACCommand adds the downlink mac-command to the queue.
	EnqueueDataDownMACCommand(ctx context.Context, in *EnqueueDataDownMACCommandRequest, opts ...grpc.CallOption) (*EnqueueDataDownMACCommandResponse, error)
	// PushDataDown pushes the given downlink payload to the node (only works for Class-C nodes).
	PushDataDown(ctx context.Context, in *PushDataDownRequest, opts ...grpc.CallOption) (*PushDataDownResponse, error)
}

type networkServerClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServerClient(cc *grpc.ClientConn) NetworkServerClient {
	return &networkServerClient{cc}
}

func (c *networkServerClient) CreateNodeSession(ctx context.Context, in *CreateNodeSessionRequest, opts ...grpc.CallOption) (*CreateNodeSessionResponse, error) {
	out := new(CreateNodeSessionResponse)
	err := grpc.Invoke(ctx, "/ns.NetworkServer/CreateNodeSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) GetNodeSession(ctx context.Context, in *GetNodeSessionRequest, opts ...grpc.CallOption) (*GetNodeSessionResponse, error) {
	out := new(GetNodeSessionResponse)
	err := grpc.Invoke(ctx, "/ns.NetworkServer/GetNodeSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) UpdateNodeSession(ctx context.Context, in *UpdateNodeSessionRequest, opts ...grpc.CallOption) (*UpdateNodeSessionResponse, error) {
	out := new(UpdateNodeSessionResponse)
	err := grpc.Invoke(ctx, "/ns.NetworkServer/UpdateNodeSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) DeleteNodeSession(ctx context.Context, in *DeleteNodeSessionRequest, opts ...grpc.CallOption) (*DeleteNodeSessionResponse, error) {
	out := new(DeleteNodeSessionResponse)
	err := grpc.Invoke(ctx, "/ns.NetworkServer/DeleteNodeSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) GetRandomDevAddr(ctx context.Context, in *GetRandomDevAddrRequest, opts ...grpc.CallOption) (*GetRandomDevAddrResponse, error) {
	out := new(GetRandomDevAddrResponse)
	err := grpc.Invoke(ctx, "/ns.NetworkServer/GetRandomDevAddr", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) EnqueueDataDownMACCommand(ctx context.Context, in *EnqueueDataDownMACCommandRequest, opts ...grpc.CallOption) (*EnqueueDataDownMACCommandResponse, error) {
	out := new(EnqueueDataDownMACCommandResponse)
	err := grpc.Invoke(ctx, "/ns.NetworkServer/EnqueueDataDownMACCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServerClient) PushDataDown(ctx context.Context, in *PushDataDownRequest, opts ...grpc.CallOption) (*PushDataDownResponse, error) {
	out := new(PushDataDownResponse)
	err := grpc.Invoke(ctx, "/ns.NetworkServer/PushDataDown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetworkServer service

type NetworkServerServer interface {
	// CreateNodeSession creates the given node-session.
	// The DevAddr must contain the same NwkID as the configured NetID.
	// Node-sessions will expire automatically after the configured TTL.
	CreateNodeSession(context.Context, *CreateNodeSessionRequest) (*CreateNodeSessionResponse, error)
	// GetNodeSession returns the node-session matching the given DevEUI.
	GetNodeSession(context.Context, *GetNodeSessionRequest) (*GetNodeSessionResponse, error)
	// UpdateNodeSession updates the given node-session.
	UpdateNodeSession(context.Context, *UpdateNodeSessionRequest) (*UpdateNodeSessionResponse, error)
	// DeleteNodeSession deletes the node-session matching the given DevAddr.
	DeleteNodeSession(context.Context, *DeleteNodeSessionRequest) (*DeleteNodeSessionResponse, error)
	// GetRandomDevAddr returns a random DevAddr taking the NwkID prefix into account.
	GetRandomDevAddr(context.Context, *GetRandomDevAddrRequest) (*GetRandomDevAddrResponse, error)
	// EnqueueDataDownMACCommand adds the downlink mac-command to the queue.
	EnqueueDataDownMACCommand(context.Context, *EnqueueDataDownMACCommandRequest) (*EnqueueDataDownMACCommandResponse, error)
	// PushDataDown pushes the given downlink payload to the node (only works for Class-C nodes).
	PushDataDown(context.Context, *PushDataDownRequest) (*PushDataDownResponse, error)
}

func RegisterNetworkServerServer(s *grpc.Server, srv NetworkServerServer) {
	s.RegisterService(&_NetworkServer_serviceDesc, srv)
}

func _NetworkServer_CreateNodeSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).CreateNodeSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ns.NetworkServer/CreateNodeSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).CreateNodeSession(ctx, req.(*CreateNodeSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_GetNodeSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).GetNodeSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ns.NetworkServer/GetNodeSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).GetNodeSession(ctx, req.(*GetNodeSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_UpdateNodeSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).UpdateNodeSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ns.NetworkServer/UpdateNodeSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).UpdateNodeSession(ctx, req.(*UpdateNodeSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_DeleteNodeSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).DeleteNodeSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ns.NetworkServer/DeleteNodeSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).DeleteNodeSession(ctx, req.(*DeleteNodeSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_GetRandomDevAddr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRandomDevAddrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).GetRandomDevAddr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ns.NetworkServer/GetRandomDevAddr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).GetRandomDevAddr(ctx, req.(*GetRandomDevAddrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_EnqueueDataDownMACCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueDataDownMACCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).EnqueueDataDownMACCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ns.NetworkServer/EnqueueDataDownMACCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).EnqueueDataDownMACCommand(ctx, req.(*EnqueueDataDownMACCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServer_PushDataDown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushDataDownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServerServer).PushDataDown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ns.NetworkServer/PushDataDown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServerServer).PushDataDown(ctx, req.(*PushDataDownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ns.NetworkServer",
	HandlerType: (*NetworkServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNodeSession",
			Handler:    _NetworkServer_CreateNodeSession_Handler,
		},
		{
			MethodName: "GetNodeSession",
			Handler:    _NetworkServer_GetNodeSession_Handler,
		},
		{
			MethodName: "UpdateNodeSession",
			Handler:    _NetworkServer_UpdateNodeSession_Handler,
		},
		{
			MethodName: "DeleteNodeSession",
			Handler:    _NetworkServer_DeleteNodeSession_Handler,
		},
		{
			MethodName: "GetRandomDevAddr",
			Handler:    _NetworkServer_GetRandomDevAddr_Handler,
		},
		{
			MethodName: "EnqueueDataDownMACCommand",
			Handler:    _NetworkServer_EnqueueDataDownMACCommand_Handler,
		},
		{
			MethodName: "PushDataDown",
			Handler:    _NetworkServer_PushDataDown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ns.proto",
}

func init() { proto.RegisterFile("ns.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 706 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xec, 0x56, 0xdd, 0x52, 0xd3, 0x40,
	0x14, 0x26, 0x94, 0x9f, 0x70, 0x68, 0xb1, 0xac, 0x08, 0xdb, 0x52, 0x9d, 0x18, 0x75, 0xa6, 0xe3,
	0x45, 0x1d, 0xaa, 0x2f, 0xc0, 0xb4, 0xc0, 0x30, 0x0a, 0x74, 0x16, 0x19, 0xb9, 0x5d, 0xc8, 0x56,
	0x23, 0xe9, 0x6e, 0xd9, 0x6c, 0x7f, 0x78, 0x04, 0x5f, 0xc8, 0x67, 0xf1, 0x11, 0x7c, 0x0c, 0x67,
	0x77, 0xd3, 0x1f, 0x68, 0x62, 0x6f, 0x75, 0x86, 0xbb, 0xfd, 0xce, 0xd7, 0xf3, 0x9d, 0x93, 0x9c,
	0xef, 0x64, 0x0b, 0x2e, 0x8f, 0x6b, 0x5d, 0x29, 0x94, 0x40, 0x8b, 0x3c, 0xf6, 0x7f, 0xe6, 0x00,
	0x37, 0x24, 0xa3, 0x8a, 0x9d, 0x8a, 0x80, 0x9d, 0xb3, 0x38, 0x0e, 0x05, 0x27, 0xec, 0xb6, 0xc7,
	0x62, 0x85, 0x30, 0xac, 0x06, 0xac, 0xbf, 0x1f, 0x04, 0x12, 0x3b, 0x9e, 0x53, 0xcd, 0x93, 0x11,
	0x44, 0xdb, 0xb0, 0x42, 0xbb, 0xdd, 0x83, 0x8b, 0x63, 0xbc, 0x68, 0x88, 0x04, 0xe9, 0x78, 0xc0,
	0xfa, 0x3a, 0x9e, 0xb3, 0x71, 0x8b, 0xb4, 0x12, 0x1f, 0xdc, 0x9c, 0x7f, 0x64, 0x77, 0x78, 0xc9,
	0x2a, 0x25, 0x50, 0x67, 0xb4, 0x1b, 0x5c, 0x5d, 0x74, 0xf1, 0xb2, 0xe7, 0x54, 0x0b, 0x24, 0x41,
	0xa8, 0x0c, 0xae, 0x3e, 0x35, 0xc5, 0x80, 0xe3, 0x15, 0xc3, 0x8c, 0xb1, 0x56, 0x93, 0xc3, 0x26,
	0x8b, 0xe8, 0x1d, 0x5e, 0x35, 0xd4, 0x08, 0x22, 0x0f, 0xd6, 0xe5, 0x70, 0xaf, 0x49, 0xce, 0xda,
	0xed, 0x98, 0x29, 0xec, 0x1a, 0x76, 0x3a, 0xa4, 0xeb, 0x5d, 0x1f, 0x7e, 0x0a, 0x63, 0x85, 0xd7,
	0xbc, 0x9c, 0xae, 0x67, 0x11, 0xaa, 0x82, 0x2b, 0x87, 0x5f, 0x42, 0x1e, 0x88, 0x01, 0x06, 0xcf,
	0xa9, 0x6e, 0xd4, 0xf3, 0x35, 0x1e, 0xd7, 0xc8, 0xa5, 0x8d, 0x91, 0x31, 0x8b, 0xb6, 0x60, 0x59,
	0x0e, 0xeb, 0x4d, 0x82, 0xd7, 0x8d, 0xba, 0x05, 0xa8, 0x02, 0x6b, 0x92, 0x45, 0x74, 0x78, 0xd8,
	0xe0, 0x0a, 0xe7, 0x3d, 0xa7, 0xea, 0x92, 0x49, 0x40, 0xf7, 0x45, 0x03, 0x79, 0xcc, 0x15, 0x93,
	0x7d, 0x1a, 0xe1, 0x82, 0xed, 0x6b, 0x2a, 0x84, 0x6a, 0x80, 0x42, 0x1e, 0x2b, 0x1a, 0x45, 0x54,
	0x85, 0x82, 0x9f, 0x50, 0xf9, 0x35, 0xe4, 0x78, 0xc3, 0x73, 0xaa, 0x0e, 0x49, 0x61, 0xfc, 0x5d,
	0x28, 0xa5, 0xcc, 0x2d, 0xee, 0x0a, 0x1e, 0x33, 0xff, 0x1d, 0x3c, 0x3b, 0x62, 0x2a, 0x65, 0xa2,
	0x93, 0xf9, 0x38, 0xd3, 0xf3, 0xf1, 0x7f, 0xe7, 0x60, 0xfb, 0x61, 0x86, 0xd5, 0x7a, 0x34, 0xc1,
	0xbf, 0x6b, 0x02, 0xf3, 0x46, 0xaf, 0x3e, 0x4b, 0xca, 0x63, 0xfc, 0xc4, 0xbe, 0x83, 0x04, 0x6a,
	0x46, 0x0d, 0x5b, 0x62, 0xc0, 0x24, 0x2e, 0x5a, 0x26, 0x81, 0x66, 0xe3, 0x2f, 0xba, 0xc1, 0xe3,
	0xc6, 0xff, 0x87, 0x1b, 0x9f, 0x32, 0xb7, 0x64, 0xe3, 0xeb, 0x80, 0x9b, 0x2c, 0x62, 0xa9, 0x43,
	0xcd, 0x5a, 0xfa, 0x5d, 0x28, 0xa5, 0xe4, 0x24, 0x82, 0x25, 0xd8, 0x39, 0x62, 0x8a, 0x50, 0x1e,
	0x88, 0x4e, 0xd3, 0x7a, 0x20, 0xd1, 0xf3, 0x3f, 0x00, 0x9e, 0xa5, 0xe6, 0x7d, 0x2d, 0x7c, 0x0e,
	0xde, 0x01, 0xbf, 0xed, 0xb1, 0x1e, 0x6b, 0x52, 0x45, 0xf5, 0x54, 0x4f, 0xf6, 0x1b, 0x0d, 0xd1,
	0xe9, 0x50, 0x1e, 0xcc, 0xe9, 0x14, 0xbd, 0x00, 0x68, 0xcb, 0x4e, 0x8b, 0xde, 0x45, 0x82, 0x06,
	0xc6, 0x80, 0x2e, 0x99, 0x8a, 0x20, 0x04, 0x4b, 0x01, 0x55, 0x34, 0xb1, 0xa0, 0x39, 0xfb, 0xaf,
	0xe0, 0xe5, 0x5f, 0xea, 0x25, 0x4f, 0xf9, 0xc3, 0x81, 0xa7, 0xad, 0x5e, 0xfc, 0x6d, 0xf4, 0x93,
	0x79, 0x8d, 0x8c, 0x0a, 0x2d, 0x4e, 0x0a, 0x69, 0x1f, 0x5c, 0x0b, 0xde, 0x0e, 0x65, 0x87, 0x05,
	0xa6, 0x03, 0x97, 0x4c, 0x02, 0xda, 0x3b, 0xed, 0x96, 0x90, 0xca, 0x6c, 0x41, 0x81, 0x58, 0xa0,
	0x75, 0xb4, 0xb7, 0x93, 0x0d, 0x30, 0x67, 0x7f, 0x1b, 0xb6, 0xee, 0xb7, 0x62, 0x7b, 0x7c, 0x5b,
	0x01, 0x77, 0xe4, 0x49, 0xb4, 0x0a, 0x39, 0x72, 0xb9, 0x57, 0x5c, 0xb0, 0x87, 0x7a, 0xd1, 0xa9,
	0xff, 0x5a, 0x82, 0xc2, 0x29, 0x53, 0x03, 0x21, 0x6f, 0xce, 0x99, 0xec, 0x33, 0x89, 0x08, 0x6c,
	0xce, 0xdc, 0x0c, 0xa8, 0xa2, 0xad, 0x9d, 0x75, 0xd1, 0x97, 0x9f, 0x67, 0xb0, 0xc9, 0x5b, 0x5a,
	0x40, 0xc7, 0xb0, 0x71, 0xff, 0x7a, 0x40, 0x25, 0x9d, 0x92, 0x7a, 0xc9, 0x94, 0xcb, 0x69, 0xd4,
	0x58, 0x8a, 0xc0, 0xe6, 0x8c, 0x8d, 0x6d, 0x7b, 0x59, 0x5f, 0x25, 0xdb, 0x5e, 0xb6, 0xf7, 0x8d,
	0xe6, 0x8c, 0x93, 0xad, 0x66, 0xd6, 0x52, 0x58, 0xcd, 0x6c, 0xfb, 0x2f, 0xa0, 0x33, 0x28, 0x3e,
	0x74, 0x39, 0xda, 0x4d, 0x9e, 0x2c, 0x6d, 0x2d, 0xca, 0x95, 0x74, 0x72, 0x2c, 0xf8, 0x1d, 0x4a,
	0x99, 0x86, 0x44, 0xaf, 0x75, 0xf2, 0xbc, 0xfd, 0x28, 0xbf, 0x99, 0xf3, 0xab, 0x71, 0xad, 0x06,
	0xe4, 0xa7, 0xbd, 0x84, 0x76, 0x74, 0x62, 0x8a, 0xd1, 0xcb, 0x78, 0x96, 0x18, 0x89, 0x5c, 0xad,
	0x98, 0xbf, 0x89, 0xef, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x25, 0x9b, 0x8b, 0x32, 0x0a,
	0x00, 0x00,
}
