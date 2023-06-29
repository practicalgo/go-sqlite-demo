# Project to demo SQLite + Go

[![Build and Test](https://github.com/practicalgo/go-sqlite-demo/actions/workflows/main.yml/badge.svg)](https://github.com/practicalgo/go-sqlite-demo/actions/workflows/main.yml)


This repository contains an illustration of how to interact with a SQLite
database using `database/sql` and `https://pkg.go.dev/modernc.org/sqlite`.

The example is intentionally chosen to losely match that of the official Go Project
tutorial, [Accessing a relational database](https://go.dev/doc/tutorial/database-access)
which uses MySQL as the database server.

## Why SQLite?

SQLite is perfect for [many use cases](https://www.sqlite.org/whentouse.html).

My choice of using SQLite for this demo is that it is perfect to illustrate how
we can persist data to a SQL database from Go.

## Why modernc.org/sqlite?

It's CGO free which means, we don't have to worry about installation woes
(potential) on different operating systems and architecture

## Brief explanations of the code 

See [this blog post](https://practicalgobook.net/posts/go-sqlite-no-cgo/).
