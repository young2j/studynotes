/*
 * File: string-cctype.cc
 * Created Date: 2022-08-29 01:58:28
 * Author: ysj
 * Description:  字符操作
 */

#include <cctype>
#include <iostream>
#include <string>
using std::cout;
using std::endl;
using std::string;

int main() {
  string s = "123...Hello, world!";

  for (auto c : s) {
    if (std::isalpha(c)) {
      cout << c << " 是一个字母 ";
      if (std::islower(c)) {
        char c_upper = std::toupper(c);
        cout << "大写为 " << c_upper << endl;
      }
      if (std::isupper(c)) {
        char c_upper = std::tolower(c);
        cout << "小写为 " << c_upper << endl;
      }
    }

    if (std::isdigit(c)) {
      cout << c << " 是一个数字" << endl;
    }

    if (std::ispunct(c)) {
      cout << c << " 是一个标点符号" << endl;
    }
  }

  // 使用范围迭代
  for (auto &c : s) {
    c = std::toupper(c);
  }
  cout << "全大写 " << s << endl;

  // 使用下标迭代
  for (decltype(s.size()) i = 0; i != s.size() && !std::isspace(s[i]); i++) {
    s[i] = std::tolower(s[i]);
  }
  cout << "一半小写 " << s << endl;

  return 0;
}
