FROM go:1.22.2 AS builder

COPY . .

FROM debian:12-slim

CMD ["cukka", "--port", "8080"]
