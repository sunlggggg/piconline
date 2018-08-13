# go使用json

## toJson
```go
res := map[string]interface{}{
        "status":  statue,
        "message": userId,
}
bytesRes, _ := json.Marshal(res)
```
