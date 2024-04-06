# ASTRA

For this project, I have interpreted the diagram question to be sending weather data from various cities to a server at weather station.

We store the received data locally on on /tmp/asta directory. We have monitoring setup on this folder wherein whenever a new file is created in this folder, 
it is read and updated in db.

sample payload :
[
    {
    "city": "Jaipur",
    "temp": 39.5,
    "humidity": 60,
    "timestamp": "2023-04-06T14:30:00Z"
},{
    "city": "Udaipur",
    "temp": 41.0,
    "humidity": 54.9,
    "timestamp": "2024-04-05T14:00:00Z"
}
]


sample curls:

curl -X POST http://localhost:8080/ingest \
-H "Content-Type: application/json" \
-d '[{
    "city": "Jaipur",
    "temp": 39.5,
    "humidity": 60,
    "timestamp": "2023-04-06T14:30:00Z"
},{
    "city": "Udaipur",
    "temp": 41.0,
    "humidity": 54.9,
    "timestamp": "2024-04-05T14:00:00Z"
}]'


curl -X POST http://localhost:8080/ingest \
-H "Content-Type: application/json" \
-d '{
    "city": "Jaipur",
    "temp": 35.5,
    "humidity": 60,
    "timestamp": "2024-04-06T14:30:00Z"
}'




DOCKER

command to build image : 
docker build -t <your-image-name> .

command to run as container :
docker run -p 8080:8080 <your-image-name>
