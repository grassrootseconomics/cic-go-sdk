package chain

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3"
	"github.com/stretchr/testify/assert"
)

func TestProvider_TokensBalance(t *testing.T) {
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
		ctx             context.Context
		ownerAddress    common.Address
		tokensAddresses []common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "A random members (min dust available) balances",
			fields: fields{
				EthClient: p.EthClient,
				Signer:    p.Signer,
			},
			args: args{
				ctx:          context.Background(),
				ownerAddress: w3.A("0x4e956b5De3c33566c596754B4fa0ABd9F2789578"),
				tokensAddresses: []common.Address{
					w3.A("0x02b0DF387A3A68AA3134668752dd82bE70B7dE1C"),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provider{
				EthClient: tt.fields.EthClient,
				Signer:    tt.fields.Signer,
			}
			got, err := p.TokensBalance(tt.args.ctx, tt.args.ownerAddress, tt.args.tokensAddresses)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.TokensBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !assert.Greater(t, got[0].Int64(), int64(0)) {
				t.Errorf("Provider.TokensBalance() = %v, want greater than 0", got[0])
			}
		})
	}
}
