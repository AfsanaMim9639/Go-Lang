package main

import (
	"crypto"
    "crypto/aes"
    //"crypto/cipher"
    "crypto/hmac"
    "crypto/md5"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "crypto/sha512"
    //"crypto/x509"
    //"encoding/hex"
    "fmt"
    "log"
)

func main() {
    // -----------------------------
    // 1️⃣ SHA256 & MD5 Hashing
    data := []byte("Hello Crypto World!")
    sha256Hash := sha256.Sum256(data)
    md5Hash := md5.Sum(data)
    fmt.Printf("SHA256: %x\n", sha256Hash)
    fmt.Printf("MD5: %x\n\n", md5Hash)

    // -----------------------------
    // 2️⃣ HMAC with SHA256
    key := []byte("mysecretkey")
    h := hmac.New(sha256.New, key)
    h.Write(data)
    hmacSignature := h.Sum(nil)
    fmt.Printf("HMAC SHA256: %x\n\n", hmacSignature)

    // -----------------------------
    // 3️⃣ AES Encryption/Decryption
    aesKey := make([]byte, 32) // 256-bit key
    _, err := rand.Read(aesKey)
    if err != nil {
        log.Fatal(err)
    }

    block, err := aes.NewCipher(aesKey)
    if err != nil {
        log.Fatal(err)
    }

    plaintext := []byte("This is a secret message")
    ciphertext := make([]byte, len(plaintext))
    block.Encrypt(ciphertext, plaintext)
    fmt.Printf("AES Encrypted: %x\n", ciphertext)

    decrypted := make([]byte, len(ciphertext))
    block.Decrypt(decrypted, ciphertext)
    fmt.Printf("AES Decrypted: %s\n\n", decrypted)

    // -----------------------------
    // 4️⃣ RSA Key Pair Generation & Signing
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        log.Fatal(err)
    }
    publicKey := &privateKey.PublicKey

    message := []byte("RSA signing example")
    hashed := sha512.Sum512(message)

    signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA512, hashed[:])
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("RSA Signature: %x\n", signature)

    // Verify RSA signature
    err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA512, hashed[:], signature)
    if err != nil {
        fmt.Println("RSA verification failed")
    } else {
        fmt.Println("RSA verification succeeded\n")
    }

    // -----------------------------
    // 5️⃣ Secure Random Number
    randomBytes := make([]byte, 16)
    _, err = rand.Read(randomBytes)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Secure Random Bytes: %x\n", randomBytes)
}
