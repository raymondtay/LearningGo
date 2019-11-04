#include <stdio.h>
#include "callC.h"

void cHello() {
  printf("Hello from the C-world!\n");
}

void printMessage(char* message) {
  printf("Go send me %s\n", message);
}

