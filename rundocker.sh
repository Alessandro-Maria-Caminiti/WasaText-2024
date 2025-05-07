#!/bin/bash

# Esegui i container contemporaneamente
FROM golang:1.17
sudo docker run --rm -p 3000:3000 wasa-text-backend:latest &
sudo docker run --rm -p 8080:80 wasa-text-frontend:latest &

# Attendi la terminazione dei processi in background
wait