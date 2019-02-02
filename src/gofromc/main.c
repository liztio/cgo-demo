#include <stdlib.h>
#include <stdio.h>

#include "_cgo_export.h"

typedef struct {
  float x;
  float y;
} Point;

Point *getPoint() {
  Point *pt = malloc(sizeof(Point));
  pt->x = 10;
  pt->y = 12;
  return pt;
}

void usePoint(Point *pt) {
  printf("x: %f, y: %f\n", pt->x, pt->y);
}

int main() {
  char *name = "FOSDEM";
  // Greet(name);

  Point *pt = getPoint();
  usePoint(pt);

  return 0; // exit code
}

