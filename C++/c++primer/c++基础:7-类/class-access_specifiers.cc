/*
 * File: class-access_specifiers.cc
 * Created Date: 2022-09-15 04:50:03
 * Author: ysj
 * Description:  class类-访问控制
 */

#include <istream>
#include <string>
class Sales_data {
private:
  // 数据成员
  std::string bookNo;
  unsigned units_sold = 0;
  double revenue = 0.0;

public:
  // 构造函数
  Sales_data() = default;                         // 默认构造函数-内联
  Sales_data(const std::string &s) : bookNo(s) {} // 构造函数-初始值列表
  Sales_data(const std::string &s, unsigned n, double p)
      : bookNo(s), units_sold(n), revenue(n * p) {} // 构造函数-初始值列表
  Sales_data(std::istream &); // 构造函数-类内声明，类外定义

  // 成员函数
  std::string isbn() const { return bookNo; }
  Sales_data &combine(const Sales_data &);

private:
  double avg_price() const;

  // 非成员接口函数--友元声明
  friend Sales_data add(const Sales_data &, const Sales_data &);
  friend std::ostream &print(std::ostream &, const Sales_data &);
  friend std::istream &read(std::istream &, Sales_data &);
};
// 声明非成员接口函数
Sales_data add(const Sales_data &, const Sales_data &);
std::ostream &print(std::ostream &, const Sales_data &);
std::istream &read(std::istream &, Sales_data &);

// 类外定义构造函数
Sales_data::Sales_data(std::istream &is) { read(is, *this); }

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
