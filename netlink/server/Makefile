ifneq ($(KERNELRELEASE),)

obj-m := netlink_server.o

else

KERNELDIR ?= /lib/modules/$(shell uname -r)/build
PWD := $(shell pwd)

default:
	$(MAKE) -C $(KERNELDIR) M=$(PWD) modules

clean:
	rm -rf .*.cmd *.ko *.o *.mod.c Module.symvers modules.order .tmp_versions
endif
