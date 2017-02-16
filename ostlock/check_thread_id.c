/*
 * gcc check_thread_id.c -lpthread -shared -fPIC -o libcheck_thread_id.so
 */

#include <pthread.h>
#include <stdio.h>

static pid_t id;

void init_id()
{
    id = pthread_self();
    fprintf(stdout, "%s: thread id = %d\n", __func__, id);
    fflush(stdout);
}

void check_id()
{
    pid_t current_id = pthread_self();
    fprintf(stdout, "%s: thread id = %d\n", __func__, current_id);
    fflush(stdout);
    // if (current_id != id)
    // {
    //     int zero = 0;
    //     int boom = 10 / zero;
    //     fprintf(stdout, "%s: boom = %d\n", __func__, boom);
    // }
}
