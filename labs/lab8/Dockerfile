FROM golang:1.16-alpine AS build

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY mongo.go ./
RUN CGO_ENABLED=0 go build -o /bin/mongo

FROM scratch
COPY --from=build /bin/mongo /bin/mongo
ENTRYPOINT ["/bin/mongo"]
