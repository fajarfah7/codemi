package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Loker struct {
	No          int
	Tipe        string
	NoIdentitas string
}

var dataLoker = []*Loker{}
var panjangLoker = 0

func main() {
	loker := Loker{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Silahkan mengisi jumlah loker dengan cara: \ninit (jumlah)")

	for {
		text, _ := reader.ReadString('\n')

		data := strings.Fields(text)

		data[0] = strings.ToLower(data[0])

		if data[0] == "init" {
			panjang, _ := strconv.Atoi(data[1])
			panjangLoker = panjang
			fmt.Println("Berhasil membuat loker dengan jumlah", panjang)
		} else if data[0] == "input" {
			if panjangLoker == 0 {
				fmt.Println("Silahkan mengisi jumlah loker dengan cara: \ninit (jumlah)")
			} else {
				if len(data) != 3 {
					fmt.Println("Inputan Tidak Valid !")
				} else {
					loker.Input(data[1], data[2])
				}
			}
		} else if data[0] == "status" {
			if panjangLoker == 0 {
				fmt.Println("Silahkan mengisi jumlah loker dengan cara: \ninit (jumlah)")
			} else {
				if len(data) > 1 {
					fmt.Println("Inputan Tidak Valid !")
				} else {
					Status()
				}
			}
		} else if (data[0]) == "leave" {
			if panjangLoker == 0 {
				fmt.Println("Silahkan mengisi jumlah loker dengan cara: \ninit (jumlah)")
			} else {
				if len(data) != 2 {
					fmt.Println("Inputan Tidak Valid !")
				} else {
					nomor, _ := strconv.Atoi(data[1])
					nomorLoker := nomor
					if nomorLoker == 0 {
						fmt.Println("Mohon masukan angka diatas 0 dan tidak boleh huruf")
					} else {
						Leave(nomorLoker)
					}
				}
			}
		} else if (data[0]) == "search"{
			if panjangLoker == 0 {
				fmt.Println("Silahkan mengisi jumlah loker dengan cara: \ninit (jumlah)")
			} else {
				if len(data) != 2 {
					fmt.Println("Inputan Tidak Valid !")
				} else {
					Search(data[1])
				}
			}
		} else if (data[0]) == "find"{
			if panjangLoker == 0 {
				fmt.Println("Silahkan mengisi jumlah loker dengan cara: \ninit (jumlah)")
			} else {
				if len(data) != 2 {
					fmt.Println("Inputan Tidak Valid !")
				} else {
					Find(data[1])
				}
			}
		} else if(data[0]) == "exit" {
			break
		} else if (data[0]) == "help" {
			fmt.Println("Membuat loker: \t\t\t\t\t\tinit (jumlah loker)")
			fmt.Println("Memasukan data ke loker: \t\t\t\tinput (tipe) (nomor identitas)")
			fmt.Println("Mengosongkan loker: \t\t\t\t\tleave (nomor loker)")
			fmt.Println("Mencari kartu identitas: \t\t\t\tsearch (tipe identitas)")
			fmt.Println("Mencari lokasi berdasarkan nomor identitas: \t\tfind (nomor identitas)")
			fmt.Println("Keluar: \t\t\t\t\t\texit")
			fmt.Println("Bantuan: \t\t\t\t\t\thelp")
		} else {
			fmt.Println("Perintah tidak ditemukan !")
		}
	}
}

func (loker Loker)Input(tipe, noIdentitas string) {
	var no int
	var statusAppend bool

	if len(dataLoker) != 0 {
		for i := 0; i < len(dataLoker); i++ {
			if len(dataLoker) != panjangLoker {
				no = i + 2
				statusAppend = true
			} else if dataLoker[i].Tipe == "" && dataLoker[i].NoIdentitas == "" && len(dataLoker) <= panjangLoker {
				dataLoker[i].Tipe = tipe
				dataLoker[i].NoIdentitas = noIdentitas
				fmt.Println("Karto identitas dimasukan ke loker nomor", dataLoker[i].No)
				break
			} else if i+1 == panjangLoker {
				fmt.Println("Maaf loker sudah penuh")
			}

		} 
	} else {
		no = 1
		statusAppend = true
	}

	if statusAppend == true {
		loker.No = no
		loker.Tipe = tipe
		loker.NoIdentitas = noIdentitas

		dataLoker = append(dataLoker, &loker)
		fmt.Println("Karto identitas dimasukan ke loker nomor", no)
	}
}

func Status() {
	fmt.Println("Nomor", "\t\t", "Tipe", "\t\t", "Nomor Identitas")
	for i := 0; i < len(dataLoker); i++ {
		if dataLoker[i].Tipe == "" && dataLoker[i].NoIdentitas == ""{
			continue
		} else {
			fmt.Println(dataLoker[i].No, "\t\t", dataLoker[i].Tipe, "\t\t", dataLoker[i].NoIdentitas)
		}
	}
}

func Leave(no int) {
	for i := 0; i < len(dataLoker); i++ {
		if dataLoker[i].No == no {
			dataLoker[i].Tipe = ""
			dataLoker[i].NoIdentitas = ""
			fmt.Println("Loker nomor", no, "berhasil dikosongkan")
			break
		} else if i+1 == len(dataLoker){
			fmt.Println("Maaf loker nomor", no, "tidak ada")
		}
	}
}

func Find(noIdentitas string){
	var listData []string
	noIdentitas = strings.ToLower(noIdentitas)
	for i := 0; i < len(dataLoker); i++ {
		if strings.ToLower(dataLoker[i].NoIdentitas) == noIdentitas {
			no := strconv.Itoa(dataLoker[i].No)
			listData = append(listData, no)
		}
	}
	if len(listData) == 0 {
		fmt.Println("Kartu identitas tidak ditemukan")
	} else {
		fmt.Println("Kartu identitas tersebut ada di loker",strings.Join(listData, ","))
	}
}

func Search(tipeIdentitas string){
	var listData []string
	tipeIdentitas = strings.ToLower(tipeIdentitas)
	for i := 0; i < len(dataLoker); i++ {
		if strings.ToLower(dataLoker[i].Tipe) == tipeIdentitas {
			listData = append(listData, dataLoker[i].NoIdentitas)
		}
	}
	if len(listData) == 0 {
		fmt.Println("Nomor identitas tidak ditemukan")
	} else {
		fmt.Println(strings.Join(listData, ","))
	}
}
