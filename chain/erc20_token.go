package chain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

// ERC20TokenMetadata represents a tokens's metadata
type ERC20TokenMetadata struct {
	Name               string
	Symbol             string
	Decimals           big.Int
	TotalSupply        big.Int
	Owner              common.Address
	DemurrageToken     bool
	SinkAddress        common.Address
	DemurrageAmount    big.Int
	DemurrageTimestamp big.Int
	PeriodStart        big.Int
	PeriodDuration     big.Int
	TaxLevel           big.Int
	// minter
	// allowance
}

// ERC20TokenInfo returns all getter only token metadata that requires no args, also includes all demurrage token info if applicable
func (p *Provider) ERC20TokenInfo(ctx context.Context, tokenAddress common.Address) (ERC20TokenMetadata, error) {
	var (
		erc20TokenMetdata  ERC20TokenMetadata
		tokenName          string
		tokenSymbol        string
		tokenDecimals      big.Int
		totalSupply        big.Int
		owner              common.Address
		sinkAddress        common.Address
		demurrageAmount    big.Int
		demurrageTimestamp big.Int
		periodStart        big.Int
		periodDuration     big.Int
		taxLevel           big.Int
	)

	err := p.EthClient.CallCtx(
		ctx,
		eth.CallFunc(w3.MustNewFunc("name()", "string"), tokenAddress).Returns(&tokenName),
		eth.CallFunc(w3.MustNewFunc("symbol()", "string"), tokenAddress).Returns(&tokenSymbol),
		eth.CallFunc(w3.MustNewFunc("decimals()", "uint256"), tokenAddress).Returns(&tokenDecimals),
		eth.CallFunc(w3.MustNewFunc("totalSupply()", "uint256"), tokenAddress).Returns(&totalSupply),
		eth.CallFunc(w3.MustNewFunc("owner()", "address"), tokenAddress).Returns(&owner),
	)
	if err != nil {
		return erc20TokenMetdata, err
	}

	erc20TokenMetdata.Name = tokenName
	erc20TokenMetdata.Symbol = tokenSymbol
	erc20TokenMetdata.Decimals = tokenDecimals
	erc20TokenMetdata.TotalSupply = totalSupply

	err = p.EthClient.CallCtx(
		ctx,
		eth.CallFunc(w3.MustNewFunc("sinkAddress()", "address"), tokenAddress).Returns(&sinkAddress),
	)
	if err != nil {
		erc20TokenMetdata.DemurrageToken = false
		return erc20TokenMetdata, err
	}

	erc20TokenMetdata.DemurrageToken = true
	erc20TokenMetdata.SinkAddress = sinkAddress
	err = p.EthClient.CallCtx(
		ctx,
		eth.CallFunc(w3.MustNewFunc("demurrageAmount()", "uint128"), tokenAddress).Returns(&demurrageAmount),
		eth.CallFunc(w3.MustNewFunc("demurrageTimestamp()", "uint256"), tokenAddress).Returns(&demurrageTimestamp),
		eth.CallFunc(w3.MustNewFunc("periodStart()", "uint256"), tokenAddress).Returns(&periodStart),
		eth.CallFunc(w3.MustNewFunc("periodDuration()", "uint256"), tokenAddress).Returns(&periodDuration),
		eth.CallFunc(w3.MustNewFunc("taxLevel()", "uint256"), tokenAddress).Returns(&taxLevel),
	)
	if err != nil {
		return erc20TokenMetdata, err
	}

	erc20TokenMetdata.DemurrageAmount = demurrageAmount
	erc20TokenMetdata.DemurrageTimestamp = demurrageTimestamp
	erc20TokenMetdata.PeriodStart = periodStart
	erc20TokenMetdata.PeriodDuration = periodDuration
	erc20TokenMetdata.TaxLevel = taxLevel

	return erc20TokenMetdata, nil
}

// BalanceOf Returns the amount of tokens owned by account.
func (p *Provider) ERC20BalanceOf(ctx context.Context, tokenAddress common.Address, accountAddress common.Address) (big.Int, error) {
	var balance big.Int

	err := p.EthClient.CallCtx(
		ctx,
		eth.CallFunc(w3.MustNewFunc("balanceOf(address _account)", "uint256"), tokenAddress, accountAddress).Returns(&balance),
	)
	if err != nil {
		return big.Int{}, err
	}

	return balance, nil
}

// BaseBalanceOf Returns the amount of tokens owned by account, unmodified by demurrage.
func (p *Provider) ERC20BaseBalanceOf(ctx context.Context, tokenAddress common.Address, accountAddress common.Address) (big.Int, error) {
	var balance big.Int

	err := p.EthClient.CallCtx(
		ctx,
		eth.CallFunc(w3.MustNewFunc("baseBalanceOf(address _account)", "uint256"), tokenAddress, accountAddress).Returns(&balance),
	)
	if err != nil {
		return big.Int{}, err
	}

	return balance, nil
}

// ActualPeriod Returns the demurrage period of the current block number.
func (p *Provider) ERC20ActualPeriod(ctx context.Context, tokenAddress common.Address) (big.Int, error) {
	var demurragePeriod big.Int

	err := p.EthClient.CallCtx(
		ctx,
		eth.CallFunc(w3.MustNewFunc("actualPeriod()", "uint128"), tokenAddress).Returns(&demurragePeriod),
	)
	if err != nil {
		return big.Int{}, err
	}

	return demurragePeriod, nil
}
