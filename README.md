# Osheet

Combined 'Oshi' and 'Sheet' into Osheet. In Japenese 'Oshi'(推し) means a person who you support. And sheet is sheet like datasheet.

## Tech Stack
- Language: golang
- Web Framework: Gin
- Database: postgresql

## Project Folder Structure

### main.go -> controller -> services -> model

```
|-- Oshhet-api
    |-- database
    |-- docs
    |-- v1
    |   |-- controllers
    |       |-- channelController.go
    |
    |   |-- models
    |       |-- channel.go
    |
    |   |-- services
    |       |-- channelService.go
    |-- .env.example
    |-- main.go
```

## Generate swagger document
```
swag init --parseDependency --parseInternal
```
