# Golang

## usage

URIのvalueパラメータの数値によってにゃーんの数が変わります


http://mypozi/api/nya-n?value=1

()[]


http://mypozi/api/nya-n?value=100
()[]


1回のリクエストで最大100000まで"にゃーん"生成可能です

実験で100万送った結果約1秒、1000万の場合10秒で返ってきました(ゴルーチン使わず)

## build

- centos
```

$ git clone https://github.com/MyPoZi/Nya-nAPI
$ cd Nya-nAPI
$ go build main.go
$ firewall-cmd --zone=public --add-port=8080/tcp --permanent
$ firewall-cmd --reload
$ ./main &

```