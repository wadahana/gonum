package gonum

import (
    "fmt"
    "time"
    "sort"
)

type Xts struct {

    index   []time.Time;
    names   []string;
    matrix  *Matrix;

}

// define interface{} for sort.
func (this Xts) Len() int {
    return len(this.index);
}
func (this Xts) Less(i, j int) bool {
    t1 := this.index[i];
    t2 := this.index[j];
    return t1.Before(t2);
}

func (this Xts) Swap(i, j int) {
    m := this.matrix;
    this.index[i], this.index[j] = this.index[j], this.index[i];
    m.SwapRow(i, j);
}


func NewXts(index []time.Time, m *Matrix) *Xts {
    rows := m.GetRowNum();
    cols := m.GetColumeNum();
    if len(index) != rows {
        panic("NewXts, length of index not equal to matrix row");
    }
    xts := Xts{};
    xts.index = make([]time.Time, rows, rows);
    copy(xts.index, index);
    xts.names = make([]string, cols, cols);
    for i := 0; i < cols; i++ {
        xts.names[i] = fmt.Sprintf("Colume.%d", i);
    }
    xts.matrix = m.Copy();

    sort.Sort(xts);
    return &xts;
}

func (this* Xts) String() string {
    var s string;
    m := this.matrix;
    for r := 0; r < m.row_num; r++ {
        s = s + this.index[r].Format("2006-01-02 15:04:05") + ", ";
        for c := 0; c < m.col_num; c++ {
            s = s + fmt.Sprintf("%0.4f, ", m.data[r * m.col_num + c]);
        }
        s += "\n";
    }
    return s;
}

func (this* Xts)GetRowNum() int {
    return this.matrix.GetRowNum();
}

func (this* Xts) GetColumeNum() int {
    return this.matrix.GetColumeNum();
}

func (this* Xts) GetNames() []string {
    return this.names;
}

func (this* Xts) SetNames(n ... string) {
    if len(n) < len(this.names) {
        for i := 0; i < len(n); i++ {
            this.names[i] = n[i];
        }
    }
}
