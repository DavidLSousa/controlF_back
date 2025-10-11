FROM golang:1.24-alpine AS build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/controlF cmd/main.go

FROM alpine
WORKDIR /etc/controlF
COPY --from=build /usr/local/bin/controlF /usr/local/bin/controlF
CMD ["controlF"]
