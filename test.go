package main
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)
const (
	key string = "thisis32bitlongpassphraseimusing"
)
type Event struct {
	Value string `json:"Value"`
}

/*func DecryptAES(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)
	c, err := aes.NewCipher(key)
	CheckError(err)
	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)
	s := string(pt)
	fmt.Println("DECRYPTED:", s)
	return s
} */
func Encrypt(plaintext []byte, key []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Decrypt(ciphertext []byte, key []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("malformed ciphertext")
	}

	return gcm.Open(nil,
		ciphertext[:gcm.NonceSize()],
		ciphertext[gcm.NonceSize():],
		nil,
	)
}



func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
func encryption(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO: error handling
		fmt.Printf("Error in encryption: %v", err)
	}
	evt := &Event{}
	err = json.Unmarshal(body, evt)
	if err != nil {
		// todo: error handling
		fmt.Printf("Error in encryption: %v", err)
	}
	fmt.Printf("Got message %v", evt)
	c, err := Encrypt([]byte(evt.Value), []byte(key))
	if err != nil {
		fmt.Printf("error encyrpting: %v", err)
		return
	} else {
		evt.Value = string(c)
		fmt.Printf("encyrpting: %v", evt)
	}
	if data, err := json.Marshal(evt); err != nil {
		fmt.Printf("error encyrpting: %v", err)
	} else {
		w.Write(data)
		fmt.Printf("Encrypted string: %v", data)
	}
}
func decryption(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	/*
		event := Event{Value: "IBM"}
		byteArray, err := json.Marshal(event)
		if err != nil {
			fmt.Println(err)
		}
		c := EncryptAES([]byte(key), string(byteArray))
		d := DecryptAES([]byte(key), c)
		fmt.Fprintf(w, d)
	*/
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api/encrypt", encryption).Methods("GET")
	router.HandleFunc("/api/decrypt", decryption).Methods("GET")
	fmt.Println("listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}
