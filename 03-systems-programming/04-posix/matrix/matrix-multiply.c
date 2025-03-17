#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>

/*
  A naive implementation of matrix multiplication.

  DO NOT MODIFY THIS FUNCTION, the tests assume it works correctly, which it
  currently does
*/
void matrix_multiply(double **C, double **A, double **B, int a_rows, int a_cols,
                     int b_cols)
{
  for (int i = 0; i < a_rows; i++)
  {
    for (int j = 0; j < b_cols; j++)
    {
      C[i][j] = 0;
      for (int k = 0; k < a_cols; k++)
        C[i][j] += A[i][k] * B[k][j];
    }
  }
}

typedef struct
{
  double **C;
  double **A;
  double **B;
  int row_start;
  int row_end;
  int a_cols;
  int b_cols;
} ThreadArgs;

void *thread_func(void *arg)
{
  ThreadArgs *args = (ThreadArgs *)arg;

  for (int i = args->row_start; i < args->row_end; i++)
  {
    for (int j = 0; j < args->b_cols; j++)
    {
      args->C[i][j] = 0;
      for (int k = 0; k < args->a_cols; k++)
      {
        args->C[i][j] += args->A[i][k] * args->B[k][j];
      }
    }
  }

  return NULL;
}

void parallel_matrix_multiply(double **c, double **a, double **b, int a_rows,
                              int a_cols, int b_cols)
{
  int num_threads = 10;
  pthread_t threads[num_threads];
  ThreadArgs args[num_threads];

  // divide the rows into all threads
  int rows_per_thread = a_rows / num_threads;
  for (int i = 0; i < num_threads; i++)
  {
    args[i].C = c;
    args[i].A = a;
    args[i].B = b;
    args[i].row_start = i * rows_per_thread;
    args[i].row_end = (i == num_threads - 1) ? a_rows : (i + 1) * rows_per_thread;
    args[i].a_cols = a_cols;
    args[i].b_cols = b_cols;

    pthread_create(&threads[i], NULL, thread_func, &args[i]);
  }

  // wait for all threads to finish
  for (int i = 0; i < num_threads; i++)
  {
    pthread_join(threads[i], NULL);
  }
}
