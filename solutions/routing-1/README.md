# Routing 1 - Solution

## Assignment 1

以降、解答ではIPアドレスを以下のように設定し、NICに適切に割り振っています。
各自の環境に合わせて読み替えるようにしてください。

![IP Setting](/assets/ip-setting.drawio.svg)

[rEX]
```
root@hijiki51-60000:/# attach rEX
minion@rEX:/$ config
[edit]
minion@rEX# set protocols static route 192.168.0.0/28 next-hop 192.168.0.2

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
minion@r1# set protocols static route 192.168.0.8/30 next-hop 192.168.0.6
minion@r1# set protocols static route 192.168.0.16/30 next-hop 192.168.0.6
minion@r1# commit
minion@r1# save
[edit]
minion@r1# exit
exit
```

[r2]
```
root@hijiki51-60000:/# attach r2
minion@r2:/$ config
[edit]
minion@r2# set protocols static route 0.0.0.0/0 next-hop 192.168.0.5
minion@r2# commit
minion@r2# save
[edit]
minion@r2# exit
exit
```

[r3]
```
root@hijiki51-60000:/# attach r3
minion@r3:/$ config
[edit]
minion@r3# set protocols static route 0.0.0.0/0 next-hop 192.168.0.9

minion@r3# commit
minion@r3# save
[edit]
minion@r3# exit
exit
```

[r5]
```
root@hijiki51-60000:/# attach r5
minion@r5:/$ config
[edit]
minion@r5# set protocols static route 0.0.0.0/0 next-hop 192.168.0.17
minion@r5# commit
minion@r5# save
[edit]
minion@r5# exit
exit
```
