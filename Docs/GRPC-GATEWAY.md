# Problems and Challenges

### The need for grpc-gateway
* Exposing only grpc server was enough for our inter-service communication as we can call any grpc from any backend service to another backend service.
* The problem arrived when the frontend service had to access these rpcs to create, retrieve, update and delete the questions.
* There are two ways we can access this: [grpc-web](https://github.com/grpc/grpc-web) and [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway).

### Why gateway over grpc-web
* Using grpc-web we can directly access the rpcs in the backend services. Binary data is sent over network as we are calling the rpc directly, resulting faster response and processing time.
* But in our case we will be calling the REST apis from frontend. So we just use grpc-gateway to proxy the json format data to binary and reach the rpc.
* We sole purpose of use of grpc-gateway here is because our font end was designed to call the REST apis. Otherwise, grpc-web is highly preferred where speed and processing time matters.

### What does grpc-gateway do
* grpc-gateway translates the RESTful JSON apis to work with our gRPC apis.
* More detail on how it does that can be found [here](https://github.com/grpc-ecosystem/grpc-gateway#about)
