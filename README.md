# TestGoApp

## Run the application

```
docker build . -t test-app-go:1.0
docker run -d -p 8080:8080 test-app-go:1.0 
```

## Rest APIs 
1. get account details 
      ```curl --location --request GET 'localhost:8080/accounts' ```
    

