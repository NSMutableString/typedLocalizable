package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strings"
)

func main() {
	fmt.Println("Generate typed translations from Localizable.strings")
	readFile("./nl.proj/Localizable.strings")
}

func readFile(filePath string) {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        match, _ := regexp.MatchString("\"(...+)\" = \"(...+)\";", line)
        if match {
            s := strings.Split(line, "\"")
            fmt.Println(s[1]) // key
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
