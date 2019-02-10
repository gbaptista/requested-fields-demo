# requested-fields-demo [![Maintainability](https://api.codeclimate.com/v1/badges/16dae1ff12447344545e/maintainability)](https://codeclimate.com/github/gbaptista/requested-fields-demo/maintainability)
Sample GraphQL server to show the use of the [requested-fields](https://github.com/gbaptista/requested-fields) lib. Created with [graphql-go](https://github.com/graph-gophers/graphql-go) and [chi](https://github.com/go-chi/chi).

- [Setup](#setup)
- [Schema](#schema)
- [Query](#query)
  - [Request](#request)
  - [Result](#result)
  - [Logs](#logs)

## Setup
```shell
dep init
dep ensure
go build
./requested-fields-demo
```

## Schema
```graphql
schema {
  query: Query
}

type Query {
  user: User
}

type User {
  name: String
  address: Address
}

type Address {
  city: String
  street: String
}
```

## Query

- *http://localhost:3000/graphql*

### Request:
```graphql
query {
  user {
    name
    address {
      street
      city
    }
  }
}
```

### Result:
```json
{
  "data": {
    "user": {
      "name": "Harry Potter",
      "address": {
        "street": "4 Privet Drive",
        "city": "Little Whinging"
      }
    }
  }
}
```

### Logs:
```shell
2019/02/10 19:52:16 Query.User Fields: [name address]
2019/02/10 19:52:16 User.Address Fields: [street city]
2019/02/10 19:52:16 "POST http://localhost:3000/graphql HTTP/1.1"
```
