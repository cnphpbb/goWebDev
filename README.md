# go_utils
自用GO的一些组合，大部分来源github

组合出个团队开发的基本结构。

每个目录中都有相关的自述文件。

**主要使用的第3方GO package**

|    库    |   描述   |
| ----------| --------------: |
| go get github.com/labstack/echo|REST框架(框架无限制)|
| go get github.com/gogits/gogs/modules/log|日志（已经合并）|
| go get gopkg.in/ini.v1| 配置文件|
| go get github.com/go-xorm/xorm| Xorm|
| go get github.com/cnphpbb/mysql (fork [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql))|MySQL驱动|
| go get github.com/lib/pq|PostgreSQL驱动|
| go get gopkg.in/mgo.v2| Mongodb 驱动    |
| go get -u github.com/garyburd/redigo/redis|Redis驱动|

**目录结构**

	.
	├── conf	
	│   ├── app.ini
	│   └── readme.md
	├── LICENSE
	├── models
	│   ├── base.go
	│   └── readme.md
	├── modules
	│   ├── log
	│   │   ├── conn.go
	│   │   ├── console.go
	│   │   ├── database.go
	│   │   ├── file.go
	│   │   ├── log.go
	│   │   └── smtp.go
	│   ├── readme.md
	│   └── utils
	│       └── setting.go
	├── README.md
	├── routers
	│   ├── base.go
	│   └── Readme.md
	└── server.go	//main golang file

### 联系方式
Email： moqiruyi@gmail.com
QQ: 48474