// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package services

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service StateStorage

type StateStorage interface {
	CommitStateDiff(ctx context.Context, input *CommitStateDiffInput) (*CommitStateDiffOutput, error)
	ReadKeys(ctx context.Context, input *ReadKeysInput) (*ReadKeysOutput, error)
	GetLastCommittedBlockInfo(ctx context.Context, input *GetLastCommittedBlockInfoInput) (*GetLastCommittedBlockInfoOutput, error)
	GetStateHash(ctx context.Context, input *GetStateHashInput) (*GetStateHashOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message ReadKeysInput (non serializable)

type ReadKeysInput struct {
	BlockHeight  primitives.BlockHeight
	ContractName primitives.ContractName
	Keys         [][]byte
}

func (x *ReadKeysInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,ContractName:%s,Keys:%s,}", x.StringBlockHeight(), x.StringContractName(), x.StringKeys())
}

func (x *ReadKeysInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *ReadKeysInput) StringContractName() (res string) {
	res = fmt.Sprintf("%s", x.ContractName)
	return
}

func (x *ReadKeysInput) StringKeys() (res string) {
	res = "["
	for _, v := range x.Keys {
		res += fmt.Sprintf("%x", v) + ","
	}
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ReadKeysOutput (non serializable)

type ReadKeysOutput struct {
	StateRecords []*protocol.StateRecord
}

func (x *ReadKeysOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{StateRecords:%s,}", x.StringStateRecords())
}

func (x *ReadKeysOutput) StringStateRecords() (res string) {
	res = "["
	for _, v := range x.StateRecords {
		res += v.String() + ","
	}
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetLastCommittedBlockInfoInput (non serializable)

type GetLastCommittedBlockInfoInput struct {
}

func (x *GetLastCommittedBlockInfoInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message GetLastCommittedBlockInfoOutput (non serializable)

type GetLastCommittedBlockInfoOutput struct {
	BlockHeight          primitives.BlockHeight
	BlockTimestamp       primitives.TimestampNano
	BlockProposerAddress primitives.NodeAddress
	CurrentReferenceTime primitives.TimestampSeconds
	PrevReferenceTime    primitives.TimestampSeconds
	CurrentNumKeys       primitives.StorageKeys
	CurrentSize          primitives.StorageSizeMegabyte
}

func (x *GetLastCommittedBlockInfoOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,BlockTimestamp:%s,BlockProposerAddress:%s,CurrentReferenceTime:%s,PrevReferenceTime:%s,CurrentNumKeys:%s,CurrentSize:%s,}", x.StringBlockHeight(), x.StringBlockTimestamp(), x.StringBlockProposerAddress(), x.StringCurrentReferenceTime(), x.StringPrevReferenceTime(), x.StringCurrentNumKeys(), x.StringCurrentSize())
}

func (x *GetLastCommittedBlockInfoOutput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *GetLastCommittedBlockInfoOutput) StringBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.BlockTimestamp)
	return
}

func (x *GetLastCommittedBlockInfoOutput) StringBlockProposerAddress() (res string) {
	res = fmt.Sprintf("%s", x.BlockProposerAddress)
	return
}

func (x *GetLastCommittedBlockInfoOutput) StringCurrentReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.CurrentReferenceTime)
	return
}

func (x *GetLastCommittedBlockInfoOutput) StringPrevReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.PrevReferenceTime)
	return
}

func (x *GetLastCommittedBlockInfoOutput) StringCurrentNumKeys() (res string) {
	res = fmt.Sprintf("%s", x.CurrentNumKeys)
	return
}

func (x *GetLastCommittedBlockInfoOutput) StringCurrentSize() (res string) {
	res = fmt.Sprintf("%s", x.CurrentSize)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message CommitStateDiffInput (non serializable)

type CommitStateDiffInput struct {
	ResultsBlockHeader *protocol.ResultsBlockHeader
	ContractStateDiffs []*protocol.ContractStateDiff
}

func (x *CommitStateDiffInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ResultsBlockHeader:%s,ContractStateDiffs:%s,}", x.StringResultsBlockHeader(), x.StringContractStateDiffs())
}

func (x *CommitStateDiffInput) StringResultsBlockHeader() (res string) {
	res = x.ResultsBlockHeader.String()
	return
}

func (x *CommitStateDiffInput) StringContractStateDiffs() (res string) {
	res = "["
	for _, v := range x.ContractStateDiffs {
		res += v.String() + ","
	}
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message CommitStateDiffOutput (non serializable)

type CommitStateDiffOutput struct {
	NextDesiredBlockHeight primitives.BlockHeight
}

func (x *CommitStateDiffOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{NextDesiredBlockHeight:%s,}", x.StringNextDesiredBlockHeight())
}

func (x *CommitStateDiffOutput) StringNextDesiredBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.NextDesiredBlockHeight)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateHashInput (non serializable)

type GetStateHashInput struct {
	BlockHeight primitives.BlockHeight
}

func (x *GetStateHashInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,}", x.StringBlockHeight())
}

func (x *GetStateHashInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetStateHashOutput (non serializable)

type GetStateHashOutput struct {
	StateMerkleRootHash primitives.Sha256
}

func (x *GetStateHashOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{StateMerkleRootHash:%s,}", x.StringStateMerkleRootHash())
}

func (x *GetStateHashOutput) StringStateMerkleRootHash() (res string) {
	res = fmt.Sprintf("%s", x.StateMerkleRootHash)
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums
