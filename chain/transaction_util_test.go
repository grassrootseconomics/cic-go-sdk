package chain

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/lmittmann/w3"
	"github.com/stretchr/testify/assert"
)

func TestProvider_NetworkNonce(t *testing.T) {
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
		ctx     context.Context
		address common.Address
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Test Random Acc Nonce",
			fields: fields{
				EthClient: p.EthClient,
				Signer:    p.Signer,
			},
			args: args{
				ctx:     context.Background(),
				address: w3.A("0x4e956b5De3c33566c596754B4fa0ABd9F2789578"),
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
			got, err := p.NetworkNonce(tt.args.ctx, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.NetworkNonce() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !assert.Greater(t, got, uint64(0)) {
				t.Errorf("Provider.NetworkNonce() = %v, want greater than 0", got)
			}
		})
	}
}

func TestProvider_BuildKitabuTx(t *testing.T) {
	p, err := NewProvider(rpcEndpoint)
	if err != nil {
		t.Fatal("Failed to connect to provider")
		return
	}

	// Throwaway key from https://vanity-eth.tk/
	privateKey, err := crypto.HexToECDSA("2e2838a1f752e343e310d3536f8acbd9149bb39f05c9a26db0d58becc33e8d57")
	if err != nil {
		t.Fatal(err)
	}

	sampleFunc := w3.MustNewFunc("applyDemurrageLimited(uint256)", "bool")
	input, err := sampleFunc.EncodeArgs(w3.I("25000"))
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		EthClient *w3.Client
		Signer    types.Signer
	}
	type args struct {
		privateKey *ecdsa.PrivateKey
		txData     TransactionData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Transaction
		wantErr bool
	}{
		{
			name: "Test Sample Tx Build",
			fields: fields{
				EthClient: p.EthClient,
				Signer:    p.Signer,
			},
			args: args{
				privateKey: privateKey,
				txData: TransactionData{
					To:        w3.A("0xcE43fd0a4E869D4B2B2Cc558CCE927dB77CbD21C"),
					InputData: input,
					GasLimit:  uint64(16000000),
					Nonce:     uint64(1),
				},
			},
			want:    &types.Transaction{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provider{
				EthClient: tt.fields.EthClient,
				Signer:    tt.fields.Signer,
			}
			got, err := p.BuildKitabuTx(tt.args.privateKey, tt.args.txData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.BuildKitabuTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// TODO: Add more assert fields
			if assert.NotNil(t, got) {
				assert.Equal(t, big.NewInt(6060), got.ChainId())
				// Input construction testing against eth-encode py equivalent tx
				assert.Equal(t, w3.B("0xc0ab707700000000000000000000000000000000000000000000000000000000000061a8"), input)
			}
		})
	}
}
