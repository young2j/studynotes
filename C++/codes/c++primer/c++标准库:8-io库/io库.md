# iostream

`iostream`即标准输入输出流，用于操纵`char`数据，默认关联到用户的控制台窗口。每一种流对象都对应包含一个操纵宽字符`wchar_t`类型数据的版本：

| 头文件     | 类型                | 作用                 |
| ---------- | ------------------- | -------------------- |
| <iostream> | istream, wistream   | 输入流：从流读取数据 |
| <iostream> | ostream, wostream   | 输出流：向流写入数据 |
| <iostream> | iostream, wiostream | 输入输出流：读写流   |

* io对象不能拷贝、不能赋值，因此

  * io操作的函数通常以引用的方式传递和返回流

  * io对象读写会改变其状态，所以传递和返回的引用也不能是`const`的

    ```cpp
    #include <iostream>
    #include <ostream>
    // 以引用的方式传递和返回流
    std::ostream &print(std::ostream &os) {
      os << "print some info" << std::endl;
      return os;
    }
    
    int main(int argc, const char **argv) {
      std::ostream &os = print(std::cout);
      os << "main end" << std::endl;
      return 0;
    }
    ```

* ostream缓冲刷新时机:

  * `main`函数正常结束时，缓冲刷新；

  * 缓冲区满时，缓冲刷新；

  * 使用`endl、flush、ends`操纵符显式刷新；

    ```cpp 
    std::cout << "endl输出内容和一个换行符，然后刷新缓冲" << std::endl;
    std::cout << "ends输出内容和一个空白符，然后刷新缓冲" << std::ends;
    std::cout << "flush输出内容不附加任何字符,然后刷新缓冲" << std::flush;
    ```

  * 使用`unitbuf`操纵符让流在接下来的每次写操作之后都进行一次flush操作；`nounitbuf`操纵符重置流为默认刷新机制；

    ```cpp
    std::cout << unitbuf; // 每次写操作之后都刷新缓冲
    //...
    std::cout << nounitbuf; // 重置为默认刷新机制
    
    std::cerr << "xxx"; // cerr是默认设置了unitbuf的，写到cerr的内容都是立即刷新的
    ```

  * 读写被关联的流时，关联的流缓冲区会被刷新；

    ```cpp
    // cin和cerr都默认关联到cout，当读cin、写cerr时会导致cout缓冲区被刷新
    // 关联伪代码形如：
    cin.tie(&cout);
    cerr.tie(&cout);
    
    // 关联的cout缓冲区被刷新
    cin >> val;
    cerr << "err";
    ```

# fstream

`fstream`为文件流，继承自`iostream`，具有`iostream`类型的所有行为，如可以使用`<< >>`来读写文件。文件流对象包含如下类型:

| 头文件    | 类型                | 作用                         |
| --------- | ------------------- | ---------------------------- |
| <fstream> | ifstream, wifstream | 文件输入流：从文件流读取数据 |
| <fstream> | ofstream, wofstream | 文件输出流：向文件流写入数据 |
| <fstream> | fstream, wfstream   | 文件输入输出流：读写文件流   |

## 文件模式

每个文件流都有一个关联的文件模式：

| 模式   | 含义                                     | 限制                                       |
| ------ | ---------------------------------------- | ------------------------------------------ |
| in     | 以读方式打开                             | 只能对ifstream和fstream对象设定in模式      |
| out    | 以写方式打开                             | 只能对ofstream和fstream对象设定out模式     |
| app    | 以附加方式打开，每次写前均定位到文件末尾 | app模式下，默认也是out 的                  |
| ate    | 打开文件时立即定位到文件末尾             | 可用于任何文件流对象，可与其他任何模式组合 |
| trunc  | 截断文件                                 | 只有out被设定时才能设定trunc               |
| binary | 以二进制的方式进行io                     | 可用于任何文件流对象，可与其他任何模式组合 |

* `ifstream`关联的文件默认为`in`模式;
* `ofstream`关联的文件默认为`out | trunc`模式；
* `fstream`关联的文件默认为`in | out`模式;

```cpp
ofstream out("out_file"); // 等价于
ofstream out("out_file", ofstream::out); // 等价于
ofstream out("out_file", ofstream::out|ofstream::trunc);

ofstream app("app_file", ofstream::app); // 等价于
ofstream app("app_file", ofstream::out|oftream::app)
```

## 特有操作

通过一个文件读写的示例，来说明文件流的特有操作: `open()、is_open()、close()`

```cpp
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
  // 也可以不写，当一个文件流对象被销毁(离开作用域)时，close会自动调用
  in_fs.close();
  out_fs.close();

  return 0;
}
```

# sstream

`sstream`为字符串流，同样继承自`iostream`，具有`iostream`类型的所有行为。字符串流包含如下类型：

| 头文件    | 类型                          | 作用                             |
| --------- | ----------------------------- | -------------------------------- |
| <sstream> | istringstream, wistringstream | 字符串输入流：从字符串流读取数据 |
| <sstream> | ostringstream, wostringstream | 字符串输出流：向字符串流写入数据 |
| <sstream> | stringstream, wsstringstream  | 字符串输入输出流：读写字符串流   |

## 特有操作

通过一个字符串流读写的示例，来说明字符串流的特有操作: `str()`

```cpp
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
```

> 参考书籍：《C++ Primer》第5版