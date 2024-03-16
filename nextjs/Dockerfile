# Use the official Node.js image as the base image
FROM node:14-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install --production

# Copy the rest of the application source code to the working directory
COPY . .

# Build the Next.js application
RUN npm run build

# Start a new stage from scratch
FROM node:14-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the built application from the previous stage
COPY --from=builder /app .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the Next.js application
CMD ["npm", "run", "start"]

