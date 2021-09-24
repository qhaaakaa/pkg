## 数据迁移

### mac 安装
```shell
brew install golang-migrate
```

### github地址 
https://github.com/golang-migrate/migrate


### 执行生成命令 migrate -source file://{目录名称} -database {数据库类型}://{账号}:{密码}@tcp({ip}:{端口})/{数据库名称} up {文件数}
### 生成全部
```shell
migrate -source file://migrate -database mysql://root:123456@tcp(127.0.0.1:3306)/dbname up
```

### 数据回滚全部
```shell
migrate -source file://migrate -database mysql://root:123456@tcp(127.0.0.1:3306)/dbname down
```

### 只生成book 和 user 表
```shell
migrate -source file://migrate -database mysql://root:123456@tcp(127.0.0.1:3306)/dbname up 2
```

### 数据回滚最新的2条
```shell
migrate -source file://migrate -database mysql://root:123456@tcp(127.0.0.1:3306)/dbname down 2
```


### 设置版本 跳过生成book_type data1 data2 直接生成data3 和 更新 user表字段
```shell
migrate -verbose -source file://migrate -database mysql://root:123456@tcp(127.0.0.1:3306)/dbname force 20210924083240
```

```shell
migrate -source file://migrate -database mysql://root:123456@tcp(127.0.0.1:3306)/dbname up
```

### 跳过版本 只回滚book表
```shell
migrate -verbose -source file://migrate -database mysql://root:123456@tcp(127.0.0.1:3306)/dbname force 20210924082109
```

```shell
migrate -source file://migrate -database mysql://root:123456@tcp(127.0.0.1:3306)/dbname up 1
```
