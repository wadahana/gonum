package gonum

import (
//    "fmt"
//    "time"
    "testing"
//    "reflect"
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
    403.5,
    500.6,
    677.1,
    744.3,
    801.8,
    933.1,
    108.2,
    112.1,
};

var col3 = [] float64 {
    1.0,
    2.0,
    3.0,
    4.0,
    5.0,
    6.0,
    7.0,
    8.0,
    9.0,
    10.0,
};

func testcase_NewMatrix() *Matrix {
    data := make([]float64, 0);
    for i := 0; i < 10; i++ {
        data = append(data, col1[i]);
        data = append(data, col2[i]);
        data = append(data, col3[i]);
    }
    m := NewMatrix(10,3,data);
    return m;
}

func Test_NewMatrix(t *testing.T) {

    m := testcase_NewMatrix();

    if m.GetRowNum() != 10 ||
        m.GetColumeNum() != 3 {
        t.Errorf("new matrix's row num or col num error.");
        return;
    }
    for r := 0; r < 10; r++ {
        if m.Get(r, 0) != col1[r] {
            t.Errorf("\nxxx  : %f, %f\n", m.Get(0,r), col1[r]);
            t.Errorf("1st colume data error, row index: %d\n", r);
            return;
        }
    }

    for r := 0; r < 10; r++ {
        if m.Get(r, 1) != col2[r] {
            t.Errorf("2nd colume data error, row index: %d\n", r);
            return;
        }
    }

    for r := 0; r < 10; r++ {
        if m.Get(r, 2) != col3[r] {
            t.Errorf("3rd colume data error, row index: %d\n", r);
            return;
        }
    }

    t.Logf("\n%v[%d,%d]\n", m, m.GetRowNum(), m.GetColumeNum());
    return;
}

func Test_ExtractColume(t *testing.T) {

    m := testcase_NewMatrix();

    cols := []int{0,2};
    sub := m.GetColumes(cols);

    if sub.GetRowNum() != 10 ||
        sub.GetColumeNum() != 2 {
        t.Errorf("sub matrix's row num or col num error.");
        return;
    }

    for r := 0; r < 10; r++ {
        if sub.Get(r, 0) != col1[r] {
            t.Errorf("1st colume data error, row index: %d\n", r);
            return;
        }
    }

    for r := 0; r < 10; r++ {
        if sub.Get(r, 1) != col3[r] {
            t.Errorf("2nd colume data error, row index: %d\n", r);
            return;
        }
    }
    t.Logf("\n%v[%d,%d]\n", sub, sub.GetRowNum(), sub.GetColumeNum());
    return ;
}

func Test_SwapRow(t *testing.T) {

    m := testcase_NewMatrix();
    m.SwapRow(1,6);
    if m.Get(1,0) != col1[6] ||
        m.Get(1,1) != col2[6] ||
        m.Get(1,2) != col3[6] {
        t.Errorf("swap row error.");
    }

    if m.Get(6,0) != col1[1] ||
        m.Get(6,1) != col2[1] ||
        m.Get(6,2) != col3[1] {
        t.Errorf("swap row error.");
    }

    t.Logf("\n%v[%d,%d]\n", m, m.GetRowNum(), m.GetColumeNum());
    return
}

func Test_RBind(t *testing.T) {

    m := testcase_NewMatrix();

    m1 := m.RBind(m);
    if m1.GetRowNum() != 20 || m1.GetColumeNum() != 3 {
        t.Errorf("new matrix's row num or col num error, after RBind");
    }

    t.Logf("\n%v[%d,%d]\n", m1, m1.GetRowNum(), m1.GetColumeNum());
}

func Test_CBind(t *testing.T) {

    cols := []int{0,2};

    m := testcase_NewMatrix();
    sub := m.GetColumes(cols);

    m2 := m.CBind(sub);
    if m2.GetRowNum() != 10 || m2.GetColumeNum() != 5 {
        t.Errorf("new matrix's row num or col num error, after CBind");
    }

    t.Logf("\n%v[%d,%d]\n", m2, m2.GetRowNum(), m2.GetColumeNum());
    return;
}
