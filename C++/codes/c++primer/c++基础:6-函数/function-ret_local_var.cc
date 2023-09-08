/*
 * File: function-ret_local_var.cc
 * Created Date: 2022-09-12 06:52:48
 * Author: ysj
 * Description: 函数-返回局部对象
 */

#include <iostream>
#include <string>

// ok, 返回局部对象副本
std::string ret_localvar() {
  std::string str = "hello world";
  return str;
}

// not ok, 不能返回局部对象的引用
std::string &ret_ref() {
  std::string str = "hello world";
  return str;
}

// not ok, 不能返回局部对象的指针
std::string *ret_pointer() {
  std::string str = "hello world";
  return &str;
}

int main(int argc, const char **argv) {
  std::cout << ret_localvar() << std::endl;
  return 0;
}
