#include <dirent.h>
#include <stdio.h>
#include <unistd.h>

int main(void)
{
	DIR *dir;
	struct dirent * ptr;

	dir = opendir("./");

	while ((ptr = readdir(dir)) != NULL)
		printf("d_name : %s\n", ptr->d_name);

	closedir(dir);
}
