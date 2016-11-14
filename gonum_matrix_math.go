
package gonum

import (
//    "fmt"
    "math"
    "math/cmplx"
)

const (
    AxisNone      = 0
    AxisByColume  = 1
    AxisByRow     = 2
)


func (m *Matrix) Diff(lag int, axis int) (*Matrix, error) {

    if axis == AxisByColume {
        return m.diffByColume(lag);
    }
    return nil, ErrorNotImplemented;
}

func (m *Matrix) diffByColume(lag int) (*Matrix, error) {
    if m.GetElementType() == ElementUnknown || m.data == nil {
        return nil, ErrorMatrixIsEmpty;
    }
    row_num := m.GetRowNum();
    col_num := m.GetColumeNum();
    if lag >= row_num {
        return nil, ErrorInvalidParameter;
    }

    var i_data interface{} = nil;
    size := (row_num - lag) * m.GetColumeNum();
    if m.GetElementType() == ElementFloat64 {
        new_data := make([]float64, 0, size);
        for c := 0; c < col_num; c++ {
            col, err := m.GetColumeData(c);
            if err != nil {
                return nil, err;
            }
            col_data := col.([]float64);
            for i := 0 ; i < row_num - lag; i++ {
                v1, v2 := col_data[i], col_data[i+lag];
                new_data = append(new_data, (v2 - v1));
            }
        }
        i_data = interface{}(new_data);
    } else if m.GetElementType() == ElementComplex128 {
        new_data := make([]complex128, 0, size);
        for c := 0; c < col_num; c++ {
            col, err := m.GetColumeData(c);
            if err != nil {
                return nil, err;
            }
            col_data := col.([]complex128);
            for i := 0 ; i < row_num - lag; i++ {
                v1, v2 := col_data[i], col_data[i+lag];
                new_data = append(new_data, (v2 - v1));
            }
            i_data = interface{}(new_data);
        }
    }
    new_matrix, err := NewMatrixWithData(row_num - lag, m.GetColumeNum(), i_data);
    return new_matrix, err;
}

func (m *Matrix) Log() (*Matrix, error) {
    if m.GetElementType() == ElementUnknown || m.data == nil {
        return nil, ErrorMatrixIsEmpty;
    }
    var i_data interface{} = nil;
    size := m.GetRowNum() * m.GetColumeNum();
    if m.GetElementType() == ElementFloat64{
        new_data := make([]float64, size, size);
        data, err := m.GetFloat64();
        if err != nil {
            return nil, err;
        }
        for i := 0; i < size; i++ {
            new_data[i] = math.Log(data[i]);
        }
        i_data = interface{}(new_data);
    } else if m.GetElementType() ==ElementComplex128 {
        new_data := make([]complex128, size, size);
        data, err := m.GetComplex128();
        if err != nil {
            return nil, err;
        }
        for i := 0; i < size; i++ {
            new_data[i] = cmplx.Log(data[i]);
        }
        i_data = interface{}(new_data);
    }

    new_matrix, err := NewMatrixWithData(m.GetRowNum(), m.GetColumeNum(), i_data);
    return new_matrix, err;
}

/*
func (m *Matrix) Log2() *Matrix {
    size := m.GetRowNum() * m.GetColumeNum();
    new_data := make([]float64, size, size);
    data := m.GetData();
    for i := 0; i < size; i++ {
        new_data[i] = math.Log2(data[i]);
    }
    lm := NewMatrix(m.GetRowNum(), m.GetColumeNum(), new_data);
    return lm;
}

func (m *Matrix) Log10() *Matrix {
    size := m.GetRowNum() * m.GetColumeNum();
    new_data := make([]float64, size, size);
    data := m.GetData();
    for i := 0; i < size; i++ {
        new_data[i] = math.Log10(data[i]);
    }
    lm := NewMatrix(m.GetRowNum(), m.GetColumeNum(), new_data);
    return lm;
}

func (m *Matrix) Log1p() *Matrix {
    size := m.GetRowNum() * m.GetColumeNum();
    new_data := make([]float64, size, size);
    data := m.GetData();
    for i := 0; i < size; i++ {
        new_data[i] = math.Log1p(data[i]);
    }
    lm := NewMatrix(m.GetRowNum(), m.GetColumeNum(), new_data);
    return lm;
}

func (m *Matrix) Sqrt() *Matrix {
    size := m.GetRowNum() * m.GetColumeNum();
    new_data := make([]float64, size, size);
    data := m.GetData();
    for i := 0; i < size; i++ {
        new_data[i] = math.Sqrt(data[i]);
    }
    lm := NewMatrix(m.GetRowNum(), m.GetColumeNum(), new_data);
    return lm;
}

func (m *Matrix) Pow(exp float64) *Matrix {
    size := m.GetRowNum() * m.GetColumeNum();
    new_data := make([]float64, size, size);
    data := m.GetData();
    for i := 0; i < size; i++ {
        new_data[i] = math.Pow(data[i], exp);
    }
    lm := NewMatrix(m.GetRowNum(), m.GetColumeNum(), new_data);
    return lm;
}

*/
