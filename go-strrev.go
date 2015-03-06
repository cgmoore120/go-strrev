package main

import(
      "fmt"
      "unicode/utf8"
      "encoding/json"
      "os"
      "github.com/golang/example/stringutil"
)

func main() {
  byte_map := map[string]int{"reverse" : string(stringutil.Reverse(os.Args[1]))}

  result, _ := json.Marshal(byte_map)
  fmt.Println(string(result))
}
