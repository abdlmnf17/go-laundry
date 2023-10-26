

INSERT INTO Pelanggan (Nama_Cust, No_Hp)
VALUES ('Jessica', '0812654987');

INSERT INTO Jenis_Layanan (Nama_Layanan, Satuan, Harga)
VALUES
    ('Cuci + Setrika', 'KG', 7000.00),
    ('Laundry Bedcover', 'Buah', 50000.00),
    ('Laundry Boneka', 'Buah', 25000.00);

INSERT INTO Transaksi (No_Pelanggan, Tanggal_Masuk, Tanggal_Selesai, Diterima_Oleh)
VALUES (1, '2022-08-18', '2022-08-20', 'Mirna');

INSERT INTO Detail_Transaksi (ID_Transaksi, ID_Layanan, Jumlah, Total_Harga)
VALUES (1, 1, 5, 35000.00);

INSERT INTO Detail_Transaksi (ID_Transaksi, ID_Layanan, Jumlah, Total_Harga)
VALUES (1, 2, 1, 50000.00);

INSERT INTO Detail_Transaksi (ID_Transaksi, ID_Layanan, Jumlah, Total_Harga)
VALUES (1, 3, 2, 50000.00);
