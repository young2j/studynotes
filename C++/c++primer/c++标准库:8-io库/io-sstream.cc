/*
 * File: io-sstream.cc
 * Created Date: 2022-09-19 01:48:47
 * Author: ysj
 * Description:  字符串流
 */
#include <iostream>
#include <sstream>
#include <string>
#include <vector>

std::string text = R"(morgan 32397293833 48727373221
drew 74783488881
lee 73283882111 382893238 3232323212)";

struct Person {
  std::string name;
  std::vector<std::string> phones;
};

int main(int argc, const char **argv) {
  // 将string拷贝到字符串输入流中
  std::istringstream istr(text);
  // 上面一行等价于：
  // std::istringstream istr;
  // istr.str(text);

  std::string line, word;
  std::vector<Person> persons;

  // 将text结构化为vector<Person>类型
  while (std::getline(istr, line)) {
    std::istringstream row(line);
    Person person;
    row >> person.name;
    while (row >> word) {
      person.phones.push_back(word);
    }
    persons.push_back(person);
  }
  // 输出vector<Person>的信息
  for (const auto person : persons) {
    // 用两个字符串输出流来写入正确和错误的电话号码
    std::ostringstream ok_phones, bad_phones;
    for (const auto phone : person.phones) {
      if (phone.size() != 11) {
        bad_phones << " " << phone;
      } else {
        ok_phones << " " << phone;
      }
    }
    if (bad_phones.str().empty()) {
      std::cout << person.name << " " << ok_phones.str() << " phone ok."
                << std::endl;
    } else {
      std::cerr << person.name << " "
                << "has bad phones:" << bad_phones.str() << std::endl;
    }
  }

  return 0;
};