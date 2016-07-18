#include <stdio.h>
#include <unistd.h>

#define ONE_MB (1024 * 1024)

int main()
{
	/*
	 * The number of processors configured is :4
	 * The number of processors currently online (available) is :4
	 * The pagesize: 4096
	 * The number of pages: 2012061
	 * The number of available pages: 800762
	 * The memory size: 7859 MB
	 * The number of files max opened:: 1024
	 * The number of ticks per second: 100
	 * The max length of host name: 64
	 * The max length of login name: 256
	 */

	printf("The number of processors configured is :%ld\n",
	       sysconf(_SC_NPROCESSORS_CONF));

	printf("The number of processors currently online (available) is :%ld\n",
	       sysconf(_SC_NPROCESSORS_ONLN));

	printf ("The pagesize: %ld\n", sysconf(_SC_PAGESIZE));

	printf ("The number of pages: %ld\n", sysconf(_SC_PHYS_PAGES));

	printf ("The number of available pages: %ld\n",
		sysconf(_SC_AVPHYS_PAGES));

	printf ("The memory size: %lld MB\n",
		(long long)sysconf(_SC_PAGESIZE) * (long long)sysconf(_SC_PHYS_PAGES) / ONE_MB );

	printf("The number of files max opened:: %ld\n",
	       sysconf(_SC_OPEN_MAX));

	printf("The number of ticks per second: %ld\n",
	       sysconf(_SC_CLK_TCK));

	printf("The max length of host name: %ld\n",
	       sysconf(_SC_HOST_NAME_MAX));

	printf("The max length of login name: %ld\n",
	       sysconf(_SC_LOGIN_NAME_MAX));

	return 0;
}
