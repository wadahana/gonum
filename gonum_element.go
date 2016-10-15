package gonum

import (
    "fmt"
    "reflect"
    "math"
    "math/cmplx"
)


var RealNaN float64= math.NaN();
var ComplexNaN complex128 = cmplx.NaN();

type ElementType uint;

const (
    ElementUnknown    = ElementType(reflect.Invalid)
    ElementFloat32    = ElementType(reflect.Float32)
    ElementFloat64    = ElementType(reflect.Float64)
    ElementComplex64 = ElementType(reflect.Complex64)
    ElementComplex128 = ElementType(reflect.Complex128)
)

func (t ElementType) String() string {

    if t == ElementFloat32 {
        return "ElementFloat32";
    } else if t == ElementFloat64 {
        return "ElementFloat64";
    } else if t == ElementComplex64 {
        return "ElementComplex64";
    } else if t == ElementComplex128 {
        return "ElementComplex128";
    } else {
        return "ElementUnknow";
    }
}


func elementToString(v interface{}) string {
    var result string = "NaN";
    if v != nil {
        t := elementTypeFromInterface(v);
        if t == ElementFloat32 {
            result = fmt.Sprintf("%16.4f", v.(float32));
        } else if t == ElementFloat64 {
            result = fmt.Sprintf("%16.4f", v.(float64));
        } else if t == ElementComplex64 {
            result = fmt.Sprintf("%g", v.(complex64));
        } else if t == ElementComplex128 {
            result = fmt.Sprintf("%g", v.(complex128));
        }
    }
    return result;
}

func elementTypeIsAvailable(t ElementType) bool {
    return (t == ElementFloat64 || t == ElementComplex128);
}

func elementTypeFromInterface(v interface{}) ElementType {
    t := reflect.TypeOf(v);
    k := ElementType(t.Kind());
    if  elementTypeIsAvailable(k) {
        return k
    }
    return ElementUnknown;
}

