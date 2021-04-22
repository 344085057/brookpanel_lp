#暂停使用 brook出现断流情况  暂不维护

# brookpanel_lp
brookpanel_lp

基于Beego。开发的Brook流控Web面板

体验 http://brook.ws

## Web搭建教程
0.把myBrookWeb.tar.gz上传并解压到服务器上

1.安装mysql

1.1.sql文件夹下的sql语句执行

1.2.mysql开启远程连接

```linux
		不会请Google
```

2.安装Redis(并且开启远程访问) 不安装的话也可以 登录错误次数限制功能用不了和流量使用历史记录

```linux
		不会请Google
```

3.设置服务器的时区

```linux
		不会请Google
```

4.获取reCAPTCHA v2的sitekey和secret

```linux
		不会请Google
```

5.修改conf/app.conf

5.1.配置mysql(通过步骤1获取信息)

5.2.配置redis(通过步骤2获取信息)

5.3.配置remote(随机输入即可)

5.4.配置regin_money(注册时默认金币)

5.5.配置appname(站点名称) 等等...

5.6.配置recaptcha v2验证(通过步骤4获取信息)

5.7.配置nginx （经供参考）

```nginx
	location / {
        proxy_pass          http://127.0.0.1:60002/;  #需要反代的地址
        proxy_redirect      off;
        proxy_set_header    X-Real-IP       $remote_addr;
        proxy_set_header    X-Forwarded-For $proxy_add_x_forwarded_for;
    }
```

5.你需要把项目跑起来！
```linux
		 nohup ./myBrookWeb &
```

6.你需要开始搭建服务端了

https://github.com/pingProMax/brookserver

### 说明

**默认管理员账号:admin@gmail.com 密码：123456 连接密码：brookadmin**


以上步骤无特殊说明，需必做，否则跑不起来

后端端口范围 1024-60000 因此这个项目的用户量最大为 60000-1024

服务器上的其他应用程序也`不要`占用此范围的端口

目前还没有自动删除长期未使用的用户，所以你只需要知道就好


### 站在巨人的肩上

BeeGo: https://github.com/astaxie/beego

Brook: https://github.com/txthinking/brook

Jquery

Bootstrap

Mysql

Redis
