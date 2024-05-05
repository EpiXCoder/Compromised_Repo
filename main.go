package main

import (
    "fmt"
    "net/http"
    "os/exec"
)

func main() {
    // Hardcoded credentials (fake examples)
    const password = "supersecret"
    const apiKey = "12345abcde"

    // Insecure HTTP URL
    response, err := http.Get("http://example.com")
    if err != nil {
        fmt.Println("Error fetching data:", err)
        return
    }
    defer response.Body.Close()

    fmt.Println("Fetched data:", response)

    // Potential command injection
    command := "ls"
    args := "-la"
    cmd := exec.Command(command, args)
    output, err := cmd.Output()
    if err != nil {
        fmt.Println("Error executing command:", err)
        return
    }

    fmt.Println("Command output:", string(output))
}
