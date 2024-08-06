FROM golang:1.22.5 AS builder

# Copy local code to the container image.
COPY . /app

# Create and change to the app directory.
WORKDIR /app

# Build the binary.
RUN make deps
RUN make build

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3.20.2 AS runner

# Install library to open the binary :(
RUN apk add --no-cache libc6-compat

# Set working directory.
WORKDIR /app

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/build/output/main /app

# Copy the config files.
COPY --from=builder /app/config config

# Run the web service on container startup.
CMD ["./main"]