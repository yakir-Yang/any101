#include <stdio.h>

int main()
{
	printf("1. main function\n");
}

__attribute__((destructor)) static void after1()
{
	printf("2. destructor '%s' funcion\n", __func__);
}

__attribute__((destructor)) static void after2()
{
	printf("2. destructor '%s' funcion\n", __func__);
}
