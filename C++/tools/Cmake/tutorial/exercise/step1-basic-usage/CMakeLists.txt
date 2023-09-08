# Set the minimum required version of CMake to be 3.10
cmake_minimum_required(VERSION 3.10)
# Create a project named Tutorial and Set the project version number as 1.0
project(Tutorial VERSION 1.0)

# Set the variable CMAKE_CXX_STANDARD to 11 and the variable
# CMAKE_CXX_STANDARD_REQUIRED to True
set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED True)

# Use configure_file to configure and copy TutorialConfig.h.in to
# TutorialConfig.h
configure_file(TutorialConfig.h.in TutorialConfig.h)

# Add an executable called Tutorial to the project
add_executable(Tutorial tutorial.cxx)

# Use target_include_directories to include ${PROJECT_BINARY_DIR}
target_include_directories(Tutorial PUBLIC "${PROJECT_BINARY_DIR}")


# 源文件搜索
# aux_source_directory(${PROJECT_SOURCE_DIR} SRC)
# file(GLOB SRC ${CMAKE_CURRENT_SOURCE_DIR}/*.cpp)

# 指定输出目录
# set(EXECUTABLE_OUTPUT_PATH "./bin")
# set(LIBRARY_OUTPUT_PATH "./lib")

# CMAKE_PREFIX_PATH 搜索配置文件的基本路径列表
message("CMAKE_PREFIX_PATH: ${CMAKE_PREFIX_PATH}")
# CMAKE_INSTALL_PREFIX 安装路径前缀
message("CMAKE_INSTALL_PREFIX: ${CMAKE_INSTALL_PREFIX}")
# CMAKE_MODULE_PATH 搜索模块文件的基本路径列表
  # Make project-provided Find modules available
  # list(APPEND CMAKE_MODULE_PATH "${CMAKE_CURRENT_SOURCE_DIR}/cmake")
message("CMAKE_MODULE_PATH: ${CMAKE_MODULE_PATH}")
# CMAKE_PROJECT_TOP_LEVEL_INCLUDES 设置依赖项提供程序，必须在project()之前，且整个项目只能使用一个提供程序

# -- 使用 FetchContent 从源代码下载和构建
# -- 可以使用 FIND_PACKAGE_ARGS 到 FetchContent_Declare() 选项指示它愿意通过任一方法接受依赖项。
# -- 这允许 FetchContent_MakeAvailable() 尝试首先调用 find_package() 来满足依赖关系.
# -- 如果找不到依赖项，则改为从源代码构建。
# include(FetchContent)
# FetchContent_Declare(
#   googletest
#   GIT_REPOSITORY https://github.com/google/googletest.git
#   GIT_TAG        703bd9caab50b139428cea1aaff9974ebee5742e # release-1.10.0
#   FIND_PACKAGE_ARGS NAMES GTest # 首先调用 find_package(googletest NAMES GTest)
# )
# FetchContent_Declare(
#   Catch2
#   GIT_REPOSITORY https://github.com/catchorg/Catch2.git
#   GIT_TAG        605a34765aa5d5ecbf476b4598a862ada971b0cc # v3.0.1
# )
# FetchContent_MakeAvailable(googletest Catch2)


# -- 项目可以通过 OVERRIDE_FIND_PACKAGE 关键字来强制从源代码生成特定的依赖项。
# -- 然后，对该依赖项的 find_package() 调用将重定向到 FetchContent_MakeAvailable() 。
# FetchContent_Declare(
#   Catch2
#   URL https://intranet.mycomp.com/vendored/Catch2_2.13.4_patched.tgz
#   URL_HASH MD5=abc123...
#   OVERRIDE_FIND_PACKAGE
# )

# -- The following is automatically redirected to FetchContent_MakeAvailable(Catch2)
# find_package(Catch2)