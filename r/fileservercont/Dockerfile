################################
# STEP 1 build executable binary
################################
FROM golang:alpine AS builder

# Install git, required for fetching the dependencies.
RUN apk update && apk add --no-cache git make

WORKDIR /app
COPY . .

# Fetch dependencies, using go get, download only, verbose.
RUN go get -d -v

# Build the binary.
RUN make webshare

############################
# STEP 2 build a small image
############################
FROM scratch

# Copy our static executable.
COPY --from=builder /app/webshare /app/webshare

# Run the hello binary.
ENTRYPOINT ["/app/webshare"]


