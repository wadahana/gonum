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

func Test_Math(t *testing.T) {
    m := testcase_NewMatrix();
    dm := m.Diff(1);
    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    dm = m.Diff(6);
    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    dm = m.Log2();
    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    dm = m.Log10();
    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    dm = m.Log1p();
    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    dm = m.Log();
    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    dm = m.Sqrt();
    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    dm = m.Pow(4);
    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());

    dm = m.Pow(0.25);
    t.Logf("\n%v[%d,%d]\n", dm, dm.GetRowNum(), dm.GetColumeNum());
}





/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *


> d<-c(7.0,4.0,3.0,6.0,5.0,8.0,9.0,2.0,10.0,1.0,253.3,334.4,403.5,500.6,677.1,744.3,801.8,933.1,108.2,112.1,1.0,2.0,3.0,4.0,5.0,6.0,7.0,8.0,9.0,10.0)
> m<-matrix(d,nrow=10,ncol=3)
> m
      [,1]  [,2] [,3]
 [1,]    7 253.3    1
 [2,]    4 334.4    2
 [3,]    3 403.5    3
 [4,]    6 500.6    4
 [5,]    5 677.1    5
 [6,]    8 744.3    6
 [7,]    9 801.8    7
 [8,]    2 933.1    8
 [9,]   10 108.2    9
[10,]    1 112.1   10
> diff(m,1)
      [,1]   [,2] [,3]
 [1,]   -3   81.1    1
 [2,]   -1   69.1    1
 [3,]    3   97.1    1
 [4,]   -1  176.5    1
 [5,]    3   67.2    1
 [6,]    1   57.5    1
 [7,]   -7  131.3    1
 [8,]    8 -824.9    1
 [9,]   -9    3.9    1
> diff(m,2)
     [,1]   [,2] [,3]
[1,]   -4  150.2    2
[2,]    2  166.2    2
[3,]    2  273.6    2
[4,]    2  243.7    2
[5,]    4  124.7    2
[6,]   -6  188.8    2
[7,]    1 -693.6    2
[8,]   -1 -821.0    2
> diff(m,6)
     [,1]   [,2] [,3]
[1,]    2  548.5    6
[2,]   -2  598.7    6
[3,]    7 -295.3    6
[4,]   -5 -388.5    6
> log2(m)
          [,1]     [,2]     [,3]
 [1,] 2.807355 7.984703 0.000000
 [2,] 2.000000 8.385431 1.000000
 [3,] 1.584963 8.656425 1.584963
 [4,] 2.584963 8.967514 2.000000
 [5,] 2.321928 9.403225 2.321928
 [6,] 3.000000 9.539740 2.584963
 [7,] 3.169925 9.647099 2.807355
 [8,] 1.000000 9.865888 3.000000
 [9,] 3.321928 6.757557 3.169925
[10,] 0.000000 6.808642 3.321928
> log10(m)
           [,1]     [,2]      [,3]
 [1,] 0.8450980 2.403635 0.0000000
 [2,] 0.6020600 2.524266 0.3010300
 [3,] 0.4771213 2.605844 0.4771213
 [4,] 0.7781513 2.699491 0.6020600
 [5,] 0.6989700 2.830653 0.6989700
 [6,] 0.9030900 2.871748 0.7781513
 [7,] 0.9542425 2.904066 0.8450980
 [8,] 0.3010300 2.969928 0.9030900
 [9,] 1.0000000 2.034227 0.9542425
[10,] 0.0000000 2.049606 1.0000000
> log10(m)
           [,1]     [,2]      [,3]
 [1,] 0.8450980 2.403635 0.0000000
 [2,] 0.6020600 2.524266 0.3010300
 [3,] 0.4771213 2.605844 0.4771213
 [4,] 0.7781513 2.699491 0.6020600
 [5,] 0.6989700 2.830653 0.6989700
 [6,] 0.9030900 2.871748 0.7781513
 [7,] 0.9542425 2.904066 0.8450980
 [8,] 0.3010300 2.969928 0.9030900
 [9,] 1.0000000 2.034227 0.9542425
[10,] 0.0000000 2.049606 1.0000000
> log1p(m)
           [,1]     [,2]      [,3]
 [1,] 2.0794415 5.538515 0.6931472
 [2,] 1.6094379 5.815324 1.0986123
 [3,] 1.3862944 6.002652 1.3862944
 [4,] 1.9459101 6.217803 1.6094379
 [5,] 1.7917595 6.519295 1.7917595
 [6,] 2.1972246 6.613787 1.9459101
 [7,] 2.3025851 6.688106 2.0794415
 [8,] 1.0986123 6.839583 2.1972246
 [9,] 2.3978953 4.693181 2.3025851
[10,] 0.6931472 4.728272 2.3978953
> log(m)
           [,1]     [,2]      [,3]
 [1,] 1.9459101 5.534575 0.0000000
 [2,] 1.3862944 5.812338 0.6931472
 [3,] 1.0986123 6.000176 1.0986123
 [4,] 1.7917595 6.215807 1.3862944
 [5,] 1.6094379 6.517819 1.6094379
 [6,] 2.0794415 6.612444 1.7917595
 [7,] 2.1972246 6.686859 1.9459101
 [8,] 0.6931472 6.838512 2.0794415
 [9,] 2.3025851 4.683981 2.1972246
[10,] 0.0000000 4.719391 2.3025851
> sqrt(m)
          [,1]     [,2]     [,3]
 [1,] 2.645751 15.91540 1.000000
 [2,] 2.000000 18.28661 1.414214
 [3,] 1.732051 20.08731 1.732051
 [4,] 2.449490 22.37409 2.000000
 [5,] 2.236068 26.02115 2.236068
 [6,] 2.828427 27.28186 2.449490
 [7,] 3.000000 28.31607 2.645751
 [8,] 1.414214 30.54669 2.828427
 [9,] 3.162278 10.40192 3.000000
[10,] 1.000000 10.58773 3.162278
> m^4
       [,1]         [,2]  [,3]
 [1,]  2401   4116619806     1
 [2,]   256  12504463842    16
 [3,]    81  26507828750    81
 [4,]  1296  62800540432   256
 [5,]   625 210189615237   625
 [6,]  4096 306896599227  1296
 [7,]  6561 413298860273  2401
 [8,]    16 758076017849  4096
 [9,] 10000    137059468  6561
[10,]     1    157914660 10000
> m^0.25
          [,1]     [,2]     [,3]
 [1,] 1.626577 3.989411 1.000000
 [2,] 1.414214 4.276284 1.189207
 [3,] 1.316074 4.481887 1.316074
 [4,] 1.565085 4.730126 1.414214
 [5,] 1.495349 5.101093 1.495349
 [6,] 1.681793 5.223204 1.565085
 [7,] 1.732051 5.321285 1.626577
 [8,] 1.189207 5.526906 1.681793
 [9,] 1.778279 3.225201 1.732051
[10,] 1.000000 3.253879 1.778279
*/
