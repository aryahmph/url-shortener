# Scalable URL Shortener Backend with High Availability

This project is the backend component of a scalable URL shortener with high availability, implemented using Docker, Go,
MongoDB, ZooKeeper, and Redis. It provides a reliable and efficient backend system for shortening long URLs into
shorter, more manageable ones, while ensuring high availability and fault tolerance.

## Background

Learn to build systems with production standards that have high scalability. Based on the case study learned about the
url shortener in the following
article: [System Design URL Shortening Service](https://www.geeksforgeeks.org/system-design-url-shortening-service/). I
tried to build a shorterner url system following the troubleshooting based on the guidelines in the article.

## Features

- Shorten long URLs into compact, easy-to-share links.
- High availability design to ensure system reliability.
- Scalable architecture with distributed databases and caching.
- Fault tolerance mechanisms for continuous operation.

## Prerequisites

Make sure you have the following prerequisites installed:

- Docker (version >= 23.0.1)
- Docker Compose (version >= v2.16.0)

## Configuration

Create a .env or replace .env.example / .env.example.replicated file in the project directory:

```
PORT=4001

DATABASE_USER=aryahmph
DATABASE_PASSWORD=aryahmph
DATABASE_HOSTS_PORTS=mongo1:27017
DATABASE_REPLICA_SET=-
DATABASE_NAME=url_shortener
DATABASE_MIN_POOL_SIZE=10
DATABASE_MAX_POOL_SIZE=100

ZOOKEEPER_HOSTS=zoo

REDIS_ADDRESS=redis:6379
```

## Installation

1. Clone the repository:

```
git clone https://github.com/aryahmph/url-shortener.git
```

2. Change to the project directory:

```
cd url-shortener
```

3. Start the Docker containers using Docker Compose:

```
make standalone-up
```

4. If you run replicated mode, run this command:

```
make replicated-up
docker compose exec mongo1 /scripts/setup.sh
```

5. Connect to `http://localhost:4001/api/v1`

## Usage

The backend component provides the API endpoints for URL shortening and redirection. To interact with the backend, you
will need to set up a frontend or use an API testing tool such as Postman.

Import postman collection [file](URL%20Shortener.postman_collection.json) to your Postman.

## Testing Result

This section provides an overview of the testing results obtained using [K6](https://github.com/grafana/k6), a popular
open-source load testing tool.

### Test Scenario

1. Scenario 1 : Link Creation and Get Link

* Simulates create and get link process with 1 : 10 ratio.
* Verifies the correctness of create and get link functionality.
* Measures response times for create and get link requests.

### Test Environment

The load tests were conducted in the following environment:

- Operating System: Ubuntu 20.04.5 LTS
- CPU: Intel® Core™ i5-8265U CPU @ 1.60GHz × 8
- RAM: 12GB

### Results Summary

Scenario 1: Link Creation and Get Link (Standalone each service)

- Virtual Users: 1000
- Duration : 30s
- Total Requests: 330.011
- Requests per Second (RPS): 10.979/s
- Average Response Time: 813.76µs
- 90th Percentile: 1.17ms
- Error Rate: 0%

## Todo List

- Add support for custom URL aliases and expiration of shortened URLs.
- Migrate from docker compose to Kubernetes (Still learning).
- Implement load balancing and horizontal scaling for distributing traffic across multiple backend instances.
- Implement Redis clustering for high availability and fault tolerance of caching.
- Implement Unit & Integration Testing
- Improve caching strategy.
- Improve performance testing strategy.

## Contributing

Contributions are welcome! If you would like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch.
3. Make your changes.
4. Commit your changes.
5. Push your branch to your forked repository.
6. Open a pull request.

