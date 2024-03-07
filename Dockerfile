FROM ubuntu
RUN apt-get update && apt-get -y install golang-go
RUN cd / && mkdir Dockerfile && chmod 777 -R Dockerfile/
COPY ./main.go /Dockerfile
