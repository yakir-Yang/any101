#include <stdio.h>

int main()
{
	printf("2. main function\n");
}

__attribute__((constructor)) static void before1()
{
	printf("1. constructor '%s' funcion\n", __func__);
}

__attribute__((constructor)) static void before2()
{
	printf("1. constructor '%s' funcion\n", __func__);
}
