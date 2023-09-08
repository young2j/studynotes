# 最低版本要求
cmake_minimum_required(VERSION 3.0.0...3.24)
if(${CMAKE_VERSION} VERSION_LESS 3.12)
  cmake_policy(VERSION ${CMAKE_MAJOR_VERSION}.${CMAKE_MINOR_VERSION})
endif()

# 设置一个项目
project(
  cmaker
  VERSION 0.1.0
  DESCRIPTION "this is a example for cmake study"
  HOMEPAGE_URL "https://xxxx.com"
  LANGUAGES CXX)

include(CTest)
enable_testing()

# ====== 设置变量 =======
# 本地变量
set(CPACK_PROJECT_NAME ${PROJECT_NAME})
set(CPACK_PROJECT_VERSION ${PROJECT_VERSION})
include(CPack)
# 本地变量--列表变量
set(LIST_VAR1 "v1" "v2")
set(LIST_VAR2 "v1;v2")

# 缓存变量--允许从命令行中设置变量
set(CACHE_VAR
    "v"
    CACHE STRING "description")
set(CACHE_VARIABLE
    "VALUE"
    CACHE STRING "" FORCE) # 强制设置该变量的值
mark_as_advanced(CACHE_VARIABLE) # 运行 cmake -L .. 或使用 GUI 界面的时候不会列出该变量

# 命令行选项
option(MY_OPTION "This is settable from the command line" OFF)

# 环境变量--最好避免用
set(ENV{CMAKE_ENV_VAR} "v") # $ENV{CMAKE_ENV_VAR} 获取变量

# ====== 包含目录 =======
# 生成一个库, 如果是 INTERFACE 库, 后面不能有文件名
add_library(demolib1 STATIC demolib1/demo.cc demolib1/demo.h)

# 指定编译目标的包含目录，目标必须已经通过add_executable或者add_library指定 PUBLIC -
# 任何链接到这个目标的目标也必须包含这个目录, 对于一个二进制目标没有什么含义 PRIVATE - 只影响当前目标，不影响依赖 INTERFACE -
# 只影响依赖
target_include_directories(demolib1 PUBLIC include)

# ====== 链接库 =======
add_library(demolib2 STATIC demolib2/demo.cc demolib2/demo.h)
target_link_libraries(demolib2 PUBLIC demolib1)

# ====== 编译特征 =======
target_compile_features(demolib1 PUBLIC cxx_std_11)
target_compile_features(demolib2 PUBLIC cxx_std_11)

# ====生成一个可执行文件====
add_executable(cmaker main.cc)
target_link_libraries(cmaker PUBLIC demolib2)

# =======设置属性=========
set_property(TARGET cmaker PROPERTY CXX_STANDARD 11) # 同下，但更通用
set_target_properties(cmaker PROPERTIES CXX_STANDARD 11)
# 获取属性
get_property(
  CPP_STANDARD
  TARGET cmaker
  PROPERTY CXX_STANDARD) # 同下，但更通用
get_target_property(CPP_STANDARD cmaker CXX_STANDARD)

# =======控制流程=========
# if("${VAR}")  cmake3.1+
if(VAR) 
# If variable is `ON`, `YES`, `TRUE`, `Y`, or non zero number
else()
# If variable is `0`, `OFF`, `NO`, `FALSE`, `N`, `IGNORE`, `NOTFOUND`, `""`, or ends in `-NOTFOUND`
endif()

# =======关键字=========
# 一元: NOT, TARGET, EXISTS(文件), DEFINED
# 二元: STREQUAL, AND, OR, MATCHES ( 正则表达式 ), VERSION_LESS, VERSION_LESS_EQUAL

# =====生成器表达式======
# $<KEYWORD:value> KEYWORD为1时值替换为value，否则不替换
target_include_directories(
  cmaker
  PUBLIC
  $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
  $<INSTALL_INTERFACE:include>
  )
  
# ====函数和参数控制====
function(FUNC REQUIRED_ARG)
    message(STATUS "Simple arguments: ${REQUIRED_ARG}, followed by ${ARGN}, all arguments:${ARGV}")
    set(${REQUIRED_ARG} "NEW VALUE FOR ARG" PARENT_SCOPE)
endfunction()

func("This" "Foo" "Bar")
message("Output: ${This}")

# COMPLEX_PREFIX_SINGLE = TRUE
# COMPLEX_PREFIX_ANOTHER = FALSE
# COMPLEX_PREFIX_ONE_VALUE = "value"
# COMPLEX_PREFIX_ALSO_ONE_VALUE = <UNDEFINED>
# COMPLEX_PREFIX_MULTI_VALUES = "some;other;values"
# COMPLEX_PREFIX_UNPARSED_ARGUMENTS=
function(COMPLEX)
    cmake_parse_arguments(
        COMPLEX_PREFIX
        "SINGLE;ANOTHER"
        "ONE_VALUE;ALSO_ONE_VALUE"
        "MULTI_VALUES"
        ${ARGN}
    )
endfunction()

complex(SINGLE ONE_VALUE value MULTI_VALUES some other values)