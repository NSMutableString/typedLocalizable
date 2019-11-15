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
	fmt.Println("Generate typed translations from Localizable.strings")
	readFile("./nl.proj/Localizable.strings")
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

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}