package chain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

var (
	bulkBalanceContract = w3.A("0xb9e215B789e9Ec6643Ba4ff7b98EA219F38c6fE5")
)

// TokensBalance returns balances of an address from multiple token addresses batched into a single on-chain evm call
func (p *Provider) TokensBalance(ctx context.Context, ownerAddress common.Address, tokensAddresses []common.Address) ([]*big.Int, error) {
	var balancesResults []*big.Int

	err := p.EthClient.CallCtx(
		ctx,
		eth.CallFunc(w3.MustNewFunc("tokensBalance(address owner, address[] contracts)", "uint256[]"), bulkBalanceContract, ownerAddress, tokensAddresses).Returns(&balancesResults),
	)
	if err != nil {
		return nil, err
	}

	return balancesResults, nil
}
