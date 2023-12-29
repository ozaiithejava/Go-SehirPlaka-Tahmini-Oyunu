package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Rastgele sayı üreteci için seed ayarı
	rand.Seed(time.Now().UnixNano())

	// Şehir-plaka çiftlerini içeren map
	sehirPlakalari := map[string]int{
		"Elazığ":   23,
		"Ankara":   6,
		"İstanbul": 34,
		"İzmir":    35,
		"Bursa":    16,
		"Antalya":  7,
		"Trabzon":  61,
		"Adana":    1,
		"Konya":    42,
		"Erzurum":  25,
	}

	// Oyun başlangıcı
	fmt.Println("Şehir-Plaka Tahmin Etme Oyunu!")
	fmt.Println("-------------------------------")

	puan := 0
	toplamSoru := 0

	for soru := 0; soru < 10; soru++ {
		// Rastgele bir şehir veya plaka seçme
		sorulan, dogruPlaka := rastgeleSec(sehirPlakalari)

		// Kullanıcıdan tahmin al
		fmt.Printf("%d. Soru: %s'nin plakası nedir? ", soru+1, sorulan)
		var kullaniciTahmin int
		fmt.Scan(&kullaniciTahmin)

		// Tahmin kontrolü
		if kullaniciTahmin == dogruPlaka {
			fmt.Println("Doğru!")
			puan++
		} else {
			fmt.Printf("Yanlış! Doğru plaka: %d\n", dogruPlaka)
		}

		toplamSoru++
	}

	// Oyun sonu istatistikleri
	fmt.Printf("\nOyun Bitti! Toplam Puan: %d/%d\n", puan, toplamSoru)
}

// Map'ten rastgele bir şehir-plaka çifti seçme
func rastgeleSec(sehirPlakalari map[string]int) (string, int) {
	// Map'teki anahtarları (şehirleri) slice'a aktarma
	var sehirler []string
	for sehir := range sehirPlakalari {
		sehirler = append(sehirler, sehir)
	}

	// Rastgele bir şehir seçme
	rastgeleIndex := rand.Intn(len(sehirler))
	rastgeleSehir := sehirler[rastgeleIndex]

	// Seçilen şehirin plakasını bulma
	dogruPlaka := sehirPlakalari[rastgeleSehir]

	return rastgeleSehir, dogruPlaka
}
