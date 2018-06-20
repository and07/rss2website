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



#### more

https://stackoverflow.com/questions/29852583/docker-compose-accessing-postgres-shell-psql/37766141#37766141
```
docker@default:~$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                    NAMES
b563764171e2        poc_web             "bundle exec rails s "   43 minutes ago      Up 43 minutes       0.0.0.0:3000->3000/tcp   poc_web_1
c6e480b0f26a        postgres            "/docker-entrypoint.s"   49 minutes ago      Up 49 minutes       0.0.0.0:5432->5432/tcp   poc_db_1
docker@default:~$ docker exec -it c6e480b0f26a sh
# su - postgres
No directory, logging in with HOME=/
$ psql
psql (9.5.3)
Type "help" for help.

postgres=#
```


debug

```
docker logs --tail 500 --follow --timestamps my_container_name_1
```


