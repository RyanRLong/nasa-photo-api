# Nasa Photo Api

## Overview

The Nasa Photo API starts a local server listening on port 8050. When the endpoint receives a proper request, it will query the data source for all Mars Rover images taken on the date in the request and will download those to a local directory named after the date requested.

Files themselves will be named in the format DATE_ID.jpg.

## Running

Clone the repository to your local machine.

To run the application, from the root of the local repository, enter the following on the command line to start the server.

`go run cmd/main.go`

## Building

To build the application, from the root of the local repository, enter

`go build`

_There are no external dependencies._

## Using

There is one endpoint available when starting the application. This endpoint receives a get request in the following format.

`GET http://localhost:8050/photos/yyyy/mm/dd`

| Placeholder | Type                    |
| ----------- | ----------------------- |
| yyyy        | Year in 4 digit format  |
| mm          | Month in 2 digit format |
| dd          | Day in 2 digit format   |

Example url to download photos from July 1, 2019

`GET http://localhost:8050/photos/2019/07/01`

The command in curl would be equivalent.

`curl localhost:8050/photos/2019/07/01`

## Testing

To run tests, from the _pkg_ folder type:

`go test`
