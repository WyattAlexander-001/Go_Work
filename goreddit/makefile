.PHONY: postgres adminer migrate clean

postgres:
	docker run --rm -d --name postgres_new --network mynetwork -e POSTGRES_PASSWORD=secret -p 5433:5432 postgres

adminer:
	docker run --rm -d --network mynetwork -p 8081:8080 adminer

migrate:
	docker run --rm --network mynetwork -v $(pwd)/migrations:/migrations migrate-tool -source file://migrations -database postgres://postgres:secret@postgres_new:5433/postgres?sslmode=disable up

clean:
	-docker stop focused_kapitsa || true
	-docker rm focused_kapitsa || true
	-docker stop great_diffie || true
	-docker rm great_diffie || true
	-docker stop postgres_new || true
	-docker rm postgres_new || true
	-docker network rm mynetwork || true
	docker container prune -f
	docker volume prune -f
	docker network prune -f
	docker image prune -a -f
