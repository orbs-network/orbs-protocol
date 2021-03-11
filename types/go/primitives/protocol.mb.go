// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package primitives

import (
	"bytes"
	"fmt"
)

type ProtocolVersion uint32

func (x ProtocolVersion) String() string {
	return fmt.Sprintf("%x", uint32(x))
}

func (x ProtocolVersion) Equal(y ProtocolVersion) bool {
	return x == y
}

func (x ProtocolVersion) KeyForMap() uint32 {
	return uint32(x)
}

type VirtualChainId uint32

func (x VirtualChainId) String() string {
	return fmt.Sprintf("%x", uint32(x))
}

func (x VirtualChainId) Equal(y VirtualChainId) bool {
	return x == y
}

func (x VirtualChainId) KeyForMap() uint32 {
	return uint32(x)
}

type BlockHeight uint64

func (x BlockHeight) String() string {
	return fmt.Sprintf("%x", uint64(x))
}

func (x BlockHeight) Equal(y BlockHeight) bool {
	return x == y
}

func (x BlockHeight) KeyForMap() uint64 {
	return uint64(x)
}

type TimestampNano uint64

func (x TimestampNano) String() string {
	return fmt.Sprintf("%x", uint64(x))
}

func (x TimestampNano) Equal(y TimestampNano) bool {
	return x == y
}

func (x TimestampNano) KeyForMap() uint64 {
	return uint64(x)
}

type TimestampSeconds uint32

func (x TimestampSeconds) String() string {
	return fmt.Sprintf("%x", uint32(x))
}

func (x TimestampSeconds) Equal(y TimestampSeconds) bool {
	return x == y
}

func (x TimestampSeconds) KeyForMap() uint32 {
	return uint32(x)
}

type NodeAddress []byte

func (x NodeAddress) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x NodeAddress) Equal(y NodeAddress) bool {
	return bytes.Equal(x, y)
}

func (x NodeAddress) KeyForMap() string {
	return string(x)
}

type ClientAddress []byte

func (x ClientAddress) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x ClientAddress) Equal(y ClientAddress) bool {
	return bytes.Equal(x, y)
}

func (x ClientAddress) KeyForMap() string {
	return string(x)
}

type ContractName string

func (x ContractName) String() string {
	return fmt.Sprintf(string(x))
}

func (x ContractName) Equal(y ContractName) bool {
	return x == y
}

func (x ContractName) KeyForMap() string {
	return string(x)
}

type MethodName string

func (x MethodName) String() string {
	return fmt.Sprintf(string(x))
}

func (x MethodName) Equal(y MethodName) bool {
	return x == y
}

func (x MethodName) KeyForMap() string {
	return string(x)
}

type EventName string

func (x EventName) String() string {
	return fmt.Sprintf(string(x))
}

func (x EventName) Equal(y EventName) bool {
	return x == y
}

func (x EventName) KeyForMap() string {
	return string(x)
}

type ExecutionContextId []byte

func (x ExecutionContextId) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x ExecutionContextId) Equal(y ExecutionContextId) bool {
	return bytes.Equal(x, y)
}

func (x ExecutionContextId) KeyForMap() string {
	return string(x)
}

type LeanHelixMessageContent []byte

func (x LeanHelixMessageContent) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x LeanHelixMessageContent) Equal(y LeanHelixMessageContent) bool {
	return bytes.Equal(x, y)
}

func (x LeanHelixMessageContent) KeyForMap() string {
	return string(x)
}

type MerkleTreeProof []byte

func (x MerkleTreeProof) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x MerkleTreeProof) Equal(y MerkleTreeProof) bool {
	return bytes.Equal(x, y)
}

func (x MerkleTreeProof) KeyForMap() string {
	return string(x)
}

type LeanHelixBlockProof []byte

func (x LeanHelixBlockProof) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x LeanHelixBlockProof) Equal(y LeanHelixBlockProof) bool {
	return bytes.Equal(x, y)
}

func (x LeanHelixBlockProof) KeyForMap() string {
	return string(x)
}

type PackedReceiptProof []byte

func (x PackedReceiptProof) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x PackedReceiptProof) Equal(y PackedReceiptProof) bool {
	return bytes.Equal(x, y)
}

func (x PackedReceiptProof) KeyForMap() string {
	return string(x)
}

type PackedEventsArray []byte

func (x PackedEventsArray) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x PackedEventsArray) Equal(y PackedEventsArray) bool {
	return bytes.Equal(x, y)
}

func (x PackedEventsArray) KeyForMap() string {
	return string(x)
}

type PackedArgumentArray []byte

func (x PackedArgumentArray) String() string {
	return fmt.Sprintf("%x", []byte(x))
}

func (x PackedArgumentArray) Equal(y PackedArgumentArray) bool {
	return bytes.Equal(x, y)
}

func (x PackedArgumentArray) KeyForMap() string {
	return string(x)
}

type Weight uint64

func (x Weight) String() string {
	return fmt.Sprintf("%x", uint64(x))
}

func (x Weight) Equal(y Weight) bool {
	return x == y
}

func (x Weight) KeyForMap() uint64 {
	return uint64(x)
}

type StorageKeys uint64

func (x StorageKeys) String() string {
	return fmt.Sprintf("%x", uint64(x))
}

func (x StorageKeys) Equal(y StorageKeys) bool {
	return x == y
}

func (x StorageKeys) KeyForMap() uint64 {
	return uint64(x)
}

type StorageSizeMegabyte uint64

func (x StorageSizeMegabyte) String() string {
	return fmt.Sprintf("%x", uint64(x))
}

func (x StorageSizeMegabyte) Equal(y StorageSizeMegabyte) bool {
	return x == y
}

func (x StorageSizeMegabyte) KeyForMap() uint64 {
	return uint64(x)
}
