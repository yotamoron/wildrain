#!/bin/bash

echo "Ricing."
rice -i ./src/com.wildrain/wildrain embed

echo "Installing"
go install com.wildrain/wildrain

echo "Starting server"
wildrain
