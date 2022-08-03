FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .
COPY . .
COPY local.env .
RUN go build -o oneclick main.go

FROM alpine
WORKDIR /build
COPY --from=builder /build/oneclick /build/oneclick
RUN pwd
EXPOSE 8000
CMD [ "./build/oneclick serve" ]

