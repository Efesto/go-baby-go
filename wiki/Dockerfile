FROM ubuntu:20.04

WORKDIR /app

COPY wiki .
COPY tmpl ./tmpl
COPY data ./data

EXPOSE 8080

CMD ["./wiki"]