/*
 * File: function-funcptr.cc
 * Created Date: 2022-09-13 02:15:18
 * Author: ysj
 * Description:  函数指针
 */

#include <iostream>
#include <string>
using std::string;

// 函数length_compare
bool length_compare(const string &s1, const string &s2) {
  // ...
  return s1.length() > s2.length();
}

// 声明函数指针
bool (*fptr)(const string &, const string &);

// 函数类型作为形参-自动转换为函数指针
string::size_type max_length(const string &str1, const string &str2,
                             bool f(const string &, const string &)) {
  bool b = f(str1, str2);
  if (b) {
    return str1.length();
  }
  return str2.length();
}
// 等价的声明-形参显式地声明为函数指针
string::size_type max_length(const string &str1, const string &str2,
                             bool (*f)(const string &, const string &));

int main(int argc, const char **argv) {
  // 等价初始化
  fptr = length_compare;  // 自动转换为指针
  fptr = &length_compare; // 取址符是可选的

  // 等价调用
  bool b1 = fptr("hello", "world"); // 自动转为所指的函数，再调用
  bool b2 = (*fptr)("hello", "world");        // 解引用，再调用
  bool b3 = length_compare("hello", "world"); // 直接调用
  std::cout << b1 << " " << b2 << " " << b3 << std::endl;

  string::size_type max_len = max_length("hello", "world", fptr);
  std::cout << max_len << std::endl;

  return 0;
}
