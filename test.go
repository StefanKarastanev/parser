package main
import (
    "fmt"
    "bufio"
    "os"
    "log"
    "path/filepath"
//     "go/parser"
//     "net/http"
//     "text/template/parse"
)
func main(){
// func readString
// func 
    files, err := filepath.Glob("*")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(files) // contains a list of all files in the current directory
// file, err := os.Open("./")
// if err != nil {
//     log.Fatal(err)
// }
// defer file.Close()

scanner := bufio.NewScanner(files)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
}
