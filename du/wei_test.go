package du

import (
	"github.com/hootuu/utils/errors"
	"math/big"
	"reflect"
	"testing"
)

func TestWeiVerify(t *testing.T) {
	type args struct {
		wei Wei
	}
	tests := []struct {
		name string
		args args
		want *errors.Error
	}{
		{
			name: ">18",
			args: args{
				wei: 19,
			},
			want: errors.Verify(ErrWeiInvalid),
		},
		{
			name: "=0",
			args: args{
				wei: 0,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WeiVerify(tt.args.wei); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WeiVerify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWei_Get(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		w    Wei
		args args
		want int64
	}{
		{
			name: "0:1",
			w:    0,
			args: args{f: 1},
			want: 1,
		},
		{
			name: "1:1",
			w:    1,
			args: args{f: 1},
			want: 10,
		},
		{
			name: "1:0",
			w:    1,
			args: args{f: 0},
			want: 0,
		},
		{
			name: "8:0.1",
			w:    8,
			args: args{f: 0.1},
			want: 10000000,
		},
		{
			name: "8:9.1",
			w:    8,
			args: args{f: 9.1},
			want: 910000000,
		},
		{
			name: "8:9999999.1",
			w:    8,
			args: args{f: 9999999.1},
			want: 999999910000000,
		},
		{
			name: "8:9999999.11111111",
			w:    8,
			args: args{f: 9999999.11111111},
			want: 999999911111111,
		},
		{
			name: "8:9999999.11111111222",
			w:    8,
			args: args{f: 9999999.11111111222},
			want: 999999911111111,
		},
		{
			name: "8:0.111111118888888",
			w:    8,
			args: args{f: 0.111111118888888},
			want: 11111111,
		},
		{
			name: "8:0.1111111123456",
			w:    8,
			args: args{f: 0.1111111123456},
			want: 11111111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Get(tt.args.f); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWei_Pow(t *testing.T) {
	tests := []struct {
		name string
		w    Wei
		want *big.Int
	}{
		{
			name: "0",
			w:    0,
			want: big.NewInt(1),
		},
		{
			name: "9",
			w:    9,
			want: big.NewInt(1000000000),
		},
		{
			name: "18",
			w:    18,
			want: big.NewInt(1000000000000000000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.Pow(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWei_PowF(t *testing.T) {
	tests := []struct {
		name string
		w    Wei
		want *big.Float
	}{
		{
			name: "0",
			w:    0,
			want: big.NewFloat(1),
		},
		{
			name: "9",
			w:    9,
			want: big.NewFloat(1000000000),
		},
		{
			name: "18",
			w:    18,
			want: big.NewFloat(1000000000000000000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.PowF(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PowF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWei_PowInt64(t *testing.T) {
	tests := []struct {
		name string
		w    Wei
		want int64
	}{
		{
			name: "0",
			w:    0,
			want: 1,
		},
		{
			name: "9",
			w:    9,
			want: 1000000000,
		},
		{
			name: "18",
			w:    18,
			want: 1000000000000000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.PowInt64(); got != tt.want {
				t.Errorf("PowInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWei_S(t *testing.T) {
	tests := []struct {
		name string
		w    Wei
		want string
	}{
		{
			name: "0",
			w:    0,
			want: "0",
		},
		{
			name: "18",
			w:    18,
			want: "18",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.w.S(); got != tt.want {
				t.Errorf("S() = %v, want %v", got, tt.want)
			}
		})
	}
}
