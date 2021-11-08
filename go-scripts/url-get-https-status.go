package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "net/http"
)

func main() {

    arg := os.Args[1]
    file, err := os.Open(arg)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        url := scanner.Text()
        resp,err := http.Get(url)
        if err != nil {
            log.Fatal(err)
        }
        status_code := resp.StatusCode
        if status_code == 200 {
            fmt.Println("ACTIVE -",status_code,":",url)
        } else {
            fmt.Println("CHANGED -",status_code,":",url)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
