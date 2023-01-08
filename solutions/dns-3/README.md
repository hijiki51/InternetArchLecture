# DNS-3 Solution

## Assignment 1

`/usr/share/dns/root.hints`
追記または書き換え

```
.             3600000      NS    NS.ROOT.
NS.ROOT.      3600000      A     {ルートネームサーバーの Global IP}
```

## Assignment 2

`/etc/resolv.conf`

```
nameserver {ルートネームサーバーの Global IP}
```