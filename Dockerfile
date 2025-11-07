FROM golangci/golangci-lint:v1.63

WORKDIR /app

CMD ["golangci-lint", "run"]
