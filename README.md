# Intro

The code was designed to have clear reponsabilities and to be easy to test and mock. I developed it guided to interfaces.

# Folder Structure

## **api/services**

Here is the brain, the core functionality is written here. Files here will be responsible for use cases like: Create a Transaction and an Account

## **api/controllers**

Where I store the transactionsController and accountsController. Their responsability is to get a request, call a service and return a response to the client.

## **api/models**

A model is a structure that shapes the data. It defines the column attributes and it is used for migrations. I don't store any business logic here

## **api/repositories**

Here I'm storing some structs that will make the queries in the DB. Create, Update, Delete and Find queries are wrapped inside it. This way my service doesn't need to know which adapter is being used (postgres, mysql, mongo, etc...)


# How to run

```
$ git pull https://github.com/emmanuelperotto/pismo-test.git
```

```
$ cd pismo-test
```

```
$ docker-compose up --build
```
