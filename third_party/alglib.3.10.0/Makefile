
#######################################################################################
#                                                                                     #
# Description:                                                                        #
#     厦大海西通讯工程中心，Makefeil格式以及源代码目录规范                                  #
#     $(PWD)                                                                          #
#       |--src    源代码目录，src目录下以及一级子目录中的c文件参与编译                       #
#       |--objs   编译c源码文件生成.o文件目录                                             #
# 	    |--bin    最终生成库文件或可执行文件                                              #
#                                                                                     #
# Author：                                                                            #
#     eric.woo (wadahana@gmail.com)                                                   #
#                                                                                     #
# Version:                                                                            #
#     Ver0.1.0 2010-08-25  初次制定                                                    #
#     Ver0.2.0 2010-10-23  支持src下两级目录                                            #
#     Ver0.3.0 2011-02-28  确定config.mk包含格式                                        #
#     Ver0.4.0 2011-04-23  目标文件名由目录名生成                                        #
#     Ver0.5.0 2011-06-22  增加连接外部库文件                                            #
#     Ver0.6.0 2011-10-03  增加自动生成BIN目录和objs目录，增加对cpp文件和m文件的支持         #
#     Ver0.6.0 2016-11-01  开放生成动态库，                                              #
#                                                                                     #
#######################################################################################

include config.mk


CFLAGS +=  -I$(SRC_DIR) -Wall
LDFLAGS +=

INC_DIR = $(PREFIX)/include/
SRC_DIR = ./src
OBJ_DIR = ./objs
BIN_DIR = ./bin



######### 检查库文件或可执行文件名  ##########
ifndef EXEC
	EXEC = $(shell basename `pwd`)
endif

SHARED_LIB = $(EXEC).so
STATIC_LIB = $(EXEC).a


######### 检查安装路径  ##########
ifndef PREFIX
	PREFIX = /usr
endif

DST_INC_DIR = $(PREFIX)/include/$(EXEC)


SRC_SUB_DIRS = $(shell find $(SRC_DIR) -maxdepth 1 -type d)


INC_SUB_DIRS :=  $(foreach path, $(SRC_SUB_DIRS),$(subst $(SRC_DIR), $(DST_INC_DIR), $(path)))

SRC += $(foreach path, $(SRC_SUB_DIRS), $(wildcard $(path)/*.c))
OBJS += $(foreach path, $(SRC_SUB_DIRS), $(patsubst $(path)/%.c, $(OBJ_DIR)/%.o, $(wildcard $(path)/*.c)))

CXXSRC += $(foreach path, $(SRC_SUB_DIRS), $(wildcard $(path)/*.cpp))
CXXOBJS += $(foreach path, $(SRC_SUB_DIRS), $(patsubst $(path)/%.cpp, $(OBJ_DIR)/%.o, $(wildcard $(path)/*.cpp)))

OOCSRC += $(foreach path, $(SRC_SUB_DIRS), $(wildcard $(path)/*.m))
OOCOBJS += $(foreach path, $(SRC_SUB_DIRS), $(patsubst $(path)/%.m, $(OBJ_DIR)/%.o, $(wildcard $(path)/*.m)))

SRC_HEADER += $(foreach path, $(SRC_SUB_DIRS), $(wildcard $(path)/*.h))
DST_HEADER += $(foreach path, $(SRC_HEADER),$(subst $(SRC_DIR), $(DST_INC_DIR), $(path)))



VPATH += $(SRC_SUB_DIRS)



all: _mkdir $(BIN_DIR)/$(STATIC_LIB) $(BIN_DIR)/$(SHARED_LIB)
	@echo '  [CHMOD]  755  '$(BIN_DIR)/$(SHARED_LIB) ; chmod 755 $(BIN_DIR)/$(SHARED_LIB)
	@echo '  [CHMOD]  755  '$(BIN_DIR)/$(STATIC_LIB) ; chmod 755 $(BIN_DIR)/$(STATIC_LIB)

_mkdir:
	@echo '  [MKDIR]  '$(BIN_DIR); mkdir -p $(BIN_DIR)
	@echo '  [MKDIR]  '$(OBJ_DIR); mkdir -p $(OBJ_DIR)

$(BIN_DIR)/$(SHARED_LIB): $(OBJS) $(CXXOBJS) $(OOCOBJS)
	@echo '  [LD]  '		$@;	$(CC) $(LDFLAGS) -shared -o $@ $^  $(LDFLAGS)

$(BIN_DIR)/$(STATIC_LIB): $(OBJS) $(CXXOBJS) $(OOCOBJS)
	@echo '  [AR]  '		$@;	$(AR) cru $@ $(OBJS) $(CXXOBJS)


#添加了定义$(SRC_SUB_DIRS)到VPATH中，则make可以在VPATH中搜索到%.c文件
$(OBJS) : $(OBJ_DIR)/%.o : %.c
	@echo '  [CC]  '    $@ ;	$(CC) -c $(CFLAGS) -c $< -o $@

$(CXXOBJS) : $(OBJ_DIR)/%.o : %.cpp
		@echo '  [CC]  '    $@ ;  $(CC) -c $(CFLAGS) -c $< -o $@

$(OOCOBJS) : $(OBJ_DIR)/%.o : %.m
		@echo '  [CC]  '    $@ ;  $(CC) -c $(CFLAGS) -c $< -o $@

install: _install-mkdir install-lib install-headers

_install-mkdir:
	@echo '  [MKDIR]  '$(PREFIX)/lib       ; mkdir -p $(PREFIX)/lib
	@echo '  [MKDIR]  '$(PREFIX)/include   ; mkdir -p $(PREFIX)/include

install-lib: $(BIN_DIR)/$(STATIC_LIB) #$(BIN_DIR)/$(SHARED_LIB)
	@echo '  [INSTALL]  '$(BIN_DIR)/$(SHARED_LIB) ; cp $(BIN_DIR)/$(SHARED_LIB) $(PREFIX)/lib/$(SHARED_LIB)
	@echo '  [INSTALL]  '$(BIN_DIR)/$(STATIC_LIB) ; cp $(BIN_DIR)/$(STATIC_LIB) $(PREFIX)/lib/$(STATIC_LIB)

install-headers: $(INC_SUB_DIRS) $(DST_HEADER)


# 以下进行inc下子目录中的文件拷贝 #
define PROGRAM_TEMPLATE_HEADERS
SRC_TMP = $(subst $(DST_INC_DIR),$(SRC_DIR), $(1))
$(1): $$(SRC_TMP)
endef

$(foreach prog,$(DST_HEADER), $(eval $(call PROGRAM_TEMPLATE_HEADERS,$(prog))))

$(DST_HEADER):
		@echo '  [INSTALL]  '  $@ ;install -m 777  $<  $@


###### 进行inc下子目录创建 #######

define PROGRAM_TEMPLATE_INC_SUBDIRS
$(1):
endef

$(foreach prog,$(INC_SUB_DIRS), $(eval $(call PROGRAM_TEMPLATE_INC_SUBDIRS,$(prog))))

$(INC_SUB_DIRS):
	@echo '  [MKDIR]  -p '$@ ; mkdir -p $@

distclean: clean
	@echo '  [RM]  '$(DST_INC_DIR)* ;							rm -rf $(DST_INC_DIR)*
	@echo '  [RM]  '$(BIN_DIR)/* ;								rm -rf $(BIN_DIR)/*

clean:
	@echo '  [RM]  '$(OBJ_DIR)/*.o ;				    rm -f $(OBJ_DIR)/*.o
	@echo '  [RM]  '$(BIN_DIR)/$(SHARED_LIB);		    rm -f $(BIN_DIR)/$(SHARED_LIB)
	@echo '  [RM]  '$(BIN_DIR)/$(STATIC_LIB);		    rm -f $(BIN_DIR)/$(STATIC_LIB)
	@echo '  [RM]  '$(PREFIX)/lib/$(SHARED_LIB);		rm -f $(PREFIX)/lib/$(SHARED_LIB)
	@echo '  [RM]  '$(PREFIX)/lib/$(STATIC_LIB);		rm -f $(PREFIX)/lib/$(STATIC_LIB)


#$(warning current dir is: $(shell basename `pwd`))
