/*
 * File: function-ret_arrptr.cc
 * Created Date: 2022-09-12 07:40:17
 * Author: ysj
 * Description: 函数-返回数组指针
 */

#include <iostream>

int arr[10] = {1, 2, 3};

// 1.直接声明方式
int (*ret_arrptr1())[10] {
  // ...
  return &arr;
}

// 2.类型别名方式
typedef int iarr[10];
iarr *ret_arrptr2() {
  //...
  return &arr;
}

// 3.尾置返回类型
auto ret_arrptr3() -> int (*)[10] {
  //...
  return &arr;
}

// 4.使用decltype
decltype(arr) *ret_arrptr4() {
  //...
  return &arr;
}

int main(int argc, const char **argv) {
  std::cout << ret_arrptr1() << std::endl;
  std::cout << ret_arrptr2() << std::endl;
  std::cout << ret_arrptr3() << std::endl;
  std::cout << ret_arrptr4() << std::endl;
  return 0;
}
