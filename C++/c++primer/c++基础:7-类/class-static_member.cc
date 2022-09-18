/*
 * File: class-static_member.cc
 * Created Date: 2022-09-15 08:33:38
 * Author: ysj
 * Description: class类-静态成员
 */

#include <string>
class Account {
public:
  Account() = default;
  Account(std::string owner = "", double amount = 0.0){};
  void calculate() { amount += amount * interest_rate; }
  static double rate() { return interest_rate; } // 公有静态成员函数
  static void rate(double);                      // 公有静态成员函数

private:
  std::string owner;
  double amount;
  static constexpr int period = 30; // 静态常量表达式成员，可以类内初始化
  static double interest_rate; // 私有静态数据成员
  static double init_rate();   // 私有静态成员函数
};

// 即使一个常量静态成员在类内部初始化了，通常也应该在类外部定义一下该成员
constexpr int Account::period;
// 静态成员不能在类内部初始化
double Account::interest_rate = 0.05;
// 定义静态成员函数
double Account::init_rate() { return 0.05; }
void Account::rate(double new_rate) { interest_rate = new_rate; }
