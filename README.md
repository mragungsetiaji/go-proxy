# Go Proxy

Go Proxy is a handy tool designed to bypass CORS issues when making requests from HTTPS to HTTP environments. It's particularly useful for developers working on applications that need to interact with APIs served over HTTP from a client served over HTTPS.

## Features

- **CORS Bypass**: Easily bypass Cross-Origin Resource Sharing (CORS) restrictions.
- **Simple Usage**: Quick to set up with minimal configuration.
- **Docker Support**: Available as a Docker container for easy deployment.

## Usage

### Running Locally

To use Go Proxy locally, you can run the following command in your terminal. 

```bash
go run main.go --ip 8.8.8.8
```

## Using Docker
For those who prefer Docker, Go Proxy is available as a prebuilt Docker image. You can run it using the following command:

```
docker run -p 8080:8080 mragungsetiaji/go-proxy:1.0.0-amd64 app --ip 8.8.8.8
```

This command pulls the Go Proxy image from Docker Hub and runs it, mapping the container's port 8080 to the host's port 8080.
