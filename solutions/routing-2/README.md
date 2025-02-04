# Routing 2 - Solution

## Assignment 1

各自のネットワークに合わせて変更してください。
[例: r6]
```
root@hijiki51-60000:/# attach r6
minion@r6:/$ config
[edit]
minion@r6# set protocols ospf area 0 network 192.168.0.12/30

minion@r6# set protocols ospf area 0 network 192.168.0.24/30

minion@r6# set protocols ospf area 0 network 192.168.0.28/30

minion@r6# commit
minion@r6# save
[edit]
minion@r6# exit
exit
```

OSPFではルーティングアップデートのために、隣接ルータがいる・いないに関係なく、OSPFが有効化されたインターフェースから定期的にHelloパケットが、隣接ルータにむけてマルチキャストで送信されるます。
しかし、隣接ルータがいない場合は、ルーティングアップデートの送信は必要ありません。
そこで、`passive-interface`コマンドを使用することによって、特定のI/Fからの無駄なアップデートを停止することができます。
rEXやr4では`passive-interface`を設定することにも注意してください。

[rEX]
```
minion@rEX# set protocols ospf passive-interface ens4
```

また、OSPFでもRIPやEIGRPなどのルーティングプロトコルのようにデフォルトルートを生成することができます。
インターネットあてのパケットをルーティングするために、デフォルトルートを利用することが多いです。
他には、スタブエリアから外部ネットワークへの到達性を確保するためにデフォルトルートの生成を行う場合などがあります。

RIPやEIGRPでは、デフォルトルートをスタティックルートとして設定して、再配送することでデフォルトルートの生成できます。
しかし、OSPFではこの方法ではデフォルトルートを生成できません。
スタティックルートをOSPFに再配送しても、デフォルトルートは再配送の対象外です。
OSPFでデフォルトルートを生成するためには、`default-information originate`コマンドを使います。

今回はrEXをデフォルトルートに設定します。

[rEX]
```
minion@rEX# set protocols ospf default-information originate always
```

[INFO]
特別なエリアとして
- Stub Area
- Totaly Stub Area
- NSSA
- Totaly NSSA

が存在します。

## Assignment 2

パスコストを設定することで到達経路を制御できます。

[例: r2]
```
minion@r2# set interfaces ethernet eth11 ip ospf cost 200
```

どのような経路を通っているか確認するには`traceroute`コマンドなどを利用してください。
