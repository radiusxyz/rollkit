package registry

import (
	"github.com/rollkit/rollkit/sequencing"

	local_sequencing "github.com/rollkit/rollkit/sequencing/local"
	remote_sequencing "github.com/rollkit/rollkit/sequencing/remote"
)

// this is a central registry for all Sequencing Layer Clients
var sequencingLayerClients = map[string]func() sequencing.SequencingLayerClient{
	"local":     func() sequencing.SequencingLayerClient { return &local_sequencing.SequencingLayerClient{} },
	"remote":     func() sequencing.SequencingLayerClient { return &remote_sequencing.SequencingLayerClient{} },
}

// TODO: stompesi
// GetClient returns client identified by name.
func GetSeqeucningLayerClient(name string) sequencing.SequencingLayerClient {
	f, ok := sequencingLayerClients[name]
	if !ok {
		return nil
	}
	return f()
}

// Register adds a Sequencing Layer Client to registry.
//
// If name was previously used in the registry, error is returned.
func RegisterSequencer(name string, constructor func() sequencing.SequencingLayerClient) error {
	if _, found := sequencingLayerClients[name]; !found {
		sequencingLayerClients[name] = constructor
		return nil
	}
	return &ErrAlreadyRegistered{name: name}
}

// RegisteredClients returns names of all Sequencer clients in registry.
func RegisteredSequencerClients() []string {
	registered := make([]string, 0, len(sequencingLayerClients))
	for name := range sequencingLayerClients {
		registered = append(registered, name)
	}
	return registered
}
