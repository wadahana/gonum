package gonum

import (
    "fmt"
//    "time"
    "testing"
)


var dates = [10]string {
    "2015-10-01", //7
    "2014-01-01", //4
    "2012-11-04", //3
    "2015-07-05", //6
    "2014-01-31", //5
    "2015-10-12", //8
    "2016-01-18", //9
    "2011-04-04", //2
    "2016-01-28", //10
    "2010-12-31", //1
};

var col1 = [] float64 {
    7.0,
    4.0,
    3.0,
    6.0,
    5.0,
    8.0,
    9.0,
    2.0,
    10.0,
    1.0,
};

var col2 = [] float64 {
    253.3,
    334.4,
    43.5,
    50.6,
    6.17,
    74.3,
    8.8,
    9.901,
    10.2334,
    11004.1234,
};

var col3 = [] float64 {
    1.3,
    2.4,
    3.5,
    4.6,
    5.17,
    6.3,
    7.8,
    8.901,
    9.2334,
    10.1234,
};

func Test_Matrix(t *testing.T) {
    data := make([]float64, 0);
    for i := 0; i < 10; i++ {
        data = append(data, col1[i]);
        data = append(data, col2[i]);
        data = append(data, col3[i]);
    }
    m1 := NewMatrix(10,3,data);

    fmt.Printf("Raw:\n");
    fmt.Printf("%v\n", m1);

    cols := []int{0,2};
    m2 := m1.GetColumes(cols);

    fmt.Printf("Extract Submatrix\n");
    fmt.Printf("%v\n", m2);

    m4 := m1.RBind(m1);
    fmt.Printf("After RBind:\n");
    fmt.Printf("%v\n", m4);

    m3 := m1.CBind(m2);
    fmt.Printf("After CBind:\n");
    fmt.Printf("%v\n", m3);

    m1.SwapRow(1,9);
    fmt.Printf("After Swap(1,9):\n");
    fmt.Printf("%v\n", m1);

}
