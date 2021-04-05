package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	resp := []byte("\x00" + "nasirmolla11shuvo@gmail.com" + "\x00" + "36259044")
	fmt.Println(string(resp), resp)

	sEnc := base64.StdEncoding.EncodeToString([]byte(resp))
	fmt.Println(sEnc)
}
