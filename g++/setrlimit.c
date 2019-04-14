#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/resource.h>

void *orig_stack_pointer;

#define STACK_LIMIT 10*1024*1024UL	// 10M

void blow_stack()
{
	blow_stack();
}

/*
 * Try to set the process stack size.
 *
 * 1. You need to compiled this program with debug option flag
 *      $ g++ -g setrlimit.c -o serlimit
 *
 * 2. After that, you need to run this program under GDB
 *      $ gdb ./setrlimit
 *
 * 3. Just run the program, and then GDB would revice the segme-
 *    ntation fault. Great, almost success now!
 *
 * 4. You need to print the current stack pointer address.
 *      (gdb) print (void *)$esp
 *    Also print the original stack pointer address.
 *      (gdb) print (void *)orig_stack_pointer
 *
 * 5. Dala, you can cacluate the real stack size now!
 */
int main()
{
	struct rlimit new_rlimit;
	int status;

	/*
	 * Set the process stack size limit to STACK_LIMIT.
	 */
	memset(&new_rlimit, 0, sizeof(new_rlimit));
	new_rlimit.rlim_cur = STACK_LIMIT;
	new_rlimit.rlim_max = STACK_LIMIT;

	status = setrlimit(RLIMIT_STACK, &new_rlimit);
	if (status == -1) {
		perror("setrlimit");
		exit(1);
	}

	status = getrlimit(RLIMIT_STACK, &new_rlimit);
	if (status == -1) {
		perror("getrlimit");
		exit(1);
	}

	printf("New stack limit: %d\n", (int)new_rlimit.rlim_cur);

	/*
	 * Storge the current stack point to orig_stack_pointer
	 */
	__asm__("movl %esp, orig_stack_pointer");


	/* Try to break the stack, and reveive the SIGSEGV of
	 * segmentation fault. Like
	 *
	 * Program received signal SIGSEGV, Segmentation fault.
	 * blow_stack () at setrlimit.c:13
	 * 13		blow_stack();
	 */
	blow_stack();


	/* After we meet the segment fault, then we can dump the
	 * orignal stack pointer and crash stack pointer, like
	 *
	 * $1 = (void *) 0xffffffffff5ff000
	 * (gdb) print (void *)orig_stack_pointer
	 * $2 = (void *) 0xffffc7d0
	 *
	 * Through previous information, we can cacluate the real
	 * proccess stack size, like
	 *
	 * stack_size = 0xFFFFC7D0 - 0xFF5FF000 = 0x9FD7D0 = 10M
	 */

	exit(0);
}
