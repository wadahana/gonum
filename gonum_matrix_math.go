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
    data := make([]float64, size, size);
    var i1 int = 0;
    var i2 int = 0;
    var lag2 int = lag * col_num;
    for row := 0; row < row_num - lag; row += 1 {
        for col := 0; col < col_num; col += 1 {
            v1 := m.data[i2];
            v2 := m.data[i2 + lag2];
            data[i1] = v2 - v1;
            i1 += 1;
            i2 += 1;
        }
    }
    dm := NewMatrix(row_num - lag, col_num, data);
    return dm;
}

func (m *Matrix) Log() *Matrix {
    size := m.GetRowNum() * m.GetColumeNum();
    d2 := make([]float64, size, size);
    d1 := m.GetData();
    for i := 0; i < size; i++ {
        d2[i] = math.Log(d1[i]);
    }
    lm := NewMatrix(m.GetRowNum(), m.GetColumeNum(), d2);
    return lm;
}

func (m *Matrix) Log2() *Matrix {
    size := m.GetRowNum() * m.GetColumeNum();
    d2 := make([]float64, size, size);
    d1 := m.GetData();
    for i := 0; i < size; i++ {
        d2[i] = math.Log2(d1[i]);
    }
    lm := NewMatrix(m.GetRowNum(), m.GetColumeNum(), d2);
    return lm;
}

func (m *Matrix) Log10() *Matrix {
    size := m.GetRowNum() * m.GetColumeNum();
    d2 := make([]float64, size, size);
    d1 := m.GetData();
    for i := 0; i < size; i++ {
        d2[i] = math.Log10(d1[i]);
    }
    lm := NewMatrix(m.GetRowNum(), m.GetColumeNum(), d2);
    return lm;
}

func (m *Matrix) Log1p() *Matrix {
    size := m.GetRowNum() * m.GetColumeNum();
    d2 := make([]float64, size, size);
    d1 := m.GetData();
    for i := 0; i < size; i++ {
        d2[i] = math.Log1p(d1[i]);
    }
    lm := NewMatrix(m.GetRowNum(), m.GetColumeNum(), d2);
    return lm;
}

func (m *Matrix) Sqrt() *Matrix {
    size := m.GetRowNum() * m.GetColumeNum();
    d2 := make([]float64, size, size);
    d1 := m.GetData();
    for i := 0; i < size; i++ {
        d2[i] = math.Sqrt(d1[i]);
    }
    lm := NewMatrix(m.GetRowNum(), m.GetColumeNum(), d2);
    return lm;
}

func (m *Matrix) Pow(exp float64) *Matrix {
    size := m.GetRowNum() * m.GetColumeNum();
    d2 := make([]float64, size, size);
    d1 := m.GetData();
    for i := 0; i < size; i++ {
        d2[i] = math.Pow(d1[i], exp);
    }
    lm := NewMatrix(m.GetRowNum(), m.GetColumeNum(), d2);
    return lm;
}
