FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build

CMD [ "./project_gosoft_login_register" ]