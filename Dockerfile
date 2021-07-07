FROM golang:1.16.5

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o chat ./main/main.go

RUN curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | bash
RUN apt-get update
RUN apt-get install -y migrate
# RUN migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5436/postgres?sslmode=disable' up
# RUN migrate -path ./schema -database 'postgres://postgres:qwerty@db:5432/postgres?sslmode=disable' up

CMD ["./chat"]
