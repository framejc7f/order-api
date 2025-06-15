run:
	docker-compose up --build

dbshell:
	docker exec -it $(docker-compose ps -q db) psql -U user -d orders
