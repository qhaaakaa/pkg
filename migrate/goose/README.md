## 安装goose

## cli 操作
### 安装goose
```shell
go get -u github.com/pressly/goose/v3/cmd/goose
```

[comment]: <> (### 创建配置文件，goose必须在db目录下创建dbconfig.yaml)

[comment]: <> (```yaml)

[comment]: <> (development:)

[comment]: <> (  driver: mysql)

[comment]: <> (  dsn: root:123456@tcp&#40;127.0.0.1:3306&#41;/gorm?parseTime=ture)
```

### 创建sql脚本
```shell
# 创建脚本文件
goose create create_user sql
goose create create_book sql
goose create create_book_type sql

# 执行迁移命令
goose mysql "root:123456@tcp(127.0.0.1:3306)/gorm?parseTime=true" up        

# 执行回滚命令
goose mysql "root:123456@tcp(127.0.0.1:3306)/gorm?parseTime=true" dwon        

# 指定迁移到某版本
goose mysql "root:123456@tcp(127.0.0.1:3306)/gorm?parseTime=true" up-to 20210927164519

# 指定回滚到某版本
goose mysql "root:123456@tcp(127.0.0.1:3306)/gorm?parseTime=true" down-to 0
```
