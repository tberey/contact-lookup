# Contact Lookup API (Go)


***


## A RESTful Contacts Lookup API, made in Golang & locally hosted, that allows users to performs CRUD operations on a temporary/unsaved database, using http/tcp requests.

### <i> Create, search, edit and delete Contact Records.

<br>

***

### Client Page (Front-End) Homepage: <br>
#### <b>http://localhost:\<Port\>/api/v0-1/home</b>

***

<br><br>

#### List of URL(http://localhost:<Port>/api/v0-1/) + URN (End-points), for Requests against an unstored array, that are currently available:

| URN | Action on Array/DB | Full URI (Using some port, e.g. "8080") |
|:---|:---|:---|
| <ul><li>"/home"</li></ul> | <b><u>HOME/CLIENT PAGE</u></b> | <ul><li>"http://localhost:8080/api/v0-1/home"</li></ul> |
| <ul><li>"/contacts/new"</li></ul> | <b><u>CREATE</u></b> | <ul><li>"http://localhost:8080/api/v0-1/contacts/new"</li></ul> |
| <ul><li>"/contacts/all"</li><li>"/contacts/find?id=\<id\>"</li></ul> | <b><u>READ</u></b> | <ul><li>"http://localhost:8080/api/v0-1/contacts/all"</li><li>"http://localhost:8080/api/v0-1/contacts/find?id=\<id\>"</li></ul> |
| <ul><li>"/"</li></ul> | <b><u>UPDATE</u></b> | <ul><li>"http://localhost:8080/api/v0-1/"</li></ul> |
| <ul><li>"/"</li></ul> | <b><u>DELETE</u></b> | <ul><li>"http://localhost:8080/api/v0-1/"</li></ul> |

##### POST/PUT/DELETE available via middleware.

<br>

***
***

<br>

|Version| Changes|
|:---|:---|
|Version 0.0.1 [2020-03-01]|<ul><li>Initial Commit.</li><li>Add "main.go" file.</li><li>Adds basic infrastructure and request methods, against the local unsaved sudo-databse</li><li>Add README.md</li></ul>|