/*
 * File: array-ptrop.cc
 * Created Date: 2022-09-07 04:09:09
 * Author: ysj
 * Description: 数组指针的运算
 */

#include <cstddef>
#include <iostream>
#include <iterator>
#include <string>
using std::string;

int main(int argc, const char **argv) {
  const unsigned n = 3;
  string nums[n] = {"one", "two", "three"};
  string *beg = std::begin(nums);
  string *end = std::end(nums);

  // 移动首元素指针
  string *second_num = nums + 1;
  std::cout << "nums第二个元素: " << *second_num << std::endl;

  // 指针相减，ptrdiff_t是带符号类型
  std::ptrdiff_t sz = end - beg;
  std::cout << "nums含有" << sz << "个元素" << std::endl;

  // 指针遍历
  while (beg != end) {
    std::cout << *beg << std::endl;
    ++beg;
  }

  // 数组下标操作等价于指针操作
  string third_num = nums[2]; // 等价于
  string *numsprt = nums;
  third_num = *(numsprt + 2);
  std::cout << "nums第三个元素: " << third_num << std::endl;

  string *second = &nums[1];
  string third = second[1];  // 等价于 *(second+1)
  string first = second[-1]; // 等价于 *(second-1)
  std::cout << "first: " + first << std::endl;
  std::cout << "second: " + *second << std::endl;
  std::cout << "third: " + third << std::endl;

  return 0;
}
