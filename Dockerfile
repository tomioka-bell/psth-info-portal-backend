# Start from a Go image
FROM golang:1.24.3-alpine

# Setup Work Directory
WORKDIR /app

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o golanginfo .

# Expose the port your app will run on
EXPOSE ${PORT}

ENV SMTP_HOST=${SMTP_HOST}
ENV SMTP_USERNAME=${SMTP_USERNAME}
ENV SMTP_PASSWORD=${SMTP_PASSWORD}
ENV SMTP_PORT=${SMTP_PORT}
ENV SMTP_SET_FROM=${SMTP_SET_FROM}

ENV SQLSERVER_USER=${SQLSERVER_USER}
ENV SQLSERVER_PASSWORD=${SQLSERVER_PASSWORD}
ENV SQLSERVER_DB=${SQLSERVER_DB}
ENV SQLSERVER_HOST=${SQLSERVER_HOST}
ENV SQLSERVER_PORT=${SQLSERVER_PORT}

ENV SERVER_PORT=${SERVER_PORT}
ENV TOKEN_SECRET_KEY=${TOKEN_SECRET_KEY}


# Run the application
CMD ["./golanginfo"]