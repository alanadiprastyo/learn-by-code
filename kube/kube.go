package kube

import "errors"

type Kubernetes struct {
	Deployment string
	Pod        string
}

func Name() string {
	return "Hallo Kubernetes"
}

func Umur(byear int) (int, error) {
	nowYear := 2022
	age := nowYear - byear

	return age, errors.New("ada error pada hitung umur")
}

func Makanan() (name string, harga int, err error) {
	name = "nasi goreng"
	harga = 12000
	err = errors.New("tukang masaknya lagi pulang kampung")

	return
}
