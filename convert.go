// go:generate _tools/tmpl generic.go.tmpl

package convert

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

// ToFloat64 converts argument to float64
func ToFloat64(arg interface{}) float64 {
	switch v := arg.(type) {
	case bool:
		if v {
			return -1
		}
		return 0
	case int:
		return float64(v)
	case int8:
		return float64(v)
	case int16:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint8:
		return float64(v)
	case uint16:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return float64(v)
	case complex64:
		return float64(real(v))
	case complex128:
		return float64(real(v))
	case string:
		ret, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err)
		}
		return ret
	case reflect.Value:
		return ToFloat64(v.Interface())
	case []byte:
		if len(v) > 0 && v[0] == '{' {
			return ToFloat64(ToMap(v))
		}
	case Matrix:
		r, c := v.Dims()
		if !(r == 1 && c == 1) {
			panic(fmt.Errorf("wrong dims %d %d", r, c))
		}
		return v.At(0, 0)
	case Vector:
		if v.Len() != 1 {
			panic(fmt.Errorf("wrong len %d", v.Len()))
		}
		return v.AtVec(0)
	case Dataer:
		return ToFloat64(v.Data())
	}
	var rv reflect.Value
	if v, ok := arg.(reflect.Value); ok {
		rv = v
	} else {
		rv = reflect.ValueOf(arg)
	}
	rv = reflect.Indirect(rv)
	rk := rv.Kind()
	if (rk == reflect.Float64 || rk == reflect.Float32) && rv.CanInterface() {
		return ToFloat64(rv.Interface())
	} else if (rk == reflect.Slice || rk == reflect.Array) && rv.Len() == 1 && rv.CanInterface() {
		return ToFloat64(rv.Index(0).Interface())
	} else if rk == reflect.Struct && rv.FieldByName("Value").Kind() != reflect.Invalid {
		return ToFloat64(rv.FieldByName("Value").Interface())
	} else if rk == reflect.Map {
		return ToFloat64(rv.MapIndex(reflect.ValueOf("Value")).Interface())
	}

	panic(fmt.Errorf("ToFloat64: unhandled type %T", arg))
}

// ToFloat32 converts argument to float32
func ToFloat32(arg interface{}) float32 {
	switch v := arg.(type) {
	case bool:
		if v {
			return -1
		}
		return 0
	case int:
		return float32(v)
	case int8:
		return float32(v)
	case int16:
		return float32(v)
	case int32:
		return float32(v)
	case int64:
		return float32(v)
	case uint:
		return float32(v)
	case uint8:
		return float32(v)
	case uint16:
		return float32(v)
	case uint32:
		return float32(v)
	case uint64:
		return float32(v)
	case float32:
		return float32(v)
	case float64:
		return float32(v)
	case complex64:
		return float32(real(v))
	case complex128:
		return float32(real(v))
	case string:
		ret, err := strconv.ParseFloat(v, 32)
		if err != nil {
			panic(err)
		}
		return float32(ret)
	case reflect.Value:
		return ToFloat32(v.Interface())
	case []byte:
		if len(v) > 0 && v[0] == '{' {
			return ToFloat32(ToMap(v))
		}
	case Matrix:
		r, c := v.Dims()
		if !(r == 1 && c == 1) {
			panic(fmt.Errorf("wrong dims %d %d", r, c))
		}
		return float32(v.At(0, 0))
	case Vector:
		if v.Len() != 1 {
			panic(fmt.Errorf("wrong len %d", v.Len()))
		}
		return float32(v.AtVec(0))
	case Dataer:
		return ToFloat32(v.Data())

	}
	var rv reflect.Value
	if v, ok := arg.(reflect.Value); ok {
		rv = v
	} else {
		rv = reflect.ValueOf(arg)
	}
	rv = reflect.Indirect(rv)
	rk := rv.Kind()
	if (rk == reflect.Float64 || rk == reflect.Float32) && rv.CanInterface() {
		return ToFloat32(rv.Interface())
	} else if (rk == reflect.Slice || rk == reflect.Array) && rv.Len() == 1 && rv.CanInterface() {
		return ToFloat32(rv.Index(0).Interface())
	} else if rk == reflect.Struct && rv.FieldByName("Value").Kind() != reflect.Invalid {
		return ToFloat32(rv.FieldByName("Value").Interface())
	} else if rk == reflect.Map {
		return ToFloat32(rv.MapIndex(reflect.ValueOf("Value")).Interface())
	}
	panic(fmt.Errorf("ToFloat32: unhandled type %T", arg))
}

// ToMap converts struct or map of json to map[string]interface{}
func ToMap(arg interface{}) map[string]interface{} {
	switch v := arg.(type) {

	case map[string]interface{}:
		return v
	case string:
		return ToMap([]byte(v))
	case []uint8:
		ret := make(map[string]interface{})
		err := json.Unmarshal(v, &ret)
		if err != nil {
			panic(errors.Wrapf(err, "ToMap(%T)", arg))
		}
		return ret
	}
	var rv reflect.Value
	if v, ok := arg.(reflect.Value); ok {
		rv = v
	} else {
		rv = reflect.ValueOf(arg)
	}
	rv = reflect.Indirect(rv)
	rt := rv.Type()
	if rv.Kind() == reflect.Struct {
		ret := make(map[string]interface{})
		for i := 0; i < rv.NumField(); i++ {
			n, f := rt.Field(i).Name, rv.Field(i)
			if f.CanInterface() {
				ret[n] = rv.Field(i).Interface()
			}
		}
	}
	panic(fmt.Errorf("unhandled ToMap(%T)", arg))
}
