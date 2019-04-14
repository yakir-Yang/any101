#include <stdio.h>
#include <stdlib.h>

#define spwan_rand()			\
({					\
	int r = rand() % 100;		\
	printf("Spwan number %d\n", r);	\
	r;				\
})

int main()
{
	return spwan_rand();
}
