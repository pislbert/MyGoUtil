package mem

import (
	"MyGoUtil/pt"
	"reflect"
	"testing"
)

// 普通对比
func TestMemCmp(t *testing.T) {
	type args struct {
		src interface{}
		des interface{}
	}
	para1 := 666
	var p2 *int
	p2 = &para1
	type info struct {
		name string
		id   int
		p    *int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{
				src: 5,
				des: 6,
			},
			want: false,
		},
		{
			name: "normal dif",
			args: args{
				src: 5,
				des: float64(5),
			},
			want: false,
		},
		{
			name: "struct 1",
			args: args{
				src: info{
					"hello1",
					32,
					&para1,
				},
				des: info{
					"hello2",
					31,
					&para1,
				},
			},
			want: false,
		},
		{
			name: "struct 2",
			args: args{
				src: info{
					"hello1",
					32,
					&para1,
				},
				des: info{
					"hello1",
					32,
					&para1,
				},
			},
			want: true,
		},
		{
			name: "struct 2 ptr",
			args: args{
				src: info{
					"hello1",
					32,
					&para1,
				},
				des: info{
					"hello1",
					32,
					p2,
				},
			},
			want: true,
		},
		{
			name: "struct ptr",
			args: args{
				src: &info{
					"hello1",
					32,
					&para1,
				},
				des: &info{
					"hello1",
					32,
					p2,
				},
			},
			want: true,
		},
		{
			name: "struct reflect 1",
			args: args{
				src: reflect.ValueOf(&info{
					"hello1",
					32,
					&para1,
				}),
				des: reflect.ValueOf(&info{
					"hello1",
					32,
					p2,
				}),
			},
			want: true,
		},
		{
			name: "struct reflect 2",
			args: args{
				src: reflect.ValueOf(&info{
					"hello1",
					32,
					&para1,
				}),
				des: reflect.ValueOf(&info{
					"hello1",
					31,
					p2,
				}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MemCmp(tt.args.src, tt.args.des); got != tt.want {
				t.Errorf("MemCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

type MyInter interface {
	do()
}

type AMan byte

func (o *AMan) do() {
	pt.Ptf("AMan is %c\n", *o)
}

// 深层次对比，对比2个接口类型
func TestMemCmp2(t *testing.T) {

	a := AMan('A')
	b := AMan('B')
	c := AMan('A')

	a.do()
	b.do()
	c.do()

	aI := &a
	bI := &b
	cI := &c

	type args struct {
		src interface{}
		des interface{}
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "普通对象 1",
			args: args{
				a,
				b,
			},
			want: false,
		},
		{
			name: "普通对象 2",
			args: args{
				a,
				c,
			},
			want: true,
		},
		{
			name: "接口对象 1",
			args: args{
				aI,
				bI,
			},
			want: false,
		},
		{
			name: "接口对象 2",
			args: args{
				aI,
				cI,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MemCmp(tt.args.src, tt.args.des); got != tt.want {
				t.Errorf("MemCmp() = %v, want %v", got, tt.want)
			}
		})
	}
}
