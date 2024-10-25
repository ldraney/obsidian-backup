#!/bin/bash

# Directory where the backups will be stored
BACKUP_DIR="/mnt/c/Users/drane/volumes"
mkdir -p "$BACKUP_DIR"

# Get a list of all Docker volumes
VOLUME_LIST=$(docker volume ls -q)

# Loop through each volume and back it up
for VOLUME in $VOLUME_LIST; do
    echo "Backing up volume: $VOLUME"
    docker run --rm \
        -v ${VOLUME}:/volume \
        -v ${BACKUP_DIR}:/backup \
        ubuntu tar czf /backup/${VOLUME}_backup.tar.gz -C /volume .
    
    echo "Backup of $VOLUME complete: ${BACKUP_DIR}/${VOLUME}_backup.tar.gz"
done

echo "All Docker volumes backed up!"

