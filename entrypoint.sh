#!/bin/bash

while true; do
    air
    echo "Air crashed with exit code $?. Respawning.." >&2
    sleep 1
done