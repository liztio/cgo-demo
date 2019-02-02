#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <errno.h>

#include "addlib.h"

int main(int argc, char** argv) {
    if (argc != 2) {
        fprintf(stderr, "Need exactly one argument, got %d\n", argc - 1);
        return 1;
    }

    char *end;
    errno = 0;
    uint32_t var = (uint32_t) strtol(argv[1], &end, 10);
    switch errno {
        case ERANGE:
        case EINVAL:
          fprintf(stderr, "Invalid number %s\n", argv[1]);
          return 1;
    }

    printf("C says: %d + one is %d\n", var, add_one(var));
}
