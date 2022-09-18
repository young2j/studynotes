/*
 * File: function-local_var.cc
 * Created Date: 2022-09-09 02:23:33
 * Author: ysj
 * Description: 函数局部变量
 */

#include <cstddef>
#include <iostream>

// 内置类型未初始化——产生未定义的值
void builtin() {
  int i, j;
  std::cout << "i:" << i << std::endl;
  std::cout << "j:" << j << std::endl;

  double m, n;
  std::cout << "m:" << m << std::endl;
  std::cout << "n:" << n << std::endl;

  int arr1[2], arr2[3];
  std::cout << "arr1:" << *arr1 << std::endl;
  std::cout << "arr2:" << *arr2 << std::endl;

  // 标准库类型——具有默认值
  std::string s; // 默认初始值""
  std::cout << "s:" << (s == "") << std::endl << std::endl;
  ;
}

// 局部静态对象——定义语句只执行一次初始化，直到程序终止时才会销毁，所在函数执行完毕后依然存在
size_t f() {
  static size_t c = 0;
  return ++c;
}
void static_var() {
  std::cout << "c:";
  for (size_t i = 0; i < 10; i++) {
    size_t c = f();
    std::cout << " " << c;
  }
  std::cout << std::endl << std::endl;
}

// 内置类型的局部静态变量——将初始化为0
void static_builtin() {
  static int si, sj;
  std::cout << "si:" << si << std::endl;
  std::cout << "sj:" << sj << std::endl;

  static double sm, sn;
  std::cout << "sm:" << sm << std::endl;
  std::cout << "sn:" << sn << std::endl;

  static int sarr1[2], sarr2[3];
  std::cout << "sarr1:" << *sarr1 << std::endl;
  std::cout << "sarr2:" << *sarr2 << std::endl;
}

int main() {
  std::cout << "内置类型局部变量未初始化:" << std::endl;
  builtin();
  std::cout << "静态局部变量仅初始化一次:" << std::endl;
  static_var();
  std::cout << "静态内置类型局部变量初始化为0:" << std::endl;
  static_builtin();
  return 0;
}

