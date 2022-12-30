# NAT - Solution

## Assignment 1

[rEX]
```
root@hijiki51-60000:/# attach rEX
minion@rEX:/$ config
[edit]
minion@rEX# set nat source rule 1 outbound-interface ens4 ;グローバルネットワークとの接続点
minion@rEX# set nat source rule 1 source address 192.168.XXX.0/24 ;NATを適用する送信元ネットワークの範囲

minion@rEX# set nat source rule 1 translation address masquerade

minion@rEX# commit
minion@rEX# save
[edit]
minion@rEX# exit
exit
```

[r1]
```
root@hijiki51-60000:/# attach r1
minion@r1:/$ config
[edit]
minion@r1# set protocols static route 0.0.0.0/0 next-hop 192.168.XXX.1 ;送信先ネットワークに応じて次のノードを指定

minion@r1# commit
minion@r1# save
minion@r1# exit
minion@r1:/$ exit
root@hijiki51-60000:/# ping 8.8.8.8
PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
64 bytes from 8.8.8.8: icmp_req=1 ttl=55 time=1.59 ms
64 bytes from 8.8.8.8: icmp_req=2 ttl=55 time=1.06 ms
^C
--- 8.8.8.8 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 1001ms
rtt min/avg/max/mdev = 1.068/1.332/1.596/0.264 ms
```