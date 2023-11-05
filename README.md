# Koji Kukac Hackathon - Team Burek

## Goal

Build a web app where a signed up user can:

- view available parking spots,
- view details about a parking spot,
- reserve a parking spot (partially implemented),

and where an admin can:

- view availability of spots in real time (unimplemented),
- add and modify park spots (unimplemented).

## How to run

**_Note:_** before running, make sure an API key is available for Google Maps in the frontend/.env file under the variable name VITE_GOOGLE_MAPS_API_KEY.

The simplest way to run the app using docker-compose by running:

```bash
$ docker-compose build
$ docker-compose up
```
