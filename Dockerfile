FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY . .
RUN go build -o ./dist/heridionibk .

EXPOSE 8080

CMD ./dist/heridionibk