package channel

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net"
	"net/http"
	"gotify/config"
	"gotify/protocol/keyexchange"
	"time"
)

type Channel struct {
	Config *config.Config
	connection net.Conn
	apAddress  string
}

type ApResolveResponse struct {
	ApList []string `json:"ap_list"`
}

const (
	listTimeout = 5 * time.Second
	apListUrl   = "https://apresolve.spotify.com/"
	fallbackAp  = "ap.spotify.com:80"
)

func getApList() []string {
	client := http.Client{Timeout: listTimeout}
	var response *http.Response
	var err error
	if response, err = client.Get(apListUrl); err != nil {
		log.Warn("Got error when loading ap list", err)
		return nil
	}
	var apList ApResolveResponse
	if err = json.NewDecoder(response.Body).Decode(&apList); err != nil {
		log.Error("Error parsing JSON response: ", err)
	}
	log.Debug("Loaded APs: ", apList.ApList)
	return apList.ApList
}

func getAp() string {
	aps := getApList()
	if aps == nil {
		log.Warn("Choosing fallback AP")
		return fallbackAp
	}
	randApIndex := rand.Int31n(int32(len(aps)))
	return aps[randApIndex]
}

func (channel *Channel) Connect() {
	var err error
	log.Info("Starting connect, resolving AP")
	channel.apAddress = getAp()
	log.Info("Using AP: ", channel.apAddress)
	if channel.connection, err = net.Dial("tcp", channel.apAddress); err != nil {
		log.Error("Unable to connect to AP: ", err)
		return
	}
	log.Info("Connected to ", channel.apAddress)
	channel.handshake()
}

func (channel *Channel) handshake() {
	helloBytes := channel.helloMessage()
	channel.connection.Write(helloBytes)
	lengthBytes := make([]byte, 4)
	read, _ := channel.connection.Read(lengthBytes)
	log.Debugf("Read %d/%d bytes", read, 4)
	length := binary.BigEndian.Uint32(lengthBytes) - 4
	dataBytes := make([]byte, length)
	read, _ = channel.connection.Read(dataBytes)
	log.Debugf("Read %d/%d bytes", read, length)
	channel.processAPResponseMessage(dataBytes)
}

func (channel *Channel) helloMessage() []byte {
	platform := keyexchange.Platform_PLATFORM_WIN32_X86
	product := keyexchange.Product_PRODUCT_PARTNER
	version := uint64(0x10700000000)
	knownKeys := uint32(1)
	nonce := make([]byte, 0x10)
	rand.Read(nonce)
	hello := &keyexchange.ClientHello{
		BuildInfo: &keyexchange.BuildInfo{
			Product: &product,
			Platform: &platform,
			Version: &version,
		},
		CryptosuitesSupported: []keyexchange.Cryptosuite{keyexchange.Cryptosuite_CRYPTO_SUITE_SHANNON},
		LoginCryptoHello: &keyexchange.LoginCryptoHelloUnion{
			DiffieHellman: &keyexchange.LoginCryptoDiffieHellmanHello{
				Gc: channel.Config.PrivateKey.SlimPub().GX.Bytes(),
				ServerKeysKnown: &knownKeys,
			},
		},
		ClientNonce: nonce,
		Padding: []byte{0x1e},
	}
	message, err := proto.Marshal(hello)
	if err != nil {
		log.Error("Unable to Marshal ClientHello: ", err)
		return nil
	}
	length := uint32(6 + len(message))
	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, length)
	header := []byte{0, 4}
	buffer := bytes.Buffer{}
	buffer.Write(header)
	buffer.Write(lengthBytes)
	buffer.Write(message)
	return buffer.Bytes()
}

func (channel *Channel) processAPResponseMessage(responseBytes []byte) []byte {
	apResponse := keyexchange.APResponseMessage{}
	if err := proto.Unmarshal(responseBytes, &apResponse); err != nil {
		log.Error("Unable to unmarshal APChallenge: ", err.Error())
	}
	log.Debug("APChallenge: ", apResponse.Challenge.ServerNonce)
	return nil
}

func (channel *Channel) Disconnect() {
	channel.connection.Close()
}
