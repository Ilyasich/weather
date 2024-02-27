FROM golang

WORKDIR /app

COPY . .

RUN echo "Hello from Docker"

