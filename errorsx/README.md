# errorsx HTTP Error Handling 


There are two parts to error handling:

1. Handle error and wrap it inside `internal` business logic.
2. Report error and write logs inside `handlers` package.

---
1.

```go

var errValidationFailed = errors.New("validation failed on field _name_here_")


// Inside `internal` package

func CreateCustomer(.... ) {

    // ... performing authorization check here.
    err := SomeAuthCheck(input)
    if err != nil {
        // HERE we provide context about the error in a form of unique text and also wrap error.
        return nil, fmt.Errorf("Invalid username or password: %w", errorsx.NewUnauthorizedError(err))
    }    

    // ... performing validation here.
    err := SomeValidation(input)
    if err != nil {
        // HERE we provide context about the error in a form of unique text and also wrap error.
        return nil, fmt.Errorf("Failed CreateCustomer input validation: %w", errorsx.NewBarRequestError(err))
    } 
    
    // Internal errors are not double wrapped, they just wrapped using fmt %w, see below:

    err := CallExternalService() 
    if err != nil {
        // HERE it is Internal Error reported as 500 error:
        return nil, fmt.Errorf("Failed to call external service _name_here_: %w", err)
    }

}

```

---
2.

```go

    // Then inside `handlers` package do

    response, err := customer.CreateCustomer(c.Request.Context(), svr.Db, ...other params...)
    if err != nil {
        errorsx.HandleError(c, err)  // <-- HERE
        svr.Logger.Errorf("Failed CreateCustomer: %v", err)
        return
    }

```

