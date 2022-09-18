/*
 * File: array-ndarray.cc
 * Created Date: 2022-09-07 05:55:24
 * Author: ysj
 * Description:
 */

#include <iostream>
#include <iterator>

int main(int argc, const char **argv) {
  // 三行四列
  int aa[3][4] = {
      {0, 1, 2, 3},
      {4, 5, 6, 7},
      {8, 9, 10, 11},
  };

  // p指向类型为int[4]
  int(*p)[4] = aa; // p指向aa的首元素
  p = &aa[1];      //  p指向aa的第二个元素

  // 多维数组遍历
  // 范围遍历——外层for需要使用引用类型，否则row为首元素指针，内层遍历指针是不允许的
  for (auto &row : aa) {
    for (auto &col : row) {
      std::cout << col << " ";
    }
    std::cout << std::endl;
  }

  // 数组指针遍历
  for (auto p = std::begin(aa); p != std::end(aa); ++p) {
    for (auto q = std::begin(*p); q != std::end(*p); ++q) {
      std::cout << *q << " ";
    }
    std::cout << std::endl;
  }

  return 0;
}
