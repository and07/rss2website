FROM golang

WORKDIR /go/src/rss2site
COPY main.go /go/src/rss2site
RUN go install
CMD ["rss2site"]