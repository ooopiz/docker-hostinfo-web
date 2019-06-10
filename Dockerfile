From golang:alpine as builder

RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o web-app

From alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/web-app /app/web-app
COPY --from=builder /app/template.html /app/template.html
EXPOSE 9090
CMD ["./web-app"]

