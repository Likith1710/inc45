## How to Test and Deploy

This Stack uses separate systems to test and deploy each part of our app:

- **Go Component:** This handles the backend. Check its code quality and run tests using a tool called GolangCI-Lint.
- **Next.js Component:** This is for the frontend. Make sure its code is clean using ESLint and Prettier.
- **WordPress Component:** This manages our content.  Ensure it follows WordPress coding standards using PHPCS.

## Dockerization

I've streamlined the process of running our app using Docker containers. Each part of our app has its own Dockerfile for building the container images.

## Running Locally

You can run the entire app locally using Docker Compose. Just follow these steps:

1. Make sure  Docker and Docker Compose installed and dockerswarm has initialised. 
2. Clone this repository to your machine.
3. Navigate to the root directory of the project.
4. Run `docker-compose stack deploy` in your terminal.

## Continuous Deployment

- Whenever there is a change in the Go folder within the repository, only the Go application will build the Docker image and push it to Docker Hub. 
- Finally, using the Jenkinsfile and docker-stack.yml files, the Go application will start using Docker Stack. 
- This repository is connected with the CI tool of Jenkins of `main` using webhooks.


- This Docker stack has 3 services.
    - `Go application`
    - `Nextjs Application`
    - `WordpressSite.`

- Jenkinsfile Uses the `docker secrets` to store and use sensitive information in the Jenkinsfile.

