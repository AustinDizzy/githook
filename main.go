package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/martini-contrib/auth"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

type Config struct {
	HookPath string  `json:"hookPath"`
	HookPort string  `json:"hookPort"`
	Sites    []Repos `json:"repos"`
}

type Repos struct {
	Repository string `json:"repository"`
	SecureKey  string `json:"secureKey"`
	Dir        string `json:"dir"`
}

type GHPayload struct {
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
}

var config *Config

func init() {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	c, err := ioutil.ReadFile(dir + "/config.json")
	if os.IsNotExist(err) {
		c, err = ioutil.ReadFile("/usr/local/etc/githook/githook.json")
		if os.IsNotExist(err) {
			log.Println("Config file 'githook.json' doesn't exist in " + dir + " or /usr/local/etc/githook.")
			log.Println("Exiting...")
			os.Exit(1)
		} else if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		} else {
			err = json.Unmarshal(c, &config)
			if err != nil {
				log.Println(err)
			}
			log.Println("Loaded config:", config)
		}
	} else if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	} else {
		err = json.Unmarshal(c, &config)
		if err != nil {
			log.Println(err)
		}
		log.Println("Loaded config:", config)
	}
}

func main() {
	log.Println("Starting server on port", config.HookPort+"...")
	http.HandleFunc(config.HookPath, gitHandle)
	http.ListenAndServe(config.HookPort, nil)
}

func gitHandle(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var ghPayload GHPayload
	_ = json.Unmarshal(body, &ghPayload)

	log.Println(ghPayload.Repository.FullName)
	if r.Method == "POST" && r.Header["X-Github-Event"][0] == "push" {
		for _, repo := range config.Sites {
			if repo.Repository == ghPayload.Repository.FullName {
				mac := hmac.New(sha1.New, []byte(repo.SecureKey))
				mac.Reset()
				mac.Write(body)
				macSum := fmt.Sprintf("sha1=%x", mac.Sum(nil))
				if auth.SecureCompare(r.Header["X-Hub-Signature"][0], macSum) {
					os.Chdir(repo.Dir)
					cmd := exec.Command("git", "pull")
					out, err := cmd.Output()

					if err != nil {
						log.Println(err.Error())
					}

					log.Println(string(out))
				} else {
					log.Println("UNAUTHORIZED REQUEST:", string(body))
				}
			}
		}
	}
}
