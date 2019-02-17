package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var dsha256Cmd = &cobra.Command{
	Use:   "dsha256",
	Short: "",
	Long:  "",
	Run:   hashFunc,
}

func hashFunc(c *cobra.Command, inp []string) {
	i, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	h := sha256.New()
	h.Write(i)
	hash1 := h.Sum(nil)

	h = sha256.New()
	h.Write(hash1)
	hash2 := h.Sum(nil)

	fmt.Printf("%64x\n", hash2)
}

func main() {
	dsha256Cmd.Execute()
}
