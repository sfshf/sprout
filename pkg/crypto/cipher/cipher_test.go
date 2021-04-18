package cipher

import (
	"fmt"
)

func ExampleAESCBC() {
	ciphertext, _ := AESCBCEncrypt([]byte(`one plain text`), []byte(`one secret key`))
	plaintext, _ := AESCBCDecrypt(ciphertext, []byte(`one secret key`))
	fmt.Printf("%s\n", plaintext)
	// Output:
	// one plain text
}

func ExampleAESCBCString() {
	ciphertext, _ := AESCBCEncryptString(`one plain text`, `one secret key`)
	plaintext, _ := AESCBCDecryptString(ciphertext, `one secret key`)
	fmt.Println(plaintext)
	// Output:
	// one plain text
}

func ExampleAESCBCStringIgnoreError() {
	ciphertext := AESCBCEncryptStringIgnoreError(`one plain text`, `one secret key`)
	plaintext := AESCBCDecryptStringIgnoreError(ciphertext, `one secret key`)
	fmt.Println(plaintext)
	// Output:
	// one plain text
}

func ExampleAESCFB() {
	ciphertext, _ := AESCFBEncrypt([]byte(`one plain text`), []byte(`one secret key`))
	plaintext, _ := AESCFBDecrypt(ciphertext, []byte(`one secret key`))
	fmt.Printf("%s\n", plaintext)
	// Output:
	// one plain text
}

func ExampleAESCFBString() {
	ciphertext, _ := AESCFBEncryptString(`one plain text`, `one secret key`)
	plaintext, _ := AESCFBDecryptString(ciphertext, `one secret key`)
	fmt.Println(plaintext)
	// Output:
	// one plain text
}

func ExampleAESCFBStringIgnoreError() {
	ciphertext := AESCFBEncryptStringIgnoreError(`one plain text`, `one secret key`)
	plaintext := AESCFBDecryptStringIgnoreError(ciphertext, `one secret key`)
	fmt.Println(plaintext)
	// Output:
	// one plain text
}

func ExampleAESCTR() {
	ciphertext, _ := AESCTREncrypt([]byte(`one plain text`), []byte(`one secret key`))
	plaintext, _ := AESCTRDecrypt(ciphertext, []byte(`one secret key`))
	fmt.Printf("%s\n", plaintext)
	// Output:
	// one plain text
}

func ExampleAESCTRString() {
	ciphertext, _ := AESCTREncryptString(`one plain text`, `one secret key`)
	plaintext, _ := AESCTRDecryptString(ciphertext, `one secret key`)
	fmt.Println(plaintext)
	// Output:
	// one plain text
}

func ExampleAESCTRStringIgnoreError() {
	ciphertext := AESCTREncryptStringIgnoreError(`one plain text`, `one secret key`)
	plaintext := AESCTRDecryptStringIgnoreError(ciphertext, `one secret key`)
	fmt.Println(plaintext)
	//Output:
	//one plain text
}

func ExampleAESOFB() {
	ciphertext, _ := AESOFBEncrypt([]byte(`one plain text`), []byte(`one secret key`))
	plaintext, _ := AESOFBDecrypt(ciphertext, []byte(`one secret key`))
	fmt.Printf("%s\n", plaintext)
	// Output:
	// one plain text
}

func ExampleAESOFBString() {
	ciphertext, _ := AESOFBEncryptString(`one plain text`, `one secret key`)
	plaintext, _ := AESOFBDecryptString(ciphertext, `one secret key`)
	fmt.Println(plaintext)
	// Output:
	// one plain text
}

func ExampleAESOFBStringIgnoreError() {
	ciphertext := AESOFBEncryptStringIgnoreError(`one plain text`, `one secret key`)
	plaintext := AESOFBDecryptStringIgnoreError(ciphertext, `one secret key`)
	fmt.Println(plaintext)
	//Output:
	//one plain text
}
