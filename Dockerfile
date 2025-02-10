FROM golang:1.23.5-alpine
WORKDIR /ebook
COPY . ./
RUN go mod download
RUN go build -v -o /output .
EXPOSE 8080
CMD ["/output"] 