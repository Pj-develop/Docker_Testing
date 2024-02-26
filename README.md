## Sock Shop Project and Integration Testing with Docker Compose

This repository contains the code and materials used for the demonstrations in the presentation on using Docker and Docker Compose to simplify integration testing, delivered at Dockercon 2020.

## Sock Shop Microservice Demo Application

Description:

The Sock Shop is a microservices-based demo application designed for showcasing and testing microservice and cloud-native technologies. It serves as the user-facing component of an online shop selling socks. The application is built using Spring Boot, Go kit, and Node.js, and is packaged in Docker containers.

## Deployment Platforms:

The /deploy folder includes scripts and instructions for provisioning the Sock Shop application onto various platforms. Feel free to explore and contribute to add support for your favorite platform.

## Contributing:

We welcome community contributions. Bug reports and feature requests can be submitted through GitHub issues, and contributions can be made via pull requests. Please refer to the contribution guidelines for more details.

## Integration Testing with Docker Compose

Demo Overview:

The presentation includes a demonstration of using Docker and Docker Compose to simplify integration testing. The aim is to showcase how Docker containers can be leveraged to create a consistent and reproducible environment for testing the Sock Shop application.

# Getting Started:

Clone this repository:

bash

git clone https://github.com/your-username/your-repo.git
cd your-repo
Run the integration tests using Docker Compose:

bash

docker-compose up`

This command will orchestrate the required services and containers for integration testing.

Explore the results and fine-tune your integration testing setup based on your requirements.

# Additional Resources:

For more details on the Sock Shop application design, refer to internal-docs/design.md.
Visualize the running application using Weave Scope or Weave Cloud (refer to deploy/ for deployment instructions).
