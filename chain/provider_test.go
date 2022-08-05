package chain

import (
	"testing"
)

func TestNewProvider(t *testing.T) {
	type args struct {
		rpcEndpoint string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid endpoint",
			args: args{
				rpcEndpoint: "https://rpc.sarafu.network",
			},
			wantErr: false,
		},
		{
			name: "Invalid endpoint",
			args: args{
				rpcEndpoint: "htt://broken",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewProvider(tt.args.rpcEndpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProvider_Close(t *testing.T) {
	type args struct {
		rpcEndpoint string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid endpoint Close()",
			args: args{
				rpcEndpoint: "https://rpc.sarafu.network",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := NewProvider(tt.args.rpcEndpoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err := p.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
