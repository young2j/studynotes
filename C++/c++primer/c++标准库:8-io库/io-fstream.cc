/*
 * File: io-fstream.cc
 * Created Date: 2022-09-19 01:48:30
 * Author: ysj
 * Description: 文件流
 */

#include <fstream>
#include <iostream>
#include <string>

int main(int argc, const char **argv) {
  std::string in_file = "io-fstream.cc";
  std::string out_file = in_file + ".copy";

  std::ifstream in_fs;
  in_fs.open(in_file);
  // 上述两行也等价于下面两种写法：
  // std::ifstream in_fs(in_file);
  // std::ifstream in_fs(in_file, std::ifstream::in);

  std::ofstream out_fs;
  out_fs.open(out_file);
  // 上述两行也等价于下面两种写法：
  // std::ofstream out_fs(out_file);
  // std::ofstream out_fs(out_file, std::ofstream::out | std::ofstream::trunc);

  if (in_fs.is_open()) {
    std::string line;
    while (std::getline(in_fs, line)) {
      out_fs << line << std::endl;
    }
  }

  // 关闭文件流对象
  // 也可以不写，当一个文件流对象被销毁时，close会自动调用
  in_fs.close();
  out_fs.close();

  return 0;
}
