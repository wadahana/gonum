#include "cblas.h"
#include "stdio.h"
#include "stdlib.h"
#include "math.h"

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

int main(const int argc, const char* argv[])
{
    int M = 4; // matrix A's row num
    int N = 5; // matrix B's col num
    int K = 2; // matrix A's col num
    int lda = max(1, M);
    int ldb = max(1, K);
    int ldc = max(1, M);
    cblas_dgemm(CblasColMajor, CblasNoTrans, CblasNoTrans,
            M, N, K, 1.0, matrix_A, lda, matrix_B, ldb, 1.0, matrix_C, ldc);

    for (int r = 0; r < 4; r++) {
        for (int c = 0; c < 5; c++) {
            fprintf(stdout, "%20.6f  ", matrix_C[c * 4 + r]);
        }
        fprintf(stdout, "\n");
    }
}
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
