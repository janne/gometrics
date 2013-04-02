# Gometrics

## Overview

A client/server system for tracking various values in clients,
that may or may not have access to the network at any given
time.

## Features

* Client
    * Package to be included in your own project
    * Track one or more values
    * Support for multiple clients
    * Support for offline clients with cached values
* Server
    * Exposes REST-API for the clients
    * Checks API-key
    * Preconfigured to be hosted at Heroku
    * Outputs graphs via Google Graph API
    * Support for backing up all data in CSV format

## Setup

* Create a [Heroku account](https://api.heroku.com/signup) and install the [Heroku Toolchain](https://toolbelt.heroku.com/)
* Create the Heroku app

        $ heroku create -b https://github.com/kr/heroku-buildpack-go.git
        $ heroku addons:add heroku-postgresql:dev
        $ heroku addons # Take note of URL color
        $ heroku pg:promote HEROKU_POSTGRESQL_<COLOR>_URL
        $ heroku pg:psql < setup.sql
        $ git push heroku master
        $ heroku config:set API_KEY=<YOUR API KEY>
        $ heroku open
