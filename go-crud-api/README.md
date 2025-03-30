CRUD API

This api will be used to access resources of a in memory variable that contains films, this memory will reinitiate if the api is restarted.

Handling 2 structs:
- Movies
- Directors

Director is a property of movie so we will build the struct with a pointer to the director struct.
When creating the data for movies we have to indicate director as a pointer since the data is generated beforehand.

Each handler will provide access to each method.

The API now handles all requests:

- Get all movies with "/"
- Get one movie by the param provided after the base url "/{id}" or return not found
- Delete one movie by the param provided after the base url "/{id}" or return not found
- Create one movie on the base route "/", the request body json is decoded to verify type assertiveness with the Movies struct, if it's correct adds the new movie and returns all the movieList json encoded again.
- Update one movie uses the PUT method with "/{id}" url, the param will be extracted from the url, the movie will be found and deleted, then added afterwards with the same id as it originally had, if not found will inform about it.

In the future this api should be updated so it handles Generics to avoid redundant code.
