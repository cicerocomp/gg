package gg

import "testing"


type InArrayStringData struct {
	array []string
	value string
	result bool
}

type InArrayFloat64Data struct {
	array []float64
	value float64
	result bool
}

type InArrayFloat32Data struct {
	array []float32
	value float32
	result bool
}


type InArrayInt64Data struct {
	array []int64
	value int64
	result bool
}


type InArrayInt32Data struct {
	array []int32
	value int32
	result bool
}

type InArrayInt16Data struct {
	array []int16
	value int16
	result bool
}


type InArrayInt8Data struct {
	array []int8
	value int8
	result bool
}

type InArrayIntData struct {
	array []int
	value int
	result bool
}



func TestInArrayString(t *testing.T) {
	var tests = []InArrayStringData{
		{ []string{"a", "b", "c", "d", "e", "f"}, "f", true},
		{ []string{"0", "1", "2", "3", "4", "5"}, "1", true},
		{ []string{"a", "b", "c", "d", "e", "f"}, "", false},
		{ []string{"a", "b", "c", "d", "e", "f"}, "l", false},
		{ []string{"a", "b", "c", "d", "e", ""}, "", true},
		{ []string{"a", "b", "c", "d", "e", "f"}, "1", false},
	}

	for _, pair := range tests {
		v := InArrayString(pair.value, pair.array)
		if v != pair.result {
			t.Error(
				"TestInArrayString(", pair.array,",", pair.value, ")",
				"expected", pair.result,
				"got", v,
			)
		}
	}
}


func TestInArrayFloat64(t *testing.T) {
	var tests = []InArrayFloat64Data{
		{ []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}, 1.1, true},
		{ []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}, 2.2, true},
		{ []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}, 7.7, true},
		{ []float64{1.1, 2, 3.3, 4.4, 5.5, 6.6, 7.7}, 2.0, true},
		{ []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}, 1.0, false},
		{ []float64{1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}, 1.0, true},
		{ []float64{1, 2.2, 3.3, 4.4, 5.5, 6.6, 100}, 100.001, false},
	}

	for _, pair := range tests {
		v := InArrayFloat64(pair.value, pair.array)
		if v != pair.result {
			t.Error(
				"InArrayFloat64(", pair.array,",", pair.value, ")",
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

func TestInArrayFloat32(t *testing.T) {
	var tests = []InArrayFloat32Data{
		{ []float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}, 1.1, true},
		{ []float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}, 2.2, true},
		{ []float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}, 7.7, true},
		{ []float32{1.1, 2, 3.3, 4.4, 5.5, 6.6, 7.7}, 2.0, true},
		{ []float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}, 1.0, false},
		{ []float32{1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}, 1.0, true},
		{ []float32{1, 2.2, 3.3, 4.4, 5.5, 6.6, 100}, 100.001, false},
		{ []float32{-1, 2.2, 3.3, 4.4, 5.5, 6.6, 100}, 1, false},
		{ []float32{-1, 2.2, 3.3, 4.4, 5.5, 6.6, 1}, 1, true},
	}

	for _, pair := range tests {
		v := InArrayFloat32(pair.value, pair.array)
		if v != pair.result {
			t.Error(
				"InArrayFloat32(", pair.array,",", pair.value, ")",
				"expected", pair.result,
				"got", v,
			)
		}
	}
}



func TestInArrayInt64(t *testing.T) {
	var tests = []InArrayInt64Data{
		{ []int64{1, 2, 3, 4, 5, 6, 7}, 1, true},
		{ []int64{1, 2, 3, 4, 5, 6, 7}, 4, true},
		{ []int64{1, 2, 3, 4, 5, 6, 7}, 7, true},
		{ []int64{1, 2, 3, 4, 5, 6, 7}, 100, false},
		{ []int64{1, 2, 3, 4, 5, 6, 7}, 0, false},
		{ []int64{0, 1, 2, 3, 4, 5, 6, 7}, 0, true},
		{ []int64{0, -100, 2, 3, 4, 5, 6, 7}, -100, true},
		{ []int64{0, 100, 2, 3, 4, 5, 6, 7}, -100, false},
	}

	for _, pair := range tests {
		v := InArrayInt64(pair.value, pair.array)
		if v != pair.result {
			t.Error(
				"InArrayInt64(", pair.array,",", pair.value, ")",
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

func TestInArrayInt32(t *testing.T) {
	var tests = []InArrayInt32Data{
		{ []int32{1, 2, 3, 4, 5, 6, 7}, 1, true},
		{ []int32{1, 2, 3, 4, 5, 6, 7}, 4, true},
		{ []int32{1, 2, 3, 4, 5, 6, 7}, 7, true},
		{ []int32{1, 2, 3, 4, 5, 6, 7}, 100, false},
		{ []int32{1, 2, 3, 4, 5, 6, 7}, 0, false},
		{ []int32{0, 1, 2, 3, 4, 5, 6, 7}, 0, true},
		{ []int32{0, -100, 2, 3, 4, 5, 6, 7}, -100, true},
		{ []int32{0, 100, 2, 3, 4, 5, 6, 7}, -100, false},
	}

	for _, pair := range tests {
		v := InArrayInt32(pair.value, pair.array)
		if v != pair.result {
			t.Error(
				"InArrayInt32(", pair.array,",", pair.value, ")",
				"expected", pair.result,
				"got", v,
			)
		}
	}
}


func TestInArrayInt16(t *testing.T) {
	var tests = []InArrayInt16Data{
		{ []int16{1, 2, 3, 4, 5, 6, 7}, 1, true},
		{ []int16{1, 2, 3, 4, 5, 6, 7}, 4, true},
		{ []int16{1, 2, 3, 4, 5, 6, 7}, 7, true},
		{ []int16{1, 2, 3, 4, 5, 6, 7}, 100, false},
		{ []int16{1, 2, 3, 4, 5, 6, 7}, 0, false},
		{ []int16{0, 1, 2, 3, 4, 5, 6, 7}, 0, true},
		{ []int16{0, -100, 2, 3, 4, 5, 6, 7}, -100, true},
		{ []int16{0, 100, 2, 3, 4, 5, 6, 7}, -100, false},
	}

	for _, pair := range tests {
		v := InArrayInt16(pair.value, pair.array)
		if v != pair.result {
			t.Error(
				"InArrayInt16(", pair.array,",", pair.value, ")",
				"expected", pair.result,
				"got", v,
			)
		}
	}
}


func TestInArrayInt8(t *testing.T) {
	var tests = []InArrayInt8Data{
		{ []int8{1, 2, 3, 4, 5, 6, 7}, 1, true},
		{ []int8{1, 2, 3, 4, 5, 6, 7}, 4, true},
		{ []int8{1, 2, 3, 4, 5, 6, 7}, 7, true},
		{ []int8{1, 2, 3, 4, 5, 6, 7}, 100, false},
		{ []int8{1, 2, 3, 4, 5, 6, 7}, 0, false},
		{ []int8{0, 1, 2, 3, 4, 5, 6, 7}, 0, true},
		{ []int8{0, -100, 2, 3, 4, 5, 6, 7}, -100, true},
		{ []int8{0, 100, 2, 3, 4, 5, 6, 7}, -100, false},
	}

	for _, pair := range tests {
		v := InArrayInt8(pair.value, pair.array)
		if v != pair.result {
			t.Error(
				"InArrayInt8(", pair.array,",", pair.value, ")",
				"expected", pair.result,
				"got", v,
			)
		}
	}
}


func TestInArrayInt(t *testing.T) {
	var tests = []InArrayIntData{
		{ []int{1, 2, 3, 4, 5, 6, 7}, 1, true},
		{ []int{1, 2, 3, 4, 5, 6, 7}, 4, true},
		{ []int{1, 2, 3, 4, 5, 6, 7}, 7, true},
		{ []int{1, 2, 3, 4, 5, 6, 7}, 100, false},
		{ []int{1, 2, 3, 4, 5, 6, 7}, 0, false},
		{ []int{0, 1, 2, 3, 4, 5, 6, 7}, 0, true},
		{ []int{0, -100, 2, 3, 4, 5, 6, 7}, -100, true},
		{ []int{0, 100, 2, 3, 4, 5, 6, 7}, -100, false},
	}

	for _, pair := range tests {
		v := InArrayInt(pair.value, pair.array)
		if v != pair.result {
			t.Error(
				"InArrayInt(", pair.array,",", pair.value, ")",
				"expected", pair.result,
				"got", v,
			)
		}
	}
}