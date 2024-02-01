
## gen sql go struct
### install gen tool
`go install github.com/starfishs/sql2struct@latest`
### gen sql struct command
`sql2struct --dsn "mysql://root:123456@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local" -t "users" -t "table_test"`