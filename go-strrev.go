package main

import(
      "fmt"
      "unicode/utf8"
      "encoding/json"
      "os"
)

func Reverse(s string) string {
        r := []rune(s)
        for i, j := 0, utf8.RuneCountInString(r)-1; i < utf8.RuneCountInString(r)/2; i, j = i+1, j-1 {
                r[i], r[j] = r[j], r[i]
        }
        return string(r)
}

func main() {
  byte_map := map[string]int{"reverse" : Reverse(os.Args[1])}

  result, _ := json.Marshal(byte_map)
  fmt.Println(string(result))
}

