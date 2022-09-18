/*
 * File: iterator-dereference.cc
 * Created Date: 2022-09-06 02:42:44
 * Author: ysj
 * Description: 迭代器-解引用
 */

#include <cctype>
#include <iostream>
#include <string>

int main(int argc, const char **argv) {
  std::string s("hello world");

  // 首字母大写
  if (s.begin() != s.end()) {
    auto it = s.begin();
    *it = std::toupper(*it);
  }

  std::cout << s << std::endl;

  // 所有字母大写
  for (auto it = s.begin(); it != s.end(); ++it) {
    *it = std::toupper(*it);
  }
  std::cout << s << std::endl;

  return 0;
}
