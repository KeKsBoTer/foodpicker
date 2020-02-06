FROM golang:1.11 as builder
WORKDIR /server/

COPY main.go .
COPY api.go .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -a -installsuffix nocgo -o foodpicker .


FROM gcr.io/distroless/base
WORKDIR /root/
COPY  food.txt .
COPY  template.html .
COPY --from=builder /server/foodpicker .
ENTRYPOINT [ "./foodpicker"]