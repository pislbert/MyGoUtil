package common

import (
	"MyGoUtil/pt"
	"testing"
)

func TestArrToObjs(t *testing.T) {

	a := []int{2, 4, 6, 8}
	b := ArrToObjs(a)

	for _, elem := range b {
		pt.Ptn(elem)
	}

}

func Test_FindElemInObjs(t *testing.T) {
	// pt.Ptln(FindElemInObjs(-1, []int{2, 4, 6, 8}))

	// pt.Ptln(FindElemInObjs(6, []float64{2, 4, 6, 8}))
	pt.Ptln(FindElemInObjs([]interface{}{2, 4, 6, 8}, 6))

	pt.Ptln(FindElemInObjs([]interface{}{"11", float64(4), 6, "haha"}, "haha"))

}
