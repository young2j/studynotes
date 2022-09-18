/*
 * File: string.cc
 * Created Date: 2022-08-29 12:57:52
 * Author: ysj
 * Description: 读写string
 */

#include <iostream>
#include <string>
using std::string;

int main() {
  string s1, s2;

  std::cin >> s1 >> s2;

  std::cout << "s1: " << s1 << std::endl;
  std::cout << "s2: " << s2 << std::endl;

  // 反复读取直至文件末尾
  string word;
  while (std::cin >> word) {
    std::cout << word << std::endl;
  }

  return 0;
}
