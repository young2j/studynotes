/*
 * File: function-ret_init_list.cc
 * Created Date: 2022-09-12 07:14:17
 * Author: ysj
 * Description: 函数-返回初始化列表
 */

#include <iostream>
#include <string>
#include <vector>

std::vector<std::string> ret_initlist() {
  std::string str;
  if (str.empty()) {
    return {"empty"};
  }
  return {str, "others"};
}

int main(int argc, const char **argv) {
  auto strvec = ret_initlist();
  for (auto str : strvec) {
    std::cout << str << std::endl;
  }

  return 0;
}
