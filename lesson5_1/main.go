package main

import (
	"fmt"
	"time"
)

type Isleyici interface {
	Isle() IslemSonucu
}

type DijitalSiparis struct {
	ID    int
	Tutar float64
}
type FizikselSiparis struct {
	ID    int
	Tutar float64
}

type IslemSonucu struct {
	SiparisID int
	Durum     string
	GecenSure time.Duration
}

func (digitalSiparis DijitalSiparis) Isle() IslemSonucu {
	start := time.Now()
	time.Sleep(1 * time.Second)
	return IslemSonucu{SiparisID: digitalSiparis.ID, Durum: "Başarili", GecenSure: time.Since(start)}
}

func (fizikselSiparis FizikselSiparis) Isle() IslemSonucu {
	start := time.Now()
	time.Sleep(3 * time.Second)

	return IslemSonucu{SiparisID: fizikselSiparis.ID, Durum: "Başarili", GecenSure: time.Since(start)}
}

func SiparisiIsle(isleyici Isleyici, ch chan IslemSonucu) {
	ch <- isleyici.Isle()
}

func main() {
	//ch oluşturuldu
	sonuclar := make(chan IslemSonucu)

	//sturct oluşturuldu
	siparis1 := DijitalSiparis{ID: 1, Tutar: 12.00}
	siparis2 := FizikselSiparis{ID: 2, Tutar: 12.00}

	//
	go SiparisiIsle(siparis1, sonuclar)
	go SiparisiIsle(siparis2, sonuclar)

	for i := 0; i < 2; i++ {
		fmt.Println(<-sonuclar)
	}

}

/* fonksiyonu çağırırken başına go kelimesini ekleyeceğiz.
Ancak burada bir sorun var: Bir goroutine (arka plan işlemi) return ile geriye değer döndüremez.
Çünkü o işini bitirdiğinde, ana program (main) o satırı çoktan geçmiş olur.
O değeri main fonksiyonuna güvenle taşıyabilmek için bir boru hattına, yani Channel'a ihtiyacımız var. */
