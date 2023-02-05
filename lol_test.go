package lol_test

import (
	"errors"
	"io"
	"math/big"
	"nrvn.cc/go/lol"
	"testing"
)

type nilReader struct{}

func (reader nilReader) Read(p []byte) (n int, err error) {
	return len(p), nil
}

type errReader struct{}

func (reader errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("bad reader")
}

//func TestWoot(t *testing.T) {
//	type args struct {
//		rand io.Reader
//	}
//	tests := []struct {
//		want    *big.Int
//		args    args
//		name    string
//		wantErr bool
//	}{
//		{
//			name: "nil",
//			args: args{nilReader{}},
//			want: big.NewInt(1),
//		},
//		{
//			name:    "err",
//			args:    args{errReader{}},
//			wantErr: true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := lol.Woot(tt.args.rand)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("Woot() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got.Cmp(tt.want) != 0 {
//				t.Errorf("Woot() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestWoot(t *testing.T) {
	type args struct {
		rand io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		want1   bool
		wantErr bool
	}{
		{
			name: "nil",
			args: args{nilReader{}},
			want: big.NewInt(0),
		},
		{
			name:    "err",
			args:    args{errReader{}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := lol.Woot(tt.args.rand)
			if (err != nil) != tt.wantErr {
				t.Errorf("Woot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Cmp(tt.want) != 0 {
				t.Errorf("Woot() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Woot() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
