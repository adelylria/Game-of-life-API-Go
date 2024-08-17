package server

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

// Función para descifrar la clave de sesión AES usando la clave privada RSA
func decryptSessionKey(privateKey *rsa.PrivateKey, encryptedSessionKey []byte) ([]byte, error) {
	// Descifrar la clave de sesión con RSA-OAEP
	sessionKey, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, encryptedSessionKey, nil)
	if err != nil {
		return nil, err
	}
	return sessionKey, nil
}

func decryptData(sessionKey, encryptedData []byte) ([]byte, error) {
	block, err := aes.NewCipher(sessionKey)

	if err != nil {
		return nil, err
	}
	// El vector de inicialización (IV) debe ser el mismo para cifrar y descifrar
	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCFBDecrypter(block, iv)

	// Descifrar los datos
	decryptedData := make([]byte, len(encryptedData))
	stream.XORKeyStream(decryptedData, encryptedData)

	return decryptedData, nil
}

// Función para cifrar los datos usando AES y el modo CFB
func encryptData(sessionKey, plainData []byte) ([]byte, error) {
	block, err := aes.NewCipher(sessionKey)
	if err != nil {
		return nil, err
	}

	// Crear un vector de inicialización (IV) para el cifrado
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)

	// Cifrar los datos
	encryptedData := make([]byte, len(plainData))
	stream.XORKeyStream(encryptedData, plainData)

	return encryptedData, nil
}
