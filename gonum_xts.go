package gonum

import (
    "fmt"
    "time"
    "sort"
    "math"
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
    xts := newXtsNoSort(index, m);
    sort.Sort(xts);
    return xts;
}

func newXtsNoSort(index []time.Time, m *Matrix) *Xts {
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
//    xts.matrix = m.Copy();
    xts.matrix = m;

    sort.Sort(xts);
    return &xts;
}

func (this *Xts) GetRowNum() int {
    return this.matrix.GetRowNum();
}

func (this *Xts) GetColumeNum() int {
    return this.matrix.GetColumeNum();
}

func (this *Xts) GetNames() []string {
    return this.names;
}

func (this *Xts) SetNames(n ... string) {
    if len(n) <= len(this.names) {
        for i := 0; i < len(n); i++ {
            this.names[i] = n[i];
        }
    }
}

func (this *Xts) GetColumeData(col int) []float64 {
    return this.matrix.GetColumeData(col);
}
func (this *Xts) Get(row, col int) float64 {
    return this.matrix.Get(row, col);
}

func (this *Xts) Set(row, col int, v float64) {
    this.matrix.Set(row, col, v);
}

func (this *Xts) GetColumes(cols []int) *Xts {

    sub := this.matrix.GetColumes(cols);

    xts := Xts{};
    xts.matrix = sub;
    xts.index = make([]time.Time, sub.row_num, sub.row_num);
    copy(xts.index, this.index);
    xts.names = make([]string, sub.col_num, sub.col_num);
    for i, c := range cols {
        xts.names[i] = this.names[c];
    }
    return &xts;
}

func (this *Xts) RBind(other *Xts) *Xts {
    if this.GetColumeNum() != other.GetColumeNum() {
        panic("two xts dont have same colume num.");
    }
    col_num := this.GetColumeNum();
    cols := make([][]float64, col_num, col_num);
    for i := 0; i < col_num; i++ {
        cols[i] = make([]float64, 0);
    }

    time := make([]time.Time, 0);
    size1 := len(this.index);
    size2 := len(other.index);
    var j, k int = 0, 0;
    for j < size1 && k < size2 {
        if this.index[j].Equal(other.index[k]) {
            return nil;
        } else if this.index[j].Before(other.index[k]) {
            for c := 0; c < col_num; c++ {
                v := this.matrix.Get(j, c);
                cols[c] = append(cols[c], v);
            }
            time = append(time, this.index[j]);
            j += 1
        } else {
            for c := 0; c < col_num; c++ {
                v := other.matrix.Get(k, c);
                cols[c] = append(cols[c], v);
            }
            time = append(time, other.index[k]);
            k += 1;
        }
    }
    if j < size1 {
        for i := 0; i < col_num; i++ {
            col := this.matrix.GetColumeData(i);
            cols[i] = append(cols[i], col[j:]...);
        }
        time = append(time, this.index[j:]...);
    }
    if k < size2 {
        for i := 0; i < col_num; i++ {
            col := other.matrix.GetColumeData(i);
            cols[i] = append(cols[i], col[k:]...);
        }
        time = append(time, other.index[j:]...);
    }
    data := cols[0];
    for i := 1; i < col_num; i++ {
        data = append(data, cols[i]...);
    }

    m := NewMatrix(len(time), col_num, data);
    xts := newXtsNoSort(time, m);
    return xts;
}

func (this *Xts) CBind(other* Xts) *Xts {
    col1_num := this.GetColumeNum();
    col2_num := other.GetColumeNum();
    col_num := col1_num + col2_num;
    cols := make([][]float64, col_num, col_num);
    for i := 0; i < col_num; i++ {
        cols[i] = make([]float64, 0);
    }

    time := make([]time.Time, 0);
    size1 := len(this.index);
    size2 := len(other.index);
    v_nan := math.NaN();
    var j, k int = 0, 0;
    for j < size1 && k < size2 {
        if this.index[j].Equal(other.index[k]) {
            for c := 0; c < col1_num; c++ {
                v := this.matrix.Get(j, c);
                cols[c] = append(cols[c], v);
            }
            for c := 0; c < col2_num; c++ {
                v := other.matrix.Get(k, c);
                cols[col1_num + c] = append(cols[col1_num + c], v);
            }
            time = append(time, this.index[j]);
            j += 1;
            k += 1;

        } else if this.index[j].Before(other.index[k]) {
            for c := 0; c < col1_num; c++ {
                v := this.matrix.Get(j, c);
                cols[c] = append(cols[c], v);
            }
            for c := 0; c < col2_num; c++ {
                cols[col1_num + c] = append(cols[col1_num + c], v_nan);
            }
            time = append(time, this.index[j]);
            j += 1;
        } else {
            for c := 0; c < col1_num; c++ {
                cols[c] = append(cols[c], v_nan);
            }
            for c := 0; c < col2_num; c++ {
                v := other.matrix.Get(k, c);
                cols[col1_num + c] = append(cols[col1_num + c], v);
            }
            time = append(time, other.index[k]);
            k += 1;
        }
    }
    if j < size1 {
        l_nan := createNaNSlice(size1 - j);
        for i := 0; i < col1_num; i++ {
            col := this.matrix.GetColumeData(i);
            cols[i] = append(cols[i], col[j:]...);
        }
        for i := 0; i < col2_num; i++ {
            cols[col1_num + i] = append(cols[col1_num + i], l_nan...);
        }
        time = append(time, this.index[j:]...);
    }
    if k < size2 {
        l_nan := createNaNSlice(size2 - k);
        for i := 0; i < col1_num; i++ {
            cols[i] = append(cols[i], l_nan...);
        }
        for i := 0; i < col2_num; i++ {
            col := other.matrix.GetColumeData(i);
            cols[col1_num + i] = append(cols[col1_num + i], col[k:]...);
        }
        time = append(time, other.index[k:]...);
    }

    data := cols[0];
    for i := 1; i < col_num; i++ {
        data = append(data, cols[i]...);
    }

    m := NewMatrix(len(time), col_num, data);
    xts := newXtsNoSort(time, m);
    return xts;
}

func createNaNSlice(size int) []float64 {
    v_nan := math.NaN();
    l_nan := make([]float64,0);
    for i := 0; i < size; i++ {
        l_nan = append(l_nan, v_nan);
    }
    return l_nan;
}

func (this *Xts) String() string {
    var s string;
    m := this.matrix;
    for _, name := range this.names {
        s = s + name + ", ";
    }
    s = s + "\n";

    for r := 0; r < m.row_num; r++ {
        s = s + this.index[r].Format("2006-01-02 15:04:05") + ", ";
        for c := 0; c < m.col_num; c++ {
            s = s + fmt.Sprintf("%15.4f, ", m.data[c * m.row_num + r]);
        }
        s += "\n";
    }
    return s;
}

/* * *
  check two xts has the same time index.
* * */
// func (this *Xts) HasSameTime(other *Xts) bool {
//     t1 := this.index;
//     t2 := other.index;
//     size := len(t1);
//     if size == len(t2) {
//         i := 0;
//         j := 0;
//         for i < size && j < size {
//             if t1[i].Equal(t2[j]) {
//                 return true;
//             } else if t1[i].Before(t2.[j]) {
//                 i += 1
//             } else {
//                 j += 1;
//             }
//         }
//     }
//     return false;
// }
