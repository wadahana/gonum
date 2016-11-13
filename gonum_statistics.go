package gonum

import (
//    "fmt"
    "math"
    "math/rand"
)
/*
    random serial generator
*/
func Uniform(n int, a, b float64) []float64 {
    data := make([]float64, n, n);
    d := b - a;
    for i := 0; i < n; i++ {
        data[i] = rand.Float64() * d + a;
    }
    return data;
}

func Exponential(n int, lambdav float64) []float64 {
    data := make([]float64, n, n);
    for i := 0; i < n; i++ {
        data[i] = rand.ExpFloat64() / lambdav;
    }
    return data;
}

func Normal(n int, mean, sd float64) []float64 {
    data := make([]float64, n, n);
    for i := 0; i < n; i++ {
        data[i] = rand.NormFloat64() * sd + mean;
    }
    return data;
}

func Max(data []float64) (int, float64) {
    max := data[0];
    idx := 0;
    for i, v := range data {
        if v > max {
            idx = i;
            max = v;
        }
    }
    return idx, max;
}

func Min(data []float64) (int, float64) {
    min := data[0];
    idx := 0;
    for i, v := range data {
        if v < min {
            idx = i;
            min = v;
        }
    }
    return idx, min;
}

func Sum(data []float64) float64 {
    var sum float64 = 0;
    for _, v := range data {
        sum += v;
    }
    return sum;
}

func Mean(data []float64) float64 {
    mean := Sum(data) / float64(len(data));
    return mean;
}

func Sd(data []float64) float64 {
    return math.Sqrt(Var(data));
}

func Var(data []float64) float64 {
    mean := Mean(data);
    var sum float64 = 0;
    for _, v := range data {
        d := v - mean;
        sum += d * d;
    }
    vv := sum / float64(len(data) - 1);
    return vv;
}

func Skewness(data []float64) float64 {
    mean := Mean(data);
    var sum0, sum1 float64 = 0, 0;
    for _, v := range data {
        d := v - mean;
        sum0 += d * d * d;
        sum1 += d * d;
    }
    vv0 := sum0 / float64(len(data));
    vv1 := sum1 / float64(len(data) - 1);
    sd := math.Sqrt(vv1);
    return vv0/math.Pow(sd, 3);
}

func Kurtosis(data []float64) float64 {
    mean := Mean(data);
    var sum0, sum1 float64 = 0, 0;
    for _, v := range data {
        d := v - mean;
        sum0 += d * d * d * d;
        sum1 += d * d;
    }
    vv0 := sum0 / float64(len(data));
    vv1 := sum1 / float64(len(data) - 1);
    sd := math.Sqrt(vv1);
    return vv0/math.Pow(sd, 4) - 3;
}

func Cov(dat1, dat2 []float64) float64 {
    if len(dat1) != len(dat2) {
        panic("two vector have diff element num");
    }
    var sum float64 = 0;
    e1 := Mean(dat1);
    e2 := Mean(dat2);
    num := len(dat1);
    for i := 0; i < num; i++ {
        d1 := (dat1[i] - e1);
        d2 := (dat2[i] - e2);
        sum += d1 * d2;
    }
    vv := sum / float64(num - 1);
    return vv;
}

func Cor(dat1, dat2 []float64) float64 {
    if len(dat1) != len(dat2) {
        panic("two vector have diff element num");
    }

    var sum0, sum1, sum2 float64 = 0, 0, 0;

    e1 := Mean(dat1);
    e2 := Mean(dat2);

    num := len(dat1);
    for i := 0; i < num; i++ {
        d1 := (dat1[i] - e1);
        d2 := (dat2[i] - e2);

        sum0 += d1 * d2;
        sum1 += d1 * d1;
        sum2 += d2 * d2;
    }
    vv1 := sum0 / float64(num - 1);
    vv2 := sum1 / float64(num - 1);
    vv3 := sum2 / float64(num - 1);
    return vv1 / math.Sqrt(vv2*vv3);
}
