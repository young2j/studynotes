#include "MathFunctions.h"
// include cmath
#include <cmath>
// Wrap the mysqrt include in a precompiled ifdef based on USE_MYMATH
#ifdef USE_MYMATH
#include "mysqrt.h"
#endif

namespace mathfunctions {
double sqrt(double x) {
#ifdef USE_MYMATH
  return detail::mysqrt(x);
#else
  return std::sqrt(x);
#endif
}
} // namespace mathfunctions
