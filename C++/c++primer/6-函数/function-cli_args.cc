/*
 * File: function-cli_args.cc
 * Created Date: 2022-09-09 04:54:18
 * Author: ysj
 * Description: 函数-命令行参数
 */

#include <iostream>

int main(int argc, const char **argv) {
  std::cout << "argc = " << argc << std::endl;

  while (*argv) {
    std::cout << *argv++ << std::endl;
  }

  return 0;
}
// 命令行参数 "-f xxx.cc -o xxx.exe", 输出:
// argc = 5
// ./a.out
// -f
// xxx.cc
// -o
// xxx.exe
