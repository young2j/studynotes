/*
 * File: string-geline.cc
 * Created Date: 2022-08-29 01:26:27
 * Author: ysj
 * Description:  读取整行
 */

#include <iostream>
#include <string>
using std::string;

int main() {
  string line;
  while (std::getline(std::cin, line)) {
    std::cout << line << std::endl;
  }

  return 0;
}
