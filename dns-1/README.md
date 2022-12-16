# Chapter8: DNS 1

この章ではDNS(Domain Name System)に関する基礎知識とその設定を学びます。

- [Chapter8: DNS 1](#chapter8-dns-1)
	- [Lesson](#lesson)
		- [DNS](#dns)
	- [Assignment](#assignment)
		- [1. ネームサーバーを構築してみよう](#1-ネームサーバーを構築してみよう)
		- [2. 名前解決をしてみよう](#2-名前解決をしてみよう)

## Lesson

### DNS
DNS(Domain Name System)はIPアドレスなどのリソースに対して別名をつけるアプリケーション層プロトコルです。

DNSはドメイン名空間(Domain Name Space)を管理するネームサーバーと名前解決を行うリゾルバからなります。

ドメイン名空間は一つのネームサーバーで一括管理されているのではなく、木構造のように複数ネームサーバーによって分散管理されており、名前解決を行う際は木構造の親から子へと再帰的に問い合わせを行います。

また、一番初めに問い合わせを行う、全体の木構造の親に当たるネームサーバーのことをルートネームサーバーと呼びます。

現在1600を超えるルートネームサーバーが起動しています。

https://root-servers.org/
## Assignment

**[INFO]**
nsはbind9導入済みUbuntuのインスタンスです。
### 1. ネームサーバーを構築してみよう
ネームサーバーを準備し、`server.{あなたの traQ ID}`のレコードを登録してみましょう。

<details>
<summary>ヒント1</summary>
</details>
扱うゾーンは`{あなたの traQ ID}`になるでしょう。(TLDです)
<details>

<summary>ヒント2</summary>
</details>
今回はIPv4を用いているので設定するのは`A`レコードです。
<details>
<summary>ヒント3</summary>
「bind9 Aレコード 設定」などで検索してみるといいでしょう
</details>

### 2. 名前解決をしてみよう
1.ができたら相手のHTTPサーバのIPアドレスを名前解決することで取得してみましょう。
ルートネームサーバーのIPは講師から共有されます。
<details>
<summary>ヒント</summary>
</details>
DNSへのリクエストには`dig`コマンドを用います。
<details>

***

[解答を見る](../solutions/dhcp/README.md)

[TOPへ](../README.md)
