# Go Restful API MyExpenses

## Features
The following functions are a set for creating this web APIs:
- Routing with [Fiber](https://github.com/gofiber/fiber)
- Database work support with [pgxpool](https://pkg.go.dev/github.com/jackc/pgx/v4/pgxpool)

## API Routes
| Path          | Method | Request/Exmaple               |  Desription                                           |                                    
| ------------- | ------ | ----------------------------- | ----------------------------------------------------- |
| /transaction   | POST  | {"amount": 10.5, "currency": "USD", "category": "fastfood", "type": "expense"} | Add transaction |   
| /transactions  | POST  | { "user_id": 1 }              | Chech transactions history                            |     
| /balance   | POST  | { "ID": 0, "Name": "John" }       | Check balance                                         |      
