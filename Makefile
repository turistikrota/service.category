build:
	docker build --build-arg GITHUB_USER=${TR_GIT_USER} --build-arg GITHUB_TOKEN=${TR_GIT_TOKEN} -t github.com/turistikrota/service.category . 

run:
	docker service create --name category-api-turistikrota-com --network turistikrota --secret jwt_private_key --secret jwt_public_key --env-file .env --publish 6022:6022 github.com/turistikrota/service.category:latest

remove:
	docker service rm category-api-turistikrota-com

stop:
	docker service scale category-api-turistikrota-com=0

start:
	docker service scale category-api-turistikrota-com=1

restart: remove build run
