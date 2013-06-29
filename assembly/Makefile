CROSS_COMPILE ?= arm-none-eabi-

CC := $(CROSS_COMPILE)gcc
OBJCOPY := $(CROSS_COMPILE)objcopy
OBJDUMP := $(CROSS_COMPILE)objdump

CFLAGS += -T link.ld -march=armv7-m -mcpu=cortex-m3 -mthumb
CFLAGS += -nostdlib

all:

%.bin : %.S
	$(CC) $(CFLAGS) -o $*.o $<
	$(OBJCOPY) -O binary $*.o $@
	$(OBJDUMP) -d -j .text -j .data -j .bss $*.o > $*.dump

clean:
	rm *.o *.bin *.dump
