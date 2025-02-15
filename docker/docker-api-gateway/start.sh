#!/usr/bin/env bash
source ../../resources/gateway.env

export APP_IMAGE_NAME=${APP_IMAGE_NAME}
export APP_IMAGE_VERSION=${APP_IMAGE_VERSION}
export CONTAINER_NAME=${CONTAINER_NAME}
export CONTAINER_PORT=${CONTAINER_PORT}
export VOLUME_CONF_EXT=${VOLUME_CONF_EXT}
export VOLUME_CONF_INN=${VOLUME_CONF_INN}
export LURA_PATH=${LURA_PATH}


# echo `ls $VOLUME_CONF_EXT`
docker-compose down
docker-compose up -d
