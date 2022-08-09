# cic-go-sdk
[![Go](https://github.com/grassrootseconomics/cic-go-sdk/actions/workflows/go.yml/badge.svg)](https://github.com/grassrootseconomics/cic-go-sdk/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/grassrootseconomics/cic-go-sdk/badge.svg?branch=master)](https://coveralls.io/github/grassrootseconomics/cic-go-sdk?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/grassrootseconomics/cic-go-sdk)](https://goreportcard.com/report/github.com/grassrootseconomics/cic-go-sdk)
[![DeepSource](https://deepsource.io/gh/grassrootseconomics/cic-go-sdk.svg/?label=active+issues&token=FZOAG8G8bwcn2PAWNfR5k18a)](https://deepsource.io/gh/grassrootseconomics/cic-go-sdk/?ref=repository-badge)

Go SDK to interact with GE CIC components and build custom solutions around them

[`chain`](https://pkg.go.dev/github.com/grassrootseconomics/cic-go-sdk/chain "API documentation") package
---------------------------------------------------------------------------------------------------------

The `chain` package provides some helpful chain utilities alongside a client and signer that can be used directly with [`lmittmann/w3`](https://github.com/lmittmann/w3).

- Fetch ERC20 contract metadata including demurrage specific information.
- Tokens balance that allows you to fetch a single user's token balances from multiple ERC20 contracts in a single call.
- Address conversion utilities.
- Transaction building utilities compatible with [Kitabu Chain](https://github.com/grassrootseconomics/kitabu-chain).

## License

AGPL-3.0-or-later
