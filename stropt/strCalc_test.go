package stropt

import (
	"testing"

	"github.com/pislbert/MyGoUtil/pt"
)

func Test_DivideFloat(t *testing.T) {
	s, f, err := DivideStrFloat("2", "0.0000001", 4)
	pt.Ptln(s, ";", f, ";", err)
}

func Test_SumStrFloat(t *testing.T) {
	s, f, err := SumStrFloat([]string{"1.2", "3.44", "3"}, 1)
	pt.Ptln(s, ";", f, ";", err)
}

func TestSumStrFloat(t *testing.T) {
	type args struct {
		fs   []string
		prec int
	}
	tests := []struct {
		name        string
		args        args
		wantSResult string
		wantResult  float64
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSResult, gotResult, err := SumStrFloat(tt.args.fs, tt.args.prec)
			if (err != nil) != tt.wantErr {
				t.Errorf("SumStrFloat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSResult != tt.wantSResult {
				t.Errorf("SumStrFloat() gotSResult = %v, want %v", gotSResult, tt.wantSResult)
			}
			if gotResult != tt.wantResult {
				t.Errorf("SumStrFloat() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
