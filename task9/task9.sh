#!/bin/bash

# SSH connection
SSH_USER="jmaradiaga"
SSH_HOST="192.168.122.129"
SSH_KEY="/home/jmaradiaga/.ssh/id_rsa"

# Docker container details
CONTAINER_NAME="task9"
IMAGE_NAME="helloworld"
IMAGE_TAG="latest"
EXPOSED_PORT="8080"
GO_APP_PATH="/home/jmaradiaga/DevOps/DevOpsGYMBatch3"

# log file
LOG_FILE="setup_docker_container.log"

# SSH command to run the Docker container
SSH_COMMAND="docker run -d --name ${CONTAINER_NAME} -p ${EXPOSED_PORT}:${EXPOSED_PORT} ${IMAGE_NAME}:${IMAGE_TAG}"

# Execute SSH command
ssh -i "${SSH_KEY}" "${SSH_USER}@${SSH_HOST}" "${SSH_COMMAND}" &> "${LOG_FILE}"

# Check if SSH command was successful
if [ $? -eq 0 ]; then
  echo "Docker container setup successful. Log saved to ${LOG_FILE}"
else
  echo "Docker container setup failed. Check ${LOG_FILE} for details."
fi
