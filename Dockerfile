FROM golang:1.11 as builder
WORKDIR /server/

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY main.go .
COPY api.go .

RUN GO111MODULE=auto CGO_ENABLED=0 go build -ldflags="-s -w" -a -installsuffix nocgo -o foodpicker .


FROM gcr.io/distroless/base
WORKDIR /root/
COPY  food.txt .
COPY  template.html .
COPY --from=builder /server/foodpicker .
ENTRYPOINT [ "./foodpicker"]
