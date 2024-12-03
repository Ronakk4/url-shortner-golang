FROM golang:alpine as builder 

run mkdir /builder

add . /build/
WORKDIR /build
RUN go build -o main

from alpine
RUN adduser -S -D -H -h /app appuser

user appuser
copy . /app
copy --from=builder /build/main /app/
WORKDIR /app
EXPOSE 3000
CMD [ "./main" ]
