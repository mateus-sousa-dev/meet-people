FROM golang:1.18

WORKDIR /app

COPY . ./

RUN go build -o /bin

EXPOSE 8080

CMD [ "/bin"]