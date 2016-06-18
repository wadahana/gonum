package gonum

import (
    "fmt"
)


type Matrix struct {
    row_num int;
    col_num int;
    data []float64;
}

func NewMatrix(rows, cols int, data []float64) *Matrix {
    m := Matrix{row_num : rows, col_num : cols};
    size := rows * cols;
    if size != len(data) {
        panic("NewMatrix, row * col not equal to length of data.")
    }
    // m.data = make([]float64, size, size);
    // copy(m.data, data);
    m.data = data;
    return &m;
}

func (m *Matrix) Copy() *Matrix {
    c := NewMatrix(m.row_num, m.col_num, m.data);
	return c;
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

func (m *Matrix) Get(row, col int) float64 {
	i := col * m.row_num + row;
	return m.data[i]
}

func (m *Matrix) Set(row, col int, v float64) {
    i := col * m.row_num + row;
    m.data[i] = v;
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

func (m *Matrix) String() string {
    var s string;
    for r := 0; r < m.row_num; r++ {
        for c := 0; c < m.col_num; c++ {
            s = s + fmt.Sprintf("%0.4f, ", m.data[c * m.row_num + r]);
        }
        s += "\n";
    }
    return s;
}
