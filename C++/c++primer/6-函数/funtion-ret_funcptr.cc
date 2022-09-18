/*
 * File: funtion-ret_funcptr.cc
 * Created Date: 2022-09-13 02:53:44
 * Author: ysj
 * Description: 返回函数指针
 */

#include <iostream>
#include <string>
using std::string;

// 函数length_compare
bool length_compare(const string &s1, const string &s2) {
  // ...
  return s1.length() > s2.length();
}

// 返回函数指针
// 1. 直接声明
bool (*f(const string &, const string &))(int);
// 2. 类型别名
typedef bool (*Func)(const string &, const string &);
Func f(int);
// 3. 尾置类型
auto f(int) -> bool (*)(const string &, const string &);
// 4. 使用decltype
decltype(length_compare) *f(int);

int main(int argc, const char **argv) { return 0; }
