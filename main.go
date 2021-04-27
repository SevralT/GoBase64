package main

import (
    b64 "encoding/base64"
    "fmt"
)

func main() {
    var menu int
    fmt.Println("Кодирование/Декодирование Base64 by SevralT")
    fmt.Println("Меню")
    fmt.Println("1. Кодиорвание")
    fmt.Println("2. Декодирование")
    fmt.Print("Выберите действие (1/2): ")
    fmt.Scanln(&menu)
    if menu == 1 {
        var code string
        fmt.Print("Введите текст для кодирования: ")
        fmt.Scanln(&code)
        coded := b64.StdEncoding.EncodeToString([]byte(code))
        fmt.Println(coded)
    } else if menu == 2 {
        var decode string
        fmt.Print("Введите текст для декодирования: ")
        fmt.Scanln(&decode)
        decoded, _ := b64.StdEncoding.DecodeString(decode)
        fmt.Println(string(decoded))
    }
}