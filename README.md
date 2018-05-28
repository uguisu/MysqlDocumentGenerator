# Mysql Document Generator

[![][license img]][license]

## # About
This project can automatically create a table design document of database in markdown format.

## # Include
- golang 1.9+
- mysql 5.6+

## # Get Start
### 1) Create two tables in Mysql

The first one is _GenderCode_
```sql
CREATE TABLE `GenderCode` (
	`GENDER_CODE` INT(11) NOT NULL COMMENT 'Code',
	`GENDER_NAME` VARCHAR(45) NOT NULL COMMENT 'Name' COLLATE 'utf8_unicode_ci',
	`GENDER_ORDER` TINYINT(1) NOT NULL COMMENT 'Display order',
	PRIMARY KEY (`GENDER_CODE`)
)
COMMENT='Gender code management'
COLLATE='utf8_unicode_ci'
ENGINE=InnoDB
;
```

The second one is _User_

```sql
CREATE TABLE `User` (
	`USER_ID` VARCHAR(6) NOT NULL COMMENT 'User ID=CC-xxx' COLLATE 'utf8_unicode_ci',
	`AGE` TINYINT(3) NULL DEFAULT NULL COMMENT 'age',
	`SEX` TINYINT(1) NULL DEFAULT NULL COMMENT '0=Male; 1=Female; 2=Other; 3=unknow',
	`USER_TYPE` TINYINT(1) NOT NULL COMMENT '1=system; 2=wechat; 3=qq; 4=zhihu; 5=weibo; 6=linkedin',
	PRIMARY KEY (`USER_ID`)
)
COMMENT='User basic information'
COLLATE='utf8_unicode_ci'
ENGINE=InnoDB
;
```

### 2) Build project
Use `go build` to build a executable file.

```shell
go build ./src/github.com/uguisu/main/mysqlDocumentGenerator.go ./src/github.com/uguisu/main/const-value.go
```

### 3) Execute `mysqlDocumentGenerator`

```shell
$ mysqlDocumentGenerator
```

### 4) Check the output result [Database.md](Database.md)

[license]:<http://www.apache.org/licenses/LICENSE-2.0>
[license img]:https://img.shields.io/badge/license-Apache%202-blue.svg
