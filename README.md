# go-web-app

Example of dockerized go web application

## Build the image
`make image`

specify the version:
`make image TAG=0.0.1`

## Run the application
`docker run -d -p 8080:8080 heliostech/go-web-app`
