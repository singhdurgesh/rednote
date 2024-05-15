# note: call scripts from /scripts
run:
	go run main.go server

start_worker:
	go run main.go worker

docker_rm:
	docker-compose rm -f

docker_down:
	docker-compose down

docker_build:
	docker-compose up --build -d

docker_up:
	docker-compose up -d

docker_build_up: docker_down docker_rm docker_build
