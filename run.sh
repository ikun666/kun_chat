#!/bin/bash
user/rpc/userrpc -f user/rpc/etc/user.yaml&
user/api/userapi -f user/api/etc/user.yaml&
relation/rpc/relationrpc -f relation/rpc/etc/relation.yaml&
relation/api/relationapi -f relation/api/etc/relation.yaml&
chat/api/chatapi -f chat/api/etc/chat.yaml


