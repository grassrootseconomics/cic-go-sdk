package chain

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3"
)

// KitabuChainId represents the mainnet chain ID and Network ID.
const (
	KitabuChainId = 6060
)

// Provider returns an RPC client and a legacy signer compatible with the Kitabu Chain.
type Provider struct {
	EthClient *w3.Client
	Signer    types.Signer
}

// NewProvider returns a new RPC provider given an RPC endpoint.
func NewProvider(rpcEndpoint string) (*Provider, error) {
	ethClient, err := w3.Dial(rpcEndpoint)
	if err != nil {
		return nil, err
	}

	return &Provider{
		EthClient: ethClient,
		Signer:    types.NewEIP155Signer(big.NewInt(KitabuChainId)),
	}, nil
}

// Close closes the RPC connection and cancels any in-flight requests.
func (p *Provider) Close() error {
	if err := p.EthClient.Close(); err != nil {
		return err
	}

	return nil
}
