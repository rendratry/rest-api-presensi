package main

import (
	"fmt"
	"github.com/robfig/cron"
	"golang/rest-api-presensi/entity/domain"
	"golang/rest-api-presensi/helper"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestEpoch(t *testing.T) {
	waktu := time.Now().UnixNano() / 1000000
	fmt.Println(waktu)
	fmt.Println(reflect.TypeOf(waktu))
}

func TestConvertEpoch(t *testing.T) {
	time := time.Unix(-2211750432000, 0)
	fmt.Println(time)
}

func TestCron(t *testing.T) {
	epochMillis := int64(1674482340000)

	// Convert milliseconds to time
	scheduledTime := time.Unix(0, epochMillis*int64(time.Millisecond))
	fmt.Println(scheduledTime)
	fmt.Println(scheduledTime.Format("15:04"))

	c := cron.New()
	c.AddFunc("0 0 "+scheduledTime.Format("15:04")+" 23 1 * 2023", func() {
		_, err := http.Get("https://api.telegram.org/bot5836520934:AAGQ_iQvY7Hbm5goVRPbwH-k57p25dA-Gns/SendMessage?chat_id=1359729975&text=Halorendrainitescron")
		if err != nil {
			panic(err)
		}
	})
	c.Start()
	select {}
}

func TestDataType(t *testing.T) {
	OtpTime := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println("tipe data :", reflect.TypeOf(OtpTime))
}

func TestParsing(t *testing.T) {
	jam := "7.45"
	parts := strings.Split(jam, ".")
	hour := parts[0]
	minute := parts[1]

	fmt.Println("hour:", hour)
	fmt.Println("minute:", minute)
}

func TestJWTValidate(t *testing.T) {
	iss := domain.JwtClaims{}
	iss = helper.ValidateJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzYwNDQ2NDIsImlhdCI6MTY3NTkzNjY0MiwiaXNzIjoiUmVuZHJhIiwic3ViIjoicmVuZHJhdHJ5a3VzdW1hQGdtYWlsLmNvbSJ9.qXI0WPkvHKl03K5UwTbTjO42kd5WiGu1bYUk0bttDo0", "secret-key")
	fmt.Println(iss)
	status := helper.GetIssuer(iss.Subject)
	fmt.Println(status)
}

func TestCreateJWT(t *testing.T) {
	jwt := helper.CreateNewJWT("Rendra", "rendratry1@gmail.com", "secret-key", time.Hour*24)
	fmt.Println(jwt)
}
