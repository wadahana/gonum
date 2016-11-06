package gonum

import (
//    "fmt"
//    "time"
//    "math"
    "testing"
//    "reflect"
)

var matrix_data = [] float64 {
    0.1, 0.2, 0.3, 0.4, 0.5,
    1.1, 1.2, 1.3, 1.4, 1.5,
    2.1, 2.2, 2.3, 2.4, 2.5,
    3.1, 2.2, 3.3, 3.4, 3.5,
};

var fmatrix_data_A = []float64 {
    1.0,  2.0,
    3.0,  4.0,
    11.0, 12.0,
    13.0, 14.0,
};

var fmatrix_data_B = []float64 {
    11.0, 12.0, 21.0, 22.0, 31.0,
    32.0, 41.0, 42.0, 51.0, 52.0,
};

var cmatrix_data_A = []complex128 {
    1.0 + 1.0i, 2.0 + 1.0i,
    3.0 + 1.0i, 4.0 + 1.0i,
    5.0 + 1.0i, 6.0 + 1.0i,
    7.0 + 1.0i, 8.0 + 1.0i,
};

var cmatrix_data_B = []complex128 {
    10.0+10.0i, 20.0+10.0i, 30.0+10.0i, 40.0+10.0i, 50.0+10.0i,
    60.0+10.0i, 70.0+10.0i, 80.0+10.0i, 90.0+10.0i, 100.0+10.0i,
};


func Test_Matrix(t *testing.T) {
    m0 := NewEmptyMatrix();
    t.Logf("create empty matrix: \n%v\n", m0);

    m1 := NewMatrix(5, 4, ElementUnknown);
    t.Logf("create empty matrix with dim(5,4): \n%v\n", m1);

    m2, err := NewMatrixWithData(5, 4, matrix_data);
    if err == nil {
       t.Logf("create matrix with matrix_data, dim(5,4): \n%v\n", m2);
    } else {
        t.Errorf("NewMatrixWithData fail, Error: %v", err);
        return;
    }

    err = m1.Set(2,2, 100.2);
    if err == nil {
        v := m1.Get(2,2).(float64);
        if v != 100.2 {
            t.Errorf("Empty Matrix's Set/Get test fail");
        }
    } else {
        t.Errorf("Empty Matrix.Set test fail, Error; %v\n", err);
        return;
    }

    err = m2.Set(1, 1, float64(100.1));
    if err == nil {
        v := m2.Get(1,1).(float64);
        if v != 100.1 {
            t.Errorf("Matrix's Set/Get test fail");
        }
    } else {
        t.Errorf("Matrix.Set test fail, Error: %v\n", err);
        return;
    }

    m3, err := m1.CBind(m2);
    if err != nil {
        t.Errorf("CBind fail, Error: %v\n", err);
        return ;
    }
    t.Logf("\n%v\n", m3);

    m4, err := m1.RBind(m2);
    if err != nil {
        t.Errorf("RBind fail, Error: %v\n", err);
        return ;
    }
    t.Logf("\n%v\n", m4);

    col_data, err := m4.GetColumeData(2);
    if err != nil {
        t.Errorf("Matrix.GetColumeData fail, Error: %v\n", err);
        return ;
    }
    t.Logf("\n%v \n", col_data);


    m1, err = NewMatrixWithData(4, 2, fmatrix_data_A);
    if err == nil {
        m2, err = NewMatrixWithData(2, 5, fmatrix_data_B);
    }
    if err == nil {
        m3, err = m1.Product(1.0, m2);
        if err != nil {
            t.Errorf("Matrix.Profuct fail, Error: %v\n", err);
            return ;
        }
        t.Logf("\n%v\n", m3);

        m3, err = m1.Multiply(2.0);
        if err != nil {
            t.Errorf("Matrix.Multiply fail, Error: %v\n", err);
            return ;
        }
        t.Logf("\n%v\n", m3);

        m3, err = m1.Add(2000.0);
        if err != nil {
            t.Errorf("Matrix.Add fail, Error: %v\n", err);
            return ;
        }
        t.Logf("\n%v\n", m3);

        m4, err = m1.Add(m3);
        if err != nil {
            t.Errorf("Matrix.Add fail, Error: %v\n", err);
            return ;
        }
        t.Logf("\n%v\n", m4);
    }

    m1, err = NewMatrixWithData(4, 2, cmatrix_data_A);
    if err == nil {
        m2, err = NewMatrixWithData(2, 5, cmatrix_data_B);
    }
    if err == nil {
        m3, err = m1.Product(1.0i, m2);
        if err != nil {
            t.Errorf("Matrix.Profuct fail, Error: %v\n", err);
            return ;
        }
        t.Logf("\n%v\n", m3);

        m3, err = m1.Multiply(2.0i);
        if err != nil {
            t.Errorf("Matrix.Multiply fail, Error: %v\n", err);
            return ;
        }
        t.Logf("\n%v\n", m3);

        m3, err = m1.Add(2000.0i);
        if err != nil {
            t.Errorf("Matrix.Add fail, Error: %v\n", err);
            return ;
        }
        t.Logf("\n%v\n", m3);

        m4, err = m1.Add(m3);
        if err != nil {
            t.Errorf("Matrix.Add fail, Error: %v\n", err);
            return ;
        }
        t.Logf("\n%v\n", m4);
    }


    return;
}
/*
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

    data = append(data, col1...);
    data = append(data, col2...);
    data = append(data, col3...);

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

    for c := 0; c < m1.GetColumeNum(); c++ {
        col := m1.GetColumeData(c);
        raw_col := m.GetColumeData(c);
        for i, v := range col {
            if v != raw_col[i % 10] {
                t.Errorf("new matrix's data error, [%d, %d]", i, c);
                return;
            }
        }
    }
    t.Logf("\n%v[%d,%d]\n", m1, m1.GetRowNum(), m1.GetColumeNum());
}

func Test_CBind(t *testing.T) {

    cols := []int{0,2};

    m := testcase_NewMatrix();
    sub := m.GetColumes(cols);

    m1 := m.CBind(sub);
    if m1.GetRowNum() != 10 || m1.GetColumeNum() != 5 {
        t.Errorf("new matrix's row num or col num error, after CBind");
    }

    for c := 0; c < 3; c++ {
        new_col := m1.GetColumeData(c);
        raw_col := m.GetColumeData(c);
        for i, v := range new_col {
            if v != raw_col[i] {
                t.Errorf("new matrix's data error, [%d, %d]", i, c);
                return;
            }
        }
    }

    for c := 3; c < 5; c++ {
        new_col := m1.GetColumeData(c);
        raw_col := sub.GetColumeData(c-3);
        for i, v := range new_col {
            if v != raw_col[i] {
                t.Errorf("new matrix's data error, [%d, %d]", i, c);
                return;
            }
        }
    }

    t.Logf("\n%v[%d,%d]\n", m1, m1.GetRowNum(), m1.GetColumeNum());
    return;
}


func Test_Math(t *testing.T) {
    m := testcase_NewMatrix();

    t.Log("Test matrix.Diff(1)\n");
    dm := m.Diff(1);
    if dm.GetColumeNum() != 3 || dm.GetRowNum() != 9 {
        t.Errorf("after Matrix.Diff, col num or row num is error.");
        return;
    }

    dm1_ok := [27]float64 {-3.0, -1.0, 3.0, -1.0, 3.0, 1, -7.0, 8.0, -9.0, 81.1, 69.1, 97.1, 176.5, 67.2, 57.5, 131.3, -824.9, 3.9, 1, 1, 1, 1, 1, 1, 1, 1, 1};
    d1 := dm.GetData();

    for i, v := range d1 {
        if math.Abs(dm1_ok[i] - v) > 0.0000001 {
            t.Errorf("after matrix.Diff, data error");
            return;
        }
    }
    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    t.Log("Test matrix.Diff(6)\n");
    dm = m.Diff(6);
    if dm.GetColumeNum() != 3 || dm.GetRowNum() != 4 {
        t.Errorf("after Matrix.Diff, col num or row num is error.");
        return;
    }
    dm6_ok := [12]float64{ 2, -2, 7, -5, 548.5, 598.7, -295.3, -388.5, 6, 6, 6, 6};
    d6 := dm.GetData();
    for i, v := range d6 {
        if math.Abs(dm6_ok[i] - v) > 0.0000001 {
            t.Errorf("after matrix.Diff, data error");
            return;
        }
    }

    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    t.Log("Test matrix.Log2()\n");
    dm = m.Log2();

    lg2_ok := [30]float64 {2.807355,2.000000,1.584963,2.584963,2.321928,3.000000,3.169925,1.000000,3.321928,0.000000,
                           7.984703,8.385431,8.656425,8.967514,9.403225,9.539740,9.647099,9.865888,6.757557,6.808642,
                           0.000000,1.000000,1.584963,2.000000,2.321928,2.584963,2.807355,3.000000,3.169925,3.321928};
    dlg2 := dm.GetData();
    for i, v := range dlg2 {
        if math.Abs(v - lg2_ok[i]) > 0.000001 {
            t.Errorf("matrix log2's data error")
            return;
        }
    }

    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    t.Log("Test matrix.Log10()\n");
    dm = m.Log10();
    lg10_ok := [30]float64 {0.8450980, 0.6020600, 0.4771213, 0.7781513, 0.6989700, 0.9030900, 0.9542425, 0.3010300, 1.0000000, 0.0000000,
                            2.4036350, 2.5242660, 2.6058440, 2.6994910, 2.8306530, 2.8717480, 2.9040660, 2.9699280, 2.0342270, 2.0496060,
                            0.0000000, 0.3010300, 0.4771213, 0.6020600, 0.6989700, 0.7781513, 0.8450980, 0.9030900, 0.9542425, 1.0000000};
    dlg10 := dm.GetData();
    for i, v := range dlg10 {
        if math.Abs(v - lg10_ok[i]) > 0.000001 {
            t.Errorf("matrix log10's data error");
            return;
        }
    }

    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    t.Log("Test matrix.Log1p()\n");
    dm = m.Log1p();
    lg1p_ok := [30]float64 {2.0794415, 1.6094379, 1.3862944, 1.9459101, 1.7917595, 2.1972246, 2.3025851, 1.0986123, 2.3978953, 0.6931472,
                            5.5385150, 5.8153240, 6.0026520, 6.2178030, 6.5192950, 6.6137870, 6.6881060, 6.8395830, 4.6931810, 4.7282720,
                            0.6931472, 1.0986123, 1.3862944, 1.6094379, 1.7917595, 1.9459101, 2.0794415, 2.1972246, 2.3025851, 2.3978953};
    dlg1p := dm.GetData();
    for i, v := range dlg1p {
        if math.Abs(v - lg1p_ok[i]) > 0.000001 {
            t.Errorf("matrix log1p's data error");
            return;
        }
    }
    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    t.Log("Test matrix.Log()\n");
    dm = m.Log();
    lg_ok := [30]float64 {1.9459101, 1.3862944, 1.0986123, 1.7917595, 1.6094379, 2.0794415, 2.1972246, 0.6931472, 2.3025851, 0.0000000,
                         5.5345750, 5.8123380, 6.0001760, 6.2158070, 6.5178190, 6.6124440, 6.6868590, 6.8385120, 4.6839810, 4.7193910,
                         0.0000000, 0.6931472, 1.0986123, 1.3862944, 1.6094379, 1.7917595, 1.9459101, 2.0794415, 2.1972246, 2.3025851};
    dlg := dm.GetData();
    for i, v := range dlg {
        if math.Abs(v - lg_ok[i]) > 0.000001 {
            t.Errorf("matrix log's data error");
            return;
         }
    }
    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    t.Log("Test matrix.Sqrt()\n");
    dm = m.Sqrt();
    sqrt_ok := [30]float64{ 2.645751,  2.000000,  1.732051,  2.449490,  2.236068,  2.828427,  3.000000,  1.414214,  3.162278,  1.000000,
                           15.915400, 18.286610, 20.087310, 22.374090, 26.021150, 27.281860, 28.316070, 30.546690, 10.401920, 10.587730,
                            1.000000,  1.414214,  1.732051,  2.000000,  2.236068,  2.449490,  2.645751,  2.828427,  3.000000,  3.162278};
    dsqrt := dm.GetData();
    for i, v := range dsqrt {
        if math.Abs(v - sqrt_ok[i]) > 0.00001 {
            t.Errorf("matrix sqrt's data error");
            return;
         }
    }
    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    t.Log("Test matrix.Pow(4)\n");
    dm = m.Pow(4);
    dpow4 := dm.GetData();
    dpow4_ok := [30]float64 {      2401.0,         256.0,          81.0,        1296.0,          625.0,         4096.0,         6561.0,           16.0,     10000.0,           1.0,
                             4116619806.0, 12504463842.0, 26507828750.0, 62800540432.0, 210189615237.0, 306896599227.0, 413298860273.0, 758076017849.0, 137059468.0,   157914660.0,
                                      1.0,          16.0,          81.0,         256.0,          625.0,         1296.0,         2401.0,         4096.0,      6561.0,       10000.0};
    for i, v := range dpow4 {
        if math.Abs(v - dpow4_ok[i]) > 0.5 {
            t.Errorf("%d, %f, %f\n", i, v, dpow4_ok[i]);
            t.Errorf("matrix pow's data error");
            return;
         }
    }
    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    t.Log("Test matrix.Pow(1/4)\n");
    dm = m.Pow(0.25);
    dpow1_4 := dm.GetData();
    dpow1_4_ok := [30]float64 {1.626577, 1.414214, 1.316074, 1.565085, 1.495349, 1.681793, 1.732051, 1.189207, 1.778279, 1.000000,
                               3.989411, 4.276284, 4.481887, 4.730126, 5.101093, 5.223204, 5.321285, 5.526906, 3.225201, 3.253879,
                               1.000000, 1.189207, 1.316074, 1.414214, 1.495349, 1.565085, 1.626577, 1.681793, 1.732051, 1.778279};
    for i, v := range dpow1_4 {
       if math.Abs(v - dpow1_4_ok[i]) > 0.00001 {
           t.Errorf("matrix pow's data error");
           return;
        }
    }

    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());
}

*/
