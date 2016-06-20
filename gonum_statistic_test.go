package gonum

import (
    "testing"
    "math"
)


var samples1 = []float64 {104.55066,89.58956,99.35684,113.53215,101.66452,106.20081,84.75633,99.82675,89.94354,91.17701,86.03428,101.79902,93.73638,103.59369,99.12036,101.38562,108.68519,99.25074,112.35618,101.36767};
var samples2 = []float64 {102.79489,104.55911,114.54894,103.84457,91.29889,104.71719,92.00355,109.54993,96.79654,109.85987,109.62344,79.83759,94.26286,108.39173,109.56648,88.74829,111.54637,113.08637,83.93840,88.60029};

func Test_Mean(t* testing.T) {
    mean1 := Mean(samples1);
    mean2 := Mean(samples2);
    if math.Abs(mean1 - 99.39637) > 0.00001 ||
        math.Abs(mean2 -100.8788) > 0.0001 {
        t.Errorf("mean caculate error");
        return;
    }
    t.Logf("\nmean 1 : %f\nmean 2 : %f\n", mean1, mean2);
}

func Test_Sum(t* testing.T) {
    sum1 := Sum(samples1);
    sum2 := Sum(samples2);

    if math.Abs(sum1 - 1987.927) > 0.001 ||
        math.Abs(sum2 -2017.575) > 0.001 {
        t.Errorf("sum caculate error");
        return;
    }
    t.Logf("\nsum 1 : %f\nsum 2 : %f\n", sum1, sum2);
}

func Test_VarAndSd(t* testing.T) {
    v1 := Var(samples1);
    v2 := Var(samples2);
    sd1 := Sd(samples1);
    sd2 := Sd(samples2);

    if math.Abs(v1 - 65.3467) > 0.0001 ||
        math.Abs(v2 - 111.2251) > 0.001 {
        t.Errorf("var caculate error");
        return;
    }
    t.Logf("\nvar 1 : %f\nvar 2 : %f\n", v1, v2);
    if math.Abs(sd1 - 8.083731) > 0.000001 ||
        math.Abs(sd2 - 10.54633) > 0.00001 {
        t.Errorf("sd caculate error");
        return;
    }
    t.Logf("\nsd 1 : %f\nsd 2 : %f\n", sd1, sd2);
}

func Test_CovAndCor(t* testing.T) {
    cov := Cov(samples1, samples2);
    if math.Abs(cov+8.073395) > 0.000001 {
        t.Errorf("cov caculate fail");
        return
    }

    t.Logf("\ncov : %f\n", cov);

    cor := Cor(samples1, samples2);
    if math.Abs(cor+0.09469845) > 0.000001 {
        t.Errorf("cor caculate fail");
        return
    }
    t.Logf("\ncor : %f\n", cor);
}

func Test_SkewnessAndKurtosis(t* testing.T) {
    sk := Skewness(samples1);

    if math.Abs(sk + 0.158841) > 0.000001 {
        t.Errorf("skewness caculate fail, sk:%f", sk);
        return
    }

    t.Logf("\nskewness : %f\n", sk);

    ku := Kurtosis(samples1);

    if math.Abs(ku+0.898537) > 0.000001 {
        t.Errorf("kurtosis caculate fail, ku:%f", ku);
        return
    }

    t.Logf("\nkurtosis : %f\n", ku);
}

/*
func Test_Hist(t* testing.T) {
    hist := Hist(samples1, 12);
    for i, v := range hist {
        t.Logf("i: %dl count: %d\n", i+1, v);
    }
}
*/
