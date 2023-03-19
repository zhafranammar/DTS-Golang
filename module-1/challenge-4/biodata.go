package main

import (
	"fmt"
	"os"
)

type Teman struct {
    Absen     int
    Nama      string
    Alamat    string
    Pekerjaan string
    Alasan    string
}

func getTeman(absen int) Teman {
    var teman Teman

    switch absen {
    case 1:
        teman = Teman{
            Absen:     1,
            Nama:      "John Doe",
            Alamat:    "Jl. Sudirman No. 123",
            Pekerjaan: "Software Engineer",
            Alasan:    "Ingin belajar bahasa pemrograman Golang untuk meningkatkan skill programming.",
        }
    case 2:
        teman = Teman{
            Absen:     2,
            Nama:      "Jane Doe",
            Alamat:    "Jl. Thamrin No. 456",
            Pekerjaan: "Data Scientist",
            Alasan:    "Ingin mempelajari bahasa pemrograman Golang untuk menganalisis data secara lebih efisien.",
        }
    default:
        teman = Teman{}
    }

    return teman
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Gunakan perintah: go run biodata.go <nomor_absen>")
        return
    }

    absen := os.Args[1]

    var absenInt int
    _, err := fmt.Sscanf(absen, "%d", &absenInt)
    if err != nil {
        fmt.Println("Nomor absen harus berupa angka")
        return
    }

    teman := getTeman(absenInt)

    if teman.Absen == 0 {
        fmt.Printf("Tidak ada teman dengan nomor absen %d\n", absenInt)
    } else {
        fmt.Printf("Nama: %s\n", teman.Nama)
        fmt.Printf("Alamat: %s\n", teman.Alamat)
        fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
        fmt.Printf("Alasan: %s\n", teman.Alasan)
    }
}