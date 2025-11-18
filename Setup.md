# oivan_url_shortener — Instruct / README

This instruction file is showing what do you need and how can you run the codebase at locally

## Overview

Small URL shortener service. This README shows prerequisites and description to build, run, test, and start the server

The server will receive your url encode it, and return to you as in it's format.

For example: https://google.com -> https://my.domain/r/someshortcode

## Prerequisites

- Go 1.18+
- Docker -> Postgres database
- task (go-task) for Taskfile: https://taskfile.dev
- swag (as in swaggo) -> documenting apis
- mockery -> mocks repo, svc, interfaces,... for testings

Install task:

- macOS: brew install go-task/tap/go-task
- Linux: follow https://taskfile.dev/#/installation

## Environment

You can finds env example with the file named .env.example

One thing to note is that the .env.test is an env file maded for only running tests
It is intended for mocking the load configs (which read from .env) to running the testings

## How to use task file?

Run these at the root of the project, ex: task start

```txt
- Build: task build -> This will built go into binary file
- Run: task start  -> This will start the server at locally
- Test: task test -> Run tests oc
- Provision: task provision -> Provision means setups db, swag stuff for your local server.
- Migrate DB: To be enhanced!
```

## Steps to setup to run the code and runs the api at locally

Run these in the terminal where at the root directory of the project

### **Step 1:** `task provision`

For building the docker image for the database and init swagger

### **Step 2:** `task start`

To start the server at locally oc...

### **Step 3:**

Run the api request through anything you like, I would use postman.

Method: POST

Request URL: `http://localhost:8081/url-shortener/encode`

Request body will be like this

```json
{
    "url": "https://google.com/search?q=oivan"
}
```

Then the server would response something like

```json
{
    "url": "http://localhost:8081/r/i"
}
```

