package main

import (
    "fmt"
    _ "github.com/lib/pq"
)

const (
    c1 = "c1"
    c2 = "c2"
)

var (
    _ = constInitCheck()
    v1 = variableInit("v1")
    v2 = variableInit("v2")
)

func constInitCheck() string {
    if c1 != "" {
        fmt.Println("main: const c1 has been initialized")
    }

    if c2 != "" {
        fmt.Println("main: const c2 has been initialized")
    }

    return ""
}

func variableInit(name string) string {
    fmt.Println("main: var %s ha been initialized\n", name)
    return name
}

func init() {
    fmt.Println("main: first init func invoked")
}

func init() {
    fmt.Println("main: second init func invoked")
}

func main() {
  //logrus.Println("hello, go module mode")
  //logrus.Println(uuid.NewString())
}
