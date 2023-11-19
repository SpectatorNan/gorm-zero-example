
tableName =
dir =
cache = true
gen-model:
	goctl model mysql datasource -url="root:localpwd@tcp(127.0.0.1:3306)/gormzero" -table=${tableName}  -dir=${dir} -cache=${cache} --style=go_zero --home=template