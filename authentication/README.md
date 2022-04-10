# microservice: authentication

Responsible to:

- Login (Generate JWT token)
- Refresh (Update token basing with an old one)
- Validate sessions
- Invalidate sessions

It uses JWT under the hood. It alsos uses Go subpackages.
Basing this approach on https://curity.io/resources/learn/jwt-best-practices/


## How to generate some key?

This project is currently using ECSDA P512 algorithm to auth the JWT token.
To generate a new key, type:

```bash
# Generates a new key, which should be in some var
ssh-keygen -t ecdsa -b 521
```

## TODO

- [ ] Give users the possibility to use NonTransaction elements