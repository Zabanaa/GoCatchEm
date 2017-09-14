## GoCatchEm

GoCatchEm is a RESTful API written in Go. I started this
project to learn Golang and get familiar with the tools and libraries offered by
the language. Also, I like Pokemons. It doesn't use any framework, just plain
old go, the only addition being mux for routing requests.

I intend to use this project as a sandbox to improve my go skills. The following
improvements are in the pipeline :

- [] Rate Limiting
- [] Authentication using JWT
- [] Pagination

The API returns information about all pokemons (name, types, stats etc).

## Who's it for ?

Front end developers (React, Angular, Vue crowd) who want to prototype their
ideas and want a simple way to fetch data without having to register developer
accounts on the major platforms (Like FB, Twitter ...). Or maybe you're
developing an HTTP library and want a quick way to make HTTP calls to an
external service.

## Endpoints

**Get all pokemons**
* /pokemons                                 GET

**Get pokemon by name**
* /pokemons/{name:string}                   GET

**Get pokemons by type**
* /pokemons?type={type:string}              GET

**Get pokemons by generation**
* /pokemons?generation={generation:int}     GET

_Note: the following endpoints will simulate insertion, update and deletion but
the changes won't be persisted to the database_

**Add a new pokemon**
* /pokemons                                 POST

Request body must contain (in JSON format):
- name `string`
- jp_name `string`
- types `string` (comma separated list)
- stats `object`
    - hp `int`
    - attack `int`
    - defense `int`
    - sp_atk `int`
    - sp_def `int`
    - speed `int`
- bio `string`
- generation `int`

**Update a pokemon**
* /pokemons/{name}                          PUT

Request body must contain any of the following (in JSON format):

- name `string`
- jp_name `string`
- types `string` (comma separated list)
- stats `object`
    - hp `int`
    - attack `int`
    - defense `int`
    - sp_atk `int`
    - sp_def `int`
    - speed `int`
- bio `string`
- generation `int`

**Delete a pokemon**
* /pokemons/{name}                          DELETE

## Contributing

All contributions, remarks and comments are welcome ! Just open an Issue or
Initiate a PR and we'll look into it :smile:


