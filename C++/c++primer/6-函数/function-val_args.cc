/*
 * File: function-val_args.cc
 * Created Date: 2022-09-09 03:28:11
 * Author: ysj
 * Description: 函数-值参数
 */

#include <iostream>
#include <string>

// i和s都是值拷贝，改变其值后不影响传入的实参值；
// p是指向实参i的指针，指针被拷贝，可以通过它改变i的值，实参p的值未被改变
void func(int i, std::string s, int *p) {
  *p = 100;
  p = 0;

  i = 2;
  s = "world";
}

int main(int argc, const char **argv) {
  int i = 0, *p = &i;
  std::string s = "hello";
  func(i, s, p);
  std::cout << "i:" << i << std::endl;
  std::cout << "s:" << s << std::endl;
  std::cout << "p:" << p << " *p:" << *p << std::endl;
  return 0;
}
