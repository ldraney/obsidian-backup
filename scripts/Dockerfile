# Use the official Golang image
FROM golang:1.20

# Install Google Cloud SDK
RUN apt-get update && apt-get install -y \
    apt-transport-https \
    ca-certificates \
    gnupg && \
    echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
    curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add - && \
    apt-get update && apt-get install -y google-cloud-sdk

# Set working directory inside the container
WORKDIR /app

# Copy Go script to the container
COPY backup_to_gcs.go .

# Set environment variables (these will be set in the .env file)
# ENV GCS_BUCKET=""
# ENV BACKUP_DIR=""

# Install Go dependencies
RUN go mod init backup_to_gcs && go get -u cloud.google.com/go/storage

# Build the Go application
RUN go build -o backup_to_gcs .

# Set Google Cloud authentication
# ENV GOOGLE_APPLICATION_CREDENTIALS # a directory defined in be in bash_secrets

# Service account credentials
COPY $GOOGLE_APPLICATION_CREDENTIALS /app/service-account.json

# Run the Go application
CMD ["./backup_to_gcs"]

