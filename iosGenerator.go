package main

import (
    "bytes"
)

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