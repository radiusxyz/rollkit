package sequencing

import (
	"context"

	"github.com/rollkit/rollkit/log"
)

type StatusCode uint64

// Sequencing Layer return codes.
const (
	StatusUnknown StatusCode = iota
	StatusSuccess
	StatusTimeout
	StatusError
)

// BaseResult contains basic information returned by Sequencing layer.
type BaseResult struct {
	Code StatusCode
	
	// Message may contain Sequencing layer specific information (like detailed error message, etc)
	Message string
}

// ResultGetTxOrderList contains batch of transactions hashs returned from Sequencing layer client.
type ResultGetTxOrderList struct {
	BaseResult BaseResult
	TxOrderList []string
}

type SequencingLayerClient interface {
	Init(config []byte, logger log.Logger) error
	Start() error
	Stop() error
	GetTxOrderList(ctx context.Context, rollupId string, height uint64, signature []byte) ResultGetTxOrderList
}