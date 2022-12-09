#!/bin/bash

if [ ! -f ".env.local" ]; then
    cp .env.example .env.local
fi

npm install

tail -f /dev/null