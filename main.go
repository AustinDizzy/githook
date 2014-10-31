package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"github.com/martini-contrib/auth"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	log.Println("Starting server...")
	http.HandleFunc("/github", gitHandle)
	http.ListenAndServe(":8800", nil)
}

func gitHandle(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if r.Method == "POST" && r.Header["X-Github-Event"][0] == "push" {
		mac := hmac.New(sha1.New, []byte("testinghmac"))
		mac.Reset()
		mac.Write(body)
		macSum := fmt.Sprintf("sha1=%x", mac.Sum(nil))
		log.Println(macSum)
		log.Println(r.Header["X-Hub-Signature"][0])
		log.Println(auth.SecureCompare(r.Header["X-Hub-Signature"][0], macSum))
		log.Println(r)
		os.Chdir("/var/www/austindizzy.me/")
		cmd := exec.Command("git", "pull")
		out, err := cmd.Output()

		if err != nil {
			log.Println(err.Error())
			return
		}

		log.Println(string(out))
	}
}
