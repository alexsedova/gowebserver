package main

import (
    "io"
    "fmt"
    "log"
    "net/http"
    "strconv"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world!")
}

func to_roman(n int)  string {
    if n == 1 {return "I" }
    if n == 2 {return "II"}
    if n == 3 {return "III"}
    if n == 4 {return "IV" }
    if n == 5 { return "V" }
    if n == 6 { return "VI"}
    if n == 7 { return "VII" }
    if n == 8 { return "VIII" }
    if n == 9 { return "IX" }
    if n == 10 { return "X" }
    return "This number is not yet supported.. Sorry!"
}

type romanGenerator int
func (n romanGenerator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ascii_num := r.URL.Path[7:]
    i, err := strconv.Atoi(ascii_num)
    if err != nil {
        log.Print(err)
    }
    fmt.Fprintf(w, "Here's your number: %s\n", to_roman(i))
}



func main() {
    h := http.NewServeMux()

    h.Handle("/roman/", romanGenerator(1))
    h.HandleFunc("/", hello)

    err := http.ListenAndServe(":8000", h)
    log.Fatal(err)
}
