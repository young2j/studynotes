/*
 * File: string-+.cc
 * Created Date: 2022-08-29 01:47:23
 * Author: ysj
 * Description:  string 相加
 */

#include <iostream>
#include <string>
using std::string;

int main() {
  string s1 = "hello", s2 = "world";

  string s3 = s1 + "," + s2 + '\n'; // ok
  string s4 = "hello" + "world";    // not ok
  string s5 = s1 + "," + "world";   // ok
  string s6 = "hello" + "," + s2;   // not ok

  return 0;
}
