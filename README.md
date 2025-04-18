## 前置条件

1. 需要一个服务器
2. 配置数据库 用户名 密码
   
我用的是postgres数据库

连接控制台指令
`psql -h <你的公网ip> -U postgres -d <数据库名称> -p 5432`

开放端口

3. 配置MQTT协议

我是用的是centos8 配置教程

[centos8配置MQTT](https://blog.csdn.net/yuanpan1987/article/details/107816237)

## 程序使用

根据单片机发送的数据 设置相应属性，设计相应的结构体接收数据 配置好数据库的连接和MQTT的使用

