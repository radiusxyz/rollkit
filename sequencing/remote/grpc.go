package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"google.golang.org/grpc"

	"github.com/rollkit/rollkit/log"
	"github.com/rollkit/rollkit/proto/pb/sequencer"
	"github.com/rollkit/rollkit/sequencing"
)

// SequencingLayerClient is a generic client that proxies all DA requests via gRPC.
type SequencingLayerClient struct {
	config Config

	conn   *grpc.ClientConn
	client sequencer.SequencingServiceClient

	logger log.Logger
}

// Config contains configuration options for SeqeuncingLayerClient.
type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// DefaultConfig defines default values for SeqeuncingLayerClient configuration.
var DefaultConfig = Config{
	Host: "127.0.0.1",
	Port: 7000,
}

var _ sequencing.SequencingLayerClient = &SequencingLayerClient{}

// Init sets the configuration options.
func (sequencerLayerClient *SequencingLayerClient) Init(config []byte, logger log.Logger) error {
	sequencerLayerClient.logger = logger

	if len(config) == 0 {
		sequencerLayerClient.config = DefaultConfig
		return nil
	}

	return json.Unmarshal(config, &sequencerLayerClient.config)
}

// Start creates connection to sequencer server and instantiates gRPC client.
func (sequencerLayerClient *SequencingLayerClient) Start() error {
	sequencerLayerClient.logger.Info("starting GRPC sequencer", "host", sequencerLayerClient.config.Host, "port", sequencerLayerClient.config.Port)
	var err error
	var opts []grpc.DialOption
	
	opts = append(opts, grpc.WithInsecure())
	sequencerLayerClient.conn, err = grpc.Dial(sequencerLayerClient.config.Host + ":" + strconv.Itoa(sequencerLayerClient.config.Port), opts...)
	if err != nil {
		return err
	}

	fmt.Println("stompesi - grpc Start run")

	sequencerLayerClient.client = sequencer.NewSequencingServiceClient(sequencerLayerClient.conn)
	return nil
}

// Stop closes connection to sequencer server.
func (sequencerLayerClient *SequencingLayerClient) Stop() error {
	sequencerLayerClient.logger.Info("stopoing GRPC sequencer")
	return sequencerLayerClient.conn.Close()
}

// GetTxOrder queries Sequencing layer to sequencer server for checking tx order.
func (sequencerLayerClient *SequencingLayerClient) GetTxOrderList(ctx context.Context, rollupId string, height uint64, signature []byte) sequencing.ResultGetTxOrderList {
	fmt.Println("stompesi - GetTxOrderList", "rollupId", rollupId, "height", height)
	
	resp, err := sequencerLayerClient.client.GetTxOrderList(ctx, &sequencer.RequestGetTxOrderList{
		RollupId: rollupId, 
		Height: height, 
		Signature: signature,
	})

	if err != nil {
		return sequencing.ResultGetTxOrderList{
			BaseResult: sequencing.BaseResult{
				Code: sequencing.StatusError, 
				Message: err.Error(),
			}}
	}
	return sequencing.ResultGetTxOrderList{
		BaseResult:    sequencing.BaseResult{
			Code: sequencing.StatusCode(resp.Result.Code),
			Message:  resp.Result.Message,
		},
		TxOrderList: 	 resp.TxOrderList,
	}
}