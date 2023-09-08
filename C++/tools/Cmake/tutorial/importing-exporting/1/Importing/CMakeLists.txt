cmake_minimum_required(VERSION 3.15)
project(Importing)

# specify the C++ standard
set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED True)

# Add executable
add_executable(myexe IMPORTED)

# Set imported location
set_property(TARGET myexe PROPERTY
             IMPORTED_LOCATION "../InstallMyExe/bin/myexe")

# Add custom command to create source file
add_custom_command(OUTPUT main.cc COMMAND myexe)

# Use source file
add_executable(mynewexe main.cc)


# ====== import library ===========
# import static lib
add_library(foo STATIC IMPORTED)
set_property(TARGET foo PROPERTY
             IMPORTED_LOCATION "/path/to/libfoo.a")
add_executable(myexe src1.c src2.c)
target_link_libraries(myexe PRIVATE foo)

# import shared lib on Windows
add_library(bar SHARED IMPORTED)
set_property(TARGET bar PROPERTY
             IMPORTED_LOCATION "c:/path/to/bar.dll")
set_property(TARGET bar PROPERTY
             IMPORTED_IMPLIB "c:/path/to/bar.lib")
add_executable(myexe src1.c src2.c)
target_link_libraries(myexe PRIVATE bar)

# import a library with multiple configurations
find_library(math_REL NAMES m)
find_library(math_DBG NAMES md)
add_library(math STATIC IMPORTED GLOBAL)
set_target_properties(math PROPERTIES
  IMPORTED_LOCATION "${math_REL}"
  IMPORTED_LOCATION_DEBUG "${math_DBG}"
  IMPORTED_CONFIGURATIONS "RELEASE;DEBUG"
)
add_executable(myexe src1.c src2.c)
target_link_libraries(myexe PRIVATE math)