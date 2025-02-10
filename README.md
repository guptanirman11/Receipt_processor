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
The dockerfile in both server and client is spinned up using Docker compose file. The Docker compose file establishes client and server in the same network so that cross-network issues does not arise and moreover it improves the security as well not allowing out of network services to interact with it.

## The server uses in-memory map to store receipt_id and points associated to a particula id


## Future Extensions
 The client right now can take one receipt from the json pasted but can be extended later on for an array of lists to process and return {id:points} for each receipt.