package main

import (
    "bufio"
    "database/sql"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "22092002"
    dbname   = "db-laundry"
)

var db *sql.DB

func main() {
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    var err error
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Panic(err)
    }
    defer db.Close()


    for {
        fmt.Println("===== Aplikasi Laundry =====")
        fmt.Println("1. Tabel Master Pelanggan")
        fmt.Println("2. Tabel Master Jenis Layanan")
        fmt.Println("3. Tabel Transaksi")
        fmt.Println("4. Keluar")
        fmt.Print("Pilih menu: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            handleMasterPelanggan()
        case 2:
            handleMasterJenisLayanan()
        case 3:
            handleTransaksi()
        case 4:
            return
        default:
            fmt.Println("Menu tidak valid.")
        }
    }
}

type Pelanggan struct {
    Nama string
    NoHp string
}

type JenisLayanan struct {
    Nama   string
    Satuan string
    Harga  float64
}

type Transaksi struct {
    NoPelanggan   string
    TanggalMasuk  string
    TanggalSelesai string
    DiterimaOleh  string
}

type DetailTransaksi struct {
    ID           int
    IDLayanan    int
    NamaLayanan  string
    Jumlah       int
    TotalHarga   float64
    NamaPelanggan string
}

func handleMasterPelanggan() {
    for {
        fmt.Println("===== Tabel Master Pelanggan =====")
        fmt.Println("1. Lihat Data Pelanggan")
        fmt.Println("2. Tambah Data Pelanggan")
        fmt.Println("3. Ubah Data Pelanggan")
        fmt.Println("4. Hapus Data Pelanggan")
        fmt.Println("5. Kembali ke Menu Utama")
        fmt.Print("Pilih menu: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            viewMasterPelanggan()
        case 2:
            addMasterPelanggan()
        case 3:
            updateMasterPelanggan()
        case 4:
            deleteMasterPelanggan()
        case 5:
            return
        default:
            fmt.Println("Menu tidak valid.")
        }
    }
}

func handleMasterJenisLayanan() {
    for {
        fmt.Println("===== Tabel Master Jenis Layanan =====")
        fmt.Println("1. Lihat Data Jenis Layanan")
        fmt.Println("2. Tambah Data Jenis Layanan")
        fmt.Println("3. Ubah Data Jenis Layanan")
        fmt.Println("4. Hapus Data Jenis Layanan")
        fmt.Println("5. Kembali ke Menu Utama")
        fmt.Print("Pilih menu: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            viewMasterJenisLayanan()
        case 2:
            addMasterJenisLayanan()
        case 3:
            updateMasterJenisLayanan()
        case 4:
            deleteMasterJenisLayanan()
        case 5:
            return
        default:
            fmt.Println("Menu tidak valid.")
        }
    }
}


func viewMasterPelanggan() {
    fmt.Println("===== Lihat Data Pelanggan =====")

    // Query untuk menampilkan semua data pelanggan
    query := "SELECT ID_Pelanggan, Nama_Cust, No_Hp FROM Pelanggan"

    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Printf("%-5s %-20s %-15s\n", "ID", "Nama Pelanggan", "No HP")
    for rows.Next() {
        var id int
        var nama string
        var noHp string
        if err := rows.Scan(&id, &nama, &noHp); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%-5d %-20s %-15s\n", id, nama, noHp)
    }
}

func addMasterPelanggan() {
    fmt.Println("===== Tambah Data Pelanggan =====")

    var pelanggan Pelanggan
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Masukkan Nama Pelanggan: ")
    if scanner.Scan() {
        pelanggan.Nama = scanner.Text()
    }

    fmt.Print("Masukkan No HP: ")
    if scanner.Scan() {
        pelanggan.NoHp = scanner.Text()
    }

    pelanggan.Nama = strings.TrimSpace(pelanggan.Nama)
    pelanggan.NoHp = strings.TrimSpace(pelanggan.NoHp)

    if pelanggan.Nama == "" || pelanggan.NoHp == "" {
        fmt.Println("Nama dan No HP harus diisi.")
        return
    }

    // Query untuk menambahkan data pelanggan
    query := "INSERT INTO Pelanggan (Nama_Cust, No_Hp) VALUES ($1, $2)"

    _, err := db.Exec(query, pelanggan.Nama, pelanggan.NoHp)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Data pelanggan berhasil ditambahkan.")
}

func updateMasterPelanggan() {
    fmt.Println("===== Ubah Data Pelanggan =====")

    var id int
    var pelanggan Pelanggan
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Masukkan ID Pelanggan yang akan diubah: ")
    if scanner.Scan() {
        input := scanner.Text()
        var err error
        id, err = strconv.Atoi(input)
        if err != nil {
            log.Println("ID harus berupa angka.")
            return
        }
    }

    fmt.Print("Masukkan Nama Pelanggan baru: ")
    if scanner.Scan() {
        pelanggan.Nama = scanner.Text()
    }

    fmt.Print("Masukkan No HP baru: ")
    if scanner.Scan() {
        pelanggan.NoHp = scanner.Text()
    }

    pelanggan.Nama = strings.TrimSpace(pelanggan.Nama)
    pelanggan.NoHp = strings.TrimSpace(pelanggan.NoHp)

    if pelanggan.Nama == "" || pelanggan.NoHp == "" {
        fmt.Println("Nama dan No HP harus diisi.")
        return
    }

    // Query untuk mengubah data pelanggan
    query := "UPDATE Pelanggan SET Nama_Cust = $1, No_Hp = $2 WHERE ID_Pelanggan = $3"

    _, err := db.Exec(query, pelanggan.Nama, pelanggan.NoHp, id)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Data pelanggan berhasil diubah.")
}

func deleteMasterPelanggan() {
    fmt.Println("===== Hapus Data Pelanggan =====")

    var id int

    fmt.Print("Masukkan ID Pelanggan yang akan dihapus: ")
    fmt.Scanln(&id)

    // Hapus detail transaksi terlebih dahulu
    deleteDetailTransaksiByCustomerID(id)

    // Kemudian, hapus transaksi yang terkait
    deleteTransaksiByCustomerID(id)

    // Terakhir, hapus data pelanggan itu sendiri
    deleteCustomerByID(id)

    fmt.Println("Data pelanggan berhasil dihapus.")
}

func viewMasterJenisLayanan() {
    fmt.Println("===== Lihat Data Jenis Layanan =====")

    // Query untuk menampilkan semua data jenis layanan
    query := "SELECT ID_Layanan, Nama_Layanan, Satuan, Harga FROM Jenis_Layanan"

    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Printf("%-5s %-20s %-10s %-15s\n", "ID", "Nama Layanan", "Satuan", "Harga")
    for rows.Next() {
        var id int
        var jenisLayanan JenisLayanan
        if err := rows.Scan(&id, &jenisLayanan.Nama, &jenisLayanan.Satuan, &jenisLayanan.Harga); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%-5d %-20s %-10s %-15.2f\n", id, jenisLayanan.Nama, jenisLayanan.Satuan, jenisLayanan.Harga)
    }
}

func addMasterJenisLayanan() {
    fmt.Println("===== Tambah Data Jenis Layanan =====")

    var jenisLayanan JenisLayanan
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Masukkan Nama Layanan: ")
    if scanner.Scan() {
        jenisLayanan.Nama = scanner.Text()
    }

    fmt.Print("Masukkan Satuan: ")
    if scanner.Scan() {
        jenisLayanan.Satuan = scanner.Text()
    }

    fmt.Print("Masukkan Harga: ")
    if scanner.Scan() {
        input := scanner.Text()
        var err error
        jenisLayanan.Harga, err = strconv.ParseFloat(input, 64)
        if err != nil {
            log.Println("Harga harus berupa angka.")
            return
        }
    }

    jenisLayanan.Nama = strings.TrimSpace(jenisLayanan.Nama)
    jenisLayanan.Satuan = strings.TrimSpace(jenisLayanan.Satuan)

    if jenisLayanan.Nama == "" || jenisLayanan.Satuan == "" {
        fmt.Println("Nama Layanan dan Satuan harus diisi.")
        return
    }

    // Query untuk menambahkan data jenis layanan
    query := "INSERT INTO Jenis_Layanan (Nama_Layanan, Satuan, Harga) VALUES ($1, $2, $3)"

    _, err := db.Exec(query, jenisLayanan.Nama, jenisLayanan.Satuan, jenisLayanan.Harga)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Data jenis layanan berhasil ditambahkan.")
}

func updateMasterJenisLayanan() {
    fmt.Println("===== Ubah Data Jenis Layanan =====")

    var id int
    var jenisLayanan JenisLayanan
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Masukkan ID Jenis Layanan yang akan diubah: ")
    if scanner.Scan() {
        input := scanner.Text()
        var err error
        id, err = strconv.Atoi(input)
        if err != nil {
            log.Println("ID harus berupa angka.")
            return
        }
    }

    fmt.Print("Masukkan Nama Layanan baru: ")
    if scanner.Scan() {
        jenisLayanan.Nama = scanner.Text()
    }

    fmt.Print("Masukkan Satuan baru: ")
    if scanner.Scan() {
        jenisLayanan.Satuan = scanner.Text()
    }

    fmt.Print("Masukkan Harga baru: ")
    if scanner.Scan() {
        input := scanner.Text()
        var err error
        jenisLayanan.Harga, err = strconv.ParseFloat(input, 64)
        if err != nil {
            log.Println("Harga harus berupa angka.")
            return
        }
    }

    jenisLayanan.Nama = strings.TrimSpace(jenisLayanan.Nama)
    jenisLayanan.Satuan = strings.TrimSpace(jenisLayanan.Satuan)

    if jenisLayanan.Nama == "" || jenisLayanan.Satuan == "" {
        fmt.Println("Nama Layanan dan Satuan harus diisi.")
        return
    }

    // Query untuk mengubah data jenis layanan
    query := "UPDATE Jenis_Layanan SET Nama_Layanan = $1, Satuan = $2, Harga = $3 WHERE ID_Layanan = $4"

    _, err := db.Exec(query, jenisLayanan.Nama, jenisLayanan.Satuan, jenisLayanan.Harga, id)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Data jenis layanan berhasil diubah.")
}

func deleteMasterJenisLayanan() {
    fmt.Println("===== Hapus Data Jenis Layanan =====")

    var id int

    fmt.Print("Masukkan ID Jenis Layanan yang akan dihapus: ")
    fmt.Scanln(&id)

    // Hapus detail transaksi terlebih dahulu
    deleteDetailTransaksiByServiceID(id)

    // Kemudian, hapus data jenis layanan itu sendiri
    deleteServiceByID(id)

    fmt.Println("Data jenis layanan berhasil dihapus.")
}


func handleTransaksi() {
    for {
        fmt.Println("===== Tabel Transaksi =====")
        fmt.Println("1. Lihat Data Transaksi")
        fmt.Println("2. Tambah Data Transaksi")
        fmt.Println("3. Ubah Data Transaksi")
        fmt.Println("4. Lihat Detail Transaksi")
        fmt.Println("5. Kembali ke Menu Utama")
        fmt.Print("Pilih menu: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            viewTransaksi()
        case 2:
            addTransaksi()
        case 3:
            updateTransaksi()
        case 4:
            var transaksiID int
            fmt.Print("Masukkan ID Transaksi: ")
            fmt.Scanln(&transaksiID)
            viewDetailTransaksi(transaksiID)
        case 5:
            return
        default:
            fmt.Println("Menu tidak valid.")
        }
    }
}

func addTransaksi() {
    fmt.Println("===== Tambah Data Transaksi =====")

    var transaksi Transaksi
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Masukkan No Pelanggan: ")
    if scanner.Scan() {
        transaksi.NoPelanggan = scanner.Text()
    }

    fmt.Print("Masukkan Tanggal Masuk (YYYY-MM-DD): ")
    if scanner.Scan() {
        transaksi.TanggalMasuk = scanner.Text()
    }

    fmt.Print("Masukkan Tanggal Selesai (YYYY-MM-DD): ")
    if scanner.Scan() {
        transaksi.TanggalSelesai = scanner.Text()
    }

    fmt.Print("Masukkan Diterima Oleh: ")
    if scanner.Scan() {
        transaksi.DiterimaOleh = scanner.Text()
    }

    transaksi.NoPelanggan = strings.TrimSpace(transaksi.NoPelanggan)
    transaksi.TanggalMasuk = strings.TrimSpace(transaksi.TanggalMasuk)
    transaksi.TanggalSelesai = strings.TrimSpace(transaksi.TanggalSelesai)
    transaksi.DiterimaOleh = strings.TrimSpace(transaksi.DiterimaOleh)

    if transaksi.NoPelanggan == "" || transaksi.TanggalMasuk == "" || transaksi.TanggalSelesai == "" || transaksi.DiterimaOleh == "" {
        fmt.Println("Semua kolom harus diisi.")
        return
    }

    // Query untuk menambahkan data transaksi
    query := "INSERT INTO Transaksi (No_Pelanggan, Tanggal_Masuk, Tanggal_Selesai, Diterima_Oleh) VALUES ($1, $2, $3, $4) RETURNING ID_Transaksi"

    var transaksiID int
    err := db.QueryRow(query, transaksi.NoPelanggan, transaksi.TanggalMasuk, transaksi.TanggalSelesai, transaksi.DiterimaOleh).Scan(&transaksiID)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Data transaksi berhasil ditambahkan dengan ID:", transaksiID)

    // Memasukkan detail transaksi
    addDetailTransaksi(transaksiID)
}

func addDetailTransaksi(transaksiID int) {
    fmt.Println("===== Tambah Detail Transaksi =====")

    for {
        var detail DetailTransaksi
        scanner := bufio.NewScanner(os.Stdin)

        fmt.Print("Masukkan ID Layanan: ")
        if scanner.Scan() {
            input := scanner.Text()
            var err error
            detail.IDLayanan, err = strconv.Atoi(input)
            if err != nil {
                log.Println("ID Layanan harus berupa angka.")
                return
            }
        }

        fmt.Print("Masukkan Jumlah: ")
        if scanner.Scan() {
            input := scanner.Text()
            var err error
            detail.Jumlah, err = strconv.Atoi(input)
            if err != nil {
                log.Println("Jumlah harus berupa angka.")
                return
            }
        }

        detail.TotalHarga = getHargaLayanan(detail.IDLayanan) * float64(detail.Jumlah)

        // Query untuk menambahkan data detail transaksi
        query := "INSERT INTO Detail_Transaksi (ID_Transaksi, ID_Layanan, Jumlah, Total_Harga) VALUES ($1, $2, $3, $4)"
        _, err := db.Exec(query, transaksiID, detail.IDLayanan, detail.Jumlah, detail.TotalHarga)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println("Detail transaksi berhasil ditambahkan.")

        fmt.Print("Tambahkan detail transaksi lainnya? (y/n): ")
        if scanner.Scan() {
            response := scanner.Text()
            if strings.ToLower(response) != "y" {
                return
            }
        }
    }
}

func getHargaLayanan(idLayanan int) float64 {
    // Query untuk mendapatkan harga layanan berdasarkan ID
    query := "SELECT Harga FROM Jenis_Layanan WHERE ID_Layanan = $1"

    var harga float64
    err := db.QueryRow(query, idLayanan).Scan(&harga)
    if err != nil {
        log.Fatal(err)
    }

    return harga
}

func viewTransaksi() {
    fmt.Println("===== Lihat Data Transaksi =====")

    // Query untuk menampilkan semua data transaksi
    query := "SELECT ID_Transaksi, No_Pelanggan, Tanggal_Masuk, Tanggal_Selesai, Diterima_Oleh FROM Transaksi"

    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Printf("%-5s %-12s %-15s %-15s %-20s\n", "ID", "No Pelanggan", "Tgl Masuk", "Tgl Selesai", "Diterima Oleh")

    for rows.Next() {
        var id, noPelanggan int
        var transaksi Transaksi
        if err := rows.Scan(&id, &noPelanggan, &transaksi.TanggalMasuk, &transaksi.TanggalSelesai, &transaksi.DiterimaOleh); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%-5d %-12d %-15s %-15s  %-20s\n", id, noPelanggan, transaksi.TanggalMasuk, transaksi.TanggalSelesai, transaksi.DiterimaOleh)
    }
}

func updateTransaksi() {
    fmt.Println("===== Ubah Data Transaksi =====")

    var id int
    var transaksi Transaksi
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Masukkan ID Transaksi yang akan diubah: ")
    if scanner.Scan() {
        input := scanner.Text()
        var err error
        id, err = strconv.Atoi(input)
        if err != nil {
            log.Println("ID harus berupa angka.")
            return
        }
    }

    fmt.Print("Masukkan No Pelanggan baru: ")
    if scanner.Scan() {
        transaksi.NoPelanggan = scanner.Text()
    }

    fmt.Print("Masukkan Tanggal Masuk baru (YYYY-MM-DD): ")
    if scanner.Scan() {
        transaksi.TanggalMasuk = scanner.Text()
    }

    fmt.Print("Masukkan Tanggal Selesai baru (YYYY-MM-DD): ")
    if scanner.Scan() {
        transaksi.TanggalSelesai = scanner.Text()
    }

    fmt.Print("Masukkan Diterima Oleh baru: ")
    if scanner.Scan() {
        transaksi.DiterimaOleh = scanner.Text()
    }

    transaksi.NoPelanggan = strings.TrimSpace(transaksi.NoPelanggan)
    transaksi.TanggalMasuk = strings.TrimSpace(transaksi.TanggalMasuk)
    transaksi.TanggalSelesai = strings.TrimSpace(transaksi.TanggalSelesai)
    transaksi.DiterimaOleh = strings.TrimSpace(transaksi.DiterimaOleh)

    if transaksi.NoPelanggan == "" || transaksi.TanggalMasuk == "" || transaksi.TanggalSelesai == "" || transaksi.DiterimaOleh == "" {
        fmt.Println("Semua kolom harus diisi.")
        return
    }

    // Query untuk mengubah data transaksi
    query := "UPDATE Transaksi SET No_Pelanggan = $1, Tanggal_Masuk = $2, Tanggal_Selesai = $3, Diterima_Oleh = $4 WHERE ID_Transaksi = $5"

    _, err := db.Exec(query, transaksi.NoPelanggan, transaksi.TanggalMasuk, transaksi.TanggalSelesai, transaksi.DiterimaOleh, id)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Data transaksi berhasil diubah.")
}


func viewDetailTransaksi(transaksiID int) {
    fmt.Println("===== Lihat Detail Transaksi =====")

    // Query untuk menampilkan detail transaksi berdasarkan ID Transaksi
    query := "SELECT dt.ID_Detail, dt.ID_Layanan, jl.Nama_Layanan, dt.Jumlah, dt.Total_Harga, p.Nama_Cust FROM Detail_Transaksi dt JOIN Jenis_Layanan jl ON dt.ID_Layanan = jl.ID_Layanan JOIN Transaksi t ON dt.ID_Transaksi = t.ID_Transaksi JOIN Pelanggan p ON t.No_Pelanggan = p.ID_Pelanggan WHERE dt.ID_Transaksi = $1"

    rows, err := db.Query(query, transaksiID)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    fmt.Printf("%-5s %-12s %-30s %-15s %-15s %-20s\n", "ID", "ID Layanan", "Nama Layanan", "Jumlah", "Total Harga", "Nama Pelanggan")

    for rows.Next() {
        var id, idLayanan, jumlah int
        var detail DetailTransaksi
        if err := rows.Scan(&id, &idLayanan, &detail.NamaLayanan, &jumlah, &detail.TotalHarga, &detail.NamaPelanggan); err != nil {
            log.Fatal(err)
        }
        detail.ID = id
        detail.IDLayanan = idLayanan
        detail.Jumlah = jumlah
        fmt.Printf("%-5d %-12d %-30s %-15d %-15.2f %-20s\n", detail.ID, detail.IDLayanan, detail.NamaLayanan, detail.Jumlah, detail.TotalHarga, detail.NamaPelanggan)
    }
}


func deleteDetailTransaksiByCustomerID(customerID int) {
    query := "DELETE FROM Detail_Transaksi dt " +
             "WHERE dt.ID_Transaksi IN (SELECT t.ID_Transaksi FROM Transaksi t WHERE t.No_Pelanggan = $1)"
    
    _, err := db.Exec(query, customerID)
    if err != nil {
        log.Fatal(err)
    }
}


func deleteTransaksiByCustomerID(customerID int) {
    query := "DELETE FROM Transaksi WHERE No_Pelanggan = $1"
    
    _, err := db.Exec(query, customerID)
    if err != nil {
        log.Fatal(err)
    }
}

func deleteCustomerByID(customerID int) {
    query := "DELETE FROM Pelanggan WHERE ID_Pelanggan = $1"
    
    _, err := db.Exec(query, customerID)
    if err != nil {
        log.Fatal(err)
    }
}


func deleteDetailTransaksiByServiceID(serviceID int) {
    query := "DELETE FROM Detail_Transaksi dt " +
             "WHERE dt.ID_Layanan = $1"

    _, err := db.Exec(query, serviceID)
    if err != nil {
        log.Fatal(err)
    }
}


func deleteServiceByID(serviceID int) {
    query := "DELETE FROM Jenis_Layanan WHERE ID_Layanan = $1"

    _, err := db.Exec(query, serviceID)
    if err != nil {
        log.Fatal(err)
    }
}



