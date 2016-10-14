package gonum

import (
    "fmt"
    "errors"
    "reflect"
)

var  ErrorInvalidParameter          = errors.New("Error Invalid Parameter");
var  ErrorMatrixIsEmpty             = errors.New("Error Matrix Is Empty");
var  ErrorElementTypeUnmatched      = errors.New("Error Element Type Unmatched");
var  ErrorElementTypeNotSet         = errors.New("Error Element Not Set");
type ElementType uint
const (
    ElementUnknown     = iota
    ElementFloat32    = 1
    ElementFloat64    = 2
    ElementComplex64  = 3
    ElementComplex128 = 4
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

func elementToString(t ElementType, v interface{}) string {
    var result string = "NaN";
    if v != nil {
        if t == ElementFloat32 {
            result = fmt.Sprintf("%0.4f", v.(float32));
        } else if t == ElementFloat64 {
            result = fmt.Sprintf("%0.4f", v.(float64));
        } else if t == ElementComplex64 {
            result = fmt.Sprintf("%g", v.(complex64));
        } else if t == ElementComplex128 {
            result = fmt.Sprintf("%g", v.(complex128));
        }
    }
    return result;
}

func elementTypeFromInterface(v interface{}) ElementType {
    var element_type ElementType = ElementUnknown;
    t := reflect.TypeOf(v);
    switch (t.Kind()) {
    case reflect.Float32:
        element_type = ElementFloat32;
    case reflect.Float64:
        element_type = ElementFloat64;
    case reflect.Complex64:
        element_type = ElementComplex64;
    case reflect.Complex128:
        element_type = ElementComplex128;
    }
    return element_type;
}

type Matrix struct {
    element_type ElementType;
    row_num int;
    col_num int;
    data interface{};
}

func NewMatrix(rows, cols int, t ElementType) *Matrix {
    m := Matrix{row_num : rows, col_num : cols, element_type : t, data : nil};
    return &m;
}

func NewEmptyMatrix() *Matrix {
    return NewMatrix(0, 0, ElementUnknown);
}

func NewMatrixWithData(rows, cols int, data interface{}) (*Matrix, error) {
    t := reflect.TypeOf(data);
    k := t.Kind();
    v := reflect.ValueOf(data);

    var element_type ElementType = ElementUnknown;
    var total_nums int = v.Len();
    var m *Matrix = nil;
    
    if rows * cols != total_nums {
        return nil, ErrorInvalidParameter;
    }
    if k != reflect.Slice && k != reflect.Array {
        return nil, ErrorInvalidParameter;
    }

    k2 := t.Elem().Kind();
    switch (k2) {
    case reflect.Float32:
        element_type = ElementFloat32;
        ss := make([]float32, total_nums, total_nums);
        copy(ss, data.([]float32));
        m = &Matrix{row_num : rows, col_num : cols, element_type : element_type, data : interface{}(ss)};
    case reflect.Float64:
        element_type = ElementFloat64;
        ss := make([]float64, total_nums, total_nums);
        copy(ss, data.([]float64));
        m = &Matrix{row_num : rows, col_num : cols, element_type : element_type, data : interface{}(ss)};
    case reflect.Complex64:
        element_type = ElementComplex64;
        ss := make([]complex64, total_nums, total_nums);
        copy(ss, data.([]complex64));
        m = &Matrix{row_num : rows, col_num : cols, element_type : element_type, data : interface{}(ss)};
    case reflect.Complex128:
        element_type = ElementComplex128;
        ss := make([]complex128, total_nums, total_nums);
        copy(ss, data.([]complex128));
        m = &Matrix{row_num : rows, col_num : cols, element_type : element_type, data : interface{}(ss)};
    }
    if m != nil {
        return m, nil;
    }
    return nil, ErrorInvalidParameter;
}

func (m *Matrix) String() string {

    var s string;
    if m.element_type != ElementUnknown {
        v := reflect.ValueOf(m.data);
        for r := 0; r < m.row_num; r++ {
            for c := 0; c < m.col_num; c++ {
                s = s + elementToString(m.element_type, v.Index(c * m.row_num + r).Interface()) + ", ";
            }
            s += "\n";
        }
    }
    s += fmt.Sprintf("Dim: [%d x %d], Type: %v\n", m.row_num, m.col_num, m.element_type);

    return s;
}

func (m *Matrix) GetRowNum() int {
    return m.row_num;
}

func (m *Matrix) GetColumeNum() int {
    return m.col_num;
}

func (m *Matrix) GetElementNum() int {
    return m.col_num * m.row_num;
}

func (m *Matrix) GetElementType() ElementType {
    return m.element_type;
}

func (m *Matrix) Get(row, col int) interface{} {
    if m.GetElementType() != ElementUnknown && m.data != nil {
        i := col * m.row_num + row;
        v := reflect.ValueOf(m.data);
        if (v.Len() > i) {
            return v.Index(i).Interface();
        }
    }
    return nil;
}

func (m *Matrix) createMatrixData(element_type ElementType, data interface{}) error {
    if m.data != nil {
        return;
    }

    total_nums := m.GetElementNum();
    if data != nil {
        t := reflect.TypeOf(data);
        k := t.Kind();
        v := reflect.ValueOf(data);
        // check data's Type
        if k != reflect.Slice && k != reflect.Array {
            return ErrorInvalidParameter;
        }
        // check data size
        if v.Len() > total_nums {
            return ;
        }
        k2 := t.Elem().Kind();
        if k2 == reflect.Float32 || k2 == reflect.Float64 || k2 == 
    }
    if element_type == ElementFloat32 {
        ss := make([]float32, total_nums, total_nums);
        if {
            copy(ss, data.([]float32));
        }
        m.data = interface{}(ss);
    } else if element_type == ElementFloat64 {
        ss := make([]float64, total_nums, total_nums);
        m.data = interface{}(ss);
    } else if element_type == ElementComplex64 {
        ss := make([]complex64, total_nums, total_nums);
        m.data = interface{}(ss);
    } else if element_type == ElementComplex128 {
        ss := make([]complex128, total_nums, total_nums);
        m.data = interface{}(ss);
    }
}

func (m *Matrix) Set(row, col int, val interface{}) error {
    element_type := elementTypeFromInterface(val);
    if m.data != nil {
        if m.GetElementType() == ElementUnknown {
            return ErrorElementTypeNotSet;
        }
        if element_type != m.GetElementType() {
            return ErrorElementTypeUnmatched;
        }
    } else {
        
    }

    i := col * m.row_num + row;
    v := reflect.ValueOf(m.data);
    v.Index(i).Set(reflect.ValueOf(val));

    return nil;
}

func (m *Matrix) SwapRow(i, j int) error {
    if m.GetElementType() == ElementUnknown || m.data == nil {
        return ErrorMatrixIsEmpty;
    }
    if i >= m.row_num || j >= m.row_num {
        return ErrorInvalidParameter;
    }
    element_type := m.GetElementType();
    if element_type == ElementFloat32 {
        s := m.data.([]float32);
        for c := 0; c < m.col_num; c++ {
            s[i + c * m.row_num], s[j + c * m.row_num] = s[j + c * m.row_num], s[i + c * m.row_num];
        }
    } else if element_type == ElementFloat64 {
        s := m.data.([]float64);
        for c := 0; c < m.col_num; c++ {
            s[i + c * m.row_num], s[j + c * m.row_num] = s[j + c * m.row_num], s[i + c * m.row_num];
        }
    } else if element_type == ElementComplex64 {
        s := m.data.([]complex64);
        for c := 0; c < m.col_num; c++ {
            s[i + c * m.row_num], s[j + c * m.row_num] = s[j + c * m.row_num], s[i + c * m.row_num];
        }
    } else if element_type == ElementComplex128 {
        s := m.data.([]complex128);
        for c := 0; c < m.col_num; c++ {
            s[i + c * m.row_num], s[j + c * m.row_num] = s[j + c * m.row_num], s[i + c * m.row_num];
        }
    }
    return nil;
}

func (m *Matrix) GetFloat32() ([]float32, error) {
    if m.GetElementType() == ElementUnknown || m.data == nil {
        return nil, ErrorMatrixIsEmpty;
    }
    if m.GetElementType() == ElementFloat32 {
        return nil, ErrorElementTypeUnmatched;
    }
    return m.data.([]float32), nil;
}

func (m *Matrix) GetFloat64() ([]float64, error) {
    if m.GetElementType() == ElementUnknown || m.data == nil {
        return nil, ErrorMatrixIsEmpty;
    }
    if m.GetElementType() == ElementFloat64 {
        return nil, ErrorElementTypeUnmatched;        
    }
    return m.data.([]float64), nil;
}

func (m *Matrix) GetComplex64() ([]complex64, error) {
    if m.GetElementType() == ElementUnknown || m.data == nil {
        return nil, ErrorMatrixIsEmpty;
    }
    if m.GetElementType() == ElementComplex64 {
        return nil, ErrorElementTypeUnmatched;
    }
    return m.data.([]complex64), nil;
}

func (m *Matrix) GetComplex128() ([]complex128, error) {
    if m.GetElementType() == ElementUnknown || m.data == nil {
        return nil, ErrorMatrixIsEmpty;
    }
    if m.GetElementType() != ElementComplex64 {
        return nil, ErrorElementTypeUnmatched;
    }
    return m.data.([]complex128), nil;
}
/*
            // switch (element_type) {
            // case ElementFloat32:

            // case ElementFloat64:
            // case ElementComplex64:
            // case ElementComplex128:
            // }

        if t == ElementFloat32 {
            v := reflect.ValueOf(m.data);
            result = fmt.Sprintf("%0.4f", v.(float64));
        } else if t == ElementFloat64 {
            result = fmt.Sprintf("%0.4f", v.(float64));
        } else if t == ElementComplex64 {
            result = fmt.Sprintf("%g", v.(complex64));
        } else if t == ElementComplex128 {
            result = fmt.Sprintf("%g", v.(complex128));
        }
func (m *Matrix) Copy() *Matrix {
    c := NewMatrix(m.row_num, m.col_num, m.data);
	return c;
}


func (m *Matrix) getPtr() *[]float64 {
    return &(m.data);
}


func (m *Matrix) GetColumeData(col int) []float64 {
    if col >= m.GetColumeNum() {
        panic("GetColumeData, col index large than matrix's col_num");
    }
    i := col * m.row_num;
    return m.data[i : i + m.row_num];
}

func (m *Matrix) GetData() []float64 {
    return m.data;
}

func (m *Matrix) GetColumes(cols []int) *Matrix {
    for _, c := range cols {
        if c >= m.col_num {
            panic("GetColumes(), submatrix's col is large than origin matrix's col_num");
        }
    }
    data := make([]float64, 0);
    for _, c := range cols {
        i := c * m.row_num;
        data = append(data, m.data[i : i + m.row_num]...);
    }
    sub := NewMatrix(m.row_num, len(cols), data);
    return sub;
}

func (m *Matrix) SwapRow(i, j int) {
    if i >= m.row_num || j >= m.row_num {
        panic("SwapRow(), i or j large than row nums");
    }
    for c := 0; c < m.col_num; c++ {
        m.data[i + c * m.row_num], m.data[j + c * m.row_num] = m.data[j + c * m.row_num], m.data[i + c * m.row_num];
    }
}

func (m *Matrix) RBind(other *Matrix) *Matrix {
    if m.GetColumeNum() != other.GetColumeNum() {
        panic("RBind(), two matrix's colume num is diff");
    }
    col_num := m.GetColumeNum();
    size := m.GetElementNum() + other.GetElementNum();
    data := make([]float64, 0, size);
    for i := 0; i < col_num; i++ {
        d1 := m.GetColumeData(i);
        d2 := other.GetColumeData(i)
        data = append(data, d1...);
        data = append(data, d2...);
    }

    matrix := NewMatrix(m.GetRowNum() + other.GetRowNum(), col_num, data);
    return matrix;
}

func (m *Matrix) CBind(other *Matrix) *Matrix {
    if m.GetRowNum() != other.GetRowNum() {
        panic("RBind(), two matrix's row num is diff");
    }
    data := append(m.data, other.data...);
    matrix := NewMatrix(m.GetRowNum(), m.GetColumeNum() + other.GetColumeNum(), data);
    return matrix;
}

*/


