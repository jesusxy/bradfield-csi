/*
Naive code for multiplying two matrices together.

There must be a better way!
*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void transpose_matrix(double **M, int m_rows, int m_cols)
{
  for (int mr = 0; mr < m_rows; mr++)
    for (int mc = mr + 1; mc < m_cols; mc++)
    {
      double temp = M[mr][mc];
      M[mr][mc] = M[mc][mr];
      M[mc][mr] = temp;
    }
}

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

// void fast_matrix_multiply(double **C, double **A, double **B, int a_rows,
//                           int a_cols, int b_cols)
// {
//   // TODO: write a faster implementation here!
//   matrix_init(C, a_rows, b_cols);
//   for (int i = 0; i < a_rows; i++)
//   {
//     for (int j = 0; j < b_cols; j++)
//     {
//       double sum0 = 0, sum1 = 0, sum2 = 0, sum3 = 0;
//       int k = 0;
//       for (; k < a_cols - 3; k += 4)
//       {
//         sum0 += A[i][k + 0] * B[k + 0][j];
//         sum1 += A[i][k + 1] * B[k + 1][j];
//         sum2 += A[i][k + 2] * B[k + 2][j];
//         sum3 += A[i][k + 3] * B[k + 3][j];
//       }
//       for (; k < a_cols; k++)
//       {
//         sum0 += A[i][k] * B[j][k];
//       }

//       C[i][j] += sum0 + sum1 + sum2 + sum3;
//     }
//   }
// }

/**
 * Transpose B matrix and multiply
 */
// void fast_matrix_multiply(double **C, double **A, double **B, int a_rows, int a_cols, int b_cols)
// {
//   transpose_matrix(B, a_cols, b_cols);
//   for (int ar = 0; ar < a_rows; ar++)
//   {
//     for (int bc = 0; bc < b_cols; bc++)
//     {
//       C[ar][bc] = 0;
//       for (int ac = 0; ac < a_cols; ac++)
//       {
//         C[ar][bc] += A[ar][ac] * B[bc][ac];
//       }
//     }
//   }
// }

// this "transposes" the second matrix, without transposing
// we take advantage of reading rows sequentially and calculatin the dot product of A B
void fast_matrix_multiply(double **C, double **A, double **B, int a_rows, int a_cols, int b_cols)
{
  // set C matrix values to 0
  for (int i = 0; i < a_rows; i++)
  {
    for (int j = 0; j < b_cols; j++)
    {
      C[i][j] = 0;
    }
  }

  for (int ar = 0; ar < a_rows; ar++)
  {
    for (int ac = 0; ac < a_cols; ac++)
    {
      for (int bc = 0; bc < b_cols; bc++)
      {
        C[ar][bc] += A[ar][ac] * B[ac][bc];
      }
    }
  }
}
