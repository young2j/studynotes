/*
 * File: array-cstring.cc
 * Created Date: 2022-09-07 04:59:05
 * Author: ysj
 * Description:  c风格字符串操作
 */

#include <cstring>
#include <iostream>
#include <string>

int main(int argc, const char **argv) {
  char arr1[] = {'h', 'e', 'l', 'l', 'o', '\0'};
  std::cout << "len(arr)=" << strlen(arr1) << std::endl;

  char arr2[] = " world";
  // arr2拼接到arr1
  std::strcat(arr1, arr2);

  // arr1 拷贝给 arr3
  char arr3[strlen(arr1)];
  std::strcpy(arr3, arr1);

  // 比较
  int cmp = std::strcmp(arr1, arr3);
  std::cout << "cmp: " << cmp << std::endl;

  if (cmp == 0) {
    std::cout << "arr1 == arr3" << std::endl;
  } else if (cmp > 0) {
    std::cout << "arr1 > arr3" << std::endl;
  } else {
    std::cout << "arr1 < arr3" << std::endl;
  }

  return 0;
}
