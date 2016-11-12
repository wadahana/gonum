package gonum

import (
    "fmt"
    "time"
    "sort"
    "reflect"
    "math"
    "math/cmplx"
)

type Xts struct {

    index           []time.Time;
    colume_name     []string;
    name_map        map[string]int;
    matrix          *Matrix;

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

/* * * * * * * * * * * * Xts * * * * * * * * * * * * * */

func NewXts(index []time.Time, m *Matrix) (*Xts, error) {
    row_nums := m.GetRowNum();
    col_nums := m.GetColumeNum();
    if len(index) != row_nums {
        return nil, ErrorDimUnmatched;
    }
    if m == nil || m.GetElementType() == ElementUnknown {
        return nil, ErrorElementTypeNotSet
    }

    xts := Xts{};
    xts.index = make([]time.Time, row_nums, row_nums);
    copy(xts.index, index);
    l_name := make([]string, col_nums, col_nums);
    for i := 0; i < col_nums; i++ {
        l_name[i] = fmt.Sprintf("Colume.%d", i + 1);
    }
    xts.matrix = m;
    sort.Sort(xts);
    _ = xts.SetName(l_name...);
    return &xts, nil;
}

func (this *Xts) GetRowNum() int {
    return this.matrix.GetRowNum();
}

func (this *Xts) GetColumeNum() int {
    return this.matrix.GetColumeNum();
}

func (this *Xts) GetElementNum() int {
    return this.matrix.GetElementNum();
}

func (this *Xts) GetElementType() ElementType {
    return this.matrix.GetElementType();
}

func (this *Xts) GetTimeIndex() []time.Time {
    return this.index;
}

func (this *Xts) GetMatrix() *Matrix {
    return this.matrix;
}

func (this *Xts) GetName() []string {
    return this.colume_name;
}

func (this *Xts) SetName(l_name ... string) error {
    if isDuplicate(l_name) {
        return ErrorInvalidParameter;
    }
    col_nums := this.GetColumeNum();
    if this.colume_name == nil {
        this.colume_name = make([]string, col_nums, col_nums);
    } else if len(this.colume_name) < col_nums {
        colume_name := make([]string, col_nums, col_nums);
        copy(colume_name, this.colume_name);
        this.colume_name = colume_name;
    }
    if this.name_map == nil {
        this.name_map = make(map[string]int);
    }
    if len(l_name) <= len(this.colume_name) {
        for i := 0; i < len(l_name); i++ {
            name := l_name[i];
            this.name_map[name] = i;
            this.colume_name[i] = name;
        }
    }
    return nil;
}

func (this *Xts) Get(row, col int) interface{} {
    return this.matrix.Get(row, col);
}

func (this *Xts) Set(row, col int, v interface{}) error {
    return this.matrix.Set(row, col, v);
}

func (this *Xts) GetColumeData(col int) (interface{}, error) {
    return this.matrix.GetColumeData(col);
}

func (this *Xts) String() string {
    var s string = "\t\t\t";
    m := this.matrix;
    if m.element_type == ElementUnknown {
        return "";
    }
    for _, name := range this.colume_name {
        s = s + name + ",\t";
    }
    s = s + "\n";

    v := reflect.ValueOf(m.data);
    for r := 0; r < m.row_num; r++ {
        s = s + this.index[r].Format("2006-01-02 15:04:05") + ", ";
        for c := 0; c < m.col_num; c++ {
            s = s + elementToString(v.Index(c * m.row_num + r).Interface()) + ", ";
        }
        s += "\n";
    }
    s += fmt.Sprintf("Dimension: [%d x %d], Type: %v\n", m.row_num, m.col_num, m.element_type);
    return s;
}

func (this *Xts) GetColumesByIndex(l_col_index ... int) (*Xts, error) {

    sub_matrix, err := this.matrix.GetColumesByIndex(l_col_index...);
    if err != nil {
        return nil, err;
    }
    xts, err := NewXts(this.GetTimeIndex(), sub_matrix);
    return xts, err;
}

func (this *Xts) GetColumesByName(l_name ... string) (*Xts, error) {

    l_index := this.GetColumeIndexByName(l_name);
    return this.GetColumesByIndex(l_index...);
}

func (this *Xts) GetColumeIndexByName(l_name ... string) []int {
    l_index := make([]int, 0, 0);
    for _, name := range l_name {
        if v, ok := this.name_map[name]; ok {
            l_index = append(l_index, v);
        }
    }
    return l_index;
}

func (this *Xts) RBind(other *Xts) (*Xts, error) {
    element_type := this.GetElementType();
    if element_type != other.GetElementType() { // 2 xts have diff element type
        return nil, ErrorElementTypeUnmatched;
    } else if element_type == ElementUnknown {
        return nil, ErrorElementTypeNotSet;
    } else if this.GetColumeNum() != other.GetColumeNum() {
        return nil, ErrorDimUnmatched;
    }
    total_col_num := this.GetColumeNum();
    col_list := make([]reflect.Value, total_col_num, total_col_num);
    for i := 0; i < total_col_num; i++ {
        col_list[i] = reflect.MakeSlice(reflect.TypeOf(this.matrix.data), 0, 0);
    }


    time := make([]time.Time, 0);
    size1 := len(this.index);
    size2 := len(other.index);
    var j, k int = 0, 0;
    for j < size1 && k < size2 {
        if this.index[j].Equal(other.index[k]) {
            return nil, ErrorIndexConflict;
        } else if this.index[j].Before(other.index[k]) {
            for c := 0; c < total_col_num; c++ {
                v := this.Get(j, c);
                col_list[c] = reflect.Append(col_list[c], reflect.ValueOf(v));
//                cols[c] = append(cols[c], v);
            }
            time = append(time, this.index[j]);
            j += 1
        } else {
            for c := 0; c < total_col_num; c++ {
                v := other.Get(k, c);
                col_list[c] = reflect.Append(col_list[c], reflect.ValueOf(v));
//                cols[c] = append(cols[c], v);
            }
            time = append(time, other.index[k]);
            k += 1;
        }
    }
    if j < size1 {
        for i := 0; i < total_col_num; i++ {
            col_data, err := this.matrix.GetColumeData(i);
            if err != nil {
                return nil, err;
            }
            v_col_data := reflect.ValueOf(col_data);
            col_list[i] = reflect.AppendSlice(col_list[i], v_col_data.Slice(j, v_col_data.Len()));
//            cols[i] = append(cols[i], col[j:]...);
        }
        time = append(time, this.index[j:]...);
    }
    if k < size2 {
        for i := 0; i < total_col_num; i++ {
            col_data, err := other.matrix.GetColumeData(i);
            if err != nil {
                return nil, err;
            }
            v_col_data := reflect.ValueOf(col_data);
            col_list[i] = reflect.AppendSlice(col_list[i], v_col_data.Slice(k, v_col_data.Len()));
//            cols[i] = append(cols[i], col[k:]...);
        }
        time = append(time, other.index[k:]...);
    }

    data := col_list[0];
    for i := 1; i < total_col_num; i++ {
        data = reflect.AppendSlice(data, col_list[i]);
//        data = append(data, cols[i]...);
    }

    var m * Matrix = nil;
    var x * Xts = nil;
    var err error = nil;

    m, err = NewMatrixWithData(len(time), total_col_num, data.Interface());
    if err == nil {
        x, err = NewXts(time, m);
    }
    return x, err;
}

func (this *Xts) CBind(other* Xts) (*Xts, error) {
    element_type := this.GetElementType();
    if element_type != other.GetElementType() { // 2 xts have diff element type
        return nil, ErrorElementTypeUnmatched;
    } else if element_type == ElementUnknown {
        return nil, ErrorElementTypeNotSet;
    }

    col1_num := this.GetColumeNum();
    col2_num := other.GetColumeNum();
    total_col_num := col1_num + col2_num;
    col_list := make([]reflect.Value, total_col_num, total_col_num);
    for i := 0; i < total_col_num; i++ {
        col_list[i] = reflect.MakeSlice(reflect.TypeOf(this.matrix.data), 0, 0);
    }

    time := make([]time.Time, 0);
    size1 := len(this.index);
    size2 := len(other.index);

    var j, k int = 0, 0;
    for j < size1 && k < size2 {
        if this.index[j].Equal(other.index[k]) {
            for c := 0; c < col1_num; c++ {
                v := this.matrix.Get(j, c);
                col_list[c] = reflect.Append(col_list[c], reflect.ValueOf(v));
//                cols[c] = append(cols[c], v);
            }
            for c := 0; c < col2_num; c++ {
                v := other.matrix.Get(k, c);
                col_list[col1_num + c] = reflect.Append(col_list[col1_num + c], reflect.ValueOf(v));
//                cols[col1_num + c] = append(cols[col1_num + c], v);
            }
            time = append(time, this.index[j]);
            j += 1;
            k += 1;

        } else if this.index[j].Before(other.index[k]) {
            for c := 0; c < col1_num; c++ {
                v := this.matrix.Get(j, c);
                col_list[c] = reflect.Append(col_list[c], reflect.ValueOf(v));
//                cols[c] = append(cols[c], v);
            }
            for c := 0; c < col2_num; c++ {
                if element_type == ElementFloat64 {
                    col_list[col1_num + c] = reflect.Append(col_list[col1_num + c], reflect.ValueOf(RealNaN));
                } else {
                    col_list[col1_num + c] = reflect.Append(col_list[col1_num + c], reflect.ValueOf(ComplexNaN));
                }
//                cols[col1_num + c] = append(cols[col1_num + c], v_nan);
            }
            time = append(time, this.index[j]);
            j += 1;
        } else {
            for c := 0; c < col1_num; c++ {
                if element_type == ElementFloat64 {
                    col_list[c] = reflect.Append(col_list[c], reflect.ValueOf(RealNaN));
                } else {
                    col_list[c] = reflect.Append(col_list[c], reflect.ValueOf(ComplexNaN));
                }
//                cols[c] = append(cols[c], v_nan);
            }
            for c := 0; c < col2_num; c++ {
                v := other.matrix.Get(k, c);
                col_list[col1_num + c] = reflect.Append(col_list[col1_num + c], reflect.ValueOf(v));
//                cols[col1_num + c] = append(cols[col1_num + c], v);
            }
            time = append(time, other.index[k]);
            k += 1;
        }
    }
    if j < size1 {
        l_nan := createNaNSlice(element_type, size1 - j);
        for i := 0; i < col1_num; i++ {
            col_data, err := this.matrix.GetColumeData(i);
            if err != nil {
                return nil, err;
            }
            v_col_data := reflect.ValueOf(col_data);
            col_list[i] = reflect.AppendSlice(col_list[i], v_col_data.Slice(j, v_col_data.Len()));
//            cols[i] = append(cols[i], col[j:]...);
        }
        for i := 0; i < col2_num; i++ {
            col_list[col1_num + i] = reflect.AppendSlice(col_list[col1_num + i], reflect.ValueOf(l_nan));
 //           cols[col1_num + i] = append(cols[col1_num + i], l_nan...);
        }
        time = append(time, this.index[j:]...);
    }
    if k < size2 {
        l_nan := createNaNSlice(element_type, size2 - k);
        for i := 0; i < col1_num; i++ {
            col_list[i] = reflect.AppendSlice(col_list[i], reflect.ValueOf(l_nan));
//            cols[i] = append(cols[i], l_nan...);
        }
        for i := 0; i < col2_num; i++ {
            col_data, err := other.matrix.GetColumeData(i);
            if err != nil {
                return nil, err;
            }
            v_col_data := reflect.ValueOf(col_data);
            col_list[col1_num + i] = reflect.AppendSlice(col_list[col1_num + i], v_col_data.Slice(k, v_col_data.Len()));
//            cols[col1_num + i] = append(cols[col1_num + i], col[k:]...);
        }
        time = append(time, other.index[k:]...);
    }

    data := col_list[0];
    for i := 1; i < total_col_num; i++ {
        data = reflect.AppendSlice(data, col_list[i]);
//        data = append(data, cols[i]...);
    }

    var m * Matrix = nil;
    var x * Xts = nil;
    var err error = nil;

    m, err = NewMatrixWithData(len(time), total_col_num, data.Interface());
    if err == nil {
        x, err = NewXts(time, m);
    }

    return x, err;
}


func isDuplicate(l_name []string) bool {
    _map := make(map[string]int);
    for i, name := range l_name {
        if _, ok := _map[name]; ok {
            return true;
        } else {
            _map[name] = i;
        }
    }
    return false;
}

func createNaNSlice(element_type ElementType, size int) interface{} {
    var result interface{} = nil;
    if element_type == ElementFloat64 {
        v_nan := math.NaN();
        l_nan := make([]float64,0);
        for i := 0; i < size; i++ {
            l_nan = append(l_nan, v_nan);
        }
        result = interface{}(l_nan);
    } else if element_type == ElementComplex128 {
        v_nan := cmplx.NaN();
        l_nan := make([]complex128,0);
        for i := 0; i < size; i++ {
            l_nan = append(l_nan, v_nan);
        }
        result = interface{}(l_nan);
    }
    return result;
}



