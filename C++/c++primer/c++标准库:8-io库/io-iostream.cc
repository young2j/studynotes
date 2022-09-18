/*
 * File: io-iostream.cc
 * Created Date: 2022-09-19 01:47:59
 * Author: ysj
 * Description: io流
 */

#include <ios>
#include <iostream>
#include <ostream>

std::ostream &print(std::ostream &os) {
  os << "print some info" << std::endl;
  return os;
}

int main(int argc, const char **argv) {
  std::ostream &os = print(std::cout);
  os << "main end" << std::endl;

  return 0;
}
