# Add a library called MathFunctions with sources MathFunctions.cxx and
# mysqrt.cxx
add_library(MathFunctions STATIC MathFunctions.cxx)

# Create a variable USE_MYMATH using option and set default to ON
option(USE_MYMATH "whether use my math functions" ON)

# If USE_MYMATH is ON, use target_compile_definitions to pass USE_MYMATH as a
# precompiled definition to our source files
if(USE_MYMATH)
  target_compile_definitions(MathFunctions PRIVATE "USE_MYMATH")
endif()

# When USE_MYMATH is ON, add a library for SqrtLibrary with source mysqrt.cxx
if(USE_MYMATH)
  add_library(SqrtLibrary STATIC mysqrt.cxx)
endif()

# When USE_MYMATH is ON, link SqrtLibrary to the MathFunctions Library
if(USE_MYMATH)
  target_link_libraries(MathFunctions PUBLIC SqrtLibrary)
endif()
