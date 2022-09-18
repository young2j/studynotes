/*
 * File: struct-sales_data.cc
 * Created Date: 2022-09-14 09:32:11
 * Author: ysj
 * Description: struct 类-成员函数
 */

#include <istream>
#include <ostream>
#include <string>

struct Sales_data {
  // 成员函数
  std::string isbn() const { return bookNo; } // 类内声明、类内定义
  Sales_data &combine(const Sales_data &);
  double avg_price() const;
  // 数据成员
  std::string bookNo;
  unsigned units_sold = 0;
  double revenue = 0.0;
};
// 声明非成员接口函数
Sales_data add(const Sales_data &, const Sales_data &);
std::ostream &print(std::ostream &, const Sales_data &);
std::istream &read(std::istream &, Sales_data &);

// 类外定义成员函数
double Sales_data::avg_price() const {
  if (units_sold) {
    return revenue / units_sold;
  }
  return 0;
}

Sales_data &Sales_data::combine(const Sales_data &rhs) {
  units_sold += rhs.units_sold;
  revenue += rhs.revenue;
  return *this; // 返回调用此函数的对象
}

// 定义非成员函数
std::istream &read(std::istream &is, Sales_data &item) {
  double price = 0;
  is >> item.bookNo >> item.units_sold >> price;
  item.revenue = price * item.units_sold;
  return is;
}

std::ostream &print(std::ostream &os, const Sales_data &item) {
  os << item.isbn() << " " << item.units_sold << " " << item.revenue << " "
     << item.avg_price() << std::endl;
  return os;
}