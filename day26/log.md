# ORMS && Databases

So far so good, we've been learning about how to write simple queries on MySQL, learning about how to handle http routers for a custom web server and how to create a JSON file but we've not yet touched on DataBases yet.

So, i would be looking at GORM as it is one of the Golang ORMs for connecting to database and i believe would be beginner friendly.

# How to install
To install the gorm and its driver use the following command
0. create a folder and cd into the folder
1. run `go get -u gorm.io/gorm` to install the orm
2. then run `go get -u gorm.io/driver/sqlite` to install the driver

## Basic Syntax on GORM

### Retrieve Single Objects
```Go
// Get the first record ordered by primary key
db.First(&user)
// SELECT * FROM users ORDER BY id LIMIT 1;

// Get one record, no specified order
db.Take(&user)
// SELECT * FROM users LIMIT 1;

// Get last record, ordered by primary key desc
db.Last(&user)
// SELECT * FROM users ORDER BY id DESC LIMIT 1;

result := db.First(&user)
result.RowsAffected // returns count of records found
result.Error        // returns error or nil


// Using MAPS && interface

// works because model is specified using `db.Model()`
result := map[string]interface{}{}
db.Model(&User{}).First(&result)
// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

// works with Take
result := map[string]interface{}{}
db.Table("users").Take(&result)

// check error ErrRecordNotFound
errors.Is(result.Error, gorm.ErrRecordNotFound)
```

### Retrieving Objects with primary key
```Go
db.First(&user, 10)
// SELECT * FROM users WHERE id = 10;

db.First(&user, "10")
// SELECT * FROM users WHERE id = 10;

db.Find(&users, []int{1,2,3})
// SELECT * FROM users WHERE id IN (1,2,3);
```

If Primary key is a `String` use the `?` character substitution:
```Go
db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";
```

### Retrieving all Objects
```Go
// Get all records
result := db.Find(&users)
// SELECT * FROM users;

result.RowsAffected // returns found records count, equals `len(users)`
result.Error        // returns error
```

### Update Object && record
```Go
db.Model(&User{}).Where("name = ?", batman).Update("name = ?", "lawson")
```

### Deleting records
deleting records using it id field.
```Go
db.Delete(&User{ID: 7})
```
## RESOURCES
1. [GORM - Quick Guide and installation](https://gorm.io/docs/index.html)
2. [Gorm package - go](https://pkg.go.dev/gorm.io/gorm#section-readme)
