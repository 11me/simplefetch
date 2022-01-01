# simplefetch
It is a simple golang http client.

Example:
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
