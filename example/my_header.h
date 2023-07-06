#ifndef MY_HEADER_H
#define MY_HEADER_H

#include <stdbool.h>

#define _POSIX_C_SOURCE 1
#define NOMINMAX 

enum Day {
	Saturday,
	Sunday = 0,
	Monday,
	Wednesday
};

struct Point2 {
	double x;
	double y;
};

typedef enum Day WeekDay;

typedef struct {
	int x;
	int y;
	int z;
} Point;

extern Point my_point;

unsigned char *const foo(int i, int (*func)(int));
void bar(const char*, ...);
void baz(...);


#endif
