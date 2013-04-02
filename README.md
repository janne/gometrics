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
