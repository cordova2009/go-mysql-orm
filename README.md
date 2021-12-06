# go-mysql-orm

Go-mysql-orm is a simple mysql ORM for Go.

## Drivers Support

Drivers for Go's sql package which currently support database/sql includes:

* Mysql: [github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)

## Installation

	go get github.com/cordova2009/go-mysql-orm

## Quick Start

* Create Engine

```Go
engine, err := mysql.New(driverName, dataSourceName)
```

* Define a struct

```Go
type User struct {
Id int64
Name string
Age int
}

```

* Query

```Go

users, err := mysql.Query("select * from user")

user, err := mysql.Query("select * from user where Id=?", id)

```

* Save

```Go
user := new(User)
user.Id = 899
user.Name = "zhangsan"
user.Age = 18
_, err := mysql.Save(user)


```

* Update

```Go
user := new(User)

user.Name = "zhangsan"
user.Age = 20
_, err := mysql.Update(user, "Id", 899)


```


