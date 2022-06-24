# MyheresaChallenge

## How to Run

```bash
go build -o MytheresaChallenge
```

and then:

```bash
./MytheresaChallenge [your port number]
```

There is an executable file in repository, so you can just run the last command above.

for example:

```bash
./MytheresaChallenge 3030
```

and project is up on port 3030 of your machine. You can access its endpoint from browser or postman like this:

[http://127.0.0.1:3030/api/products?category=boots&priceLessThan=80000](http://127.0.0.1:3030?category=boots&priceLessThan=80000)


## How to Test

You just need to run the command below in main directory:

```bash
go test
```
