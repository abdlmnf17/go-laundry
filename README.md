# Aplikasi Laundry oleh Abdul Manap



## Instalasi 
1.   Git clone:  git clone https://github.com/abdlmnf17/go-laundry.git

2.   Buat database di PG Admin dengan nama "db-laundry", lalu copy paste dan jalankan kueri tool untuk kode SQL di file DDL dan DML 
3.   Atur password database dan lain lain di file main.go lalu buka terminal dan jalankan.

## Penggunaan
1.   Ada 3 menu utama, yaitu Menu Master Pelanggan, Master Layanan, dan Transaksi
2.   Ketiga menu saling berhubungan satu sama lain
3.   Harap diingat jika akan melakukan aksi, harus hafal ID dari setiap ID Pelanggan, Layanan dan Transaksi
4.   Cara melihat ID nya adalah tinggal memilih menu Lihat Data dari masing-masing menu
5.   Contoh: Menu Pelanggan -> Lihat Data Pelanggan -> akan muncul list pelanggan beserta ID nya

6.   Menu Pelanggan bisa melihat daftar pelanggan laundry, bisa menambah, memperbarui, dan menghapus. Karena saling berhubungan, jika menghapus pelanggan, maka daftar transaksi yang dilakukan oleh pelanggan itu akan IKUT TERHAPUS.

7. Menu Layanan berisi daftar layanan jasa laundry, bisa menambah, memperbarui, maupun menghapus.

8. Menu Transaksi berisi daftar transaksi, tambah, dan mengubah, lalu ada bagian untuk melihat detail transaksi. Dalam menu "detail transaksi", berisi list layanan, jumlah, total harga, nama layanan, dan nama pelanggan. Sedangkan di dalam menu "lihat data transaksi" hanya ada informasi tanggal masuk dan keluar lalu orang yang menerima laundryan.

9. Untuk menghapus data-data transaksi hanya bisa dilakukan di menu pelanggan. Dengan menghapus pelanggan sesuai ID, maka transaksinya akan ikut terhapus. 

10. Disetiap masing-masing menu juga sudah ada jump navigation, agar lebih cepat dalam memilih menu maupun keluar manu.  


## Akhir Kata
Mohon maaf jika masih banyak kekurangan dan hasil kodinganya belum di optimasi sebaik mungkin. Akhir kata, Terima Kasih.
