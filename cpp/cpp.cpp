
extern "C" {
#include "cpp.h"
}
#include <iostream>

extern "C"{
extern void Logo(char *s);
}

void SayCppHello() {
    std::cout << "Say Hello From Cpp";
    Logo((char*)"Say Hello From Cpp From Go");
}