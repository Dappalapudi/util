
## Usage

JWT based package with Gin middleware.

Two primary use cases:

```go
    router.Use(auth.ValidateToken(signingKey, nil)) // as middleware
    
    // or

    tokenString, err := claims.GenerateToken(signingKey) // as token string generator
```

### Middleware

```go
    router := gin.New()
    router.Use(gin.Logger())
    router.Use(gin.Recovery())
    
    // signingKey - comes from `secrets`.
    // auth.ValidateToken middleware validates JWT token and also sets gin "Auth" value to JWT Claims.
    
    router.Use(auth.ValidateToken(signingKey, nil))

    // ------
    // Inside the Gin handlers you can access Claims using auth.GetClaims(c)
```

### Generate Token Upon Login

```go
    // signingKey - comes from `secrets`.

    claims := &Claims{
        Username: "jane doe",
        // set other fields here
    }

    tokenString, err := claims.GenerateToken(signingKey)
    if err != nil {
        // generate token failed
        // log it here and return error to the frontend    
        return
    }
    w.Header().Set("Token", tokenString)
```

### Anonymous Routes

*AnonymousRoutes* let you setup routes that will be excluded from JWT token validation. This is important for public routes that need to function without login.

Example is login route itself that will permit login to get executed.

```go
var anonymousRoutes auth.AnonymousRoutes

// this will allow POST /auth/v1/login executed without JWT token.
anonymousRoutes.Add("POST", "/auth/v1/login")

router.Use(auth.ValidateToken([]byte(key), &anonymousRoutes))  
```

Below is a default setup that requires JWT tokens for all the routes in the microservice:

```go
router.Use(auth.ValidateToken([]byte(key), nil))
```