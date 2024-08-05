# Go Proxy

Go Proxy is a versatile tool designed to bypass CORS issues when making requests from HTTPS to HTTP environments and now supports blue-green deployments. It's particularly useful for developers working on applications that need to interact with APIs served over HTTP from a client served over HTTPS, as well as those implementing blue-green deployment strategies.

## Features

- **CORS Bypass**: Easily bypass Cross-Origin Resource Sharing (CORS) restrictions.
- **Simple Usage**: Quick to set up with minimal configuration.
- **Docker Support**: Available as a Docker container for easy deployment.
- **Blue-Green Deployment**: Support for blue-green deployment strategy, allowing for zero-downtime updates.

## Usage

### Running Locally

To use Go Proxy locally, you have two options:

1. Single Proxy Mode:

```bash
go run main.go single -t http://target-server:8080 -l :8080
```

2. Blue-Green Deployment Mode:

```bash
go run main.go blue-green --blue=http://blue-server:8080 --green=http://green-server:8080 --blue-weight=100 --green-weight=0 -l :8080
```

### Command-Line Arguments

#### Single Proxy Mode

- `-t, --target`: Target server URL (default: "http://localhost:8080")
- `-l, --listen`: Listen address (default: ":8080")

#### Blue-Green Deployment Mode

- `--blue`: Blue server URL (default: "http://blue-server:8080")
- `--green`: Green server URL (default: "http://green-server:8080")
- `--blue-weight`: Blue server weight (0-100) (default: 100)
- `--green-weight`: Green server weight (0-100) (default: 0)
- `-l, --listen`: Listen address (default: ":8080")

### Docker Support

To run Go Proxy using Docker:

1. Build the Docker image:

```bash
docker build -t go-proxy .
```

2. Run the container:

For single proxy mode:
```bash
docker run -p 8080:8080 go-proxy single -t http://target-server:8080
```

For blue-green deployment mode:
```bash
docker run -p 8080:8080 go-proxy blue-green --blue=http://blue-server:8080 --green=http://green-server:8080 --blue-weight=100 --green-weight=0
```

## Blue-Green Deployment

The blue-green deployment feature allows you to run two versions of your application simultaneously and easily switch traffic between them. This enables zero-downtime deployments and quick rollbacks if issues are detected.

To use blue-green deployment:

1. Start with 100% traffic to the blue server (current production version).
2. Deploy the new version to the green server.
3. Gradually shift traffic from blue to green by adjusting the weights.
4. Once confident, shift 100% traffic to green.
5. If issues arise, quickly revert to blue by adjusting weights back.