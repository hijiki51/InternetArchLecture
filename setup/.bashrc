attach() {	
    name=$1	
    if [ "${name:0:1}" = "r" ]; then	
        docker exec -it --user 1000 $name /bin/bash	
    else	
        docker exec -it $name /bin/bash	
    fi	
}

add_nic() {
    ovs-docker add-port $1 eth$3 $2
    if [ $? != 0 ]; then
        add_nic $1 $2 $(($3+1))
    fi
}

connect() {    
    cn1=$1
    cn2=$2    
    ovs-vsctl add-br br-$1-$2 
    add_nic br-$1-$2 $1 10      
    add_nic br-$1-$2 $2 10  
}

reset_nic() {    
    ovs-docker del-ports dummy $1
}

add_server() {
    router_name=$1
    container_name=$2
    ovs-vsctl add-br br-$router_name-server
    ovs-docker add-port br-$router_name-server eth100 $router_name
    docker run -d --name $container_name --hostname=$container_name --net=none --privileged ubuntu:20.04 /bin/sh -c "while :; do sleep 1000; done"
    ovs-docker add-port br-$router_name-server eth0 $container_name 
}

nic_full_reset() {
    docker start $(docker ps -qa)
    
    seq 1 6 | xargs -I XXX docker exec rXXX bash -c "echo '127.0.0.1 rXXX' >> /etc/hosts"
    docker exec rEX bash -c "echo '127.0.0.1 rEX' >> /etc/hosts"
    # docker exec ns bash -c "echo '127.0.0.1 ns' >> /etc/hosts"
    
    reset_nic r1
    reset_nic r2
    reset_nic r3
    reset_nic r4
    reset_nic r5
    reset_nic r6
    reset_nic rEX
    # reset_nic ns    

    connect r1 r6
    connect r1 r2
    connect r2 r3
    connect r2 r5
    connect r3 r4
    connect r4 r5
    connect r5 r6
    connect r1 rEX
    connect r6 rEX
    # connect r4 ns
    
    add_nic br-r4-server r4 100
    add_nic br-rEX-server rEX 100
}

full_reset() {
    docker ps -qa | xargs docker rm -f
    docker network prune

    seq 1 6 | xargs -IXXX docker run -d --name rXXX --hostname=rXXX --net=none --privileged -v /lib/modules:/lib/modules ghcr.io/hijiki51/internetarchlecture/vyos:1.2 /sbin/init
    docker run -d --name rEX --hostname=rEX --net=host --privileged -v /lib/modules:/lib/modules ghcr.io/hijiki51/internetarchlecture/vyos:1.2 /sbin/init
    # docker run -d --name ns --hostname=ns --net=host --privileged  -v named:/etc/bind -v lib_bind:/var/lib/bind -v cache_bind:/var/cache/bind ubuntu/bind9:latest

    nic_full_reset
    add_server r4 s1
    add_server r4 s2
    add_server r4 s3
    add_server rEX sEX
}
