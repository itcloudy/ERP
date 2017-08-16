#golangERP
后台：框架采用国人谢大开发的beego

前端:webpack2+vue2+vuex+vue-router

简化数据库表设计，取消表的创建者、更新者和用户的直接关联关系，orm上为操作者的ID，而非对象

clone工程到go的src目录下，工程文件夹的名字必须为golangERP，若要修改名字需要将代码中所有golangERP修改为工程文件夹的名字

前端支持桌面和移动端，服务器根据请求区分（不全面），web_pc针对的是pc端，web_mobile针对的是移动端

增加了前后台，后台地址以admin开始

在golangERP\web_pc和golangERP\web_mobile 目录下执行:npm install & npm run build 

回到golangERP目录下执行：bee run 

默认端口为8888

默认开启了https

域名为www.hechihan.com 本机修改hosts文件

生成crt文件

gopath下src/crypto/tls/generate_cert.go

go run generate_cert.go -host www.hechihan.com

##QQ群
![](http://i.imgur.com/fxfcP6k.png)

##捐赠

#微信
![](http://i.imgur.com/ScbDcOW.jpg)

#支付宝
![](http://i.imgur.com/3zoIh5S.jpg)

