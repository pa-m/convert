package convert_test

import (
	"math"
	"testing"

	"github.com/pa-m/convert"
)

func TestToFloat64(t *testing.T) {
	ea := []struct {
		e   float64
		arg interface{}
	}{
		{0, false},
		{-1, true},
		{2, byte(2)},
		{-1, int8(-1)},
		{-1, int16(-1)},
		{-1, int32(-1)},
		{-1, int64(-1)},

		{-1, int(-1)},
		{2, (uint(2))},
		{-1.2, (float32(-1.2))},
		{-1.2, (float64(-1.2))},
		{-1.2, (complex64(-1.2 + 3.4i))},
		{-1.2, (complex128(-1.2 + 3.4i))},
		{-1.2, []float32{-1.2}},
		{-1.2, "-1.2"},
		{-1.2, map[string]string{"Value": "-1.2"}},
		{-1.2, struct{ Value float64 }{-1.2}},
		{-1.2, []byte(`{"Value":"-1.2"}`)},
		//{-1.2, mat.NewDense(1, 1, []float64{-1.2})},
		//{-1.2, mat.NewVecDense(1, []float64{-1.2})},
	}
	for _, pair := range ea {
		arg, e, a := pair.arg, pair.e, convert.ToFloat64(pair.arg)
		if math.Abs(e-a) > 1e-7 {
			t.Errorf("ToFloat64(%T): expected %v got %v", arg, e, a)
		}
	}

}

func TestToFloat32(t *testing.T) {
	ea := []struct {
		e   float32
		arg interface{}
	}{
		{0, false},
		{-1, true},
		{2, byte(2)},
		{-1, int8(-1)},
		{-1, int16(-1)},
		{-1, int32(-1)},
		{-1, int64(-1)},

		{-1, int(-1)},
		{2, (uint(2))},
		{-1.2, (float32(-1.2))},
		{-1.2, (float64(-1.2))},
		{-1.2, (complex64(-1.2 + 3.4i))},
		{-1.2, (complex128(-1.2 + 3.4i))},
		{-1.2, []float32{-1.2}},
		{-1.2, "-1.2"},
		{-1.2, map[string]string{"Value": "-1.2"}},
		{-1.2, struct{ Value float64 }{-1.2}},
		{-1.2, []byte(`{"Value":"-1.2"}`)},
	}
	for _, pair := range ea {
		arg, e, a := pair.arg, pair.e, convert.ToFloat32(pair.arg)
		if math.Abs(float64(e-a)) > 1e-7 {
			t.Errorf("ToFloat32(%T): expected %v got %v", arg, e, a)
		}
	}

}
