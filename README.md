# hashapi-challenge
This is a Challenge to build a API that returns the hash of a token.

Install instructions:
1 - Clone this repository to your favorite directory.
2 - Move the directories hash_api and hashapi_db_conn to the src libraries in
the GOPATH.
3 - run the project execution the command "go run api.go" or build the
application.
4 - Use your favorite browser to the curl command in the shell to do the
communication with the API

The API answer for the following routes:
1 - http://[your host]:[your door]/api/create_hash?t=[token]
Action: Insert the token on the database and returns a message with the hash of
the token
2 - http://[your host]:[your door]/api/return_hashs
Action: Returns a JSON with all the hash and correspondents tokens on the database
3 - http://[your host]:[your door]/api/return_text?t=[token]
Action: Returns a JSON with the token that matchs with the passed token or a
404 (not found) code.
