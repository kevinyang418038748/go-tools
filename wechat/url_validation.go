package wechat

import (
	"log"
	"net/http"
)

//func validateUrl(w http.ResponseWriter, r *http.Request) bool {
//	timestamp := strings.Join(r.Form["timestamp"], "")
//	nonce := strings.Join(r.Form["nonce"], "")
//	signatureGen := makeSignature(timestamp, nonce)
//
//	signatureIn := strings.Join(r.Form["signature"], "")
//	if signatureGen != signatureIn {
//		return false
//	}
//	echostr := strings.Join(r.Form["echostr"], "")
//	fmt.Fprintf(w, echostr)
//	return true
//}

func processRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !validateUrl(w, r) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return
	}
	log.Println("Wechat Service: validateUrl Ok!")
}

