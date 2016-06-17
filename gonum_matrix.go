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
    m.data = make([]float64, size, size);
    copy(m.data, data);
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
	i := row * m.col_num + col;
	return m.data[i]
}

func (m *Matrix) Set(row, col int, v float64) {
    i := row + m.row_num*col;
    m.data[i] = v;
}

func (m *Matrix) getPtr() *[]float64 {
    return &(m.data);
}

func (m *Matrix) GetRow(row int) []float64 {
    if row >= m.GetRowNum() {
        panic("GetRow, row index large than row_num");
    }
    i := row * m.row_num;
    return m.data[i : i + m.col_num];
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
    for r := 0; r < m.row_num; r++ {
        for _, c := range cols {
            index := r * m.col_num + c;
            data = append(data, m.data[index]);
        }
    }
    sub := NewMatrix(m.row_num, len(cols), data);
    return sub;
}

func (m *Matrix) SwapRow(i, j int) {
    if i >= m.row_num || j >= m.row_num {
        panic("SwapRow(), i or j large than row nums");
    }
    row_i := i * m.col_num;
    row_j := j * m.col_num;
    for c := 0; c < m.col_num; c++ {
        m.data[row_i + c], m.data[row_j + c] = m.data[row_j + c], m.data[row_i + c];
    }
}

func (m *Matrix) RBind(other *Matrix) *Matrix {
    if m.GetColumeNum() != other.GetColumeNum() {
        panic("RBind(), 1st matrix's colume num NOT equal to 2nd matrix");
    }
    data := append(m.data, other.data...);
    matrix := NewMatrix(m.GetRowNum() + other.GetRowNum(), m.GetColumeNum(), data);
    return matrix;
}

func (m *Matrix) CBind(other *Matrix) *Matrix {
    if m.GetRowNum() != other.GetRowNum() {
        panic("RBind(), 1st matrix's row num NOT equal to 2nd matrix");
    }
    size := m.GetElementNum() + other.GetElementNum();
    data := make([]float64, 0, size);
    c1 := m.GetColumeNum();
    c2 := other.GetColumeNum();
    i1 := 0;
    i2 := 0;
    for r := 0; r < m.GetRowNum(); r += 1 {
        data = append(data, m.data[i1 : i1 + c1]...);
        data = append(data, other.data[i2 : i2 + c2]...);
        i1 = i1 + c1;
        i2 = i2 + c2;
    }
    matrix := NewMatrix(m.GetRowNum(), c1 + c2, data);
    return matrix;
}

func (m *Matrix) String() string {
    var s string;
    for r := 0; r < m.row_num; r++ {
        for c := 0; c < m.col_num; c++ {
            s = s + fmt.Sprintf("%0.4f, ", m.data[r * m.col_num + c]);
        }
        s += "\n";
    }
    return s;
}
