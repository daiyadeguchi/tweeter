run:
	sudo docker compose up --build -d --remove-orphans

stop:
	sudo docker stop $$(sudo docker ps -a -q)