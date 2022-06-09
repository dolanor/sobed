#include <stdio.h>
#include <stdlib.h>

#include "greet.h"

char* greet(char* name) {
	char* g = malloc(256);
	sprintf(g, "Hello, %s", name);
	return g;
}
