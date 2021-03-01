package main
import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "crypto/aes"
    "encoding/hex"

    "github.com/gorilla/mux"
  )

type Event struct {
    ID          string `json:"ID"`
    Name        string `json:"Name"`
}

func EncryptAES(key []byte, plaintext string) string {

    c, err := aes.NewCipher(key)
    CheckError(err)

    out := make([]byte, len(plaintext))

    c.Encrypt(out, []byte(plaintext))

    return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) string {
    ciphertext, _ := hex.DecodeString(ct)

    c, err := aes.NewCipher(key)
    CheckError(err)

    pt := make([]byte, len(ciphertext))
    c.Decrypt(pt, ciphertext)

    s := string(pt)
    fmt.Println("DECRYPTED:", s)

  return s

}

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func homeLink(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome home!")
}

func encryption(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  event := Event{ID: "1", Name: "IBM"}
  byteArray, err := json.Marshal(event)
  if err != nil {
    fmt.Println(err)
   }
  key := "thisis32bitlongpassphraseimusing"
  c := EncryptAES([]byte(key), string(byteArray))
  fmt.Fprintf(w, c)
  fmt.Println(c)


}

func decryption(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  event := Event{ID: "1", Name: "IBM"}
  byteArray, err := json.Marshal(event)
  if err != nil {
    fmt.Println(err)
   }
  key := "thisis32bitlongpassphraseimusing"
  c := EncryptAES([]byte(key), string(byteArray))
  d := DecryptAES([]byte(key), c)
  fmt.Fprintf(w, d)
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", homeLink)
    router.HandleFunc("/api/encrypt", encryption).Methods("GET")
    router.HandleFunc("/api/decrypt", decryption).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router))
}
