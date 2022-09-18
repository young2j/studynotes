/*
 * File: expression-*beg++.cc
 * Created Date: 2022-09-07 07:13:44
 * Author: ysj
 * Description:
 */

#include <iostream>
#include <string>

int main(int argc, const char **argv) {
  std::string str = "hello";
  auto beg = str.begin();
  // ++和--运算符优先级高于解引用运算符
  // *beg++ 本质是*(beg++), ++将beg移到下一位，然后返回前一位的副本，再解引用
  while (beg != str.end()) {
    std::cout << *beg++ << std::endl;
  }

  return 0;
}
