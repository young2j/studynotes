/*
 * File: function-constexpr.cc
 * Created Date: 2022-09-12 08:55:55
 * Author: ysj
 * Description:  常量表达式函数
 */

#include <cstddef>
constexpr std::size_t scale(std::size_t sz) { return 2 * sz; }

int main(int argc, const char **argv) {
  int arr1[scale(1)] = {1, 2}; // scale(1) 返回的是常量表达式
  int i = 1;
  int arr2[scale(i)] = {1, 2}; // scale(i)返回的不是常量表达式

  return 0;
}
