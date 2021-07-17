package xslice

import (
	"reflect"
)

// SplitToChunks split a slice to some chunks, support []T which T can be any type
// usage:
// s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// i := SplitToChunks(s, 3)
// ss := i.([][]int)
func SplitToChunks(slice interface{}, chunkSize int) interface{} {
	sliceType := reflect.TypeOf(slice)
	sliceVal := reflect.ValueOf(slice)
	length := sliceVal.Len()
	if sliceType.Kind() != reflect.Slice {
		panic("parameter must be []T")
	}
	n := 0
	if length%chunkSize > 0 {
		n = 1
	}
	SST := reflect.MakeSlice(reflect.SliceOf(sliceType), 0, length/chunkSize+n)
	st, ed := 0, 0
	for st < length {
		ed = st + chunkSize
		if ed > length {
			ed = length
		}
		SST = reflect.Append(SST, sliceVal.Slice(st, ed))
		st = ed
	}
	return SST.Interface()
}
