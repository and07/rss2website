# rss2website

## 0. Clone this repo into your GOPATH

Clone it into `$GOPATH/src/github.com/and07`.

## 1. Build Docker image with program in it

```
sudo docker build -t and07/rss2site .
```

## 2. Show image 
```
sudo docker images
```

## 3. Run it

Now that we built the image, let's run it.

```
sudo docker run -p 3000:3000 --rm and07/rss2site
```

Surf to: http://127.0.0.1:3000



