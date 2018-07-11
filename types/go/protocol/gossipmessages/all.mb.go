// AUTO GENERATED FILE (by membufc proto compiler v0.0.14)
package gossipmessages

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message Header

// reader

type Header struct {
	// ProtocolVersion primitives.ProtocolVersion
	// VirtualChainId primitives.VirtualChainId
	// RecipientPublicKeys []primitives.Ed25519Pkey
	// RecipientMode RecipientsListMode
	// Topic HeaderTopic
	// NumPayloads uint32

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _Header_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint32,membuffers.TypeBytesArray,membuffers.TypeUint16,membuffers.TypeUnion,membuffers.TypeUint32,}
var _Header_Unions = [][]membuffers.FieldType{{membuffers.TypeUint16,membuffers.TypeUint16,membuffers.TypeUint16,}}

func HeaderReader(buf []byte) *Header {
	x := &Header{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _Header_Scheme, _Header_Unions)
	return x
}

func (x *Header) IsValid() bool {
	return x._message.IsValid()
}

func (x *Header) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *Header) ProtocolVersion() primitives.ProtocolVersion {
	return primitives.ProtocolVersion(x._message.GetUint32(0))
}

func (x *Header) RawProtocolVersion() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *Header) MutateProtocolVersion(v primitives.ProtocolVersion) error {
	return x._message.SetUint32(0, uint32(v))
}

func (x *Header) VirtualChainId() primitives.VirtualChainId {
	return primitives.VirtualChainId(x._message.GetUint32(1))
}

func (x *Header) RawVirtualChainId() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *Header) MutateVirtualChainId(v primitives.VirtualChainId) error {
	return x._message.SetUint32(1, uint32(v))
}

func (x *Header) RecipientPublicKeysIterator() *HeaderRecipientPublicKeysIterator {
	return &HeaderRecipientPublicKeysIterator{iterator: x._message.GetBytesArrayIterator(2)}
}

type HeaderRecipientPublicKeysIterator struct {
	iterator *membuffers.Iterator
}

func (i *HeaderRecipientPublicKeysIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *HeaderRecipientPublicKeysIterator) NextRecipientPublicKeys() primitives.Ed25519Pkey {
	return primitives.Ed25519Pkey(i.iterator.NextBytes())
}

func (x *Header) RawRecipientPublicKeysArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *Header) RecipientMode() RecipientsListMode {
	return RecipientsListMode(x._message.GetUint16(3))
}

func (x *Header) RawRecipientMode() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *Header) MutateRecipientMode(v RecipientsListMode) error {
	return x._message.SetUint16(3, uint16(v))
}

type HeaderTopic uint16

const (
	HEADER_TOPIC_TRANSACTION_RELAY HeaderTopic = 0
	HEADER_TOPIC_BLOCK_SYNC HeaderTopic = 1
	HEADER_TOPIC_LEAN_HELIX HeaderTopic = 2
)

func (x *Header) Topic() HeaderTopic {
	return HeaderTopic(x._message.GetUint16(4))
}

func (x *Header) IsTopicTransactionRelay() bool {
	is, _ := x._message.IsUnionIndex(4, 0, 0)
	return is
}

func (x *Header) TransactionRelay() TransactionsRelayMessageType {
	_, off := x._message.IsUnionIndex(4, 0, 0)
	return TransactionsRelayMessageType(x._message.GetUint16InOffset(off))
}

func (x *Header) MutateTransactionRelay(v TransactionsRelayMessageType) error {
	is, off := x._message.IsUnionIndex(4, 0, 0)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *Header) IsTopicBlockSync() bool {
	is, _ := x._message.IsUnionIndex(4, 0, 1)
	return is
}

func (x *Header) BlockSync() BlockSyncMessageType {
	_, off := x._message.IsUnionIndex(4, 0, 1)
	return BlockSyncMessageType(x._message.GetUint16InOffset(off))
}

func (x *Header) MutateBlockSync(v BlockSyncMessageType) error {
	is, off := x._message.IsUnionIndex(4, 0, 1)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *Header) IsTopicLeanHelix() bool {
	is, _ := x._message.IsUnionIndex(4, 0, 2)
	return is
}

func (x *Header) LeanHelix() LeanHelixMessageType {
	_, off := x._message.IsUnionIndex(4, 0, 2)
	return LeanHelixMessageType(x._message.GetUint16InOffset(off))
}

func (x *Header) MutateLeanHelix(v LeanHelixMessageType) error {
	is, off := x._message.IsUnionIndex(4, 0, 2)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *Header) RawTopic() []byte {
	return x._message.RawBufferForField(4, 0)
}

func (x *Header) NumPayloads() uint32 {
	return x._message.GetUint32(5)
}

func (x *Header) RawNumPayloads() []byte {
	return x._message.RawBufferForField(5, 0)
}

func (x *Header) MutateNumPayloads(v uint32) error {
	return x._message.SetUint32(5, v)
}

// builder

type HeaderBuilder struct {
	ProtocolVersion primitives.ProtocolVersion
	VirtualChainId primitives.VirtualChainId
	RecipientPublicKeys []primitives.Ed25519Pkey
	RecipientMode RecipientsListMode
	Topic HeaderTopic
	NumPayloads uint32
	TransactionRelay TransactionsRelayMessageType
	BlockSync BlockSyncMessageType
	LeanHelix LeanHelixMessageType

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *HeaderBuilder) arrayOfRecipientPublicKeys() [][]byte {
	res := make([][]byte, len(w.RecipientPublicKeys))
	for i, v := range w.RecipientPublicKeys {
		res[i] = v
	}
	return res
}

func (w *HeaderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUint32(buf, uint32(w.ProtocolVersion))
	w._builder.WriteUint32(buf, uint32(w.VirtualChainId))
	w._builder.WriteBytesArray(buf, w.arrayOfRecipientPublicKeys())
	w._builder.WriteUint16(buf, uint16(w.RecipientMode))
	w._builder.WriteUnionIndex(buf, uint16(w.Topic))
	switch w.Topic {
	case HEADER_TOPIC_TRANSACTION_RELAY:
		w._builder.WriteUint16(buf, uint16(w.TransactionRelay))
	case HEADER_TOPIC_BLOCK_SYNC:
		w._builder.WriteUint16(buf, uint16(w.BlockSync))
	case HEADER_TOPIC_LEAN_HELIX:
		w._builder.WriteUint16(buf, uint16(w.LeanHelix))
	}
	w._builder.WriteUint32(buf, w.NumPayloads)
	return nil
}

func (w *HeaderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *HeaderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *HeaderBuilder) Build() *Header {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return HeaderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type RecipientsListMode uint16

const (
	RECIPIENT_LIST_MODE_BROADCAST RecipientsListMode = 0
	RECIPIENT_LIST_MODE_LIST RecipientsListMode = 1
	RECIPIENT_LIST_MODE_ALL_BUT_LIST RecipientsListMode = 2
)

