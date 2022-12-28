#!/bin/bash

sudo apt update
sudo apt upgrade -y
sudo apt install -y ca-certificates curl gnupg lsb-release
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt install -y openvswitch-common openvswitch-switch docker-ce docker-ce-cli containerd.io docker-compose-plugin


if [ ! -f ~/.commands ]; then
    curl https://raw.githubusercontent.com/hijiki51/InternetArchLecture/main/setup/.bashrc >> ~/.commands
    echo "source ~/.commands" >> ~/.bashrc
    source ~/.commands
fi

source ~/.commands
seq 1 6 | xargs -IXXX docker run -d --name rXXX --hostname=rXXX --net=none --privileged -v /lib/modules:/lib/modules 2stacks/vyos:latest /sbin/init
docker run -d --name rEX --hostname=rEX --net=host --privileged -v /lib/modules:/lib/modules 2stacks/vyos:latest /sbin/init
# docker run -d --name ns --hostname=ns --net=host --privileged  -v named:/etc/bind -v lib_bind:/var/lib/bind -v cache_bind:/var/cache/bind ubuntu/bind9:latest


nic_full_reset
add_server r4 s1
add_server r4 s2
add_server r4 s3
add_server rEX sEX
