/*
 * File: vector-.cc
 * Created Date: 2022-09-06 02:11:14
 * Author: ysj
 * Description: vector 基本操作
 */

#include <iostream>
#include <vector>
using std::vector;

int main() {
  vector<int> ivec;

  if (ivec.empty()) {
    for (int i = 0; i < 10; ++i) {
      ivec.push_back(i);
    }
  }

  for (auto &v : ivec) {
    v *= v;
  }

  for (vector<int>::size_type ix = 0; ix != ivec.size(); ++ix) {
    std::cout << ivec[ix] << std::endl;
  }

  return 0;
}
