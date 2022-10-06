ImageBuild:
	docker build -t roomreservation-golang .
Start:
	docker run -p 8080:8080 --name RR-Golang roomreservation-golang
