// cmd/smartcontractsengine/main.go
package main

import (
"flag"
"log"
"os"

"smartcontractsengine/internal/smartcontractsengine"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := smartcontractsengine.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
