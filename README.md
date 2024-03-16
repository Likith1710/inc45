## How We Test and Deploy

This Stack uses separate systems to test and deploy each part of our app:

- **Go Component:** This handles the backend. We check its code quality and run tests using a tool called GolangCI-Lint.
- **Next.js Component:** This is for the frontend. We make sure its code is clean using ESLint and Prettier.
- **WordPress Component:** This manages our content. We ensure it follows WordPress coding standards using PHPCS.

## Dockerization

We've streamlined the process of running our app using Docker containers. Each part of our app has its own Dockerfile for building the container images.

## Running Locally

You can run our entire app locally using Docker Compose. Just follow these steps:

1. Make sure you have Docker and Docker Compose installed.
2. Clone this repository to your machine.
3. Navigate to the root directory of the project.
4. Run `docker-compose up` in your terminal.

## Continuous Deployment

Whenever there is a change in the Go folder within the repository, only the Go application will build the Docker image and push it to Docker Hub. Finally, using the Jenkinsfile and docker-stack.yml files, the Go application will start using Docker Stack. This repository is connected with the CI tool using webhooks.

