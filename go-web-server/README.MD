Web server

The purpose of this very light weight web server is to provide routing for static files and to show the simplicity of setting up a server with a few lines.

By accessing the / route, as by default it looks for the index.html as an entrypoint in the folder where the FileServer is pointing, the index homepage will be provided.

Apart from the base route we have 2 more

- Test Route: /test provides a readable response with the fmt package, this route only allows get requests
- Form Route: /form.html accesing this url a form appear on screen, the action will determine where do we get redirected when submitting the form, specified to /form in which the form handler will take the values from the response form values and will be able to print it to screen.

This go server doesn't use imported libraries, only standard libraries.
