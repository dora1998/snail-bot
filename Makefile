.PHONY: deploy
deploy:
	docker-compose --context conoha -f docker-compose.prod.yaml stop
	docker-compose --context conoha -f docker-compose.prod.yaml up --build -d