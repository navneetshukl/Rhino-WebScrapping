FROM golang:1.19-alpine


WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download


COPY *.go ./


RUN CGO_ENABLED=0 GOOS=linux go build -o /webscrapper


EXPOSE 8080


CMD ["/webscrapper"]