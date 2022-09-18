/*
 * File: function-arr_args.cc
 * Created Date: 2022-09-09 04:03:52
 * Author: ysj
 * Description:  函数-数组参数
 */

#include <iostream>
#include <iterator>

void print(const int *arr, size_t n) {
  while (n > 0) {
    std::cout << *arr++;
    --n;
  }
  std::cout << std::endl;
  std::cout << "----" << std::endl;
}
// 数组参数形式1
void func1(const int *arr, size_t n) { print(arr, n); }

// 数组参数形式2
void func2(const int arr[], size_t n) { print(arr, n); }

// 数组参数形式3
void func3(const int arr[10], size_t n) { print(arr, n); }

// 数组引用参数
void func4(const int (&arr)[3]) {
  for (auto elem : arr) {
    std::cout << elem;
  }
  std::cout << std::endl;
  std::cout << "----" << std::endl;
}

int main(int argc, const char **argv) {
  int arr[] = {1, 2, 3};
  size_t n = std::end(arr) - std::begin(arr);
  func1(arr, n);
  func2(arr, n);
  func3(arr, n);
  func4(arr);

  return 0;
}

// 123
// ----
// 123
// ----
// 123
// ----
// 123
// ----
