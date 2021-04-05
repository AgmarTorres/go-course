go mod init modulo

go build

./modulo

go run main.go

Se uma visão começa com letra minusctula ela não pode ser importada por outros pacotes

go get github.com/badoux/checkmail

go mod tidy // remove dependencias que não estão sendo utilizadas
