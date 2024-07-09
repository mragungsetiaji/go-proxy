FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk update --no-cache && apk add --no-cache tzdata
RUN apk add --no-cache ca-certificates 
RUN apk add --no-cache git

# Create a group and user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /build

COPY . .

RUN go build -o app . && \mv app /usr/bin

FROM golang:alpine

# Copy the user and group files
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Tell docker that all future commands should run as the appuser user
USER appuser

COPY --from=builder /usr/bin/app /usr/bin/app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Jakarta /usr/share/zoneinfo/Asia/Jakarta


CMD ["/usr/bin/app"]