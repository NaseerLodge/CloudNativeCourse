FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY movieapi/ movieapi/
COPY movieserver/ movieserver/
WORKDIR /app/movieserver
RUN CGO_ENABLED=0 go build -o /bin/movieserver

FROM scratch
COPY --from=build /bin/movieserver /bin/movieserver
ENTRYPOINT ["/bin/movieserver"]