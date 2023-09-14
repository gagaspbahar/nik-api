#!make
SQLITE_URL = file:./pkg/database/data-wilayah.db
# export $(shell sed 's/=.*//' .env)

migrate-new:
	migrate create -ext sql -dir ./pkg/database/sql/migrations -seq ${name}

migrate-up:
	migrate -database ${SQLITE_URL} -path migrations up ${n}

migrate-down:
	migrate -database ${SQLITE_URL} -path migrations down ${n}

migrate-force:
	migrate -database ${SQLITE_URL} -path migrations force ${version}