/*
 * File: iterator-arithmetic.cc
 * Created Date: 2022-09-06 03:14:38
 * Author: ysj
 * Description: 迭代器的算术运算
 */
#include <iostream>
#include <string>
using std::string;

int main() {
  const string text = "a b c d e f g h i j k l m n o";
  const char target = 'h';
  string::const_iterator beg = text.cbegin(), end = text.cend();
  string::difference_type half_len = (end - beg) / 2;
  string::const_iterator mid = beg + half_len;

  while (mid != end && *mid != target) {
    if (target < *mid) {
      end = mid;
    } else {
      beg = mid + 1;
    }
    mid = beg + (end - beg) / 2;
  }

  std::cout << *mid << std::endl;
  // 如果mid==end则没有找到目标元素
  std::cout << (mid == end) << std::endl;

  return 0;
}
