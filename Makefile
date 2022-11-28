startPostgres:
	sudo docker container start postgres

createDB:
	sudo docker exec -it postgres createdb --username=postgres --owner=postgres kurdi_go

dropDB:
	sudo docker exec -it postgres dropdb --username=postgres kurdi_go

migrateUp:
	migrate -path infrastructure/database/migrations -database "postgresql://postgres:123456789@10.5.0.2:5432/kurdi_go?sslmode=disable" -verbose up

migrateDown:
	migrate -path infrastructure/database/migrations -database "postgresql://postgres:123456789@10.5.0.2:5432/kurdi_go?sslmode=disable" -verbose down

# make createMigration name=migration_name
createMigration:
	migrate create -ext sql -dir infrastructure/database/migrations -seq $(name)

buildUp:	startPostgres createDB migrateUp

