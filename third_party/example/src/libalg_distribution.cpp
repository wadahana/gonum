#include <libalg/alglibinternal.h>
#include <libalg/alglibmisc.h>
#include <libalg/linalg.h>
#include <libalg/statistics.h>
#include <libalg/dataanalysis.h>
#include <libalg/specialfunctions.h>
#include <libalg/solvers.h>
#include <libalg/optimization.h>
#include <libalg/diffequations.h>
#include <libalg/fasttransforms.h>
#include <libalg/integration.h>
#include <libalg/interpolation.h>

//using namespace alglib;


bool create_uniform_distrubtion(double * data, int n) {
    bool result = false;
    alglib_impl::ae_state alglib_env_state;
    alglib_impl::hqrndstate hqrnd_state;

    alglib_impl::ae_state_init(&alglib_env_state);
    _hqrndstate_init(&hqrnd_state, &alglib_env_state);
    try {
        alglib_impl::hqrndrandomize(const_cast<alglib_impl::hqrndstate*>(&hqrnd_state), &alglib_env_state);

        for (int i = 0; i < n; i++) {
            data[i] = alglib_impl::hqrnduniformr(const_cast<alglib_impl::hqrndstate*>(&hqrnd_state), &alglib_env_state);
        }
        alglib_impl::ae_state_clear(&alglib_env_state);
        result = true;
    } catch(alglib_impl::ae_error_type) {
        fprintf(stderr, "%s\n", alglib_env_state.error_msg);
    }
    return result;
}

bool create_normal_distrubtion(double * data, int n) {
    bool result = false;
    alglib_impl::ae_state alglib_env_state;
    alglib_impl::hqrndstate hqrnd_state;

    alglib_impl::ae_state_init(&alglib_env_state);
    _hqrndstate_init(&hqrnd_state, &alglib_env_state);
    try {
        alglib_impl::hqrndrandomize(const_cast<alglib_impl::hqrndstate*>(&hqrnd_state), &alglib_env_state);

        for (int i = 0; i < n; i++) {
            data[i] = alglib_impl::hqrndnormal(const_cast<alglib_impl::hqrndstate*>(&hqrnd_state), &alglib_env_state);
        }
        alglib_impl::ae_state_clear(&alglib_env_state);
        result = true;
    } catch(alglib_impl::ae_error_type) {
        fprintf(stderr, "%s\n", alglib_env_state.error_msg);
    }
    return result;
}

bool create_exponential_distrubtion(double * data, int n, double lambdav) {
    bool result = false;
    alglib_impl::ae_state alglib_env_state;
    alglib_impl::hqrndstate hqrnd_state;

    alglib_impl::ae_state_init(&alglib_env_state);
    _hqrndstate_init(&hqrnd_state, &alglib_env_state);
    try {
        alglib_impl::hqrndrandomize(const_cast<alglib_impl::hqrndstate*>(&hqrnd_state), &alglib_env_state);

        for (int i = 0; i < n; i++) {
            data[i] = alglib_impl::hqrndexponential(const_cast<alglib_impl::hqrndstate*>(&hqrnd_state), lambdav, &alglib_env_state);
        }
        alglib_impl::ae_state_clear(&alglib_env_state);
        result = true;
    } catch(alglib_impl::ae_error_type) {
        fprintf(stderr, "%s\n", alglib_env_state.error_msg);
    }
    return result;
}

static void __print_vector(double *vector, int size) {
    for(int i = 0; i < size; i++) {
        printf("%4.2f, ", vector[i]);
    }
    printf("\n");
}

extern "C" void test_distribution() {

    double data[256];
    int total = 256;

    printf("uniform distrubution : \n");
    create_uniform_distrubtion(data, total);
    __print_vector(data, total);

    printf("normal distrubution : \n");
    create_normal_distrubtion(data, total);
    __print_vector(data, total);

    printf("normal exponential : \n");
    create_exponential_distrubtion(data, total, 1);
    __print_vector(data, total);

    double v = alglib::normaldistribution(0);
    printf("v: %f\n", v);
    return ;

}