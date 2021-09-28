|  | cli命令 | 迁移 | 回滚 | 版本控制 | 支持数据库 | 优缺点说明 | github星数 | Go Report Card评级
| ------ | ------ | ------ | ------ | ------ | ------ | ------ | ------ | ------ |
| golang-migrate | 是 | 是 | 是 | 是 | PostgreSQL，PGX，SQLite，Redshift，SQLite3，SQLCipher，MySQL/ MariaDB，MongoDB，MS SQL Server，Cassandra，CrateDB 等更多查看 https://github.com/golang-migrate/migrate| 支持数据库驱动多、支持版本控制、跳版本迁移回滚、支持命令、go脚本执行、操作历史记录、使用人数多 | 7.2k | A+                                            
| goose | 是 | 是 | 是 | 是 |postgres，mysql，sqlite3，mssql，redshift | 支持数据库驱动多、支持版本控制、跳版本迁移回滚、支持命令 | 1.9k | A+
| gormigrate | 否| 是 | 是 | 是 | PostgreSQL，MySQL，SQLite ，Microsoft SQL Server | gorm的迁移功能上支持了数据回滚，版本控制功能，历史记录，迁移/回滚到制定版本，很好兼容gorm；缺点支持数据库驱动不多，使用人不多，不支持命令操作 | 675 | A+
| gorm | 否 | 是 | 否 | 否 | MySQL、PostgreSQL、SQLite、SQL Server | gorm自带迁移工具,优点：简洁，兼容项目里使用gorm。缺点：没有版本控制，支持数据库驱动不多。| | 
| darwin | 否 | 是 | 否 | 否 | MySQL、PostgreSQL、SQLite、SQL Server | 优点：代码简洁，记录操作历史。缺点：只支持向上迁移。| 121 | A+ 