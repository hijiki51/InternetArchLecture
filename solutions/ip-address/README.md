# IP Address - Solution

## Assignment 1

[rEX]
```
root@hijiki51-60000:/# attach rEX
minion@rEX:/$ config
[edit]
minion@rEX# set interfaces ethernet eth10 address 192.168.XXX.1/30
[edit]
minion@rEX# commit
[edit]
minion@rEX# save
Done
[edit]
minion@rEX# exit
exit
minion@rEX:/$ exit
exit
```

[r1]
```
root@hijiki51-60000:/home/hijiki51# attach r1
minion@r1:/$ config
[edit]
minion@r1# set interfaces ethernet eth12 address 192.168.XXX.2/30
[edit]
minion@r1# commit
[edit]
minion@r1# save
Done
[edit]
minion@r1# exit
exit
minion@r1:/$ exit
exit
```