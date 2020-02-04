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
## API'S
### /api/v1/activity 
    To see all activities

### /api/v1/activity/persons
    Time spent per person

### /api/v1/activity/persons?duration=10
    People who have spent more than 10 seconds
 
### /api/v1/activity/persons/count 
    Number of people who spent any time
    
### /api/v1/activity/persons/count?duration=10
    Number of people who spent more than 10 units of time
 
### /api/v1/activity?breakout=day 
    Time spent per day
 
### /api/v1/activity?breakout=hour
     Time spent per day
 
### /api/v1/activity/sessions
     session is defined if there has been more than a 30 minute break
