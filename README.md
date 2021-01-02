# life-server

The server to the daily diary keeping application written in Go.

## Table of Contents
- [Endpoints](#Endpoints)

## Endpoints

All endpoints have the prefix `api/v1`

### Get all entries

`GET /entry` gets a list of all diary entries.

### Get an entry

`GET /entry/{id}` gets an entry with a given ID. Returns a JSON blob with a diary entry.

### Create an entry

`POST /entry` creates a new entry. Returns the entry and an ID


### Delete an entry 

`DELETE /entry/{id}` deletes an entry. Returns OK if successful


### Update an entry

`UPDATE /entry/{id}` updates an entry. 

