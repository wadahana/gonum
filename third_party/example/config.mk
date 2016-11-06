


EXEC                                            =
CC                                              = $(CROSS_PREFIX)gcc
ifeq ($(PLATFORM), ios-arm)
AR                                              = ar
else
AR                                              = $(CROSS_PREFIX)ar
endif
LD                                              = $(CROSS_PREFIX)gcc
RANLIB                                          = $(CROSS_PREFIX)ranlib
STRIP                                           = $(CROSS_PREFIX)strip
CFLAGS                                          = -I../target/include
LDFLAGS                                         = -lstdc++
EXTRA_OBJS                                      = ../target/lib/libopenblas.a ../target/lib/libalg.a
