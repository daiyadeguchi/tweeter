dev:
	sudo docker compose up --build -d

stop:
	sudo docker stop $$(sudo docker ps -a -q)