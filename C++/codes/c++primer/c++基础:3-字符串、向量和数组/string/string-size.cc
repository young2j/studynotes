/*
 * File: string-size.cc
 * Created Date: 2022-08-29 01:31:16
 * Author: ysj
 * Description:  string 长度
 */

#include <iostream>
#include <string>
using std::string;

int main() {
  string line;
  while (getline(std::cin, line)) {
    if (line.empty()) {
      std::cout << "input is empty" << std::endl;
    } else {
      string::size_type sz = line.size();
      std::cout << "input size is " << sz << std::endl;
    }
  }
  return 0;
}
