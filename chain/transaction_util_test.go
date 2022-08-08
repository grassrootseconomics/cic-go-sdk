package chain

import (
	"context"
	"crypto/ecdsa"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3"
)

func TestProvider_NetworkNonce(t *testing.T) {
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
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
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
			if got != tt.want {
				t.Errorf("Provider.NetworkNonce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProvider_BuildKitabuTx(t *testing.T) {
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
		// TODO: Add test cases.
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Provider.BuildKitabuTx() = %v, want %v", got, tt.want)
			}
		})
	}
}
