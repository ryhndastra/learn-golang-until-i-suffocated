package main

import "fmt"

func main() {
	// if nama := "resa"; nama == "rey" {
	// 	fmt.Println("halo")
	// } else {
	// 	fmt.Println("bukan")
	// }

	// var age = 19
	// if age >= 20 {
	// 	fmt.Println("lu udah tua bego")
	// } else {
	// 	fmt.Println("yahaha bocil")
	// }

	// standar for-loop (3-Component Loop)
	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// gaya while loop (hanya kondisi)
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println(counter)
	// 	counter++
	// }

	// infinite loop
	// for {
	// // Loop selamanya sampai dipaksa berhenti dengan 'break'
	// if conditionMet {
	//     break
	// }

	// di go switch case gausah pake break, soalnya udah ada auto break bawaan go nya langsung
	// role := "user"
	// switch role {
	// case "admin":
	// 	fmt.Println("Full access")
	// case "editor", "author":
	// 	fmt.Println("Write Access")
	// default:
	// 	fmt.Println("Read Only")
	// }

	// Tagless Switch (Alternatif Pengganti if-else if Beruntun)
	// grade := 90
	// switch {
	// case grade >= 90:
	// 	fmt.Println("A")
	// case grade >= 75:
	// 	fmt.Println("B")
	// default:
	// 	fmt.Println("C")
	// }

	if statusCode := 404; statusCode == 200 {
		fmt.Println("status ok: data berhasil di get")
	} else if statusCode == 404 {
		fmt.Printf("Not Found %d\n", statusCode)
	} else {
		fmt.Printf("Server Error: kode %d\n", statusCode)
	}

	// fmt.Println(statusCode)

	for i := 1; i <= 3; i++ {
		fmt.Print(i)
	}

	fmt.Println()

	countdown := 3
	for countdown > 0 {
		fmt.Printf("Countdown: %d\n", countdown)
		countdown--
	}
	fmt.Println("MISI SELESAI")

	num := 0
	for {
		num++
		if num%2 == 0 {
			continue
		}

		if num == 5 {
			fmt.Printf("ODD : %d (Break!)\n", num)
			break
		}

		fmt.Printf("ODD : %d\n", num)
	}

	totalBelanja := 250000
	isMember := true

	fmt.Printf("Total Belanja: Rp %d | Status Member: %t\n", totalBelanja, isMember)

	switch {
	case isMember && totalBelanja >= 200000:
		fmt.Println("Selamat! Anda mendapatkan diskon 20%")
	case isMember && totalBelanja >= 100000:
		fmt.Println("Selamat! Anda mendapatkan diskon 10%")
	case totalBelanja >= 200000:
		fmt.Println("Selamat! Anda mendapatkan diskon 5%")
	default:
		fmt.Println("Tidak Ada Diskon")
	}

	myBalance := 50000
	topupAmount := 100000

	fmt.Printf("Saldo Awal : Rp %d\n", myBalance)
	fmt.Printf("Alamat Memori : %p\n", &myBalance)
	topUpEwallet(&myBalance, topupAmount)
	fmt.Printf("Top Up sebesar : %d\n", topupAmount)
	fmt.Printf("Saldo Akhir : %d\n", myBalance)

}

func topUpEwallet(balance *int, amount int) {
	*balance += amount
}
