# cgen
cgen is RCD internal golang framework scaffold

## Installation

    go install github.com/41443658/cgen@latest 

## Initializing a project

    cgen init . 
    cgen init /var/www/blog/

## Creating a model

    cgen model -d "root:password@tcp(127.0.0.1:3306)/db_name?charset=utf8mb4&parseTime=True&loc=Local" -t bundle

这里-d代表连接到数据库的dsn信息， -t代表数据表的名字,通过gentool自动生成表结构的结构体

## Creating a api

    cgen api bundle

## Creating a service

    cgen service bundle

## Creating a dao

    cgen dao bundle