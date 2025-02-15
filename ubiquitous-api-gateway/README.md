# Ubiquitous API gateway

This project serves an API gateway service, which depends on [LURA](github.com/luraproject/lura).

## Ubiquitous-alchemy-server

1. configuration:

   - [x] [POST] check_connection
   - [x] [GET] conn
   - [x] [POST] conn
   - [x] [PUT] conn
   - [x] [DELETE] conn

1. schema:

   - [ ] [GET] table_list
   - [ ] [GET] column_list
   - [ ] [POST] table_create
   - [ ] [POST] table_alter
   - [ ] [POST] table_drop
   - [ ] [POST] table_rename
   - [ ] [POST] table_truncate
   - [ ] [POST] index_create
   - [ ] [POST] index_drop
   - [ ] [POST] foreign_key_create
   - [ ] [POST] foreign_key_drop

1. query:

   - [ ] [POST] select

1. auth:

   - [ ] [POST] invitation
   - [ ] [POST] registration
   - [ ] [POST] login
   - [ ] [GET] check_alive
   - [ ] [DELETE] logout
