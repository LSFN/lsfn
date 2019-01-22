// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/sensor.proto

package lsfn

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

// This gives the difference between a control that is boolean, one that's an int, and one that's a float.
type SensorDescription_SensorType int32

const (
	// Galactic Coordinates defined as a position in the Galaxy.
	SensorDescription_GalacticCoordinates SensorDescription_SensorType = 0
	// Radar
	SensorDescription_VolumetricRadar SensorDescription_SensorType = 1
	// Identify-Friend-Foe Detector
	SensorDescription_IFFScanner SensorDescription_SensorType = 2
)

var SensorDescription_SensorType_name = map[int32]string{
	0: "GalacticCoordinates",
	1: "VolumetricRadar",
	2: "IFFScanner",
}

var SensorDescription_SensorType_value = map[string]int32{
	"GalacticCoordinates": 0,
	"VolumetricRadar":     1,
	"IFFScanner":          2,
}

func (x SensorDescription_SensorType) String() string {
	return proto.EnumName(SensorDescription_SensorType_name, int32(x))
}

func (SensorDescription_SensorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a01333b65c596309, []int{0, 0}
}

// Describe a Sensor. Sensors are things that the ship knows. For example it's galactic position is a sensor, radar is a
// sensor. Sensors aren't changed, only described.
type SensorDescription struct {
	Id         string                       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string                       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SensorType SensorDescription_SensorType `protobuf:"varint,3,opt,name=sensorType,proto3,enum=lsfn.SensorDescription_SensorType" json:"sensorType,omitempty"`
	// Only provided for Volumetric Radar type
	RadarRange *VolumetricRange `protobuf:"bytes,4,opt,name=radarRange,proto3" json:"radarRange,omitempty"`
	// Only provided for IFF Detector type
	IffDetectorRange     *VolumetricRange `protobuf:"bytes,5,opt,name=iffDetectorRange,proto3" json:"iffDetectorRange,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SensorDescription) Reset()         { *m = SensorDescription{} }
func (m *SensorDescription) String() string { return proto.CompactTextString(m) }
func (*SensorDescription) ProtoMessage()    {}
func (*SensorDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_a01333b65c596309, []int{0}
}

func (m *SensorDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SensorDescription.Unmarshal(m, b)
}
func (m *SensorDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SensorDescription.Marshal(b, m, deterministic)
}
func (m *SensorDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SensorDescription.Merge(m, src)
}
func (m *SensorDescription) XXX_Size() int {
	return xxx_messageInfo_SensorDescription.Size(m)
}
func (m *SensorDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_SensorDescription.DiscardUnknown(m)
}

var xxx_messageInfo_SensorDescription proto.InternalMessageInfo

func (m *SensorDescription) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SensorDescription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SensorDescription) GetSensorType() SensorDescription_SensorType {
	if m != nil {
		return m.SensorType
	}
	return SensorDescription_GalacticCoordinates
}

func (m *SensorDescription) GetRadarRange() *VolumetricRange {
	if m != nil {
		return m.RadarRange
	}
	return nil
}

func (m *SensorDescription) GetIffDetectorRange() *VolumetricRange {
	if m != nil {
		return m.IffDetectorRange
	}
	return nil
}

type GalacticCoordinate struct {
	// Okay we're totally cheating here and using X, Y, and Z. This should be more
	// complex. https://en.wikipedia.org/wiki/Galactic_coordinate_system Regardless, these are *global* coordinates and
	// are not relative to the ship like the IFF and Radar coordinates are.
	X                    float64  `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float64  `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float64  `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GalacticCoordinate) Reset()         { *m = GalacticCoordinate{} }
func (m *GalacticCoordinate) String() string { return proto.CompactTextString(m) }
func (*GalacticCoordinate) ProtoMessage()    {}
func (*GalacticCoordinate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a01333b65c596309, []int{1}
}

func (m *GalacticCoordinate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GalacticCoordinate.Unmarshal(m, b)
}
func (m *GalacticCoordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GalacticCoordinate.Marshal(b, m, deterministic)
}
func (m *GalacticCoordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GalacticCoordinate.Merge(m, src)
}
func (m *GalacticCoordinate) XXX_Size() int {
	return xxx_messageInfo_GalacticCoordinate.Size(m)
}
func (m *GalacticCoordinate) XXX_DiscardUnknown() {
	xxx_messageInfo_GalacticCoordinate.DiscardUnknown(m)
}

var xxx_messageInfo_GalacticCoordinate proto.InternalMessageInfo

func (m *GalacticCoordinate) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *GalacticCoordinate) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *GalacticCoordinate) GetZ() float64 {
	if m != nil {
		return m.Z
	}
	return 0
}

type ShipRelativeCoordinate struct {
	// XYZ coordinate centered on the ship at `(0, 0, 0)`. Scale the same as volumetric data.
	X                    float64  `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float64  `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float64  `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShipRelativeCoordinate) Reset()         { *m = ShipRelativeCoordinate{} }
func (m *ShipRelativeCoordinate) String() string { return proto.CompactTextString(m) }
func (*ShipRelativeCoordinate) ProtoMessage()    {}
func (*ShipRelativeCoordinate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a01333b65c596309, []int{2}
}

func (m *ShipRelativeCoordinate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShipRelativeCoordinate.Unmarshal(m, b)
}
func (m *ShipRelativeCoordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShipRelativeCoordinate.Marshal(b, m, deterministic)
}
func (m *ShipRelativeCoordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShipRelativeCoordinate.Merge(m, src)
}
func (m *ShipRelativeCoordinate) XXX_Size() int {
	return xxx_messageInfo_ShipRelativeCoordinate.Size(m)
}
func (m *ShipRelativeCoordinate) XXX_DiscardUnknown() {
	xxx_messageInfo_ShipRelativeCoordinate.DiscardUnknown(m)
}

var xxx_messageInfo_ShipRelativeCoordinate proto.InternalMessageInfo

func (m *ShipRelativeCoordinate) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *ShipRelativeCoordinate) GetY() float64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *ShipRelativeCoordinate) GetZ() float64 {
	if m != nil {
		return m.Z
	}
	return 0
}

type IFFPing struct {
	Location             *ShipRelativeCoordinate `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Name                 string                  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *IFFPing) Reset()         { *m = IFFPing{} }
func (m *IFFPing) String() string { return proto.CompactTextString(m) }
func (*IFFPing) ProtoMessage()    {}
func (*IFFPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_a01333b65c596309, []int{3}
}

func (m *IFFPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IFFPing.Unmarshal(m, b)
}
func (m *IFFPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IFFPing.Marshal(b, m, deterministic)
}
func (m *IFFPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IFFPing.Merge(m, src)
}
func (m *IFFPing) XXX_Size() int {
	return xxx_messageInfo_IFFPing.Size(m)
}
func (m *IFFPing) XXX_DiscardUnknown() {
	xxx_messageInfo_IFFPing.DiscardUnknown(m)
}

var xxx_messageInfo_IFFPing proto.InternalMessageInfo

func (m *IFFPing) GetLocation() *ShipRelativeCoordinate {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *IFFPing) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type IFFReadout struct {
	Pings                []*IFFPing `protobuf:"bytes,1,rep,name=pings,proto3" json:"pings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *IFFReadout) Reset()         { *m = IFFReadout{} }
func (m *IFFReadout) String() string { return proto.CompactTextString(m) }
func (*IFFReadout) ProtoMessage()    {}
func (*IFFReadout) Descriptor() ([]byte, []int) {
	return fileDescriptor_a01333b65c596309, []int{4}
}

func (m *IFFReadout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IFFReadout.Unmarshal(m, b)
}
func (m *IFFReadout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IFFReadout.Marshal(b, m, deterministic)
}
func (m *IFFReadout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IFFReadout.Merge(m, src)
}
func (m *IFFReadout) XXX_Size() int {
	return xxx_messageInfo_IFFReadout.Size(m)
}
func (m *IFFReadout) XXX_DiscardUnknown() {
	xxx_messageInfo_IFFReadout.DiscardUnknown(m)
}

var xxx_messageInfo_IFFReadout proto.InternalMessageInfo

func (m *IFFReadout) GetPings() []*IFFPing {
	if m != nil {
		return m.Pings
	}
	return nil
}

type VolumetricRange struct {
	// This is a little complex as the volumentric data is just given as a 3D array. So this gives you the offests to
	// figure out where the ship is in that (as it's centered on the ship). So if `upY` is 20 and `downY` is 40 then you
	// expect a y dimention in the volumetric data of size 60, where the ship is 20 from the top.
	LeftX                int32    `protobuf:"varint,1,opt,name=leftX,proto3" json:"leftX,omitempty"`
	RightX               int32    `protobuf:"varint,2,opt,name=rightX,proto3" json:"rightX,omitempty"`
	UpY                  int32    `protobuf:"varint,3,opt,name=upY,proto3" json:"upY,omitempty"`
	DownY                int32    `protobuf:"varint,4,opt,name=downY,proto3" json:"downY,omitempty"`
	ForwardsZ            int32    `protobuf:"varint,5,opt,name=forwardsZ,proto3" json:"forwardsZ,omitempty"`
	BackwardsZ           int32    `protobuf:"varint,6,opt,name=backwardsZ,proto3" json:"backwardsZ,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VolumetricRange) Reset()         { *m = VolumetricRange{} }
func (m *VolumetricRange) String() string { return proto.CompactTextString(m) }
func (*VolumetricRange) ProtoMessage()    {}
func (*VolumetricRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a01333b65c596309, []int{5}
}

func (m *VolumetricRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumetricRange.Unmarshal(m, b)
}
func (m *VolumetricRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumetricRange.Marshal(b, m, deterministic)
}
func (m *VolumetricRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumetricRange.Merge(m, src)
}
func (m *VolumetricRange) XXX_Size() int {
	return xxx_messageInfo_VolumetricRange.Size(m)
}
func (m *VolumetricRange) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumetricRange.DiscardUnknown(m)
}

var xxx_messageInfo_VolumetricRange proto.InternalMessageInfo

func (m *VolumetricRange) GetLeftX() int32 {
	if m != nil {
		return m.LeftX
	}
	return 0
}

func (m *VolumetricRange) GetRightX() int32 {
	if m != nil {
		return m.RightX
	}
	return 0
}

func (m *VolumetricRange) GetUpY() int32 {
	if m != nil {
		return m.UpY
	}
	return 0
}

func (m *VolumetricRange) GetDownY() int32 {
	if m != nil {
		return m.DownY
	}
	return 0
}

func (m *VolumetricRange) GetForwardsZ() int32 {
	if m != nil {
		return m.ForwardsZ
	}
	return 0
}

func (m *VolumetricRange) GetBackwardsZ() int32 {
	if m != nil {
		return m.BackwardsZ
	}
	return 0
}

type VolumetricData struct {
	Palnes               []*VolumetricDataPlane `protobuf:"bytes,1,rep,name=palnes,proto3" json:"palnes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *VolumetricData) Reset()         { *m = VolumetricData{} }
func (m *VolumetricData) String() string { return proto.CompactTextString(m) }
func (*VolumetricData) ProtoMessage()    {}
func (*VolumetricData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a01333b65c596309, []int{6}
}

func (m *VolumetricData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumetricData.Unmarshal(m, b)
}
func (m *VolumetricData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumetricData.Marshal(b, m, deterministic)
}
func (m *VolumetricData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumetricData.Merge(m, src)
}
func (m *VolumetricData) XXX_Size() int {
	return xxx_messageInfo_VolumetricData.Size(m)
}
func (m *VolumetricData) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumetricData.DiscardUnknown(m)
}

var xxx_messageInfo_VolumetricData proto.InternalMessageInfo

func (m *VolumetricData) GetPalnes() []*VolumetricDataPlane {
	if m != nil {
		return m.Palnes
	}
	return nil
}

// Okay three dimentional volumetric data is hard okay! Each cell is a float that describes the intensity of the
// thinig in that cell from 0.0f to 1.0f.
type VolumetricDataPlane struct {
	Lines                []*VolumetricDataPlaneLine `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *VolumetricDataPlane) Reset()         { *m = VolumetricDataPlane{} }
func (m *VolumetricDataPlane) String() string { return proto.CompactTextString(m) }
func (*VolumetricDataPlane) ProtoMessage()    {}
func (*VolumetricDataPlane) Descriptor() ([]byte, []int) {
	return fileDescriptor_a01333b65c596309, []int{6, 0}
}

func (m *VolumetricDataPlane) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumetricDataPlane.Unmarshal(m, b)
}
func (m *VolumetricDataPlane) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumetricDataPlane.Marshal(b, m, deterministic)
}
func (m *VolumetricDataPlane) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumetricDataPlane.Merge(m, src)
}
func (m *VolumetricDataPlane) XXX_Size() int {
	return xxx_messageInfo_VolumetricDataPlane.Size(m)
}
func (m *VolumetricDataPlane) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumetricDataPlane.DiscardUnknown(m)
}

var xxx_messageInfo_VolumetricDataPlane proto.InternalMessageInfo

func (m *VolumetricDataPlane) GetLines() []*VolumetricDataPlaneLine {
	if m != nil {
		return m.Lines
	}
	return nil
}

type VolumetricDataPlaneLine struct {
	Cell                 []float32 `protobuf:"fixed32,1,rep,packed,name=cell,proto3" json:"cell,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *VolumetricDataPlaneLine) Reset()         { *m = VolumetricDataPlaneLine{} }
func (m *VolumetricDataPlaneLine) String() string { return proto.CompactTextString(m) }
func (*VolumetricDataPlaneLine) ProtoMessage()    {}
func (*VolumetricDataPlaneLine) Descriptor() ([]byte, []int) {
	return fileDescriptor_a01333b65c596309, []int{6, 0, 0}
}

func (m *VolumetricDataPlaneLine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VolumetricDataPlaneLine.Unmarshal(m, b)
}
func (m *VolumetricDataPlaneLine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VolumetricDataPlaneLine.Marshal(b, m, deterministic)
}
func (m *VolumetricDataPlaneLine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VolumetricDataPlaneLine.Merge(m, src)
}
func (m *VolumetricDataPlaneLine) XXX_Size() int {
	return xxx_messageInfo_VolumetricDataPlaneLine.Size(m)
}
func (m *VolumetricDataPlaneLine) XXX_DiscardUnknown() {
	xxx_messageInfo_VolumetricDataPlaneLine.DiscardUnknown(m)
}

var xxx_messageInfo_VolumetricDataPlaneLine proto.InternalMessageInfo

func (m *VolumetricDataPlaneLine) GetCell() []float32 {
	if m != nil {
		return m.Cell
	}
	return nil
}

// Describe the state of a Sensor. This is a message sent from the environment to the vessel to describe the current
// state.
type SensorState struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to SensorTypeValue:
	//	*SensorState_GalacticCoordinate
	//	*SensorState_RadarData
	//	*SensorState_IffData
	SensorTypeValue      isSensorState_SensorTypeValue `protobuf_oneof:"SensorTypeValue"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *SensorState) Reset()         { *m = SensorState{} }
func (m *SensorState) String() string { return proto.CompactTextString(m) }
func (*SensorState) ProtoMessage()    {}
func (*SensorState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a01333b65c596309, []int{7}
}

func (m *SensorState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SensorState.Unmarshal(m, b)
}
func (m *SensorState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SensorState.Marshal(b, m, deterministic)
}
func (m *SensorState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SensorState.Merge(m, src)
}
func (m *SensorState) XXX_Size() int {
	return xxx_messageInfo_SensorState.Size(m)
}
func (m *SensorState) XXX_DiscardUnknown() {
	xxx_messageInfo_SensorState.DiscardUnknown(m)
}

var xxx_messageInfo_SensorState proto.InternalMessageInfo

func (m *SensorState) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type isSensorState_SensorTypeValue interface {
	isSensorState_SensorTypeValue()
}

type SensorState_GalacticCoordinate struct {
	GalacticCoordinate *GalacticCoordinate `protobuf:"bytes,2,opt,name=galacticCoordinate,proto3,oneof"`
}

type SensorState_RadarData struct {
	RadarData *VolumetricData `protobuf:"bytes,3,opt,name=radarData,proto3,oneof"`
}

type SensorState_IffData struct {
	IffData *IFFReadout `protobuf:"bytes,4,opt,name=iffData,proto3,oneof"`
}

func (*SensorState_GalacticCoordinate) isSensorState_SensorTypeValue() {}

func (*SensorState_RadarData) isSensorState_SensorTypeValue() {}

func (*SensorState_IffData) isSensorState_SensorTypeValue() {}

func (m *SensorState) GetSensorTypeValue() isSensorState_SensorTypeValue {
	if m != nil {
		return m.SensorTypeValue
	}
	return nil
}

func (m *SensorState) GetGalacticCoordinate() *GalacticCoordinate {
	if x, ok := m.GetSensorTypeValue().(*SensorState_GalacticCoordinate); ok {
		return x.GalacticCoordinate
	}
	return nil
}

func (m *SensorState) GetRadarData() *VolumetricData {
	if x, ok := m.GetSensorTypeValue().(*SensorState_RadarData); ok {
		return x.RadarData
	}
	return nil
}

func (m *SensorState) GetIffData() *IFFReadout {
	if x, ok := m.GetSensorTypeValue().(*SensorState_IffData); ok {
		return x.IffData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SensorState) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SensorState_GalacticCoordinate)(nil),
		(*SensorState_RadarData)(nil),
		(*SensorState_IffData)(nil),
	}
}

func init() {
	proto.RegisterEnum("lsfn.SensorDescription_SensorType", SensorDescription_SensorType_name, SensorDescription_SensorType_value)
	proto.RegisterType((*SensorDescription)(nil), "lsfn.SensorDescription")
	proto.RegisterType((*GalacticCoordinate)(nil), "lsfn.GalacticCoordinate")
	proto.RegisterType((*ShipRelativeCoordinate)(nil), "lsfn.ShipRelativeCoordinate")
	proto.RegisterType((*IFFPing)(nil), "lsfn.IFFPing")
	proto.RegisterType((*IFFReadout)(nil), "lsfn.IFFReadout")
	proto.RegisterType((*VolumetricRange)(nil), "lsfn.VolumetricRange")
	proto.RegisterType((*VolumetricData)(nil), "lsfn.VolumetricData")
	proto.RegisterType((*VolumetricDataPlane)(nil), "lsfn.VolumetricData.plane")
	proto.RegisterType((*VolumetricDataPlaneLine)(nil), "lsfn.VolumetricData.plane.line")
	proto.RegisterType((*SensorState)(nil), "lsfn.SensorState")
}

func init() { proto.RegisterFile("api/proto/sensor.proto", fileDescriptor_a01333b65c596309) }

var fileDescriptor_a01333b65c596309 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xae, 0xb3, 0xa6, 0xfb, 0x77, 0xf2, 0xd3, 0x65, 0xde, 0x18, 0x51, 0x35, 0x41, 0x15, 0x6e,
	0x7a, 0x81, 0x3a, 0x11, 0x98, 0xc4, 0x15, 0x12, 0xa3, 0x2a, 0xdd, 0xae, 0x90, 0x8b, 0xc6, 0xb6,
	0x3b, 0x2f, 0x71, 0x3b, 0x0b, 0xcf, 0x8e, 0x1c, 0x97, 0xad, 0x7b, 0x1a, 0x6e, 0x78, 0x10, 0xde,
	0x84, 0x47, 0x41, 0x76, 0xb2, 0xa5, 0xac, 0x85, 0x0b, 0xee, 0x7c, 0x8e, 0xbf, 0xef, 0xf8, 0x9c,
	0x2f, 0xdf, 0x09, 0xec, 0xd2, 0x9c, 0xef, 0xe7, 0x5a, 0x19, 0xb5, 0x5f, 0x30, 0x59, 0x28, 0xdd,
	0x77, 0x01, 0x6e, 0x8a, 0x62, 0x22, 0xe3, 0x1f, 0x1e, 0x6c, 0x8d, 0x5d, 0x7a, 0xc0, 0x8a, 0x54,
	0xf3, 0xdc, 0x70, 0x25, 0x71, 0x1b, 0x3c, 0x9e, 0x45, 0xa8, 0x8b, 0x7a, 0x1b, 0xc4, 0xe3, 0x19,
	0xc6, 0xd0, 0x94, 0xf4, 0x8a, 0x45, 0x9e, 0xcb, 0xb8, 0x33, 0x3e, 0x04, 0x28, 0xeb, 0x7d, 0x9a,
	0xe7, 0x2c, 0x5a, 0xeb, 0xa2, 0x5e, 0x3b, 0x89, 0xfb, 0xb6, 0x68, 0x7f, 0xa9, 0x60, 0x95, 0xb1,
	0x48, 0xb2, 0xc0, 0xc2, 0x07, 0x00, 0x9a, 0x66, 0x54, 0x13, 0x2a, 0xa7, 0x2c, 0x6a, 0x76, 0x51,
	0x2f, 0x48, 0x1e, 0x97, 0x35, 0x4e, 0x94, 0x98, 0x5d, 0x31, 0xa3, 0x79, 0xea, 0x2e, 0xc9, 0x02,
	0x10, 0xbf, 0x83, 0x90, 0x4f, 0x26, 0x03, 0x66, 0x58, 0x6a, 0x54, 0x45, 0xf6, 0xff, 0x46, 0x5e,
	0x82, 0xc7, 0xc7, 0x00, 0x75, 0x4f, 0xf8, 0x09, 0x6c, 0x7f, 0xa0, 0x82, 0xa6, 0x86, 0xa7, 0xef,
	0x95, 0xd2, 0x19, 0x97, 0xd4, 0xb0, 0x22, 0x6c, 0xe0, 0x6d, 0xd8, 0x5c, 0xac, 0x95, 0x51, 0x1d,
	0x22, 0xdc, 0x06, 0x38, 0x1a, 0x0e, 0xc7, 0x29, 0x95, 0x92, 0xe9, 0xd0, 0x8b, 0xdf, 0x02, 0x5e,
	0x66, 0xe3, 0xff, 0x01, 0xdd, 0x38, 0x09, 0x11, 0x41, 0x37, 0x36, 0x9a, 0x3b, 0xf9, 0x10, 0x41,
	0x73, 0x1b, 0xdd, 0x3a, 0xc9, 0x10, 0x41, 0xb7, 0xf1, 0x21, 0xec, 0x8e, 0x2f, 0x79, 0x4e, 0x98,
	0xa0, 0x86, 0x7f, 0x65, 0xff, 0x54, 0xe3, 0x33, 0xac, 0x1f, 0x0d, 0x87, 0x1f, 0xb9, 0x9c, 0xe2,
	0x37, 0xf0, 0x9f, 0x50, 0x29, 0xb5, 0xba, 0x3b, 0x6e, 0x90, 0xec, 0x55, 0x9f, 0x65, 0xe5, 0x23,
	0xe4, 0x1e, 0xbd, 0xea, 0x33, 0xc7, 0x2f, 0xdd, 0xb0, 0x84, 0xd1, 0x4c, 0xcd, 0x0c, 0x7e, 0x0e,
	0x7e, 0xce, 0xe5, 0xb4, 0x88, 0x50, 0x77, 0xad, 0x17, 0x24, 0x8f, 0xca, 0xc2, 0xd5, 0xcb, 0xa4,
	0xbc, 0x8b, 0xbf, 0xa3, 0xdf, 0x55, 0xb3, 0x9f, 0x6c, 0x07, 0x7c, 0xc1, 0x26, 0xe6, 0xd4, 0x75,
	0xe4, 0x93, 0x32, 0xc0, 0xbb, 0xd0, 0xd2, 0x7c, 0x7a, 0x69, 0x4e, 0xdd, 0x93, 0x3e, 0xa9, 0x22,
	0x1c, 0xc2, 0xda, 0x2c, 0x3f, 0x73, 0xd3, 0xf9, 0xc4, 0x1e, 0x2d, 0x3f, 0x53, 0xd7, 0xf2, 0xcc,
	0x99, 0xc4, 0x27, 0x65, 0x80, 0xf7, 0x60, 0x63, 0xa2, 0xf4, 0x35, 0xd5, 0x59, 0x71, 0xee, 0x1c,
	0xe0, 0x93, 0x3a, 0x81, 0x9f, 0x02, 0x5c, 0xd0, 0xf4, 0x4b, 0x75, 0xdd, 0x72, 0xd7, 0x0b, 0x99,
	0xf8, 0x1b, 0x82, 0x76, 0xdd, 0xe7, 0x80, 0x1a, 0x8a, 0x13, 0x68, 0xe5, 0x54, 0x48, 0x76, 0x37,
	0x60, 0xe7, 0xa1, 0x9f, 0x2c, 0xaa, 0x9f, 0x0b, 0x2a, 0x19, 0xa9, 0x90, 0x9d, 0x73, 0xf0, 0x5d,
	0x02, 0x1f, 0x80, 0x2f, 0x78, 0xcd, 0x7d, 0xf6, 0x67, 0x6e, 0xdf, 0xe2, 0x48, 0x89, 0xee, 0x74,
	0xa0, 0x69, 0x0f, 0x56, 0xfd, 0x94, 0x09, 0xe1, 0xd8, 0x1e, 0x71, 0xe7, 0xf8, 0x27, 0x82, 0xa0,
	0xf4, 0xe9, 0xd8, 0x58, 0x43, 0x3c, 0x5c, 0xcc, 0x63, 0xc0, 0xd3, 0x25, 0xeb, 0x39, 0x31, 0x83,
	0x24, 0x2a, 0xdf, 0x5f, 0xb6, 0xe6, 0xa8, 0x41, 0x56, 0xb0, 0xf0, 0x6b, 0xd8, 0x70, 0x3b, 0x66,
	0xdb, 0x74, 0xd2, 0x07, 0xc9, 0xce, 0xaa, 0x11, 0x46, 0x0d, 0x52, 0x03, 0xf1, 0x0b, 0x58, 0xb7,
	0xcb, 0x65, 0x39, 0xe5, 0xfe, 0x86, 0xf7, 0x9e, 0xa8, 0x4c, 0x33, 0x6a, 0x90, 0x3b, 0xc8, 0xe1,
	0x16, 0x6c, 0xd6, 0x6b, 0x77, 0x42, 0xc5, 0x8c, 0x5d, 0xb4, 0xdc, 0xef, 0xe8, 0xd5, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1e, 0xe7, 0x01, 0xd9, 0xa8, 0x04, 0x00, 0x00,
}
