startPostgres:
	sudo docker container start postgres

createDB:
	sudo docker exec -it postgres createdb --username=postgres --owner=postgres kurdi_go

dropDB:
	sudo docker exec -it postgres dropdb --username=postgres kurdi_go

migrateUp:
	migrate -path database/migrations -database "postgresql://postgres:123456789@localhost:5432/kurdi_go?sslmode=disable" -verbose up

migrateDown:
	migrate -path database/migrations -database "postgresql://postgres:123456789@localhost:5432/kurdi_go?sslmode=disable" -verbose down

createMigration:
	migrate create -ext sql -dir database/migrations -seq $(name)

buildUp:	startPostgres createDB migrateUp

