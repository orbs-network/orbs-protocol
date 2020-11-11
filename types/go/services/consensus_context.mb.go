// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package services

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service ConsensusContext

type ConsensusContext interface {
	RequestNewTransactionsBlock(ctx context.Context, input *RequestNewTransactionsBlockInput) (*RequestNewTransactionsBlockOutput, error)
	RequestNewResultsBlock(ctx context.Context, input *RequestNewResultsBlockInput) (*RequestNewResultsBlockOutput, error)
	ValidateTransactionsBlock(ctx context.Context, input *ValidateTransactionsBlockInput) (*ValidateTransactionsBlockOutput, error)
	ValidateResultsBlock(ctx context.Context, input *ValidateResultsBlockInput) (*ValidateResultsBlockOutput, error)
	RequestOrderingCommittee(ctx context.Context, input *RequestCommitteeInput) (*RequestCommitteeOutput, error)
	RequestBlockProofOrderingCommittee(ctx context.Context, input *RequestBlockProofCommitteeInput) (*RequestBlockProofCommitteeOutput, error)
	RequestValidationCommittee(ctx context.Context, input *RequestCommitteeInput) (*RequestCommitteeOutput, error)
	ValidateBlockReferenceTime(ctx context.Context, input *ValidateBlockCommitteeInput) (*ValidateBlockCommitteeOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewTransactionsBlockInput (non serializable)

type RequestNewTransactionsBlockInput struct {
	CurrentBlockHeight      primitives.BlockHeight
	MaxBlockSizeKb          uint32
	MaxNumberOfTransactions uint32
	PrevBlockHash           primitives.Sha256
	PrevBlockTimestamp      primitives.TimestampNano
	BlockProposerAddress    primitives.NodeAddress
	PrevBlockReferenceTime  primitives.TimestampSeconds
}

func (x *RequestNewTransactionsBlockInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,MaxBlockSizeKb:%s,MaxNumberOfTransactions:%s,PrevBlockHash:%s,PrevBlockTimestamp:%s,BlockProposerAddress:%s,PrevBlockReferenceTime:%s,}", x.StringCurrentBlockHeight(), x.StringMaxBlockSizeKb(), x.StringMaxNumberOfTransactions(), x.StringPrevBlockHash(), x.StringPrevBlockTimestamp(), x.StringBlockProposerAddress(), x.StringPrevBlockReferenceTime())
}

func (x *RequestNewTransactionsBlockInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
	return
}

func (x *RequestNewTransactionsBlockInput) StringMaxBlockSizeKb() (res string) {
	res = fmt.Sprintf("%v", x.MaxBlockSizeKb)
	return
}

func (x *RequestNewTransactionsBlockInput) StringMaxNumberOfTransactions() (res string) {
	res = fmt.Sprintf("%v", x.MaxNumberOfTransactions)
	return
}

func (x *RequestNewTransactionsBlockInput) StringPrevBlockHash() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockHash)
	return
}

func (x *RequestNewTransactionsBlockInput) StringPrevBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockTimestamp)
	return
}

func (x *RequestNewTransactionsBlockInput) StringBlockProposerAddress() (res string) {
	res = fmt.Sprintf("%s", x.BlockProposerAddress)
	return
}

func (x *RequestNewTransactionsBlockInput) StringPrevBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockReferenceTime)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewTransactionsBlockOutput (non serializable)

type RequestNewTransactionsBlockOutput struct {
	TransactionsBlock *protocol.TransactionsBlockContainer
}

func (x *RequestNewTransactionsBlockOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{TransactionsBlock:%s,}", x.StringTransactionsBlock())
}

func (x *RequestNewTransactionsBlockOutput) StringTransactionsBlock() (res string) {
	res = x.TransactionsBlock.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewResultsBlockInput (non serializable)

type RequestNewResultsBlockInput struct {
	CurrentBlockHeight     primitives.BlockHeight
	PrevBlockHash          primitives.Sha256
	TransactionsBlock      *protocol.TransactionsBlockContainer
	PrevBlockTimestamp     primitives.TimestampNano
	BlockProposerAddress   primitives.NodeAddress
	PrevBlockReferenceTime primitives.TimestampSeconds
}

func (x *RequestNewResultsBlockInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,PrevBlockHash:%s,TransactionsBlock:%s,PrevBlockTimestamp:%s,BlockProposerAddress:%s,PrevBlockReferenceTime:%s,}", x.StringCurrentBlockHeight(), x.StringPrevBlockHash(), x.StringTransactionsBlock(), x.StringPrevBlockTimestamp(), x.StringBlockProposerAddress(), x.StringPrevBlockReferenceTime())
}

func (x *RequestNewResultsBlockInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
	return
}

func (x *RequestNewResultsBlockInput) StringPrevBlockHash() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockHash)
	return
}

func (x *RequestNewResultsBlockInput) StringTransactionsBlock() (res string) {
	res = x.TransactionsBlock.String()
	return
}

func (x *RequestNewResultsBlockInput) StringPrevBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockTimestamp)
	return
}

func (x *RequestNewResultsBlockInput) StringBlockProposerAddress() (res string) {
	res = fmt.Sprintf("%s", x.BlockProposerAddress)
	return
}

func (x *RequestNewResultsBlockInput) StringPrevBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockReferenceTime)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message RequestNewResultsBlockOutput (non serializable)

type RequestNewResultsBlockOutput struct {
	ResultsBlock *protocol.ResultsBlockContainer
}

func (x *RequestNewResultsBlockOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ResultsBlock:%s,}", x.StringResultsBlock())
}

func (x *RequestNewResultsBlockOutput) StringResultsBlock() (res string) {
	res = x.ResultsBlock.String()
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsBlockInput (non serializable)

type ValidateTransactionsBlockInput struct {
	CurrentBlockHeight     primitives.BlockHeight
	TransactionsBlock      *protocol.TransactionsBlockContainer
	PrevBlockHash          primitives.Sha256
	PrevBlockTimestamp     primitives.TimestampNano
	BlockProposerAddress   primitives.NodeAddress
	PrevBlockReferenceTime primitives.TimestampSeconds
}

func (x *ValidateTransactionsBlockInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,TransactionsBlock:%s,PrevBlockHash:%s,PrevBlockTimestamp:%s,BlockProposerAddress:%s,PrevBlockReferenceTime:%s,}", x.StringCurrentBlockHeight(), x.StringTransactionsBlock(), x.StringPrevBlockHash(), x.StringPrevBlockTimestamp(), x.StringBlockProposerAddress(), x.StringPrevBlockReferenceTime())
}

func (x *ValidateTransactionsBlockInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
	return
}

func (x *ValidateTransactionsBlockInput) StringTransactionsBlock() (res string) {
	res = x.TransactionsBlock.String()
	return
}

func (x *ValidateTransactionsBlockInput) StringPrevBlockHash() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockHash)
	return
}

func (x *ValidateTransactionsBlockInput) StringPrevBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockTimestamp)
	return
}

func (x *ValidateTransactionsBlockInput) StringBlockProposerAddress() (res string) {
	res = fmt.Sprintf("%s", x.BlockProposerAddress)
	return
}

func (x *ValidateTransactionsBlockInput) StringPrevBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockReferenceTime)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateTransactionsBlockOutput (non serializable)

type ValidateTransactionsBlockOutput struct {
}

func (x *ValidateTransactionsBlockOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateResultsBlockInput (non serializable)

type ValidateResultsBlockInput struct {
	CurrentBlockHeight     primitives.BlockHeight
	ResultsBlock           *protocol.ResultsBlockContainer
	PrevBlockHash          primitives.Sha256
	TransactionsBlock      *protocol.TransactionsBlockContainer
	PrevBlockTimestamp     primitives.TimestampNano
	BlockProposerAddress   primitives.NodeAddress
	PrevBlockReferenceTime primitives.TimestampSeconds
}

func (x *ValidateResultsBlockInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,ResultsBlock:%s,PrevBlockHash:%s,TransactionsBlock:%s,PrevBlockTimestamp:%s,BlockProposerAddress:%s,PrevBlockReferenceTime:%s,}", x.StringCurrentBlockHeight(), x.StringResultsBlock(), x.StringPrevBlockHash(), x.StringTransactionsBlock(), x.StringPrevBlockTimestamp(), x.StringBlockProposerAddress(), x.StringPrevBlockReferenceTime())
}

func (x *ValidateResultsBlockInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
	return
}

func (x *ValidateResultsBlockInput) StringResultsBlock() (res string) {
	res = x.ResultsBlock.String()
	return
}

func (x *ValidateResultsBlockInput) StringPrevBlockHash() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockHash)
	return
}

func (x *ValidateResultsBlockInput) StringTransactionsBlock() (res string) {
	res = x.TransactionsBlock.String()
	return
}

func (x *ValidateResultsBlockInput) StringPrevBlockTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockTimestamp)
	return
}

func (x *ValidateResultsBlockInput) StringBlockProposerAddress() (res string) {
	res = fmt.Sprintf("%s", x.BlockProposerAddress)
	return
}

func (x *ValidateResultsBlockInput) StringPrevBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockReferenceTime)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateResultsBlockOutput (non serializable)

type ValidateResultsBlockOutput struct {
}

func (x *ValidateResultsBlockOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// message RequestCommitteeInput (non serializable)

type RequestCommitteeInput struct {
	CurrentBlockHeight     primitives.BlockHeight
	RandomSeed             uint64
	MaxCommitteeSize       uint32
	PrevBlockReferenceTime primitives.TimestampSeconds
}

func (x *RequestCommitteeInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,RandomSeed:%s,MaxCommitteeSize:%s,PrevBlockReferenceTime:%s,}", x.StringCurrentBlockHeight(), x.StringRandomSeed(), x.StringMaxCommitteeSize(), x.StringPrevBlockReferenceTime())
}

func (x *RequestCommitteeInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
	return
}

func (x *RequestCommitteeInput) StringRandomSeed() (res string) {
	res = fmt.Sprintf("%v", x.RandomSeed)
	return
}

func (x *RequestCommitteeInput) StringMaxCommitteeSize() (res string) {
	res = fmt.Sprintf("%v", x.MaxCommitteeSize)
	return
}

func (x *RequestCommitteeInput) StringPrevBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockReferenceTime)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message RequestCommitteeOutput (non serializable)

type RequestCommitteeOutput struct {
	NodeAddresses            []primitives.NodeAddress
	NodeRandomSeedPublicKeys []primitives.Bls1PublicKey
	Weights                  []primitives.Weight
}

func (x *RequestCommitteeOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{NodeAddresses:%s,NodeRandomSeedPublicKeys:%s,Weights:%s,}", x.StringNodeAddresses(), x.StringNodeRandomSeedPublicKeys(), x.StringWeights())
}

func (x *RequestCommitteeOutput) StringNodeAddresses() (res string) {
	res = "["
	for _, v := range x.NodeAddresses {
		res += fmt.Sprintf("%s", v) + ","
	}
	res += "]"
	return
}

func (x *RequestCommitteeOutput) StringNodeRandomSeedPublicKeys() (res string) {
	res = "["
	for _, v := range x.NodeRandomSeedPublicKeys {
		res += fmt.Sprintf("%s", v) + ","
	}
	res += "]"
	return
}

func (x *RequestCommitteeOutput) StringWeights() (res string) {
	res = "["
	for _, v := range x.Weights {
		res += fmt.Sprintf("%s", v) + ","
	}
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message RequestBlockProofCommitteeInput (non serializable)

type RequestBlockProofCommitteeInput struct {
	CurrentBlockHeight     primitives.BlockHeight
	PrevBlockReferenceTime primitives.TimestampSeconds
}

func (x *RequestBlockProofCommitteeInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{CurrentBlockHeight:%s,PrevBlockReferenceTime:%s,}", x.StringCurrentBlockHeight(), x.StringPrevBlockReferenceTime())
}

func (x *RequestBlockProofCommitteeInput) StringCurrentBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.CurrentBlockHeight)
	return
}

func (x *RequestBlockProofCommitteeInput) StringPrevBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockReferenceTime)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message RequestBlockProofCommitteeOutput (non serializable)

type RequestBlockProofCommitteeOutput struct {
	NodeAddresses []primitives.NodeAddress
	Weights       []primitives.Weight
}

func (x *RequestBlockProofCommitteeOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{NodeAddresses:%s,Weights:%s,}", x.StringNodeAddresses(), x.StringWeights())
}

func (x *RequestBlockProofCommitteeOutput) StringNodeAddresses() (res string) {
	res = "["
	for _, v := range x.NodeAddresses {
		res += fmt.Sprintf("%s", v) + ","
	}
	res += "]"
	return
}

func (x *RequestBlockProofCommitteeOutput) StringWeights() (res string) {
	res = "["
	for _, v := range x.Weights {
		res += fmt.Sprintf("%s", v) + ","
	}
	res += "]"
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateBlockCommitteeInput (non serializable)

type ValidateBlockCommitteeInput struct {
	BlockHeight            primitives.BlockHeight
	PrevBlockReferenceTime primitives.TimestampSeconds
}

func (x *ValidateBlockCommitteeInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,PrevBlockReferenceTime:%s,}", x.StringBlockHeight(), x.StringPrevBlockReferenceTime())
}

func (x *ValidateBlockCommitteeInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *ValidateBlockCommitteeInput) StringPrevBlockReferenceTime() (res string) {
	res = fmt.Sprintf("%s", x.PrevBlockReferenceTime)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ValidateBlockCommitteeOutput (non serializable)

type ValidateBlockCommitteeOutput struct {
}

func (x *ValidateBlockCommitteeOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{}")
}

/////////////////////////////////////////////////////////////////////////////
// enums
