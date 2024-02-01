# Todo App

## Table of contents

1. [Project's layout](#projects-layout)
2. [Stack](#stack)
3. [Commmit style](#commit-style)

## Stack

- go
- mongodb

## Project's layout

### Cmd

Includes main file.  

### Internal

Includes golang packages that usage only in this repository.

#### Internal/app

Projet's start point

#### Internal/config

Config reader

#### Internal/models

Includes projet's models

#### Services

Merge all branch of projet's

#### Storage/mongo

Includes files for using mongodb

##### Storage/mongo/converter

Includes function for convert models to mongo objects and mongo objects to models

##### Storage/mongo/object

Includes mongo object

#### Transport

Includes http router

##### Transport/todo

Router for todo model

##### Transport/user

Router for user model

##### Transport/jwtToken

Packages for usage jwt tokens

### Pkg

Includes golang packages that can be add in another repository.

#### Pgk/hashing

Includes function for hashing string

#### Pkg/middleware

Includes some middleware for http route

### Docs

Includes swagger documentation

### Config

Includes config.yml file.

## Commit Style

### Commit type

#### add

When added new feature into project.

#### fix

When repair project's bug.

#### refactor

When change some code.

### Commit message

*type*: message *in file/packages/files*
