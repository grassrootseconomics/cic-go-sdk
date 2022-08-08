package chain

import "testing"

func TestChecksumToSarafuAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test Checksum2Sarafu address conversion",
			args: args{
				address: "0x02b0DF387A3A68AA3134668752dd82bE70B7dE1C",
			},
			want:    "02b0df387a3a68aa3134668752dd82be70b7de1c",
			wantErr: false,
		},
		{
			name: "Test Bad address Checksum2Sarafu address conversion",
			args: args{
				address: "0x02b0DF387A3A68AA3134668752dd82bE70B7dE1C567890",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ChecksumToSarafuAddress(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChecksumToSarafuAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ChecksumToSarafuAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSarafuAddressToChecksum(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Sarafu2Checksum address conversion",
			args: args{
				address: "02b0df387a3a68aa3134668752dd82be70b7de1c",
			},
			want: "0x02b0DF387A3A68AA3134668752dd82bE70B7dE1C",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SarafuAddressToChecksum(tt.args.address); got != tt.want {
				t.Errorf("SarafuAddressToChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
