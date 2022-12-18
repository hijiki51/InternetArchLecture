#!/bin/bash

sudo apt update
sudo apt upgrade -y
sudo apt install -y openvswitch-common openvswitch-switch lxc docker.io

lxd init --auto
lxc profile device remove default eth0

if [ ! -f ~/.commands ]; then
    curl https://raw.githubusercontent.com/hijiki51/InternetArchLecture/main/setup/.bashrc >> ~/.commands
    echo "source ~/.commands" >> ~/.bashrc
    source ~/.commands
fi

source ~/.commands
seq 1 6 | xargs -IXXX docker run -d --name rXXX --hostname=rXXX --net=none --privileged -v /lib/modules:/lib/modules 2stacks/vyos:latest /sbin/init
docker run -d --name rEX --hostname=rEX --net=host --privileged -v /lib/modules:/lib/modules 2stacks/vyos:latest /sbin/init
docker run -d --name ns --hostname=ns --net=host --privileged  -v named:/etc/bind -v lib_bind:/var/lib/bind -v cache_bind:/var/cache/bind ubuntu/bind9:latest


nic_full_reset
add_server r4 s1
add_server r4 s2
add_server r4 s3
add_server rEX sEX
