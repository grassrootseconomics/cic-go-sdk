package chain

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3/module/eth"
)

// TransactionData represents a prepared tx data ready to be signed
type TransactionData struct {
	To        common.Address
	InputData []byte
	GasLimit  uint64
	Nonce     uint64
}

// NetworkNonce returns the nonce for the given address
func (p *Provider) NetworkNonce(ctx context.Context, address common.Address) (uint64, error) {
	var nonce uint64

	err := p.EthClient.CallCtx(
		ctx,
		eth.Nonce(address, nil).Returns(&nonce),
	)
	if err != nil {
		return 0, err
	}

	return nonce, nil
}

// BuildKitabuTx returns a signed transaction compatible with Kitabu chain, it is specific to contrtact execution
func (p *Provider) BuildKitabuTx(privateKey *ecdsa.PrivateKey, txData TransactionData) (*types.Transaction, error) {
	tx, err := types.SignNewTx(privateKey, p.Signer, &types.LegacyTx{
		To:       &txData.To,
		Nonce:    txData.Nonce,
		Data:     txData.InputData,
		Gas:      txData.GasLimit,
		GasPrice: big.NewInt(1),
	})
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// BuildGasTransferTx returns a signed transaction compatible with Kitabu chain, it is specific to gas transfers
func (p *Provider) BuildGasTransferTx(privateKey *ecdsa.PrivateKey, txData TransactionData, value *big.Int) (*types.Transaction, error) {
	tx, err := types.SignNewTx(privateKey, p.Signer, &types.LegacyTx{
		Value:    value,
		To:       &txData.To,
		Nonce:    txData.Nonce,
		Gas:      21000,
		GasPrice: big.NewInt(1),
	})
	if err != nil {
		return nil, err
	}

	return tx, nil
}
