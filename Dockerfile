From golang:latest

RUN go get github.com/oxtoacart/bpool && go get github.com/julienschmidt/httprouter && go get github.com/gosimple/slug && go get github.com/mmcdole/gofeed && go get github.com/and07/rss2website/proto

WORKDIR /go/src/rss2site
COPY . /go/src/rss2site
RUN go install
CMD ["rss2site"]

#RUN mkdir /app 
#ADD . /app/ 
#WORKDIR /app 
#RUN go build -o main . 
#CMD ["/app/main"]

EXPOSE 3000
