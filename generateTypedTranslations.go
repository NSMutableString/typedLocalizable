package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strings"
    "bytes"
    "io/ioutil"
    "github.com/iancoleman/strcase"
)

func main() {
    validateArguments()
    file := os.Args[1]
    readFile(file)
}

func validateArguments() {
    if len(os.Args) != 2 {
        fmt.Println("Usage:", os.Args[0], "FILE")
        os.Exit(0)
    }
}

func readFile(filePath string) {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var translationsBuffer bytes.Buffer
    codeGenerator := iOSCodeGenerator {
        buffer: &translationsBuffer,
    }
    codeGenerator.writeHeader()
    codeGenerator.writeContainingStructStart()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        match, _ := regexp.MatchString("\"(...+)\" = \"(...+)\";", line)
        if match {
            parts := strings.Split(line, "\"")
            key := parts[1]
            propertyName := strcase.ToLowerCamel(key)
            codeGenerator.writeTranslationKeyLine(key, propertyName)
        }
    }

    codeGenerator.writeContainingStructEnd()

    fileData := []byte(translationsBuffer.String())
    ioutil.WriteFile("Translations.swift", fileData, 0644)
    fmt.Println("Code generation: SUCCESS")

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}