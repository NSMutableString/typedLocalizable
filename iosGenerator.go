package main

import (
    "bytes"
)

type codeGenerator interface {
    writeHeader(buffer *bytes.Buffer)
    writeTranslationKeyLine(buffer *bytes.Buffer, key string, propertyName string)
    writeContainingStruct(buffer *bytes.Buffer, name string)
    writeContainingStructStart(buffer *bytes.Buffer)
    writeContainingStructEnd(buffer *bytes.Buffer)
}

type iOSCodeGenerator struct {}

func(generator iOSCodeGenerator) writeHeader(buffer *bytes.Buffer) {
    buffer.WriteString("//\n")
    buffer.WriteString("//  Translations.swift\n")
    buffer.WriteString("//\n")
    buffer.WriteString("//  Generated code that contains the available keys used in the Localizable.strings\n")
    buffer.WriteString("//  Copyright © 2019. All rights reserved.\n")
    buffer.WriteString("//\n")
    buffer.WriteString("\n")
}

func(generator iOSCodeGenerator) writeTranslationKeyLine(buffer *bytes.Buffer, key string, propertyName string) {
    buffer.WriteString("\tstatic let " + propertyName + " = \"" +  key + "\"\n")
}

func(generator iOSCodeGenerator) writeContainingStruct(buffer *bytes.Buffer, name string) {
    buffer.WriteString("public struct Translations {\n")
}

func(generator iOSCodeGenerator) writeContainingStructStart(buffer *bytes.Buffer) {
    buffer.WriteString("public struct Translations {\n")
}

func(generator iOSCodeGenerator) writeContainingStructEnd(buffer *bytes.Buffer) {
    buffer.WriteString("}\n")
}