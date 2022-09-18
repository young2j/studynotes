/*
 * File: function-inline.cc
 * Created Date: 2022-09-12 08:42:41
 * Author: ysj
 * Description: 函数-内联
 */
#include <iostream>
#include <string>
using std::string;

inline const string &func(const string &s1, const string &s2) {
  // ...
  return s1.size() < s2.size() ? s1 : s2;
}

int main(int argc, const char **argv) {
  string s1("hello"), s2("world");
  std::cout << func(s1, s2) << std::endl; // 编译时会展开为:
  // std::cout << (s1.size() < s2.size() ? s1 : s2) << std::endl;

  return 0;
}
