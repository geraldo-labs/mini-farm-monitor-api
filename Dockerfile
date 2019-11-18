FROM golang:1.13.4-alpine3.10 as builder
LABEL maintainer="Geraldo Andrade <hi@geraldoandrade.com>"

WORKDIR /builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/web-api.sh ./cmd


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /builder/bin/web-api.sh .

EXPOSE 1323
CMD ["./web-api.sh"] 