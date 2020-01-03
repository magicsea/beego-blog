beego-blog
==========
blog use beego framework

## install

```
go get github.com/ulricqin/beego-blog/goutils
go get github.com/magicsea/beego-blog
cd beego-blog && modify conf/app.conf
bee run
```

## admin 

```
url: localhost/me
username: magicsea
password: 111111
```

## 改动
- 因为编译不过，移除qiniu相关依赖
- 稍微修改了展示页