#!/bin/bash

EXEC_COMMAND="root/demo/go-client"
export CGO_LDFLAGS=-L/root/openssl/lib
export C_INCLUDE_PATH=/root/openssl
export CPLUS_INCLUDE_PATH=/root/openssl

# 编译目标文件
#go build -o /root/demo/go-client

# 加载环境变量
source /root/demo/configure

#[ -z "$ORGNAMEHASH" ] && echo "you need write orghash in configure" && exit 1


cat >> /etc/hosts <<EOF
$EIP node-$ORGNAMEHASH-0.node-$ORGNAMEHASH.default.svc.cluster.local
$EIP node-$ORGNAMEHASH-1.node-$ORGNAMEHASH.default.svc.cluster.local
EOF

PEERS="node-$ORGNAMEHASH-0,node-$ORGNAMEHASH-1"
PEER0="node-$ORGNAMEHASH-0"
PEER1="node-$ORGNAMEHASH-1"
CONSENSUS_PEERS="node-$ORGNAMEHASH-0"
CONFIG_DIR="/root/demo/config"

# 发起交易
/$EXEC_COMMAND contract send -n $PEER0 -s $CONSENSUS_PEERS -c $CHAIN_NAME -g $CONFIG_DIR/$SDK_CONFIG_FILE  -t $CONTRACT_NAME -f initMarble  -a "$1;$2"

# 查询链上交易记录,只能在某一个节点上查询
/$EXEC_COMMAND contract query -n $PEER0 -c $CHAIN_NAME -g $CONFIG_DIR/$SDK_CONFIG_FILE  -t $CONTRACT_NAME -f getMarble  -a "$3"
