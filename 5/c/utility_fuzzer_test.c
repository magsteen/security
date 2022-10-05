#include "utility.h"
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

int LLVMFuzzerTestOneInput() {
  char* start_value = "Foo<&>Bar";
  char* end_value = swap(start_value);
  free(end_value);

  return 0;
}
