package gonum

import (
    "fmt"
    "reflect"
)

/*
#cgo CFLAGS: -Wall -Wno-unused-variable -Wno-unused-function -I"./third_party/target/include"
#cgo LDFLAGS: -L./third_party/target/lib -lopenblas


#include <complex.h>
#include <cblas.h>

#define max(a, b) (a > b ? a : b)

void cgo_cblas_dgemm(double* m1, double *m2, double *m3, int m1_row, int m1_col, int m2_col, double alpha, double beta) {

    int lda = max(1, m1_row);
    int ldb = max(1, m1_col);
    int ldc = max(1, m1_row);

    cblas_dgemm(CblasColMajor, CblasNoTrans, CblasNoTrans,
        m1_row, m2_col, m1_col, alpha, m1, lda, m2, ldb, beta, m3, ldc);
    return;
}

void cgo_cblas_zgemm(double* m1, double *m2, double *m3, int m1_row, int m1_col, int m2_col, double complex alpha, double complex beta) {

    int lda = max(1, m1_row);
    int ldb = max(1, m1_col);
    int ldc = max(1, m1_row);

    cblas_zgemm(CblasColMajor, CblasNoTrans, CblasNoTrans,
            m1_row, m2_col, m1_col, (double *)&alpha, m1, lda, m2, ldb, (double *)&beta, m3, ldc);
    return;
}

*/
import "C"
import "unsafe"

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

func NewMatrixWithData(rows, cols int, values interface{}) (*Matrix, error) {
    t := reflect.TypeOf(values);
    k := t.Kind();
    v := reflect.ValueOf(values);

    var element_type ElementType = ElementUnknown;
    var element_nums int = v.Len();
    var m *Matrix = nil;

    if rows * cols != element_nums {
        return nil, ErrorInvalidParameter;
    }
    if k != reflect.Slice && k != reflect.Array {
        return nil, ErrorInvalidParameter;
    }

    element_type = ElementType(t.Elem().Kind());
    data, err := createDataSlice(element_type, element_nums);
    if err == nil {
        switch (element_type) {
        case ElementFloat64:
            copy(data.([]float64), values.([]float64));
        case ElementComplex128:
            copy(data.([]complex128), values.([]complex128));
        }
        m = &Matrix{row_num : rows, col_num : cols, element_type : element_type, data : data};
    }
    return m, err;
}

func (m *Matrix) String() string {

    var s string;
    if m.element_type != ElementUnknown {
        v := reflect.ValueOf(m.data);
        for r := 0; r < m.row_num; r++ {
            for c := 0; c < m.col_num; c++ {
                s = s + elementToString(v.Index(c * m.row_num + r).Interface()) + ", ";
            }
            s += "\n";
        }
    }
    s += fmt.Sprintf("Dimension: [%d x %d], Type: %v\n", m.row_num, m.col_num, m.element_type);
    return s;
}

func (m *Matrix) Copy() *Matrix {
    new_matrix, _ := NewMatrixWithData(m.GetRowNum(), m.GetColumeNum(), m.data);
    return new_matrix;
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

func (m *Matrix) Set(row, col int, value interface{}) error {

    element_type := elementTypeFromInterface(value);
    element_nums := m.GetElementNum();

    i := col * m.row_num + row;
    if i >= m.GetElementNum() {
        return ErrorInvalidParameter;
    }

    if m.data == nil {
        if m.GetElementType() != ElementUnknown && m.GetElementType() != element_type {
            return ErrorElementTypeUnmatched;
        }
        var err error = nil;
        m.data, err = createDataSlice(element_type, element_nums);
        if err != nil {
            return err;
        }
        m.element_type = element_type;
    }
    if m.GetElementType() == ElementUnknown {
        return ErrorElementTypeNotSet;
    }

    if m.GetElementType() != element_type {
        return ErrorElementTypeUnmatched;
    }

    v := reflect.ValueOf(m.data);
    v.Index(i).Set(reflect.ValueOf(value));

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
    if element_type == ElementFloat64 {
        s := m.data.([]float64);
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
//  cgo_cblas_dgemm(double* m1, double *m2, double *m3, int m1_row, int m1_col, int m2_col, double alpha, double beta)
func (m1 *Matrix) Product(alpha interface{}, m2 *Matrix) (*Matrix, error) {

    if m1.GetColumeNum() != m2.GetRowNum() {
        return nil, ErrorDimUnmatched;
    }

    var err error = nil;
    var new_matrix *Matrix = nil;
    if m1.GetElementType() == ElementFloat64 {
        if m2.GetElementType() != ElementFloat64 {
            return nil, ErrorElementTypeUnmatched;
        }
        var f_alpha float64 = 1.0;
        if alpha != nil {
            if elementTypeFromInterface(alpha) != ElementFloat64 {
                return nil, ErrorElementTypeUnmatched;
            }
            f_alpha = alpha.(float64);
        }
        var data1 []float64 = nil;
        var data2 []float64 = nil;
        total := m1.GetRowNum() * m2.GetColumeNum();

        data1, err = m1.GetFloat64();
        if err == nil {
            data2, err = m2.GetFloat64();
        }
        if err == nil {
            data3 := make([]float64, total, total);

            C.cgo_cblas_dgemm((*C.double)(unsafe.Pointer(&data1[0])), (*C.double)(unsafe.Pointer(&data2[0])), (*C.double)(unsafe.Pointer(&data3[0])),
                            (C.int)(m1.GetRowNum()), (C.int)(m1.GetColumeNum()), (C.int)(m2.GetColumeNum()), C.double(f_alpha), C.double(0.0));

            new_matrix, err = NewMatrixWithData(m1.GetRowNum(), m2.GetColumeNum(), data3);
        }

    } else if m1.GetElementType() == ElementComplex128 {

       if m2.GetElementType() != ElementComplex128 {
            return nil, ErrorElementTypeUnmatched;
        }
        var c_alpha complex128 = 1.0;
        var c_beta complex128 = 0.0;
        if alpha != nil {
            if elementTypeFromInterface(alpha) != ElementComplex128 {
                return nil, ErrorElementTypeUnmatched;
            }
            c_alpha = alpha.(complex128);
        }
        var data1 []complex128 = nil;
        var data2 []complex128 = nil;
        total := m1.GetRowNum() * m2.GetColumeNum();

        data1, err = m1.GetComplex128();
        if err == nil {
            data2, err = m2.GetComplex128();
        }
        if err == nil {
            data3 := make([]complex128, total, total);

            C.cgo_cblas_zgemm((*C.double)(unsafe.Pointer(&data1[0])), (*C.double)(unsafe.Pointer(&data2[0])), (*C.double)(unsafe.Pointer(&data3[0])),
                            (C.int)(m1.GetRowNum()), (C.int)(m1.GetColumeNum()), (C.int)(m2.GetColumeNum()), (C.complexdouble)(c_alpha), (C.complexdouble)(c_beta));

            new_matrix, err = NewMatrixWithData(m1.GetRowNum(), m2.GetColumeNum(), data3);
        }
    }
    return new_matrix, err;
}

func (m *Matrix) Multiply(v interface{}) (*Matrix, error) {
    if elementTypeFromInterface(v) != m.GetElementType() {
        return nil, ErrorElementTypeUnmatched;
    }
    var err error = nil;
    var new_matrix *Matrix = nil;
    var total_nums int = m.GetElementNum();
    if m.GetElementType() == ElementFloat64 {
        var src []float64;
        src, err = m.GetFloat64();
        if err == nil {
            dst := make([]float64, total_nums, total_nums);
            vv := v.(float64);
            for i := 0; i < total_nums; i++ {
                dst[i] = src[i] * vv;
            }
            new_matrix, err = NewMatrixWithData(m.GetRowNum(), m.GetColumeNum(), dst);
        }
    } else if m.GetElementType() == ElementComplex128 {
        var src []complex128;
        src, err = m.GetComplex128();
        if err == nil {
            dst := make([]complex128, total_nums, total_nums);
            vv := v.(complex128);
            for i := 0; i < total_nums; i++ {
                dst[i] = src[i] * vv;
            }
            new_matrix, err = NewMatrixWithData(m.GetRowNum(), m.GetColumeNum(), dst);
        }
    }

    return new_matrix, err;
}

func (m *Matrix) Add(v interface{}) (*Matrix, error) {
    t := reflect.TypeOf(v);
    if t.Kind() == reflect.Ptr && t.Elem().Name() == "Matrix" {
        vv := v.(*Matrix);
        if vv.GetRowNum() != m.GetRowNum() || vv.GetColumeNum() != m.GetColumeNum() {
            return nil, ErrorDimUnmatched;
        }

        var err error = nil;
        var new_matrix *Matrix = nil;
        var total_nums int = m.GetElementNum();

        if vv.GetElementType() == ElementFloat64 && m.GetElementType() == ElementFloat64 {

            var data1 []float64 = nil;
            var data2 []float64 = nil;

            data1, err = m.GetFloat64();
            if err == nil {
                data2, err = vv.GetFloat64();
            }
            if err == nil {
                data3 := make([]float64, total_nums, total_nums);
                for i := 0; i < total_nums; i++ {
                    data3[i] = data1[i] + data2[i];
                }
                new_matrix, err = NewMatrixWithData(m.GetRowNum(), m.GetColumeNum(), data3);
            }

        } else if vv.GetElementType() == ElementComplex128 && m.GetElementType() == ElementComplex128 {

            var data1 []complex128 = nil;
            var data2 []complex128 = nil;

            data1, err = m.GetComplex128();
            if err == nil {
                data2, err = vv.GetComplex128();
            }
            if err == nil {
                data3 := make([]complex128, total_nums, total_nums);
                for i := 0; i < total_nums; i++ {
                    data3[i] = data1[i] + data2[i];
                }
                new_matrix, err = NewMatrixWithData(m.GetRowNum(), m.GetColumeNum(), data3);
            }
        }
        return new_matrix, err;

    } else {

        var err error = nil;
        var new_matrix *Matrix = nil;
        var total_nums int = m.GetElementNum();
        var v_type ElementType = elementTypeFromInterface(v);

        if v_type == ElementFloat64 && m.GetElementType() == ElementFloat64 {
            var data1 []float64 = nil;
            var data2 []float64 = nil;
            vv := v.(float64);
            data1, err = m.GetFloat64();
            if err == nil {
                data2 = make([]float64, total_nums, total_nums);
                for i := 0; i < total_nums; i++ {
                    data2[i] = data1[i] + vv;
                }
                new_matrix, err = NewMatrixWithData(m.GetRowNum(), m.GetColumeNum(), data2);
            }

        } else if v_type == ElementComplex128 && m.GetElementType() == ElementComplex128 {
            vv := v.(complex128);
            var data1 []complex128 = nil;
            var data2 []complex128 = nil;

            data1, err = m.GetComplex128();
            if err == nil {
                data2 = make([]complex128, total_nums, total_nums);
                for i := 0; i < total_nums; i++ {
                    data2[i] = data1[i] + vv;
                }
                new_matrix, err = NewMatrixWithData(m.GetRowNum(), m.GetColumeNum(), data2);
            }
        }
        return new_matrix, err;
    }
    return nil, ErrorElementTypeUnmatched;
}

func (m *Matrix) Transpose() *Matrix {
    row_num := m.GetRowNum();
    col_num := m.GetColumeNum();
    total_nums := m.GetElementNum();
    data := reflect.MakeSlice(reflect.TypeOf(m.data), total_nums, total_nums);
    i := 0;
    fmt.Printf("%d,%d", row_num, col_num);
    for r := 0; r < row_num; r++ {
        for c := 0; c < col_num; c++ {
            v := data.Index(i);
            vv := m.Get(r, c);
            v.Set(reflect.ValueOf(vv));
            i += 1;
        }
    }
    new_matrix, _ := NewMatrixWithData(m.GetColumeNum(), m.GetRowNum(), data.Interface());
    return new_matrix;
}

func (m *Matrix) RBind(other *Matrix) (*Matrix, error) {
    if m.GetColumeNum() != other.GetColumeNum() {
        return nil, ErrorDimUnmatched;
    } else if m.GetElementType() != other.GetElementType() {
        return nil, ErrorElementTypeUnmatched;
    } else if m.GetElementType() == ElementUnknown || m.data == nil {
        return nil, ErrorElementTypeNotSet;
    }

    col_num := m.GetColumeNum();
    data := reflect.MakeSlice(reflect.TypeOf(m.data), 0, 0);
    for i := 0; i < col_num; i++ {
        d1, err := m.GetColumeData(i);
        if err != nil {
            return nil, err;
        }
        d2, err := other.GetColumeData(i);
        if err != nil {
            return nil, err;
        }
        data = reflect.AppendSlice(data, reflect.ValueOf(d1));
        data = reflect.AppendSlice(data, reflect.ValueOf(d2));
    }

    matrix, err := NewMatrixWithData(m.GetRowNum() + other.GetRowNum(), col_num, data.Interface());
    return matrix, err;
}

func (m *Matrix) CBind(other *Matrix) (*Matrix, error) {
    if m.GetRowNum() != other.GetRowNum() {
        return nil, ErrorDimUnmatched;
    } else if m.GetElementType() != other.GetElementType() { // 2 matrix have diff element type
        return nil, ErrorElementTypeUnmatched;
    } else if m.GetElementType() == ElementUnknown || m.data == nil {
        return nil, ErrorElementTypeNotSet;
    }
    data := reflect.MakeSlice(reflect.TypeOf(m.data), 0, 0);
    data = reflect.AppendSlice(data, reflect.ValueOf(m.data));
    data = reflect.AppendSlice(data, reflect.ValueOf(other.data));
  //  fmt.Printf("%d\n", data);
    matrix, err := NewMatrixWithData(m.GetRowNum(), m.GetColumeNum() + other.GetColumeNum(), data.Interface());
    return matrix, err;
}

func (m *Matrix) GetColumesByIndex(l_col_index ... int) (*Matrix, error) {
    for _, c := range l_col_index {
        if c >= m.GetColumeNum() {
            return nil, ErrorInvalidParameter;
        }
    }

    data := reflect.MakeSlice(reflect.TypeOf(m.data), 0, 0);
    for _, c := range l_col_index {
        v := reflect.ValueOf(m.data);
        i := c * m.GetRowNum();
        col := v.Slice(i, i + m.GetRowNum());
        data = reflect.AppendSlice(data, col);
    }
    sub_matrix, err := NewMatrixWithData(m.GetRowNum(), len(l_col_index), data.Interface());
    return sub_matrix, err;
}

func (m *Matrix) GetColumeData(col int) (interface{}, error) {
    if col >= m.GetColumeNum() {
        return nil, ErrorInvalidParameter;
    }
    if m.data == nil || m.GetElementType() == ElementUnknown {
        return nil, ErrorMatrixIsEmpty;
    }
    v := reflect.ValueOf(m.data);
    i := col * m.row_num;
    result := v.Slice(i, i + m.row_num);
    return result.Interface(), nil;
}

func (m *Matrix) GetFloat64() ([]float64, error) {
    if m.GetElementType() == ElementUnknown || m.data == nil {
        return nil, ErrorMatrixIsEmpty;
    }
    if m.GetElementType() == ElementFloat64 {
        return m.data.([]float64), nil;
    }
    return nil, ErrorElementTypeUnmatched;
}

func (m *Matrix) GetComplex128() ([]complex128, error) {
    if m.GetElementType() == ElementUnknown || m.data == nil {
        return nil, ErrorMatrixIsEmpty;
    }
    if m.GetElementType() == ElementComplex128 {
        return m.data.([]complex128), nil;
    }
    return nil, ErrorElementTypeUnmatched;
}

func createDataSlice(element_type ElementType, element_nums int) (interface{}, error) {

    var data interface{} = nil;
    var err error = nil
    if element_nums > 0 {
        if element_type == ElementFloat64 {
            ss := make([]float64, element_nums, element_nums);
            data = interface{}(ss);
        } else if element_type == ElementComplex128 {
            ss := make([]complex128, element_nums, element_nums);
            data = interface{}(ss);
        } else {
            err = ErrorInvalidParameter
        }
    } else {
        err = ErrorMatrixIsEmpty;
    }
    return data, err;
}

/*

   if m.data != nil {
        return;
    }
    element_type ElementType

    total_nums := m.GetElementNum();
    if data != nil {
        t := reflect.TypeOf(data);
        k := reflect.TypeOf(data).Kind();
        v := reflect.ValueOf(data);

        // make sure data is slice or array
        if k != reflect.Slice && k != reflect.Array {
            return ErrorInvalidParameter;
        }
        // check data size
        if v.Len() > total_nums {
            return ErrorInvalidParameter;
        }

        element_type = t.Elem().Kind();
        if element_type != m.GetElementType() {
            return ErrorInvalidParameter;
        }
    } else {

    }

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

func (m *Matrix) GetRowData(row int) []float64 {
    if row >= m.GetRowNum() {
        panic("GetRowData, row index large than matrix's row_num");
    }
    col_nums := m.GetColumeNum();
    rdata := make([]float64, col_nums, col_nums);
    for i := 0; i < col_nums; i++ {
        rdata[i] = m.Get(row, i);
    }
    return rdata;
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


