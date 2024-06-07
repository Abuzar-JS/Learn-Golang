#include <iostream>
#include <cstdlib>
#include <ctime>
#include <omp.h>
#define N 10
void generateRandomMatrix(int matrix[N][N]) {
for (int i = 0; i < N; ++i) {
for (int j = 0; j < N; ++j) {
matrix[i][j] = rand() % 100; 
}
}
}
void multiplyMatrices(int result[N][N], int matrix1[N][N], int matrix2[N][N]) {
#pragma omp parallel for
for (int i = 0; i < N; ++i) {
for (int j = 0; j < N; ++j) {
result[i][j] = 0;
for (int k = 0; k < N; ++k) {
result[i][j] += matrix1[i][k] * matrix2[k][j];
}
}
}
}

void printMatrix(int matrix[N][N]) {
for (int i = 0; i < N; ++i) {
for (int j = 0; j < N; ++j) {
std::cout << matrix[i][j] << " ";
}
std::cout << std::endl;
}
}
int main() {

srand(time(nullptr));
int matrix1[N][N];
int matrix2[N][N];
int result[N][N];

generateRandomMatrix(matrix1);
generateRandomMatrix(matrix2);

multiplyMatrices(result, matrix1, matrix2);

std::cout << "Resulting Matrix:" << std::endl;
printMatrix(result);
return 0;
}







