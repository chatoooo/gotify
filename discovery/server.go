package discovery

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"encoding/json"
	"fmt"
	"gotify/dh"
)

var privateKey *dh.PrivateKey

type Info struct {
	Status int `json:"status"`
	StatusString string `json:"statusString"`
	SpotifyError int `json:"spotifyError"`
	Version string `json:"version"`
	DeviceId string `json:"deviceId"`
	RemoteName string `json:"remoteName"`
	ActiveUser string `json:"activeUser"`
	PublicKey string `json:"publicKey"`
	DeviceType string `json:"deviceType"`
	LibraryVersion string `json:"libraryVersion"`
	AccountReq string `json:"accountReq"`
	BrandDisplayName string `json:"brandDisplayName"`
	ModelDisplayName string `json:"modelDisplayName"`
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Query()["action"][0] {
	case "getInfo":
		handleGetInfo(w,r)
	case "setUser":
		handleSetUser(w,r)
	}
}

func handleGetInfo(w http.ResponseWriter, r *http.Request) {
	var slimpubkey []byte = privateKey.SlimPub().GX.Bytes()
	var dhPublicB64 []byte = make([]byte, base64.StdEncoding.EncodedLen(len(slimpubkey)))
	base64.StdEncoding.Encode(dhPublicB64, slimpubkey)
	fmt.Println(r.Header)
	info := Info{
		Status: 101,
		StatusString: "OK",
		SpotifyError: 0,
		Version: "2.1.0",
		DeviceId: "be8e4713-7c73-4ad9-ac8e-1d9086ed0582",
		RemoteName: "GoSpot",
		ActiveUser: "",
		PublicKey: string(dhPublicB64),
		DeviceType: "SPEAKER",
		LibraryVersion: "0.1.0",
		AccountReq: "PREMIUM",
		BrandDisplayName: "gospot",
		ModelDisplayName: "gospot",
	}
	fmt.Println("handling getInfo")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func handleSetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling setUser")
}

func StartServer(port int) {
	privateKey,_ = dh.MakeKey(rand.Reader, dh.Group1)
	http.HandleFunc("/", handleReq) // set router
	go func() {
		http.ListenAndServe(":42424", nil)
	}()
}
