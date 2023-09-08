/*
 * File: class-access_specifiers2.cc
 * Created Date: 2022-09-15 05:36:33
 * Author: ysj
 * Description: class类型-访问控制、类型成员、可变成员
 */

#include <cstddef>
#include <ostream>
#include <string>
class Screen {
public:
  // 类型成员—先声明后使用
  typedef std::string::size_type pos;

  // 构造函数
  Screen() = default; // 默认构造
  Screen(pos ht, pos wd, char c)
      : height(ht), width(wd), contents(ht * wd, c){}; // 重载构造

  // 成员函数
  char get() const { return contents[cursor]; } // 隐式内联--定义在内部
  inline char get(pos ht, pos wd) const; // 显式内联--类内, 成员函数重载
  Screen &move(pos r, pos c);            // 返回*this，类外显式内联
  Screen &set(char);                     // 返回*this，类外显式内联
  Screen &set(pos, pos, char);           // 返回*this，类外显式内联

  Screen &display(std::ostream &os) {
    do_display(os);
    return *this;
  } // 非常量版本
  const Screen &display(std::ostream &os) const {
    do_display(os);
    return *this;
  } // 常量版本

  void incr_access() const; // const对象内也能修改access_ctr；

private:
  pos cursor = 0;
  pos height = 0, width = 0;
  std::string contents;
  mutable std::size_t access_ctr; // 可变数据成员, 记录成员函数被调用的次数

  // 显示screen的内容
  void do_display(std::ostream &os) const { os << contents; }
};

// 成员函数-类外显式内联
inline Screen &Screen::move(pos r, pos c) {
  pos row = r * width;
  cursor = row + c;
  return *this;
}

// 成员函数-重载
char Screen::get(pos r, pos c) const {
  pos row = r * width;
  return contents[row + c];
}

// 修改可变数据成员
void Screen::incr_access() const { ++access_ctr; }

// 返回*this，类外显式内联
inline Screen &Screen::set(char c) {
  contents[cursor] = c;
  return *this;
}
inline Screen &Screen::set(pos row, pos col, char ch) {
  contents[row * width + col] = ch;
  return *this;
}