#!/bin/bash

go build -o fake .
docker build -t fake:latest .