# Go-crud

## CRUD API

This repository provides a simple Go API that implements CRUD (Create, Read, Update, Delete) operations for managing person data using an in-memory database.

## Details:

1. API path `/person`:
    * **GET** `/person` or `/person/${personId}` return all persons or person with corresponding `personId`
    * **POST** `/person` is used to create record about new person and store it in database
    * **PUT** `/person/${personId}` is used to update record about existing person
    * **DELETE** `/person/${personId}` is used to delete record about existing person from database
2. Persons are stored as `objects` that have following properties:
    * `id` — unique identifier (`string`, `uuid`) generated on server side
    * `name` — person's name (`string`, **required**)
    * `age` — person's age (`number`, **required**)
    * `hobbies` — person's hobbies (`array` of `strings` or empty `array`, **required**)
3. Non-existing endpoints (e.g. `/some-non/existing/resource`) are handled.
4. Internal server errors are handled and processed correctly.
5. The api is accesible by frontend apps hosted on a different domain (cross-site resource sharing)