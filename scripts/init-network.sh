#!/bin/bash
NETWORK_NAME="cluster"

# Check if the network already exists
if docker network inspect $NETWORK_NAME &> /dev/null; then
    # The network already exists
    echo "Network '$NETWORK_NAME' already exists, skipping creation."
else
    # Create the external network
    docker network create $NETWORK_NAME
    echo "Network '$NETWORK_NAME' created."
fi
