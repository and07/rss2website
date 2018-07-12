FROM golang:latest

#RUN git config --global url."git@gitlab.com:".insteadOf "https://gitlab.com/"
#RUN go get github.com/oxtoacart/bpool 
#RUN go get github.com/julienschmidt/httprouter 
#RUN go get github.com/gosimple/slug 
#RUN go get github.com/mmcdole/gofeed 
#RUN go get gitlab.com/and07/rss2website/proto

WORKDIR /go/src/gitlab.com/and07/rss2website
COPY . /go/src/gitlab.com/and07/rss2website
RUN go install
CMD ["rss2website"]

#RUN mkdir /app 
#ADD . /app/ 
#WORKDIR /app 
#RUN go build -o main . 
#CMD ["/app/main"]

EXPOSE 3000
