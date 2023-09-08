#include <iostream>
#include <string>
#include <yajl/yajl_parse.h>

static int yajl_callback_null(void *ctx) {
  std::cout << "null" << std::endl;
  return 1;
}

static int yajl_callback_boolean(void *ctx, int boolVal) {
  std::cout << (boolVal ? "true" : "false") << std::endl;
  return 1;
}

static int yajl_callback_integer(void *ctx, long long integerVal) {
  std::cout << integerVal << std::endl;
  return 1;
}

static char *yajl_callback_string(void *ctx, const unsigned char *str,
                                  size_t len) {
  std::cout << std::string((const char *)str, len) << std::endl;
  return nullptr;
}

int main() {
  std::string json = "{\"key\": 123}";
  return 0;
}
