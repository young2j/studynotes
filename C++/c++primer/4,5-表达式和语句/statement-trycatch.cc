/*
 * File: statement-trycatch.cc
 * Created Date: 2022-09-08 03:35:09
 * Author: ysj
 * Description:
 */

#include <iostream>
#include <stdexcept>
#include <string>
int main(int argc, const char **argv) {
  std::string item1, item2;
  while (std::cin >> item1 >> item2) {
    try {
      if (item1 != item2) {
        throw std::runtime_error("not the same input items!");
      }
    } catch (std::runtime_error err) {
      std::cout << err.what() << "\nTry again? Enter y or n" << std::endl;
      char c;
      std::cin >> c;
      if (!std::cin || c == 'n') {
        break;
      }
    }
  }
  
  return 0;
}
