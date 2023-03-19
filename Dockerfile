#dockerfile && docker-compose skipped.


#FROM golang:1.20
#
#WORKDIR /app
#
#RUN ls -la
#
#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download
#
#COPY . .
#WORKDIR /cmd
#
#CMD ["./main"]