package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"servertestgo/models"
	"time"
)

// por medio de este paquee interactuamos con okx

// función para crear firma conforme a los estándares de okx, los datos de secret, apy key y demás , se leen desde .env
func sign(method, path, body string) (string, string) {
	format := "2006-01-02T15:04:05.999Z07:00"
	t := time.Now().UTC().Format(format)
	ts := fmt.Sprint(t)
	fmt.Println("Timestamp:" + ts)
	s := ts + method + path + body
	fmt.Println("sign:" + s)
	p := []byte(s)
	h := hmac.New(sha256.New, []byte(os.Getenv("secretkey")))
	h.Write(p)
	return ts, base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// Do the http request to the server
func EjecutarConsulta(method, path, urlfilters, bodyJson string) (*http.Response, error) {
	var u string
	u = fmt.Sprintf("%s%s%s", os.Getenv("OKX_API_URL"), path, urlfilters)
	var (
		r   *http.Request
		err error
		j   []byte
	)
	if bodyJson == "{}" {
		bodyJson = ""
	}
	r, err = http.NewRequest(method, u, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")

	//aquí ponemos los header de seguridad
	timestamp, sign := sign(method, path, urlfilters+bodyJson)
	r.Header.Add("OK-ACCESS-KEY", os.Getenv("apikey"))
	r.Header.Add("OK-ACCESS-PASSPHRASE", os.Getenv("PASSPHRASE"))
	r.Header.Add("OK-ACCESS-SIGN", sign)
	r.Header.Add("OK-ACCESS-TIMESTAMP", timestamp)
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Print(err.Error())
	}

	return resp, err
}

func SolicitarBookOrder(Currency string) {
	resp, err := EjecutarConsulta("GET", "/api/v5/market/books", "?instId="+Currency, "")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	var dataorder models.OrderBook
	json.NewDecoder(resp.Body).Decode(&dataorder)
	fmt.Println(dataorder)
}
