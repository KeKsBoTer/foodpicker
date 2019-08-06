FROM golang:1.11 as builder
WORKDIR /server/

COPY  main.go .
COPY  food.txt .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -a -installsuffix nocgo -o foodpicker .


FROM gcr.io/distroless/base
WORKDIR /root/
COPY --from=builder /server/foodpicker .
ENTRYPOINT [ "./foodpicker"]