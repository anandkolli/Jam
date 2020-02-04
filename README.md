# Serivce exposing various endpoints to provide metrics

## Pre-requisites
1) Postgress must be accessible

## To build
1) Clone the repo
```
https://github.com/anandkolli/Jam.git
cd Jam
```
2) Build docker
```
docker build -t <image-name> .
```

3) Run the docker
```
docker run -p 9090:9090 -e POSTGRESSADDR="<ip-addr-post-gress>" <image-name>
```
