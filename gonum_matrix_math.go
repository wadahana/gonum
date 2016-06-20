package gonum

import (
    //"fmt"
    "math"
)

func (m *Matrix) Diff(lag int) *Matrix {
    row_num := m.GetRowNum();
    col_num := m.GetColumeNum();

    if lag >= row_num {
        panic("matrix.diff, lag large than row_num.");
    }
    size := (row_num - lag) * m.GetColumeNum();
    data := make([]float64, 0, size);

    for c := 0; c < col_num; c++ {
        col := m.GetColumeData(c);
        for i := 0 ; i < row_num - lag; i++ {
            v1, v2 := col[i], col[i+lag];
            data = append(data, (v2 - v1));
        }
    }
    dm := NewMatrix(row_num - lag, col_num, data);
    return dm;
}

func (m *Matrix) Log() *Matrix {
    size := m.GetRowNum() * m.GetColumeNum();
    new_data := make([]float64, size, size);
    data := m.GetData();
    for i := 0; i < size; i++ {
        new_data[i] = math.Log(data[i]);
    }
    lm := NewMatrix(m.GetRowNum(), m.GetColumeNum(), new_data);
    return lm;
}

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
