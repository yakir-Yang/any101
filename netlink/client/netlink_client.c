#include <asm/types.h>
#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/socket.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>

#include <linux/netlink.h>
#include <linux/socket.h>

#define MAX_PAYLOAD 1024 // maximum payloadsize
#define NETLINK_TEST 31  // 自定义的协议

int main(int argc, char* argv[])
{
	struct sockaddr_nl src_addr, dest_addr;
	struct nlmsghdr *nlh = NULL;
	struct iovec iov;
	struct msghdr msg;
	int sock_fd;
	int retval;
	int state;
	int state_smg = 0;

	// 1. Create a socket file
	sock_fd = socket(PF_NETLINK, SOCK_RAW, NETLINK_TEST);
	if (sock_fd == -1) {
		printf("error getting socket: %s\n", strerror(errno));
		return -1;
	}

	// 2.1 Create source nl sockaddr
	memset(&src_addr, 0, sizeof(src_addr));
	src_addr.nl_family = AF_NETLINK;
	src_addr.nl_pid = getpid();  // A：设置源端端口号
	src_addr.nl_groups = 0;

	// 2.2 Binding
	retval = bind(sock_fd, (struct sockaddr*)&src_addr, sizeof(src_addr));
	if (retval < 0) {
		printf("bind failed: %s", strerror(errno));
		close(sock_fd);
		return -1;
	}

	// 3.1 Create nl mssage for iov
	nlh = (struct nlmsghdr *)malloc(NLMSG_SPACE(MAX_PAYLOAD));
	if (!nlh) {
		printf("malloc nlmsghdr error!\n");
		close(sock_fd);
		return -1;
	}
	nlh->nlmsg_len = NLMSG_SPACE(MAX_PAYLOAD);
	nlh->nlmsg_pid = getpid(); // C：设置源端口
	nlh->nlmsg_flags = 0;
	strcpy(NLMSG_DATA(nlh), "Hello you!");

	iov.iov_base = (void *)nlh;
	iov.iov_len = NLMSG_SPACE(MAX_PAYLOAD);

	// 3.2 Create dest nl sockaddr
	memset(&dest_addr, 0, sizeof(dest_addr));
	dest_addr.nl_family = AF_NETLINK;
	dest_addr.nl_pid = 0; // B：设置目的端口号
	dest_addr.nl_groups = 0;

	// 3.3 Create mssage (dst nl sockaddr + iov)
	memset(&msg, 0, sizeof(msg));
	msg.msg_name = (void *)&dest_addr;
	msg.msg_namelen = sizeof(dest_addr);
	msg.msg_iov = &iov;
	msg.msg_iovlen = 1;

	// 3.4 Send message
	printf("state_smg\n");
	state_smg = sendmsg(sock_fd, &msg, 0);
	if (state_smg == -1) {
		printf("get error sendmsg = %s\n", strerror(errno));
	}

	// 4.1 Receive message
	printf("waiting received!\n");

	memset(nlh, 0, NLMSG_SPACE(MAX_PAYLOAD));

	while (1) {
		printf("In while recvmsg\n");
		state = recvmsg(sock_fd, &msg, 0);
		if (state < 0) {
			printf("state = %d\n", state);
		}

		printf("Received message: %s\n", (char *)NLMSG_DATA(nlh));
	}

	close(sock_fd);

	return 0;

}
