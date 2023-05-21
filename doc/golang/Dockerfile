FROM golang:1.9

RUN mkdir /echo
COPY main.go /echo
CMD ["go", "run", "/echo/main.go"]

RUN mkdir /echo2
COPY sub.go /echo2
CMD ["go", "run", "/echo2/sub.go"]
CMD ["go", "run", "/echo/main.go"]