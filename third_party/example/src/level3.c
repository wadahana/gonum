
#include "cblas.h"

#define max(a, b) (a > b ? a : b)
/* * * * * * * * * *
 * A:
 *  1.0, 11.0
 *  2.0, 12.0
 *  3.0, 13.0
 *  4.0, 14.0
 * B:
 * 11.0, 21.0, 31.0, 41.0, 51.0
 * 12.0, 22.0, 32.0, 42.0, 52.0
 *
 * C:
 *   0, 0, 0, 0, 0
 *   0, 0, 0, 0, 0
 *   0, 0, 0, 0, 0
 *   0, 0, 0, 0, 0
 *
 *     http://www.math.utah.edu/software/lapack/lapack-blas/dgemm.html
 *
 * * * */
double matrix_A[8] = {1.0, 2.0, 3.0, 4.0, 11.0, 12.0, 13.0, 14.0};
double matrix_B[10] = {11.0, 12.0, 21.0, 22.0, 31.0, 32.0, 41.0, 42.0, 51.0, 52.0};
double matrix_C[20] = {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0};

/*

m1<-c(1.0, 2.0, 3.0, 4.0, 11.0, 12.0, 13.0, 14.0)
dim(m1)<-c(4,2)
m2<-c(11.0, 12.0, 21.0, 22.0, 31.0, 32.0, 41.0, 42.0,51.0, 52.0)
dim(m2)<-c(2,5)
m1%*%m2
     [,1] [,2] [,3] [,4] [,5]
[1,]  143  263  383  503  623
[2,]  166  306  446  586  726
[3,]  189  349  509  669  829
[4,]  212  392  572  752  932

*/

static void __test_real()
{
    int M = 4; // matrix A's row num
    int N = 5; // matrix B's col num
    int K = 2; // matrix A's col num
    int lda = max(1, M);
    int ldb = max(1, K);
    int ldc = max(1, M);

    fprintf(stdout, "\ncblas_dgemm test:\n");

    cblas_dgemm(CblasColMajor, CblasNoTrans, CblasNoTrans,
            M, N, K, 1.0, matrix_A, lda, matrix_B, ldb, 0.0, matrix_C, ldc);

    for (int r = 0; r < 4; r++) {
        for (int c = 0; c < 5; c++) {
            fprintf(stdout, "%.2f,  ", matrix_C[c * 4 + r]);
        }
        fprintf(stdout, "\n");
    }
}

double complex cmatrix_A[8] = {1.0 + 1.0i, 2.0 + 1.0i, 3.0 + 1.0i, 4.0 + 1.0i, 5.0 + 1.0i, 6.0 + 1.0i, 7.0 + 1.0i, 8.0 + 1.0i};
double complex cmatrix_B[10] = {10.0 + 10.0i, 20.0 + 10.0i, 30.0 + 10.0i, 40.0 + 10.0i, 50.0 + 10.0i, 60.0 + 10.0i, 70.0 + 10.0i, 80.0 + 10.0i, 90.0+10.0i, 100.0+10.i};
double complex cmatrix_C[20];


/*
m1<-c(1.0 + 1.0i, 2.0 + 1.0i, 3.0 + 1.0i, 4.0 + 1.0i, 5.0 + 1.0i, 6.0 + 1.0i, 7.0 + 1.0i, 8.0 + 1.0i);
dim(m1)<-c(4,2)
m2<-c(10.0 + 10.0i, 20.0 + 10.0i, 30.0 + 10.0i, 40.0 + 10.0i, 50.0 + 10.0i, 60.0 + 10.0i, 70.0 + 10.0i, 80.0 + 10.0i, 90.0+10.0i, 100.0+10.i);
dim(m2)<-c(2,5)
m1%*%m2
         [,1]     [,2]     [,3]     [,4]      [,5]
[1,]  90+ 90i 210+130i 330+170i 450+210i  570+250i
[2,] 120+110i 280+150i 440+190i 600+230i  760+270i
[3,] 150+130i 350+170i 550+210i 750+250i  950+290i
[4,] 180+150i 420+190i 660+230i 900+270i 1140+310i

*/
static void __test_complex()
{
    int M = 4; // matrix A's row num
    int N = 5; // matrix B's col num
    int K = 2; // matrix A's col num
    int lda = max(1, M);
    int ldb = max(1, K);
    int ldc = max(1, M);

    double complex alpha = 1.0;
    double complex beta = 0.0;
    fprintf(stdout, "\ncblas_zgemm test:\n");

    cblas_zgemm(CblasColMajor, CblasNoTrans, CblasNoTrans,
            M, N, K, (double *)&alpha, (double *)cmatrix_A, lda, (double *)cmatrix_B, ldb, (double *)&beta, (double *)cmatrix_C, ldc);

    for (int r = 0; r < 4; r++) {
        for (int c = 0; c < 5; c++) {
            double complex v = cmatrix_C[c * 4 + r];
            fprintf(stdout, "%.2f+%.2fi,  ", creal(v), cimag(v));
        }
        fprintf(stdout, "\n");
    }

}
void test_level3() {
    __test_real();
    __test_complex();
}