package main

import "fmt"

const NMAX = 100

type Tenant struct {
	ID, Transaction int
	Name            string
	TotalIncome     float64
}

type TenantsArray [NMAX]Tenant

var Tenants TenantsArray
var tenantCount int

func main() {
	tenantCount = 0
	var choice int
	fmt.Println("--------------------------Hai, Selamat Datang!------------------------")
	fmt.Println("---------------------Aplikasi Manajemen Kantin TEL-U------------------")
	fmt.Println("--------------------------Login sebagai Admin-------------------------")
	fmt.Println()
	for choice != 7 {
		fmt.Println("1. Tambah Tenant")
		fmt.Println("2. Edit Tenant")
		fmt.Println("3. Hapus Tenant")
		fmt.Println("4. Rekam Transaksi")
		fmt.Println("5. Laba Admin")
		fmt.Println("6. Data Tenant")
		fmt.Println("7. Keluar")
		fmt.Println()
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			AddTenant(&Tenants, &tenantCount)
		case 2:
			EditTenant(&Tenants, tenantCount)
		case 3:
			DeleteTenant(&Tenants, &tenantCount)
		case 4:
			RecordTransaction(&Tenants, tenantCount)
		case 5:
			DisplayAdminProfit(&Tenants, tenantCount)
		case 6:
			DisplayTenantsData(&Tenants, tenantCount)
		case 7:
			fmt.Println("Selamat tinggal!")
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

func AddTenant(tenants *TenantsArray, tenantCount *int) {
    /*I.S. array tenants sebagian diisi dengan elemen tenantCount,tenantCount       menunjukkan jumlah penyewa saat ini.
      F.S Jika tenant ditambahkan, array tenant sekarang menyertakan tenant baru, dan jumlah penyewa bertambah 1, tetapi jika larik sudah penuh tenant dan tenantCount tetap tidak berubah. */
	if *tenantCount < NMAX {
		var newID int
		var newName string
		fmt.Print("Masukkan ID Tenant: ")
		fmt.Scan(&newID)
		fmt.Print("Masukkan Nama Tenant: ")
		fmt.Scan(&newName)

		tenants[*tenantCount] = Tenant{
			ID:          newID,
			Name:        newName,
			Transaction: 0,
			TotalIncome: 0.0,
		}

		*tenantCount++
		fmt.Println("Tenant Berhasil Ditambahkan.")
		fmt.Println()
	} else {
		fmt.Println("Kapasitas Tenant Telah Penuh.")
	}
}

func EditTenant(tenants *TenantsArray, tenantCount int) {
    /* I.S.array tenant diurutkan berdasarkan ID dalam urutan menaik dan berisi elemen tenantCount, tenantCount menunjukkan jumlah penyewa saat ini. 
      F.S. Jika tenant dengan ID yang diberikan ditemukan maka nama tenant yang sesuai diperbarui berdasarkan masukan pengguna. Pesan “Penyewa Telah Diperbarui.” dicetak. Namun, jika penyewa dengan ID yang diberikan tidak ditemukan pesan “Tenant Tidak Ditemukan.” dicetak. */
	var id int
	fmt.Print("Masukkan ID Tenant untuk Diedit: ")
	fmt.Scan(&id)
	var found int = -1
	var med int
	var kr int = 0
	var kn int = tenantCount - 1
	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if id < tenants[med].ID {
			kn = med - 1
		} else if id > tenants[med].ID {
			kr = med + 1
		} else {
			found = med
		}
	}

	if found != -1 {
		var newName string
		fmt.Print("Masukkan Nama Baru Tenant: ")
		fmt.Scan(&newName)
		tenants[found].Name = newName
		fmt.Println("Nama Tenant Telah Diperbarui.")
		fmt.Println()
	} else {
		fmt.Println("Tenant Tidak Ditemukan.")
		fmt.Println()
	}
}

func DeleteTenant(tenants *TenantsArray, tenantCount *int) {
    /* I.S. array tenants diurutkan berdasarkan ID dalam urutan menaik dan berisi elemen tenantCount, tenantCount menunjukkan jumlah penyewa saat ini.
      F.S. jika tenant dengan ID yang diberikan ditemukan dan dihapus, array tenants sekarang berisi satu penyewa lebih sedikit dan diurutkan berdasarkan ID dalam urutan menaik jumlah penyewa dikurangi dengan 1. Namun, Jika tenant dengan ID yang diberikan tidak ditemukan, maka array tenants dan tenantCount tetap tidak berubah */
	var id int
	fmt.Print("Masukkan ID Tenant untuk Dihapus: ")
	fmt.Scan(&id)

	foundIndex := -1
	for i := 0; i < *tenantCount; i++ {
		if tenants[i].ID == id {
			foundIndex = i
		}
	}

	if foundIndex != -1 {
		for j := foundIndex; j < *tenantCount-1; j++ {
			tenants[j] = tenants[j+1]
		}
		*tenantCount--
		fmt.Println("Tenant Telah Dihapus.")
		fmt.Println()

		for i := 1; i < *tenantCount; i++ {
			key := tenants[i]
			j := i - 1
			for j >= 0 && tenants[j].ID > key.ID {
				tenants[j+1] = tenants[j]
				j = j - 1
			}
			tenants[j+1] = key
		}
	} else {
		fmt.Println("Tenant Tidak Ditemukan.")
		fmt.Println()
	}
}

func RecordTransaction(tenants *TenantsArray, tenantCount int) {
    /* I.S. array tenants diurutkan berdasarkan ID dalam urutan menaik dan berisi     elemen tenantCount, tenantCount menunjukkan jumlah penyewa saat ini.
      F.S. Jika tenant dengan ID yang diberikan ditemukan maka jumlah transaksi   dan total pendapatan untuk penyewa tersebut diperbarui. Namun, jika    tenant dengan ID yang diberikan tidak ditemukan maka array tenant dan  jumlah tenant tetap tidak berubah.*/
	var id int
	var amount float64
	fmt.Print("Masukkan ID Tenant: ")
	fmt.Scan(&id)

	foundIndex := -1
	for i := 0; i < tenantCount; i++ {
		if tenants[i].ID == id {
			foundIndex = i
		}
	}

	if foundIndex != -1 {
		fmt.Print("Masukkan Jumlah Transaksi: ")
		fmt.Scan(&amount)
		tenants[foundIndex].Transaction++
		tenants[foundIndex].TotalIncome += amount
		fmt.Println("Transaksi berhasil dicatat.")
		fmt.Println()
	} else {
		fmt.Println("Tenant Tidak Ditemukan.")
		fmt.Println()
	}
}

func DisplayAdminProfit(tenants *TenantsArray, tenantCount int) {
    /* I.S. array tenants diurutkan berdasarkan ID dalam urutan menaik dan berisi     elemen tenantCount, tenantCount menunjukkan jumlah penyewa saat ini.
      F.S.  susunan tenants dan tenantCount tetap tidak berubah. */
	var totalProfit float64 = 0.0
	for i := 0; i < tenantCount; i++ {
		totalProfit += tenants[i].TotalIncome * 0.25
	}
	fmt.Printf("Total Admin Profit: Rp%.2f\n", totalProfit)
	fmt.Println()
}

func DisplayTenantsData(tenants *TenantsArray, tenantCount int) {
    /* I.S. array tenants diurutkan berdasarkan ID dalam urutan menaik dan berisi elemen tenantCount, tenantCount menunjukkan jumlah penyewa saat ini.
    F.S. array tenants tetap diurutkan berdasarkan transaksi dalam urutan menaik, dan tenantCount tetap tidak berubah. */
	for i := 0; i < tenantCount-1; i++ {
		minIndex := i
		for j := i + 1; j < tenantCount; j++ {
			if tenants[j].Transaction < tenants[minIndex].Transaction {
				minIndex = j
			}
		}
		if minIndex != i {
			tenants[i], tenants[minIndex] = tenants[minIndex], tenants[i]
		}
	}

	fmt.Println("Data Tenant (Diurutkan berdasarkan Transaksi):")
	for i := 0; i < tenantCount; i++ {
		fmt.Printf("ID: %d, Nama: %s, Transaksi: %d, Total Pendapatan: Rp%.2f\n",
			tenants[i].ID, tenants[i].Name, tenants[i].Transaction, tenants[i].TotalIncome)
			fmt.Println()
	}
}