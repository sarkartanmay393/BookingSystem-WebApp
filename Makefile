BuildDocker:
	docker build -t roomreservation-golang .
StartDocker:
	docker run -p 8080:8080 --name RR-Golang roomreservation-golang
Run:
	chmod +x run.sh
	./run.sh