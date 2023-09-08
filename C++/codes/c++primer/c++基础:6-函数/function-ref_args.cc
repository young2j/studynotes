/*
 * File: function-ref_args.cc
 * Created Date: 2022-09-09 03:52:50
 * Author: ysj
 * Description: 函数-引用参数
 */

#include <iostream>
#include <string>

void func(const int &i, std::string &s) {
  // i = 2; // 不能改变i引用的对象值
  s = "world";
}

int main(int argc, const char **argv) {
  int i = 0;
  std::string s = "hello";
  func(i, s);
  std::cout << "i:" << i << std::endl;
  std::cout << "s:" << s << std::endl;
  return 0;
}
