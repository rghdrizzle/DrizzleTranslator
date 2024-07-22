FROM golang:alpine3.19 AS build

WORKDIR /usr/src/app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o drizzletranslator


FROM scratch AS dev

COPY --from=build /usr/src/app .


EXPOSE 3000

CMD ["/drizzletranslator"]