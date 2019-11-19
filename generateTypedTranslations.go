package main

import (
    "bufio"
    "bytes"
    "errors"
    "fmt"
    "github.com/iancoleman/strcase"
    "io/ioutil"
    "log"
    "os"
    "path"
    "regexp"
    "strings"
)

func main() {
    validateArguments()
    file := os.Args[1]
    outputDir := os.Args[2]
    keys := readKeysFromLocalizableFile(file)
    writeFile(keys, outputDir)
}

func validateArguments() {
    if len(os.Args) != 3 {
        fmt.Println("Usage:", os.Args[0], "FILE OUTPUT_DIR")
        os.Exit(0)
    }
}

func readKeysFromLocalizableFile(filePath string) []string {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    translationKeys := []string{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        key, err := extractKeyFromLine(line)
        if err == nil {
            translationKeys = append(translationKeys, key)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return translationKeys
}

func writeFile(translationKeys []string, outputDir string) {
    var translationsBuffer bytes.Buffer
    codeGenerator := iOSCodeGenerator{
        buffer: &translationsBuffer,
    }
    codeGenerator.writeHeader()
    codeGenerator.writeContainingStructStart()
    for _, key := range translationKeys {
        propertyName := strcase.ToLowerCamel(key)
        codeGenerator.writeTranslationKeyLine(key, propertyName)
    }
    codeGenerator.writeContainingStructEnd()

    fileData := []byte(translationsBuffer.String())
    outputFile := path.Join(path.Dir(outputDir), "Translations.swift")
    ioutil.WriteFile(outputFile, fileData, 0644)
    fmt.Println("Code generation: SUCCESS")
}

func extractKeyFromLine(line string) (string, error) {
    match, _ := regexp.MatchString("\"(...+)\" = \"(...+)\";", line)
    if match {
        parts := strings.Split(line, "\"")
        return parts[1], nil
    }
    return "", errors.New("Line did not match regex")
}
