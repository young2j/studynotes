/*
 * File: function-var_args.cc
 * Created Date: 2022-09-09 05:20:08
 * Author: ysj
 * Description: 函数-可变参数
 */

#include <initializer_list>
#include <iostream>
#include <string>

void init_lst(std::initializer_list<std::string> strs) {
  std::cout << "共" << strs.size() << "个参数: " << std::endl;

  for (auto s : strs) {
    std::cout << s << std::endl;
  }
}

void ellipsis(int a, ...) {}

int main(int argc, const char **argv) {
  init_lst({"a", "b", "c"});
  ellipsis(1, 2, 3);
  return 0;
}
