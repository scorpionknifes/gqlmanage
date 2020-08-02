GraphQL API Server with MongoDB
=======

### Introduction

Learning [gqlgen](https://github.com/99designs/gqlgen) to build full GraphQL servers with MongoDB. The following project is built to manage multiple devices in multiple rooms. This GraphQL API is designed to be secure and requires the user to login to manage devices. A React Frontend is being built to use this API.

The following Golang Project uses:
- [MongoDB](https://github.com/mongodb/mongo-go-driver)
- [dataloaden](https://github.com/vektah/dataloaden)
- [chi](https://github.com/go-chi/chi)
- [jwt-go](https://github.com/dgrijalva/jwt-go)

### Instructions

Rename [.env-example](.env-example) to .env and setup variables
```
# MongoDB server config
MONGODB_URL=mongodb://localhost/
MONGODB_DATABASE=gqlmanage

# Server Port
PORT=8000
CORS=http://localhost:8000

# JWT Token Settings
JWT_SECRET=my_jwt_secret
JWT_ISSUER=my_issuer
```

Need to insert first user to access API (TODO)


##### Run server

- default port :8000
- GraphQL Playground default url http://localhost:8080

```bash
go run ./main.go
# OR
go run github.com/scorpionknifes/gqlmanage
```

### Notes

##### Resolvers

Custom resolvers for devices and rooms are created to allow recursive and nested GraphQL calls.

Example of a recursive query.

```graphql
{
  rooms{
    devices{
      room{
        devices{
         ... 
        }
      }
    }
  }
}
```

##### Authentication

Login with username and password to get JWT Token to access other routes. JWT Token has an expire time of 7 days.

```graphql
mutation login{
  login(input:{
    username: "username"
    password: "password"
  }){
    authToken{
      accessToken
    }
  }
}
```

Example result getting JWT token

```graphql
{
  "data": {
    "login": {
      "authToken": {
        "accessToken": "JWT_TOKEN_HERE"
      }
    }
  }
}
```

Use JWT token to authenticate. Add Authorization Bearer token in Http Header 
```json
{
  "Authorization": "Bearer JWT_TOKEN_HERE"
}
```

##### Dataloader

[Read Here](https://gqlgen.com/reference/dataloaders/) for more info about dataloader in gqlgen.

Dataloader solves repeated queries. The dataloader is used in custom resolver for devices.

Example of repeated queries - room needs to be called multiple times for each device that exist

```graphql
{
  devices{
    room{
      ...
    } 
  }
}
```

##### Subscriptions (with Redis)

[Read Subscription with Redis](https://github.com/99designs/gqlgen/issues/846)

[Read Passing token using Subscription](https://github.com/99designs/gqlgen/issues/691#issuecomment-503352009)

Websocket subscription is secured with jwt by passing the token through "graphql-ws" header. A different jwt extract method is used.

Redis pub/sub for running multiple instances.

Example of using a jwt token using Apollo Client in ReactJS
```js
const subscriptionClient = new SubscriptionClient(
    "ws://localhost:8000/query",
    {
        reconnect: true,
    },
    null,
    [ "graphql-ws", 'JWT TOKEN' ])

const wsLink = new WebSocketLink(subscriptionClient);
```
