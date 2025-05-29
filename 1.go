Program Sudah Di Tambah:

package main

import "fmt"

type CryptoAsset struct {
	Nama       string
	Algoritma  string
	Difficulty int // tingkat kesulitan
	Reward     int // estimasi reward per mining
}

var assets []CryptoAsset
var totalHasilMining int

func tambahAsset() {
	var nama, algo string
	var diff, reward int
	fmt.Print("Nama Aset: ")
	fmt.Scan(&nama)
	fmt.Print("Algoritma: ")
	fmt.Scan(&algo)
	fmt.Print("Tingkat Kesulitan (1-100): ")
	fmt.Scan(&diff)
	fmt.Print("Reward (per mining): ")
	fmt.Scan(&reward)

	assets = append(assets, CryptoAsset{nama, algo, diff, reward})
	fmt.Println("Aset berhasil ditambahkan.")
}

func ubahAsset() {
	var nama string
	fmt.Print("Masukkan nama aset yang ingin diubah: ")
	fmt.Scan(&nama)
	for i, asset := range assets {
		if asset.Nama == nama {
			fmt.Print("Algoritma baru: ")
			fmt.Scan(&assets[i].Algoritma)
			fmt.Print("Tingkat Kesulitan baru: ")
			fmt.Scan(&assets[i].Difficulty)
			fmt.Print("Reward baru: ")
			fmt.Scan(&assets[i].Reward)
			fmt.Println("Aset berhasil diubah.")
			return
		}
	}
	fmt.Println("Aset tidak ditemukan.")
}

func hapusAsset() {
	var nama string
	fmt.Print("Masukkan nama aset yang ingin dihapus: ")
	fmt.Scan(&nama)
	for i, asset := range assets {
		if asset.Nama == nama {
			assets = append(assets[:i], assets[i+1:]...)
			fmt.Println("Aset berhasil dihapus.")
			return
		}
	}
	fmt.Println("Aset tidak ditemukan.")
}

func hitungEstimasi() {
	var dayaKomputasi int
	fmt.Print("Masukkan daya komputasi Anda (1-100): ")
	fmt.Scan(&dayaKomputasi)

	for _, asset := range assets {
		waktu := asset.Difficulty * 100 / dayaKomputasi
		fmt.Printf("Aset: %s | Algoritma: %s | Estimasi waktu: %d detik | Reward: %d\n", asset.Nama, asset.Algoritma, waktu, asset.Reward)
		totalHasilMining += asset.Reward
	}
}

func sequentialSearch(nama string) int {
	for i, asset := range assets {
		if asset.Nama == nama {
			return i
		}
	}
	return -1
}

func binarySearch(nama string) int {
	insertionSortByNama()
	low, high := 0, len(assets)-1
	for low <= high {
		mid := (low + high) / 2
		if assets[mid].Nama == nama {
			return mid
		} else if assets[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func selectionSortByDifficulty() {
	n := len(assets)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if assets[j].Difficulty < assets[min].Difficulty {
				min = j
			}
		}
		assets[i], assets[min] = assets[min], assets[i]
	}
}

func insertionSortByReward() {
	for i := 1; i < len(assets); i++ {
		temp := assets[i]
		j := i - 1
		for j >= 0 && assets[j].Reward > temp.Reward {
			assets[j+1] = assets[j]
			j--
		}
		assets[j+1] = temp
	}
}

func insertionSortByNama() {
	for i := 1; i < len(assets); i++ {
		temp := assets[i]
		j := i - 1
		for j >= 0 && assets[j].Nama > temp.Nama {
			assets[j+1] = assets[j]
			j--
		}
		assets[j+1] = temp
	}
}

func tampilkanLaporan() {
	fmt.Println("Laporan Hasil Mining Total:")
	fmt.Printf("Total reward: %d\n", totalHasilMining)
}

func tampilkanAset() {
	fmt.Println("Daftar Aset Kripto:")
	for _, asset := range assets {
		fmt.Printf("- %s | Algo: %s | Difficulty: %d | Reward: %d\n", asset.Nama, asset.Algoritma, asset.Difficulty, asset.Reward)
	}
}

func main() {
	assets = []CryptoAsset{
		{"Bitcoin", "SHA-256", 90, 6},
		{"Ethereum", "Ethash", 70, 2},
		{"Litecoin", "Scrypt", 50, 25},
		{"Monero", "CryptoNight", 65, 4},
		{"Dogecoin", "Scrypt", 40, 10},
		{"Dash", "X11", 55, 3},
		{"Zcash", "Equihash", 75, 10},
		{"Cardano", "Ouroboros", 30, 5},
		{"Polkadot", "Nominated Proof-of-Stake", 35, 8},
		{"Ripple", "XRP Ledger", 25, 1},
	}
	var pilihan int
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Aset")
		fmt.Println("2. Ubah Aset")
		fmt.Println("3. Hapus Aset")
		fmt.Println("4. Hitung Estimasi Mining")
		fmt.Println("5. Cari Aset (Sequential)")
		fmt.Println("6. Cari Aset (Binary)")
		fmt.Println("7. Urutkan Aset (Difficulty - Selection Sort)")
		fmt.Println("8. Urutkan Aset (Reward - Insertion Sort)")
		fmt.Println("9. Tampilkan Laporan")
		fmt.Println("10. Tampilkan Semua Aset")
		fmt.Println("0. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahAsset()
		case 2:
			ubahAsset()
		case 3:
			hapusAsset()
		case 4:
			hitungEstimasi()
		case 5:
			var nama string
			fmt.Print("Masukkan nama aset: ")
			fmt.Scan(&nama)
			idx := sequentialSearch(nama)
			if idx != -1 {
				fmt.Println("Aset ditemukan:", assets[idx])
			} else {
				fmt.Println("Aset tidak ditemukan.")
			}
		case 6:
			var nama string
			fmt.Print("Masukkan nama aset: ")
			fmt.Scan(&nama)
			idx := binarySearch(nama)
			if idx != -1 {
				fmt.Println("Aset ditemukan:", assets[idx])
			} else {
				fmt.Println("Aset tidak ditemukan.")
			}
		case 7:
			selectionSortByDifficulty()
			fmt.Println("Aset diurutkan berdasarkan tingkat kesulitan.")
			tampilkanAset()
		case 8:
			insertionSortByReward()
			fmt.Println("Aset diurutkan berdasarkan reward.")
			tampilkanAset()
		case 9:
			tampilkanLaporan()
		case 10:
			tampilkanAset()
		case 0:
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
