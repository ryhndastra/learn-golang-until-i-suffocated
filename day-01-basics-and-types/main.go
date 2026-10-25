package main

import "fmt"

const (
	RoleJunior = iota //iota buat bikin enum auto increment otomatis bawaan go
	RoleMid
	RoleSenior
)

func main() {
	name := "rey"
	var experienceYears int = 3
	dailyRate := 45.5

	currentRole := RoleJunior
	// age := 21
	// isMobileDev := true

	// fmt.Printf("Halo nama ku %s, dan umurku %d, Mobile dev %t.\n", name, age, isMobileDev)

	// var userLevel int
	// var accountBalance float64
	// var isVerified bool
	// var userBio string

	// fmt.Printf("Zero Values -> level : %v | Balance : %v | Verified : %v | Bio : %q \n", userLevel, accountBalance, isVerified, userBio)

	// itemPrice := 150000
	// taxRate := 0.11
	// taxAmount := float64(itemPrice) * taxRate

	// totalPrice := itemPrice + int(taxAmount)
	// fmt.Printf("Harga Barang : Rp %d\n", itemPrice)
	// fmt.Printf("Pajak (11%%) : Rp %.2f\n", taxAmount)
	// fmt.Printf("total Akhir : Rp %d",totalPrice)

	fmt.Println("========================================")
	fmt.Println("DEVELOPER STATS CARD")
	fmt.Println("========================================")

	fmt.Printf("Name : %s\n", name)
	fmt.Printf("Role Lever : %d\n", currentRole)
	fmt.Printf("Experience : %d\n", experienceYears)
	fmt.Printf("Hourly Rate : $%.2f\n", dailyRate)
	fmt.Printf("Status 90D  : Semangat Belajar Golang!\n")
	fmt.Println("========================================")

}
