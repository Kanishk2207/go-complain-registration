to start the app server,
run following commands:

docker-compose up --build -d

go run main.go

your app will start on localhost:8080

the routes are:

POST /signup
POST /login
POST /users/:userID/complaints
GET /users/:userID/complaints
GET /admin/:userID/complaints
GET /users/:userID/complaints/:complaintID
PUT /users/:userID/complaints/:complaintID/mark-resolved
