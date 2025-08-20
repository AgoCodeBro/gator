# Gator

## Requirements

This project require that you have the following installed:
    
    - Go
    - Postgres

## Installation

After installing the requirements, open a terminal and run the following

```
go install github.com/AgoCodeBro/gator
```

## Set Up

On your first usage use `./gator register {yourName}` To add a user and log in. Most commands require a user to be logged in.
Then in a seprate window use `.gator agg {timeInterval}` and leave it running to fetch posts from your feeds at the set time interval

## Usage

Here are some common commands

### Add a feed

```
./gator addFeed {feedName} {url}
```

### Follow a feed

```
./gator follow {url}
```

### Browse posts

```
./gator browse (optional){limitNumber of posts (default: 2_}
```

