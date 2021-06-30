module github.com/dimitarkovachev/golang-at-ocado/proj/sort/sorting-service

go 1.16

replace github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen => ../gen

require (
	github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen v0.0.0-20210622173111-d85be6bed5d6
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.38.0
)
