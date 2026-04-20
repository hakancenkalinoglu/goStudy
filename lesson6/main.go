package main

import (
	"encoding/json"
	"fmt"
	"io" // Veriyi okumak için bu paketi eklememiz gerekiyor
	"net/http"
	"time"
)

type APIResponse struct {
	Records []Record `json:"Records"`
}

type Record struct {
	Copmany CompanyData `json:"Company"`
}

type CompanyData struct {
	CompanyName string `json:"name"`
	Users       []User `json:"Users"`
}

type User struct {
	UserName string `json:"name"`
}

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	response, err := client.Get("https://mocki.io/v1/3ef708ca-aa87-4dfa-b071-a86346dac9e2")

	if err != nil {
		fmt.Println("Error", err)
		return
	}
	// BEST PRACTICE: Hata (err) kontrolünden hemen sonra Body'yi kapatmayı defer etmelisiniz.
	// Böylece aşağıda başka bir hata olsa bile kaynağın kapatılacağından emin oluruz.
	//(memory leak) önlemek için en iyi pratik (best practice) defer
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error ", err)
		return
	}

	//-------- UNMARSHALING (DÖNÜŞTÜRME) --------
	var result APIResponse

	// json.Unmarshal, bodyBytes içindeki veriyi alır ve result değişkeninin içine yerleştirir.
	// NOT: result değişkeninin "adresini" (&) veriyoruz ki fonksiyon içini doldurabilsin!
	err = json.Unmarshal(bodyBytes, &result)

	if err != nil {
		fmt.Println("JSON dönüştürme hatası:", err)
		return
	}

	//------- DÖNGÜ İLE VERİLERİ ÇEK ---------
	for _, record := range result.Records {
		fmt.Printf("Şirket adı: %s\n", record.Copmany.CompanyName)

		for index, user := range record.Copmany.Users {
			fmt.Printf("%d - User Name: %s \n", index+1, user.UserName)
		}
	}

}
