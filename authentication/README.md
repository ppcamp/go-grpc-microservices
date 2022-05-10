# microservice: authentication

Responsible to:

- Login (Generate JWT token)
- Refresh (Update token basing with an old one)
- Validate sessions
- Invalidate sessions

It uses JWT under the hood. It alsos uses Go subpackages.
Basing this approach on https://curity.io/resources/learn/jwt-best-practices/


## How it works?


```mermaid
sequenceDiagram
  autonumber

  User -->> +Api: Wants to create login
  Api ->> -User: Create a deactivated login and return token to enable
  Note right of User: this token should be sent to user email

  User -->> +Api: Want to activate account
  Api ->> -User: Check sent token and activate login

  User -->> +Api: Request to recover login
  Api ->> -User: Create a secret to recover the login and return it
  Note left of Api: this secret should be sent to user email

  User -->> +Api: Request to update with secret
  Api ->> -User: Validate secret and update password

  User -->> +Api: When logged, tries to update password
  Api ->> -User: Validate JWT and tries to update


  User -->> +Api: Tries to login
  Api ->> -User: Api validate user and return token

  User -->> +OthersServices: wants to get resource
  OthersServices ->> -Api: check if JWT is valid
  Api ->> +OthersServices: allow
  OthersServices ->> -User: return resource

  User -->> +Api: invalidate JWTs
  Api ->> -User: remove all logged JWTs

  User -->> +Api: wants to refresh current JWT
  Api ->> -User: return a refreshed JWT
```


## Database

Note that users don't belong to current database and it's response depends on
the user's crud to remove them.

```mermaid
erDiagram
  users }o--|{ passwords : has
  passwords {
      serial id
      varchar user_id
      varchar user_password
      bool active
      timestamp created_at
      timestamp updated_at
  }
  users {
      serial id

  }
```

## How to generate some key?

This project is currently using ECSDA P512 algorithm to auth the JWT token.
To generate a new key, type:

```bash
# Generates a new key, which should be in some var
ssh-keygen -t ecdsa -b 521
```


Install

```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2
```

## TODO

- [ ] Allow repositories the possibility to use NonTransaction connection
- [ ] Migrations should be runned from each microservice, not from a given cli


## See

- https://github.com/golang-migrate/migrate/tree/master/cmd/migrate


