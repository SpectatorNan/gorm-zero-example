
## generate model script
```shell
goctl model mysql datasource -url="root:localpwd@tcp(localhost:3306)/gormzero" -table="users"  -dir="services/model" -cache=true --style=goZero --home ./template
```