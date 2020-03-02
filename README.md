# Contact Lookup API (Go)


***


## A RESTful Contacts Lookup API, made in Golang & locally hosted, that allows users to performs CRUD operations on a temporary/unsaved database, using tcp connection requests.

### <i> Create, search, edit and delete Contact Records.

<br>

***

### Client Page (Front-End) Homepage: <br>
#### <b>http://localhost:<Port\>/api/v0-1/home</b>

***

<br><br>

#### List of URL(http://localhost:<Port\>/api/v0-1/) + URN (End-points), for Requests against an unstored array, that are currently available:

| URN | Action on Array/DB | Full URI (Using some port, e.g. "8080") |
|:---|:---|:---|
| <ul><li>"/home"</li></ul> | <b><u>HOME/CLIENT PAGE</u></b> | <ul><li>"http://localhost:8080/api/v0-1/home"</li></ul> |
| <ul><li>"/contacts/new"</li></ul> | <b><u>CREATE (POST)</u></b> | <ul><li>"http://localhost:8080/api/v0-1/contacts/new"</li></ul> |
| <ul><li>"/contacts/all"</li><li>"/contacts/find?id=\<id\>"</li></ul> | <b><u>READ (GET)</u></b> | <ul><li>"http://localhost:8080/api/v0-1/contacts/all"</li><li>"http://localhost:8080/api/v0-1/contacts/find?id=<id\>"</li></ul> |
| <ul><li>"/update?id=\<id\>"</li></ul> | <b><u>UPDATE (PUT)</u></b> | <ul><li>"http://localhost:8080/api/v0-1/update?id=<id\>"</li></ul> |
| <ul><li>"/deleteAll"</li><li>"/delete?id=<id\>"</li></ul> | <b><u>DELETE</u></b> | <ul><li>"http://localhost:8080/api/v0-1/deleteAll"</li><li>"http://localhost:8080/api/v0-1/delete?id=<id\>"</li></ul> |


##### POST/PUT/DELETE available via middleware.

<br>

***
***

<br>

|Version| Changes|
|:---|:---|
|Version 0.0.1 [2020-03-01]|<ul><li>Initial Commit.</li><li>Add "main.go" file.</li><li>Adds basic infrastructure and request methods, against the local unsaved sudo-databse</li><li>Add README.md</li></ul>|
|Version 0.0.2 [2020-03-02]|<ul><li>Add three new request methods and accompanying handler functions: PUT, DELETE(x2)</li><li>Add new dir "Screenshots", and Add image files to new dir.</li><li>Update README.md</li></ul>|