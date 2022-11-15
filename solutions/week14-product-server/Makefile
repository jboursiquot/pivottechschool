default:

run-server:
	go run main.go -file ../../cmd/server/products.json

arg ?= -i

get-products:
		@curl $(arg) -X GET http://localhost:8080/products

id ?= 1

get-product:
		@curl $(arg) -X GET http://localhost:8080/products/$(id)

name ?= "Pencil"
price ?= 2
update-product:
		@curl $(arg) -X PUT http://localhost:8080/products/$(id) \
			-d '{"name": $(name), "price": $(price)}'

create-product:
		curl $(arg) -X POST http://localhost:8080/products \
			-d '{"name": $(name), "price": $(price)}'

delete-product:
		@curl $(arg) -X DELETE http://localhost:8080/products/$(id)