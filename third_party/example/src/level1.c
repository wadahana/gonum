#include "cblas.h"
#include "stdio.h"
#include "stdlib.h"
#include "math.h"


double vector_A[8] = {1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0};
double vector_B[8] = {10.0, 20.0, 30.0, 40.0, 50.0, 60.0, 70.0, 80.0};

void print_vector(double *vector, int size) {
    for(int i = 0; i < size; i++) {
        printf("%4.2f, ", vector[i]);
    }
    printf("\n");
}

void test_float() {

    double x[64];
    double y[64];
    /*
        a<-c(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0)
        b<-c(10.0, 20.0, 30.0, 40.0, 50.0, 60.0, 70.0, 80.0)
        1. sum(a) = 36
        2. a*2+b = {12 24 36 48 60 72 84 96}
        3. a%*%b = 2040
        4. sqrt(sum(a^2)) = 14.28286
    */


    // sum(x)
    cblas_dcopy(8, vector_A, 1, x, 1);
    double result = cblas_dasum(8, x, 1);
    printf("cblas_dasum : \n\t%16.4f\n", result);

    // y = x*a + y
    cblas_dcopy(8, vector_A, 1, x, 1);
    cblas_dcopy(8, vector_B, 1, y, 1);
    cblas_daxpy(8, 2, x, 1, y, 1);
    printf("cblas_saxpy : \n\t");
    print_vector(y, 8);

    // y = x.y    点乘 / 内积
    cblas_dcopy(8, vector_A, 1, x, 1);
    cblas_dcopy(8, vector_B, 1, y, 1);
    result = cblas_ddot(8, x, 1, y, 1);
    printf("cblas_ddot : \n\t%16.4f\n", result);

    // sqrt(x.Trans(x))  欧式距离
    cblas_dcopy(8, vector_A, 1, x, 1);
    result = cblas_dnrm2(8, x, 1);
    printf("cblas_dnrm2 : \n\t%16.4f\n", result);

    /*
        | x |   | c   s |   | x |
        |   | = |       | x |   |
        | y |   | s  -c |   | y |

    */
    cblas_dcopy(8, vector_A, 1, x, 1);
    cblas_dcopy(8, vector_B, 1, y, 1);
    cblas_drot(8, x, 1, y, 1, 1, 1);
    printf("cblas_drot : \n");
    printf("\tx: ");
    print_vector(x, 8);
    printf("\ty: ");
    print_vector(y, 8);

    // x = a*x
    cblas_dcopy(8, vector_A, 1, x, 1);
    cblas_dscal(8, 2, x, 1);
    printf("cblas_dscal : \n");
    printf("\tx: ");
    print_vector(x, 8);

    // x, y = y, x
    cblas_dcopy(8, vector_A, 1, x, 1);
    cblas_dcopy(8, vector_B, 1, y, 1);
    cblas_dswap(8, x, 1, y, 1);
    printf("cblas_dswap : \n");
    printf("\tx: ");
    print_vector(x, 8);
    printf("\ty: ");
    print_vector(y, 8);

    cblas_dcopy(8, vector_A, 1, x, 1);
    int i_max = cblas_idamax(8, x, 1);
//    int i_min = cblas_idamin(8, x, 1);
    printf("cblas_idamax : \n\t%d\n", i_max);
//    printf("cblas_idamin : \n\t%d\n", i_min);
}



double complex vector_CA[8] = {1.0 + 1.0i, 2.0 + 1.0i, 3.0 + 1.0i, 4.0 + 1.0i, 5.0 + 1.0i, 6.0 + 1.0i, 7.0 + 1.0i, 8.0 + 1.0i};
double complex vector_CB[8] = {10.0 + 10.0i, 20.0 + 10.0i, 30.0 + 10.0i, 40.0 + 10.0i, 50.0 + 10.0i, 60.0 + 10.0i, 70.0 + 10.0i, 80.0 + 10.0i};

void print_complex_vector(double complex * vector, int size) {
    for(int i = 0; i < size; i++) {
        printf("%4.2f+%4.2fi, ", creal(vector[i]), cimag(vector[i]));
    }
    printf("\n");
}

void test_complex() {
    double complex x[64];
    double complex y[64];
    double complex alpha = 2;

    // |creal(x[1])| + |cimag(x[1])| + |creal(x[2])| + |cimag(x[2])| + .... + |creal(x[n])| + |cimag(x[n])|
    cblas_zcopy(8, (double *)vector_CA, 1, (double *)x, 1);
    double result = cblas_dzasum(8, (double *)x, 1);
    printf("cblas_dzasum : \n\t%0.2f+%0.2fi\n", creal(result), cimag(result));

    // y = x*a + y
    cblas_zcopy(8, (double *)vector_CA, 1, (double *)x, 1);
    cblas_zcopy(8, (double *)vector_CB, 1, (double *)y, 1);
    cblas_zaxpy(8, (double *)&alpha, (double *)x, 1, (double *)y, 1);
    printf("cblas_zaxpy : \n\t");
    print_complex_vector(y, 8);

    /*
                        |    1    |
     Trans([1, 2+3i]) = |         |
                        |  2 - 3i |
    */
    cblas_zcopy(8, (double *)vector_CA, 1, (double *)x, 1);
    result = cblas_dznrm2(8, (double *)x, 1);
    printf("cblas_dznrm2 : \n\t%0.2f\n", result);

    cblas_zcopy(8, (double *)vector_CA, 1, (double *)x, 1);
    cblas_zcopy(8, (double *)vector_CB, 1, (double *)y, 1);
    cblas_zdrot(8, x, 1, y, 1, 1, 1);
    printf("cblas_zdrot : \n");
    printf("\tx: ");
    print_complex_vector(x, 8);
    printf("\ty: ");
    print_complex_vector(y, 8);

}

void test_level1() {
    test_float();
    test_complex();
}