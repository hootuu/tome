package du

import (
	"fmt"
	"github.com/hootuu/utils/errors"
	"reflect"
	"testing"
)

func TestDu_Add(t *testing.T) {
	type fields struct {
		self Du
	}
	type args struct {
		other Du
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Du
		wantErr *errors.Error
	}{
		{
			name: "1.1+1.1",
			fields: struct {
				self Du
			}{self: FromFloat64(10, 1.1)},
			args:    struct{ other Du }{other: FromFloat64(10, 1.1)},
			want:    FromFloat64(10, 2.2),
			wantErr: nil,
		},
		{
			name: "1.9+1.2",
			fields: struct {
				self Du
			}{self: FromFloat64(10, 1.2)},
			args:    struct{ other Du }{other: FromFloat64(10, 1.9)},
			want:    FromFloat64(10, 3.1),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.self
			got, got1 := d.Add(tt.args.other)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantErr) {
				t.Errorf("Add() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}

func TestDu_Divide(t *testing.T) {
	type fields struct {
		self Du
	}
	type args struct {
		div int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Du
		wantErr *errors.Error
	}{
		{
			name: "1.1/1",
			fields: struct {
				self Du
			}{self: FromFloat64(10, 1.1)},
			args:    struct{ div int64 }{div: 1},
			want:    FromFloat64(10, 1.1),
			wantErr: nil,
		},
		{
			name: "1.9/3",
			fields: struct {
				self Du
			}{self: FromFloat64(10, 1.9)},
			args:    struct{ div int64 }{div: 3},
			want:    FromFloat64(10, 0.6333333333),
			wantErr: nil,
		},
		{
			name: "-1.9/3",
			fields: struct {
				self Du
			}{self: FromFloat64(10, -1.9)},
			args:    struct{ div int64 }{div: 3},
			want:    MustFromString("-10:0.6333333334"),
			wantErr: nil,
		},
		{
			name: "9.99/3",
			fields: struct {
				self Du
			}{self: FromFloat64(10, 9.99)},
			args:    struct{ div int64 }{div: 3},
			want:    FromFloat64(10, 9.99/3),
			wantErr: nil,
		},
		{
			name: "9.89/3",
			fields: struct {
				self Du
			}{self: FromFloat64(10, 9.89)},
			args:    struct{ div int64 }{div: 3},
			want:    MustFromString("+10:3.2966666666"),
			wantErr: nil,
		},
		{
			name: "-9.89/3",
			fields: struct {
				self Du
			}{self: FromFloat64(10, -9.89)},
			args:    struct{ div int64 }{div: 3},
			want:    FromFloat64(10, -9.89/3),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.self
			got, got1 := d.Divide(tt.args.div)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Divide() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantErr) {
				t.Errorf("Divide() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}

func TestDu_DivideF(t *testing.T) {
	type fields struct {
		self Du
	}
	type args struct {
		div float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Du
		wantErr *errors.Error
	}{
		{
			name: "1.1/1.1",
			fields: struct {
				self Du
			}{self: FromFloat64(10, 1.1)},
			args:    struct{ div float64 }{div: 1.1},
			want:    FromFloat64(10, 1.0),
			wantErr: nil,
		},
		{
			name: "1.9/3.2",
			fields: struct {
				self Du
			}{self: FromFloat64(10, 1.9)},
			args:    struct{ div float64 }{div: 3.2},
			want:    FromFloat64(10, 1.9/3.2),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.self
			got, got1 := d.DivideF(tt.args.div)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Divide() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantErr) {
				t.Errorf("Divide() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}

func TestDu_Multiply(t *testing.T) {
	type fields struct {
		self Du
	}
	type args struct {
		mul int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Du
		wantErr *errors.Error
	}{
		{
			name: "1.1x1",
			fields: struct {
				self Du
			}{self: FromFloat64(10, 1.1)},
			args:    struct{ mul int64 }{mul: 1},
			want:    FromFloat64(10, 1.1),
			wantErr: nil,
		},
		{
			name: "1.9x9",
			fields: struct {
				self Du
			}{self: FromFloat64(10, 1.9)},
			args:    struct{ mul int64 }{mul: 9},
			want:    FromFloat64(10, 1.9*9),
			wantErr: nil,
		},
		{
			name: "99999x999999.2",
			fields: struct {
				self Du
			}{self: FromFloat64(10, 999999.2)},
			args: struct{ mul int64 }{mul: 99999},
			want: Du{
				W: 10,
				L: 99998920000,
				R: 8000000000,
			},
			wantErr: nil,
		},
		{
			name: "-1.9x9",
			fields: struct {
				self Du
			}{self: FromFloat64(10, -1.9)},
			args:    struct{ mul int64 }{mul: 9},
			want:    FromFloat64(10, -1.9*9),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.self
			got, got1 := d.Multiply(tt.args.mul)
			fmt.Println("multi.got ", got.String())
			fmt.Println("multi.want ", tt.want.String())
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiply() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantErr) {
				t.Errorf("Multiply() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}

func TestDu_MultiplyF(t *testing.T) {
	type fields struct {
		self Du
	}
	type args struct {
		mul float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Du
		wantErr *errors.Error
	}{
		{
			name: "1.1x1.1",
			fields: struct {
				self Du
			}{self: FromFloat64(10, 1.1)},
			args:    struct{ mul float64 }{mul: 1.1},
			want:    FromFloat64(10, 1.21),
			wantErr: nil,
		},
		{
			name: "1.9x9.9",
			fields: struct {
				self Du
			}{self: FromFloat64(10, 1.9)},
			args:    struct{ mul float64 }{mul: 9.9},
			want:    FromFloat64(10, 1.9*9.9),
			wantErr: nil,
		},
		{
			name: "99999.9x999999.2",
			fields: struct {
				self Du
			}{self: FromFloat64(10, 999999.2)},
			args: struct{ mul float64 }{mul: 99999.9},
			want: Du{
				W: 10,
				L: 99999820000,
				R: 800000000,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.self
			got, got1 := d.MultiplyF(tt.args.mul)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiplyF() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantErr) {
				t.Errorf("Multiply() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}

func TestDu_Subtract(t *testing.T) {
	type fields struct {
		self Du
	}
	type args struct {
		other Du
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Du
		wantErr *errors.Error
	}{
		{
			name: "1.9-1.8",
			fields: fields{
				self: FromFloat64(10, 1.9),
			},
			args: args{
				other: FromFloat64(10, 1.8),
			},
			want:    FromFloat64(10, 1.9-1.8),
			wantErr: nil,
		},
		{
			name: "1.9-2.5",
			fields: fields{
				self: FromFloat64(10, 1.9),
			},
			args: args{
				other: FromFloat64(10, 2.5),
			},
			want:    FromFloat64(10, 1.9-2.5),
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.self
			got, got1 := d.Subtract(tt.args.other)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantErr) {
				t.Errorf("Subtract() got1 = %v, want %v", got1, tt.wantErr)
			}
		})
	}
}

func TestFromByte(t *testing.T) {
	type args struct {
		bArr []byte
	}
	tests := []struct {
		name  string
		args  args
		want  Du
		want1 *errors.Error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FromByte(tt.args.bArr)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromByte() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FromByte() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFromFloat64(t *testing.T) {
	type args struct {
		w Wei
		f float64
	}
	tests := []struct {
		name string
		args args
		want Du
	}{
		{
			name: "0.099809",
			args: args{
				w: 8,
				f: 0.099809,
			},
			want: MustFromString("+8:0.099809"),
		},
		{
			name: "-0.099809",
			args: args{
				w: 8,
				f: -0.099809,
			},
			want: MustFromString("-8:0.099809"),
		},
		{
			name: "998.0099809",
			args: args{
				w: 8,
				f: 998.0099809,
			},
			want: MustFromString("+8:998.0099809"),
		},
		{
			name: "-998.099809",
			args: args{
				w: 8,
				f: -998.099809,
			},
			want: MustFromString("-8:998.099809"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromFloat64(tt.args.w, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromInt64(t *testing.T) {
	type args struct {
		w Wei
		d int64
	}
	tests := []struct {
		name string
		args args
		want Du
	}{
		{
			name: "9998",
			args: args{
				w: 18,
				d: 9998,
			},
			want: Du{
				W: 18,
				N: false,
				L: 9998,
				R: 0,
			},
		},
		{
			name: "-9998",
			args: args{
				w: 18,
				d: -9998,
			},
			want: Du{
				W: 18,
				N: true,
				L: 9998,
				R: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromInt64(tt.args.w, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromLR(t *testing.T) {
	type args struct {
		w   Wei
		neg bool
		iL  int64
		iR  int64
	}
	tests := []struct {
		name string
		args args
		want Du
	}{
		{
			name: "1.01111",
			args: args{
				w:   8,
				neg: false,
				iL:  1,
				iR:  Wei(8).Get(0.01111),
			},
			want: Du{
				W: 8,
				N: false,
				L: 1,
				R: Wei(8).Get(0.01111),
			},
		},
		{
			name: "-1.0001111",
			args: args{
				w:   8,
				neg: true,
				iL:  1,
				iR:  Wei(8).Get(0.0001111),
			},
			want: Du{
				W: 8,
				N: true,
				L: 1,
				R: Wei(8).Get(0.0001111),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromLR(tt.args.w, tt.args.neg, tt.args.iL, tt.args.iR); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name  string
		args  args
		want  Du
		want1 *errors.Error
	}{
		{
			name: "+18:0.36",
			args: struct{ str string }{str: "+18:0.36"},
			want: Du{
				W: 18,
				N: false,
				L: 0,
				R: Wei(18).Get(0.36),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FromString(tt.args.str)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromString() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FromString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIsPositiveOrDecimalNumber(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid number",
			args: struct{ str string }{str: "abc"},
			want: false,
		},
		{
			name: "invalid decimal",
			args: struct{ str string }{str: "abc.def"},
			want: false,
		},
		{
			name: "valid number",
			args: struct{ str string }{str: "9989"},
			want: true,
		},
		{
			name: "valid number: -number",
			args: struct{ str string }{str: "-9989"},
			want: true,
		},
		{
			name: "valid dec",
			args: struct{ str string }{str: "0.98"},
			want: true,
		},
		{
			name: "valid dec: - des",
			args: struct{ str string }{str: "-0.98"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPositiveOrDecimalNumber(tt.args.str); got != tt.want {
				t.Errorf("IsPositiveOrDecimalNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDu_Bytes(t *testing.T) {
	type fields struct {
		self Du
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name:   "1111",
			fields: struct{ self Du }{self: FromInt64(8, 9998)},
			want:   []byte("+8:9998"),
		},
		{
			name:   "-1111",
			fields: struct{ self Du }{self: FromInt64(8, -9998)},
			want:   []byte("-8:9998"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.self
			if got := d.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDu_String(t *testing.T) {
	type fields struct {
		self Du
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "+8:1.01",
			fields: fields{
				self: FromFloat64(8, 1.01),
			},
			want: "+8:1.01",
		},
		{
			name: "+18:111.0199",
			fields: fields{
				self: Du{
					W: 18,
					N: false,
					L: 111,
					R: 19900000000000000,
				},
			},
			want: "+18:111.0199",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.self
			if got := d.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDu_View(t *testing.T) {
	type fields struct {
		self Du
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "1.01",
			fields: struct{ self Du }{self: Du{
				W: 18,
				N: false,
				L: 1,
				R: 1 * Wei(16).PowInt64(),
			}},
			want: "1.01",
		},
		{
			name: "-91.000198",
			fields: struct{ self Du }{self: Du{
				W: 18,
				N: true,
				L: 91,
				R: 198 * Wei(12).PowInt64(),
			}},
			want: "-91.000198",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.self
			if got := d.View(); got != tt.want {
				t.Errorf("View() = %v, want %v", got, tt.want)
			}
		})
	}
}
