#include <stdio.h>
#include "usedByC.h"

int main(int args, char **argv) {
    GoInt x = 12;
    GoInt y = 23;
    printf("About to call a Go function!\n");
    Printmsg();

    GoInt r = Mul(x, y);
    printf("%d * %d Product: %d\n",(int)x, (int)y, (int)r);
    printf("DONE;\n");
    return 0;
}

// gcc -o willUseGo willUseGo.c ./usedByC.o