# Development

Contains small utilities for converting or parsing values via Alfred workflows. 

## Compiling

To create a binary file use `make build`.
The default output directory is `/dist`.

~~~sh
make build
~~~

## Testing

For testing run `make test` with operation and query arguments.

~~~sh
make test operation=dice query='6'
~~~

~~~sh
make test operation=base64enc query='{"test":123}'
~~~
