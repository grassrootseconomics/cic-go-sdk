package chain

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3"
	"github.com/stretchr/testify/assert"
)

var rpcEndpoint = "https://rpc.sarafu.network"

func TestProvider_ERC20TokenInfo(t *testing.T) {
	p, err := NewProvider(rpcEndpoint)
	if err != nil {
		t.Fatal("Failed to connect to provider")
		return
	}

	type fields struct {
		EthClient *w3.Client
		Signer    types.Signer
	}
	type args struct {
		ctx          context.Context
		tokenAddress common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ERC20TokenMetadata
		wantErr bool
	}{
		{
			name: "Test SRF",
			fields: fields{
				EthClient: p.EthClient,
				Signer:    p.Signer,
			},
			args: args{
				ctx:          context.Background(),
				tokenAddress: w3.A("0x02b0DF387A3A68AA3134668752dd82bE70B7dE1C"),
			},
			want: ERC20TokenMetadata{
				Symbol:         "SRF",
				DemurrageToken: true,
				PeriodDuration: *big.NewInt(604800),
			},
			wantErr: false,
		},
		{
			name: "Test Non ERC20",
			fields: fields{
				EthClient: p.EthClient,
				Signer:    p.Signer,
			},
			args: args{
				ctx:          context.Background(),
				tokenAddress: w3.A("0x000000000000000000000000000000000000000A"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			p := &Provider{
				EthClient: tt.fields.EthClient,
				Signer:    tt.fields.Signer,
			}
			got, err := p.ERC20TokenInfo(tt.args.ctx, tt.args.tokenAddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.ERC20TokenInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if assert.NotNil(t, got) {
				assert.Equal(t, tt.want.DemurrageToken, got.DemurrageToken)
				assert.Equal(t, tt.want.Symbol, got.Symbol)
				assert.Equal(t, tt.want.PeriodDuration, got.PeriodDuration)
			}
		})
	}
}
