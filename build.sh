#!/bin/bash

DOCKER_BUILDKIT=1 docker build --tag stridezone:wic1 -f Dockerfile.wic1 .