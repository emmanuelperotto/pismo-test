# Intro

The code was designed to have clear reponsabilities and to be easy to test and mock. I developed it guided to interfaces.

# Folder Structure

## **src/services**
![image]()

Here is the brain, the core functionality is written here. Files here will be responsible for use cases like: Create a Transaction and an Account

## **src/controllers**
![image]()

Where I store the transactionsController and accountsController. Their responsability is to get a request, call a service and return a response to the client.


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

**PS**:

You need to send the params in camelCase: "amountCents" instead of "amount_cents"
