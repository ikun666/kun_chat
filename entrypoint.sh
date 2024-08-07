#!/bin/bash
#set -x
#******************************************************************************
# @file    : entrypoint.sh
# @author  : simon
# @date    : 2018-08-28 15:18:43
#
# @brief   : entry point for manage service start order
# history  : init
#******************************************************************************
 
: ${SLEEP_SECOND:=2}
 
wait_for() {
    echo Waiting for $1 to listen on $2...
    while ! nc -z $1 $2; do echo $1 $2 waiting...; sleep $SLEEP_SECOND; done
}
 
#declare DEPENDS
#declare CMD
 
DEPENDS=''
CMD=''
while getopts "d:c:" arg
do
    case $arg in
        d)
            DEPENDS=$OPTARG
            ;;
        c)
            CMD=$OPTARG
            ;;
        ?)
            echo "unkonw argument"
            exit 1
            ;;
    esac
done
 
for var in ${DEPENDS//,/}
do
    host=${var%:*}
    port=${var#*:}
    wait_for $host $port
done
 
eval $CMD