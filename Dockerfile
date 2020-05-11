FROM golang:1.10

WORKDIR /jsonValidator
COPY . /jsonValidator

EXPOSE 8080

ENTRYPOINT ["go", "run", "jsonValidator.go"]