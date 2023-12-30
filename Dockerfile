# Use an official Golang runtime as a parent image
FROM golang:1.17

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Install any needed packages specified in the go.mod file
RUN go mod download

# Make port 5001 available to the world outside this container
EXPOSE 5001

# Define the command to run your app using CMD which defaults to sh -c
CMD ["go", "run", "."]