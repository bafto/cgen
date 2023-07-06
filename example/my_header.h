#include <stdbool.h>

struct Point2 {
	double x;
	double y;
};

typedef struct {
	int x;
	int y;
	int z;
} Point;

extern Point my_point;

unsigned char *const foo(int i, int (*func)(int));

