MYSQL Book management system
This api covers a full api crud for a single entity.
The book management system allows creation of a book, update, and  deletion, as well as fetching a single book or multiple without restriction.
The project structure follows a specific guideline as shown below:
 - Server setup in the cmd/main folder, in the main file we start the server, creating a router via mux library and implementing the routes to this router. After that we provide the router to the http/net package to listen and serve requests. The rest of files will be stored in separate folder called pkg.
 -As mentioned, the routes are going to be our entrypoint for requests, in this case there's the base route with "/" and "/books", calling books with an id if needed, this is what will allow interactions with our database table for books. Can be found in the routes folder. 
 - Inside routes we're also declaring function handlers per route, we can also specify a certain method per route allowed. This functions are held in the controllers.
 -The controllers is the first stop of the external requests to the api, they're in charge of handling the header and parsing of the bodies. Apart from that, they do not interact with the database layer, this is a job of the models/services.
 - A model is needed per entity declared, if the gorm is added in the struct declaration, a database migration should occur with the pertaining restrictions and extra details. Within this model we will be able to stablish a connection to the database, not directly but we can say that the connection is passed to them to achieve whatever query is needed. 
 - This database connection comes from a different layer, the config folder holds the app.go file which stablishes the connection to the database in this case with mysql locally, and passes an open connection pool for the database to the service.

 Client -> Api -> Route -> Controller -> Service -> Database 
 And then the data travels back.

 - Good to mention that the data in runtime of the go app is understandable by the app itself but in order to travel it needs to be encoded/decoded or parsed to JSON format, we take care of this in the utils. Though this should be updated since the function to read from the request is currently deprecated.

 This covers a very basic integrations of a lighweight crud api, in the future project a more extense api should be done, providing authentication, authorization, third party auth, validation of dtos, gRPC, sockets, etc.