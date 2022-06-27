package main

import (
	"net/http"
	"strings"
	"time"
)

func checkNetwork() string {
	const sampleUrl = "http://222.222.222.222"
	result := ""

	client := http.Client{Timeout: time.Duration(5 * time.Second)}
	// if user is not authorized in Nandos, any request should be redirected to local host
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		location := req.Response.Header.Get("Location")
		if !strings.Contains(location, sampleUrl) {
			result = location
		}
		return nil
	}
	client.Get(sampleUrl)

	return result
}

func logIn(urlRedirect string) error {
	client := http.Client{Timeout: time.Duration(5 * time.Second)}
	req, _ := http.NewRequest("POST", "http://192.168.100.30:8880/guest/s/default/login?t=1656313299343", nil)
	req.Header.Add("Cookie", "ec=J_DjUvO3mtMv_9oDll3MbRJdp6SSLXrnK715PH5WBfSgNu3uoqN2eVKTklTmQXEdtUNUcGbkRIao3fdiD945UKjD2EsC_iu_OUrxoDL7rWwtUn_si100kodiUyxxYLOobS2sRG4-ZrD6keDl_dowYnLJL-hgg9zqg545FzGR0owu_FBoe9FBXq_svm-XT69Q; unifi-portal-tos_undefined=true; NG_TRANSLATE_LANG_KEY=%22en%22")
	_, err := client.Do(req)
	return err
}
