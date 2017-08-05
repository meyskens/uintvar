package uintvar

import (
	"encoding/hex"
	"testing"
)

func Test_hasContinueByte(t *testing.T) {
	type args struct {
		in byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{
				in: 0x88,
			},
			want: true,
		},
		{
			name: "true",
			args: args{
				in: 0x08,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasContinueByte(tt.args.in); got != tt.want {
				t.Errorf("hasContinueByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	bytes, _ := hex.DecodeString("850c")
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		args    args
		wantU   uint64
		wantN   int
		wantErr bool
	}{
		{
			name: "test 2 bytes",
			args: args{
				in: bytes,
			},
			wantU:   uint64(652),
			wantN:   2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotU, gotN, err := Parse(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotU != tt.wantU {
				t.Errorf("Parse() gotU = %v, want %v", gotU, tt.wantU)
			}
			if gotN != tt.wantN {
				t.Errorf("Parse() gotN = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_prefixWithZeros(t *testing.T) {
	type args struct {
		s string
		l int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test add",
			args: args{
				s: "test",
				l: 8,
			},
			want: "0000test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixWithZeros(tt.args.s, tt.args.l); got != tt.want {
				t.Errorf("prefixWithZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
