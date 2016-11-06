

PLATFORM=macosx-x64

ifeq ($(PLATFORM), android-arm)

CROSS_PREFIX = arm-linux-androideabi-
ARCH         = arm
CFLAGS       += --sysroot=$(SYSROOT) -Wall -O3
CFLAGS       += -Wno-unused-variable
LDFLAGS      += --sysroot=$(SYSROOT)

else ifeq ($(PLATFORM), mingw-x86)

CROSS_PREFIX = i486-mingw32-
ARCH         = i386
CFLAGS       += -Wall -O3
CFLAGS       += -Wno-unused-variable
LDFLAGS      +=


else ifeq ($(PLATFORM), linux-x86)

ARCH         = i386
CFLAGS       += -Wall -O3 -m32 -fPIC
CFLAGS       += -Wno-unused-variable
LDFLAGS      += -m32 -fPIC
LDFLAGS      += -lpthread

else ifeq ($(PLATFORM), macosx-x64)

ARCH         = x86_64
CFLAGS       += -DAE_CPU=AE_INTEL -DAE_COMPILER=AE_GNUC
CFLAGS       += -Wall -O3 -m64 -fPIC -Wno-unused-variable
LDFLAGS      += -m64 -fPIC -lstdc++

endif

PREFIX                                          = ../target
EXEC                                            = libalg

CC                                              = $(CROSS_PREFIX)gcc
AR                                              = $(CROSS_PREFIX)ar
LD                                              = $(CROSS_PREFIX)gcc
RANLIB                                          = $(CROSS_PREFIX)ranlib
STRIP                                           = $(CROSS_PREFIX)strip
CFLAGS                                          +=


