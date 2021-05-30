.PHONY: postgres adminer migrate
postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=dima postgres

adminer:
	docker run --rm -ti --network host adminer

migrate:
	migrate -source file://migrations \ 
			-database postgres://postgres:dima@locakhost/postgres?sslmode=disable up