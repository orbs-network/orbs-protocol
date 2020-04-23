// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package services

import (
	"github.com/orbs-network/go-mock"
	"github.com/orbs-network/orbs-spec/types/go/services/gossiptopics"
)

/////////////////////////////////////////////////////////////////////////////
// service Gossip

type MockGossip struct {
	mock.Mock
	gossiptopics.MockTransactionRelay
	gossiptopics.MockBlockSync
	gossiptopics.MockHeaderSync
	gossiptopics.MockLeanHelix
	gossiptopics.MockBenchmarkConsensus
}
