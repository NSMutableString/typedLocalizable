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

    writeHeader(&translationsBuffer)
    writeContainingStructStart(&translationsBuffer)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        match, _ := regexp.MatchString("\"(...+)\" = \"(...+)\";", line)
        if match {
            parts := strings.Split(line, "\"")
            key := parts[1]
            propertyName := strcase.ToLowerCamel(key)
            writeTranslationKeyLine(&translationsBuffer, key, propertyName)
        }
    }

    writeContainingStructEnd(&translationsBuffer)

    fileData := []byte(translationsBuffer.String())
    ioutil.WriteFile("Translations.swift", fileData, 0644)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func writeHeader(buffer *bytes.Buffer) {
    buffer.WriteString("//\n")
    buffer.WriteString("//  Translations.swift\n")
    buffer.WriteString("//\n")
    buffer.WriteString("//  Generated code that contains the available keys used in the Localizable.strings\n")
    buffer.WriteString("//  Copyright © 2019. All rights reserved.\n")
    buffer.WriteString("//\n")
    buffer.WriteString("\n")
}

func writeTranslationKeyLine(buffer *bytes.Buffer, key string, propertyName string) {
    buffer.WriteString("\tstatic let " + propertyName + " = \"" +  key + "\"\n")
}

func writeContainingStruct(buffer *bytes.Buffer, name string) {
    buffer.WriteString("public struct Translations {\n")
}

func writeContainingStructStart(buffer *bytes.Buffer) {
    buffer.WriteString("public struct Translations {\n")
}

func writeContainingStructEnd(buffer *bytes.Buffer) {
    buffer.WriteString("}\n")
}