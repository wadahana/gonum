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
	i := row + m.row_num*col;
	return m.data[i]
}

func (m *Matrix) Set(row, col int, v float64) {
    i := row + m.row_num*col;
    m.data[i] = v;
}

func (m *Matrix) GetRawData() []float64 {
    return m.data;
}

func (m *Matrix) SwapRow(i, j int) {
    if i >= m.row_num || j >= m.row_num {
        panic("SwapRow, i or j large than row nums");
    }
    row_i := i * m.col_num;
    row_j := j * m.col_num;
    for c := 0; c < m.col_num; c++ {
        m.data[row_i + c], m.data[row_j + c] = m.data[row_j + c], m.data[row_i + c];
    }
}

func (m* Matrix) String() string {
    var s string;
    for r := 0; r < m.row_num; r++ {
        for c := 0; c < m.col_num; c++ {
            s = s + fmt.Sprintf("%0.4f, ", m.data[r * m.col_num + c]);
        }
        s += "\n";
    }
    return s;
}
