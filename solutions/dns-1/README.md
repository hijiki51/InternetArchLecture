# DNS-1 Solution

## Assignment 1(例)

NICには事前にIPアドレスを割り当てておきます。

まずはサーバーに割り振られたIPアドレスを確認します。（DHCPであるため手動確認）
```
root@s1:/# ip address
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
61: ens4@if62: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether 7a:53:1d:45:76:56 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 192.168.0.129/28 brd 192.168.0.143 scope global dynamic ens4
       valid_lft 86332sec preferred_lft 86332sec
```

ゾーンの設定をします。
`/etc/bind/named.conf.local`
```
//
// Do any local configuration here
//

// Consider adding the 1918 zones here, if they are not used in your
// organization
//include "/etc/bind/zones.rfc1918";

zone "hijiki51" IN {
        type master;
        file "/etc/bind/rr/hijiki51";
};
```

各種リソースレコードの設定をします。
`/etc/bind/rr/hijiki51`

```
$TTL 60
@       IN      SOA ns.{自分のtraQ ID}. root.{自分のtraQ ID}. (
                        1;
                        600;
                        600;
                        600;
                        600;
                );
        IN      NS ns.{自分のtraQ ID}.

ns      IN      A 192.168.0.38
server  IN      A 192.168.0.129
```
<!-- 講師用
```
$TTL 60
.               IN      SOA     ns.root. ns.root. (
                                3;
                                600;
                                600;
                                600;
                                600;
                        );
.               IN      NS      ns.root.
ns.root.        IN      A       {ルートネームサーバーのGlobal IP}
hijiki51.       IN      NS      ns.hijiki51.
ns.hijiki51.    IN      A       {受講者のGlobal IP}
``` -->

## Assignment 2

`dig @localhost server.{あなたのtraQ ID}`
