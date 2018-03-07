# stackerr


```go
err := stackerr.New()
err.Pushf("Invalid User Name: %v", input.Username)
err.Pushf("Invalid Phone Number: %v", input.PhoneNumber)
return err.Err() // Wrap errors with error interface, or nil
```
