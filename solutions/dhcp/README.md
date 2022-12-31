# DHCP - Solution

## Assignment 1

[r4]
```
root@hijiki51-60000:/# attach r4
minion@r4:/$ config
[edit]
minion@r4# set interfaces ethernet eth100 address 192.168.0.142/28

minion@r4# set protocols ospf area 0 network 192.168.0.128/28

minion@r4# set protocols ospf passive-interface eth100

minion@r4# set service dhcp-server shared-network-name dhcp_scope_01 subnet 192.168.0.128/28 default-router 192.168.0.142 ; 送信先ネットワークに対してデフォルトルート(今回はDHCPホストサーバー)を設定
minion@r4# set service dhcp-server shared-network-name dhcp_scope_01 subnet 192.168.0.128/28 range 0 start 192.168.0.129
minion@r4# set service dhcp-server shared-network-name dhcp_scope_01 subnet 192.168.0.128/28 range 0 stop 192.168.0.139 ; DHCPで使用するネットワークとその中で割り振る範囲を設定

minion@r4# commit
minion@r4# save
[edit]
minion@r4# exit
exit

minion@r4/$ show dhcp server statistics

Pool                      Pool size   # Leased    # Avail
----                      ---------   --------    -------
dhcp_scope_01             10          0           10
```

`show dhcp server statistics`で割り振りが行われているか確認可能です。

[s1~s3]
```
root@s1:~# dhclient ens4
```
でDHCPの再リースが可能です。

割り振り前後で`ip address`コマンドの結果を比較してみると良いでしょう。
