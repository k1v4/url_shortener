all: docker_build docker_run

docker_build:
	docker build -t url_shortener:last .

docker_run:
	docker run -p 8080:8080 -p 50051:50051 --name url_shortener url_shortener:last /app --db "postgres"