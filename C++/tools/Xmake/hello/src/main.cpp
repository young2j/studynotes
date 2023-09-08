#include <iostream>

using namespace std;

template <typename... T> auto sum(T... t) { return (t + ...); }

int main(int argc, char **argv) {
  cout << sum(1, 2, 3, 4, 5) << endl;
  return 0;
}
