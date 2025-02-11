# Project Receipt_processor

## Business Requirement
A receipt processor application to calculate the number of points earned by a valid receipt according to the constraints and tell the number of points.

## System Requirements
### Functional Requirements
A box to add json receipt
A button to calculate points and show points earned after that
Application should be deployed using Docker containers with Golang as the preferred backend coding language
Application to use in-memory storage so no need to retain the data of past receipts added.

Non-Functional Requirements
The project focuses on developing a scalable application architecture.

High Level Architecture

The application uses Client Server Architecture where client is responsible to take the receipt and server processes the list to calculate the points and store. 

After that with next Get request it fetches the points for that list.

Users can send a previous recepit as well to check the points and if a valid receipt is added it returns the points otherwise return the error Not Found.

## Rest API Framework used for Golang
I am using Gin which is a REST APi based framework.

## File structure

Receipt_Process/
|── client/
|   |── index.html
|   |── style.css
|   |── script.js
|   |── receipt.json
|── server/
|   |── receipt_processor_service.go
|── models/
|   |── models.go
|── go.mod
|── go.sum
|── README.md

The web server has client and server files.

## Run instructions
The Dockerfile loads and run the backend go file and run a nginx server.
The Nginx server hosts the static files and reverse proxy the api calls to the backend at prot 8080 where the gin server is running and listening.
 cmds to run the service
 `docker build -t receipt_service .`
 `docker run -p 80:80 -p 8080:8080 receipt_service`

 then you got localhost:80 on web browser.

 There you have to paste the json receipt. If the json receipt is correct it will return the reeipt id and also by default auto fill the receipt id to calculate points text area. After that you can click on Calculate your points button to get you points if the receipt id is valid.

## The server uses in-memory map to store receipt_id and points associated to a particula id


## Future Extensions
 The client right now can take one receipt from the json pasted but can be extended later on for an array of lists to process and return {id:points} for each receipt.


 References
 1) https://gin-gonic.com/docs/quickstart/ 
 2) https://stackoverflow.com/questions/10075304/nginx-fails-to-load-css-files
 3) https://github.com/nginx/nginx/blob/master/conf/nginx.conf