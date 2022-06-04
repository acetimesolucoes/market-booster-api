# Use nxgo/cli as the base image to do the build
FROM nxgo/cli as builder

# Create app directory
WORKDIR /workspace

# Copy package.json and the lock file
COPY package.json /workspace/

# Install app dependencies
RUN npm install

# Copy source files
COPY . .

# Build apps
RUN npm run build api

# This is the stage where the final production image is built
FROM golang:1.18-alpine as final

# Copy over artifacts from builder image
COPY --from=builder /workspace/dist/apps/api /workspace/api

# Set environment variables
ENV PORT=3000
ENV HOST=0.0.0.0

# Expose default port
EXPOSE 3000

# Start server
CMD [ "/workspace/api" ]

# FROM golang:latest

# RUN go install github.com/canthefason/go-watcher/cmd/watcher@latest

# WORKDIR /opt/acetimesolucoes

# COPY ./apps/api /opt/acetimesolucoes

# ENTRYPOINT /go/bin/watcher