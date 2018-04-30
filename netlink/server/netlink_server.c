// linux-4.4.0-62-generic

#include <linux/init.h>
#include <linux/module.h>
#include <linux/timer.h>
#include <linux/time.h>
#include <linux/types.h>

#include <net/sock.h>
#include <net/netlink.h>

#define NETLINK_TEST 31
#define MAX_MSGSIZE 1024

struct sock *nl_sk = NULL;
int flag = 0;
int err;

// 向用户态进程回发消息
void sendnlmsg(char *message, int pid)
{
	struct sk_buff *skb;
	struct nlmsghdr *nlh;
	int slen;

	if (!message || !nl_sk)
		return;

	printk(KERN_ERR "pid:%d\n", pid);

	skb = nlmsg_new(MAX_MSGSIZE, 0);
	if (!skb)
		printk(KERN_ERR "my_net_link:alloc_skb error\n");

	nlh = nlmsg_put(skb, 0, 0, NLMSG_DONE, MAX_MSGSIZE, 0);

	//NETLINK_CB(skb).pid = 0;
	NETLINK_CB(skb).dst_group = 0;

	slen = strlen(message);

	memcpy(nlmsg_data(nlh), message, slen + 1);

	printk("my_net_link:send message '%s'\n", (char*)nlmsg_data(nlh));

	nlmsg_unicast(nl_sk, skb, pid);
}

// 接收用户态发来的消息
void nl_data_ready(struct sk_buff *_skb)
{
	struct sk_buff *skb;
	struct completion cmpl;
	struct nlmsghdr *nlh;
	char str[100];
	int i = 3;
	int pid;

	printk("begin data_ready\n");

	skb = skb_get(_skb);
	if (skb->len >= NLMSG_SPACE(0))	{
		nlh = nlmsg_hdr(skb);

		memcpy(str, NLMSG_DATA(nlh), sizeof(str));

		printk("Message received:%s\n", str) ;

		pid = nlh->nlmsg_pid;

		// 我们使用completion做延时，每3秒钟向用户态回发一个消息
		while (i--) {
			init_completion(&cmpl);
			wait_for_completion_timeout(&cmpl, 1 * HZ);
			sendnlmsg("I am from kernel!", pid);
		}

		flag = 1;
		kfree_skb(skb);
	}
}

// Initialize netlink
int netlink_init(void)
{
	struct netlink_kernel_cfg cfg = {
		.input		= nl_data_ready,
	};

	nl_sk = netlink_kernel_create(&init_net, NETLINK_TEST, &cfg);
	if (!nl_sk) {
		printk(KERN_ERR "my_net_link: create netlink socketerror.\n");
		return 1;
	}

	printk("my_net_link_4: create netlink socket ok.\n");

	return 0;
}

static void netlink_exit(void)
{
	netlink_kernel_release(nl_sk);

	printk("my_net_link: self module exited\n");
}

module_init(netlink_init);
module_exit(netlink_exit);
MODULE_AUTHOR("yilong");
MODULE_AUTHOR("Kuankuan Yang");
MODULE_LICENSE("GPL");
