run:
	docker compose up -d #--build
stop:
	docker compose down
	#docker rmi apps-app_mpa
	#docker rmi apps-app_spa
	#docker rmi nginx
	#docker rmi src-app
	#nginx - s reload
	#docker stop postgres
sudo_run_docker:
	sudo systemctl start docker
rm_all:
	docker stop $$(docker ps -a -q)
	docker rm $$(docker ps -a -q)