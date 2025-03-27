CRUD API

The CRUD api will serve to allow the main http requests onto the 2 structs in this api:
- Movies
- Directors

Director ir a property of movie so we will build the struct with a pointer to the director struct.
When creating the data for movies we have to indicate director as a pointer since the data is generated beforehand.

Currently there's a get method that encodes the variable movieList based on the struct json structure provided and returns it, can be obtained via the getMoviesHandler function and the base route "/", each handler will provide access to each method. 