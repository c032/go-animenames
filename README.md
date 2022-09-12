# go-animenames

Parser for anime names that follow common naming conventions.

## Example

```go
package main

import (
	"fmt"

	"github.com/c032/go-animenames"
)

func main() {
	const name = "[Kantai] Eighty Six (86) - 23 (1920x1080 AC3) [05BD70FE].mkv"
	anime, err := animenames.Parse(name)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Title: %s\n", anime.Title)
	fmt.Printf("Group: %s\n", anime.Group)
	fmt.Printf("Episode: %d\n", anime.Episode)
	fmt.Printf("CRC32: %s\n", anime.CRC32)
	fmt.Printf("IsBD: %#v\n", anime.IsBD)
}
```

Output:

```
Title: Eighty Six (86)
Group: Kantai
Episode: 23
CRC32: 05BD70FE
IsBD: false
```

## License

Apache 2.0
