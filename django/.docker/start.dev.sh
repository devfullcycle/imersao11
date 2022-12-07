#!/bin/bash

if [ ! -f ".env" ]; then
  cp .env.example .env
fi

pipenv install
pipenv run python manage.py migrate
pipenv run python manage.py loaddata initial_data

tail -f /dev/null