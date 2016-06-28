package gonum

import (
    "math"
)

func (this *Xts) Diff(lag int) *Xts {
    row_num := this.GetRowNum();
    col_num := this.GetColumeNum();

    data := make([]float64, 0, 0);
    nan := math.NaN();

    for c := 0; c < col_num; c++ {

        for i := 0; i < lag; i++ {
            data = append(data, nan);
        }

        col := this.GetColumeData(c);

        for i := lag ; i < row_num; i++ {
            v1, v2 := col[i-lag], col[i];
            data = append(data, (v2 - v1));
        }
    }
    dm := NewMatrix(row_num, col_num, data);
    xts := NewXts(this.index, dm);
    return xts;
}
