/*
 * File: function-overload_const_cast.cc
 * Created Date: 2022-09-12 08:26:08
 * Author: ysj
 * Description: 函数重载-const_cast应用
 */

#include <string>
using std::string;

const string &func(const string &s1, const string &s2) {
  // ...
  return s1.size() < s2.size() ? s1 : s2;
}

// 使用const_cast复用逻辑
string &func(string &s1, string &s2) {
  auto &s =
      func(const_cast<const string &>(s1), const_cast<const string &>(s2));

  return const_cast<string &>(s);
}

int main(int argc, const char** argv) {
    return 0;
}
