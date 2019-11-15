package main

import (
    "bytes"
)

type codeGenerator interface {
    writeHeader()
    writeContainingStructStart()
    writeTranslationKeyLine(key string, propertyName string)
    writeContainingStructEnd()
}

type iOSCodeGenerator struct {
    buffer *bytes.Buffer
}

func(generator iOSCodeGenerator) writeHeader() {
    generator.buffer.WriteString("//\n")
    generator.buffer.WriteString("//  Translations.swift\n")
    generator.buffer.WriteString("//\n")
    generator.buffer.WriteString("//  Generated code that contains the available keys used in the Localizable.strings\n")
    generator.buffer.WriteString("//  Copyright © 2019. All rights reserved.\n")
    generator.buffer.WriteString("//\n")
    generator.buffer.WriteString("\n")
}

func(generator iOSCodeGenerator) writeContainingStructStart() {
    generator.buffer.WriteString("public struct Translations {\n")
}

func(generator iOSCodeGenerator) writeTranslationKeyLine(key string, propertyName string) {
    generator.buffer.WriteString("\tstatic let " + propertyName + " = \"" +  key + "\"\n")
}

func(generator iOSCodeGenerator) writeContainingStructEnd() {
    generator.buffer.WriteString("}\n")
}