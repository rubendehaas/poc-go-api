FROM golang:1.12

WORKDIR /go/src/app
COPY . .

RUN go get github.com/pilu/fresh
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/context
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/joho/godotenv

CMD [ "fresh" ]

EXPOSE 8080