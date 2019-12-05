FROM golang:1.13-alpine

WORKDIR /helloworld

COPY . .

RUN go install -v .

ENV LISTEN ":8080"
ENV NAME "Default"
EXPOSE 8080

CMD ["helloworld"]
