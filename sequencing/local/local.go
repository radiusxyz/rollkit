package local

import (
	"context"
	"encoding/json"

	"github.com/rollkit/rollkit/log"
	"github.com/rollkit/rollkit/sequencing"
)

// Config contains configuration options for SeqeuncingLayerClient.
type Config struct {}

var DefaultConfig = Config{}

// SequencingLayerClient is intended only for usage in tests.
type SequencingLayerClient struct {
	logger   log.Logger
	config   Config
}

var _ sequencing.SequencingLayerClient = &SequencingLayerClient{}

// Init is called once to allow Sequencing client to read configuration and initialize resources.
func (sequencerLayerClient *SequencingLayerClient) Init(config []byte, logger log.Logger) error {
	sequencerLayerClient.logger = logger
	
	if len(config) == 0 {
		sequencerLayerClient.config = DefaultConfig
		return nil
	}
	
	return json.Unmarshal(config, &sequencerLayerClient.config)
}

// Start implements SequencingLayerClient interface.
func (sequencerLayerClient *SequencingLayerClient) Start() error {
	sequencerLayerClient.logger.Debug("Local Sequencing Layer Client starting")
	return nil
}

// Stop implements SequencingLayerClient interface.
func (sequencerLayerClient *SequencingLayerClient) Stop() error {
	sequencerLayerClient.logger.Debug("Local Sequencing Layer Client stopped")
	return nil
}

// GetTxOrder queries Sequencing layer to check tx order
func (sequencerLayerClient *SequencingLayerClient) GetTxOrder(ctx context.Context, rollupId string, height uint64, signature []byte) sequencing.ResultGetTxOrderList {
	return sequencing.ResultGetTxOrderList{
		BaseResult: sequencing.BaseResult{Code: sequencing.StatusSuccess}, 
		TxOrderList: nil,
	}
}