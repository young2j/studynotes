/*
 * File: class-constructor_explicit.cc
 * Created Date: 2022-09-15 07:27:36
 * Author: ysj
 * Description: class类-类类型隐式转换与显式构造函数
 */

#include <string>
class Sales_data {
public:
  // 数据成员
  std::string bookNo;

  // 构造函数
  Sales_data() = default; // 默认构造函数-内联
  // Sales_data(const std::string &s)
  //     : bookNo(s) {} // 构造函数-初始值列表：转换构造函数
  explicit Sales_data(const std::string &s)
      : bookNo(s) {} // 构造函数-初始值列表：显式构造函数

  // 成员函数
  Sales_data &combine(const Sales_data &);
};