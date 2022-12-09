comando de migração do banco de dados do Golang

migrate -source file:///go/app/sql/migrations -database 'mysql://root:root@tcp(mysql:3306)/cartola' up