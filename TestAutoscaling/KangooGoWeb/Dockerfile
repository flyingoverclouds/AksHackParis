FROM golang:1.10

WORKDIR /app

# Copy the current directory contents into the container at /app
ADD KangooGoWeb /app


# Make port 80 available to the world outside this container
EXPOSE 8080

# Define environment variable
ENV LISTENINGPORT ":8080"

# Run app.py when the container launches
ENTRYPOINT [ "/app/KangooGoWeb" ]