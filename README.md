# simplefetch
It is a simple golang http client.

## Exmaples

1) Simple GET
```golang
import fetch "github.com/11me/simplefetch"

func main() {

  res, err := fetch.Get(fetch.Options{
    URL: "http://example.com",
  })
  if err != nil {
    // handle error
  }
  // handle response
}
```

2) Simple POST
```golang
import fetch "github.com/11me/simplefetch"

func main() {

  res, err := fetch.Post(fetch.Options{
    URL: "http://example.com",
    Data: fetch.Data{
      "id": 1,
      "name": "lime",
    },
  })

  if err != nil { // handle error }
  // handle response
}
```
