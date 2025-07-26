package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	fmt.Print("Nama: ")
	nama := bacaInput()

	fmt.Print("Umur: ")
	umur, err := strconv.Atoi(bacaInput())

	if err != nil {
		log.WithError(err).Error("Input umur salah")
		fmt.Println(">> Error: Umur harus angka! (example : 18,19,20)")
		return
	}

	if umur >= 18 {
		log.Info("Sesuai umur")
		fmt.Printf(">> Selamat datang, %s!\n", nama)
		return
	}

	if umur < 18 {
		err := fmt.Errorf("umur %d belum boleh", umur)
		log.WithError(err).Warning("Akses ditolak")
		fmt.Println(">> Error: Umur tidak valid (minimal 18 tahun)")
		return
	}

}

func bacaInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

