# DNS-2 Solution

## Assignment

### 親側の設定

```
$TTL 60
... (省略)
{子のtraQ ID}                  IN      A {子のネームサーバーのGlobal IP}

sub                            IN      NS ns.{子のtraQ ID}.{親のtraQ ID}.
ns.{子のtraQ ID}.{親のtraQ ID}. IN      A {子のネームサーバーのGlobal IP}
```
```
zone "{子のtraQ ID}.{親のtraQ ID}" IN {
        type master;
        file "/etc/bind/rr/{子のtraQ ID}";
};
```


### 子側の設定
```
zone "{子のtraQ ID}.{親のtraQ ID}" IN {
        type master;
        file "/etc/bind/rr/{親のtraQ ID}";
};
```

```
$TTL 60
@       IN      SOA ns.{子のtraQ ID}.{親のtraQ ID}. ns.{子のtraQ ID}.{親のtraQ ID}. (
        2;
        600;
        600;
        600;
        600;
);

        IN      NS ns.{子のtraQ ID}.{親のtraQ ID}.
ns      IN      A {子のネームサーバーのGlobal IP}
```