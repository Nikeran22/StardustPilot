// cmd/stardustpilot/main.go
package main

import (
"flag"
"log"
"os"

"stardustpilot/internal/stardustpilot"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := stardustpilot.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
