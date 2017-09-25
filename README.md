## GoCatchEm

GoCatchEm is a RESTful API written in Go. I started this
project to learn Golang and get familiar with the tools and libraries offered by
the language. Also, I like Pokemons. It doesn't use any framework, just plain
old go, the only addition being mux for routing requests.

I intend to use this project as a sandbox to improve my go skills. The following
improvements are in the pipeline :

- [ ] Rate Limiting
- [ ] Pagination
- [ ] Logging
- [ ] Return a location header with a URL to the "created resource"
- [ ] Update stats in PUT request (Embedded struct reflection)

The API returns information about all pokemons (name, types, stats etc).

## Who's it for ?

Front end developers (React, Angular, Vue crowd) who want to prototype their
ideas and want a simple way to fetch data without having to register developer
accounts on the major platforms (Like FB, Twitter ...). Or maybe you're
developing an HTTP library and want a quick way to make HTTP calls to an
external service.

## Endpoints

**Get all pokemons**
* `GET` /pokemons

**Get pokemon by name**
* `GET` /pokemons/{name:string}

**Get pokemons by type**
* `GET` /pokemons?type={type:string}

**Get pokemons by generation**
* `GET` /pokemons?generation={generation:int}

_Note: the following endpoints will simulate insertion, update and deletion but
the changes won't be persisted to the database_

**Add a new pokemon**
* `POST` /pokemons

Request body must contain (in JSON format):
```javascript
{
    "name": type string,
    "jp_name": type string,
    "types": type string (comma separated),
    "stats": {,
        "hp" type int,
        "attack" type int,
        "defense" type int,
        "sp_attack" type int,
        "sp_defense" type int,
        "speed" type int,
    },,
    "bio": type string,
    "generation": type int
}
```

**Update a pokemon**
* `PUT` /pokemons/{name}

Request body must contain any of the following (in JSON format):

```javascript
{
    "name": type string,
    "jp_name": type string,
    "types": type string (comma separated),
    "stats": {,
        "hp" type int,
        "attack" type int,
        "defense" type int,
        "sp_attack" type int,
        "sp_defense" type int,
        "speed" type int,
    },,
    "bio": type string,
    "generation": type int
}
```

**Delete a pokemon**
* `DELETE` /pokemons/{name}

## Contributing

All contributions, remarks and comments are welcome ! Just open an Issue or
Initiate a PR and we'll look into it :smile:


