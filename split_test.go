package xslice

import (
	"fmt"
	"testing"
)

func TestSplitToChunks(t *testing.T) {
	var si = []int{1, 2, 3, 4, 5}
	i := SplitToChunks(si, 3)
	ssi := i.([][]int)
	fmt.Println(ssi) // [[1 2 3] [4 5]]

	var sm = []map[string]string{
		{"k1": "v1"},
		{"k2": "v2"},
		{"k3": "v3"},
		{"k4": "v4"},
		{"k5": "v5"},
	}
	ri := SplitToChunks(sm, 2)
	ssm := ri.([][]map[string]string)
	fmt.Println(ssm) // [[map[k1:v1] map[k2:v2]] [map[k3:v3] map[k4:v4]] [map[k5:v5]]]
}
