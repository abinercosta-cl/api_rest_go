FROM golang:1.22.6

#set working directory
WORKDIR /go/src/app

#copy the source code
COPY . .

#expose the port
EXPOSE 8000

#build the Go app
RUN go build -o main cmd/main.go

#Run the executable
CMD [ "./main" ]

