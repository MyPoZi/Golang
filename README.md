# Golang

## usage

URIのvalueパラメータの数値によってにゃーんの数が変わります.


http://mypozi.com:8080/api/nya-n?value=1

![value1](https://github.com/MyPoZi/Nya-nAPI/blob/master/value1.png)


http://mypozi.com:8080/api/nya-n?value=100
![value100](https://github.com/MyPoZi/Nya-nAPI/blob/master/value100.png)


1回のリクエストで最大100000まで"にゃーん"生成可能です.

実験で100万送った結果約1秒、1000万の場合10秒で返ってきました(ゴルーチン使わず).

## build

- centos
- go 1.11.1
```
$ git clone https://github.com/MyPoZi/Nya-nAPI
$ cd Nya-nAPI
$ go build main.go
$ firewall-cmd --zone=public --add-port=8080/tcp --permanent
$ firewall-cmd --reload
$ nohup ./main &
```
通常の./mainだと起動したユーザがログアウト後プロセスが消えます.
nohupだとlogout後も起動するがpsに表示されないため、止め方が知っている限り再起動しか無いです^^;

docker-composeの場合はユーザログアウト後もコンテナが建っています.
また止め方も
```
$ docker-compose stop
```
で簡単に止められます.
以下はdockerの場合のbuild方法です.

```
$ git clone https://github.com/MyPoZi/Nya-nAPI
$ cd Nya-nAPI
$ go build main.go
$ firewall-cmd --zone=public --add-port=8080/tcp --permanent
$ firewall-cmd --reload
$ docker-compose up &
```


## memo
sudo yum install git-all
sudo yum install -y golang

特に設定していないとサーバ自体を停止したらプロセス消えるのでアクセスできなくなる.