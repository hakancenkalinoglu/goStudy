package main

import (
    "fmt"
    "math"
)

// 1. Interface Tanımlaması
// "Alan" hesaplayabilen her şey bir Shape'tir.
type Shape interface {
    Area() float64
}

// 2. Struct Tanımlamaları
type Rectangle struct {
    Width  float64
    Height float64
}

type Circle struct {
    Radius float64
}

// 3. Interface Metotlarını Uygulama (Implicit Implementation)

// Rectangle için Area metodu
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Circle için Area metodu
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// 4. Polimorfizm (Çok Biçimlilik) Fonksiyonu
// Bu fonksiyon Shape arayüzünü kabul eder. 
// Yani Area() metoduna sahip herhangi bir struct bu fonksiyona gönderilebilir.
func PrintArea(s Shape) {
    fmt.Printf("Bu şeklin alanı: %.2f\n", s.Area())
}

func main() {
    rect := Rectangle{Width: 5, Height: 4}
    circ := Circle{Radius: 3}

    // Her ikisi de Shape interface'ini sağladığı için aynı fonksiyona parametre olabilirler.
    PrintArea(rect) // Çıktı: Bu şeklin alanı: 20.00
    PrintArea(circ) // Çıktı: Bu şeklin alanı: 28.27
}