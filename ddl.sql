CREATE TABLE Pelanggan (
    ID_Pelanggan SERIAL PRIMARY KEY,
    Nama_Cust VARCHAR(255) NOT NULL,
    No_Hp VARCHAR(15) NOT NULL
);

CREATE TABLE Jenis_Layanan (
    ID_Layanan SERIAL PRIMARY KEY,
    Nama_Layanan VARCHAR(255) NOT NULL,
    Satuan VARCHAR(15) NOT NULL,
    Harga NUMERIC(10, 2) NOT NULL
);


CREATE TABLE Transaksi (
    ID_Transaksi SERIAL PRIMARY KEY,
    No_Pelanggan INT REFERENCES Pelanggan(ID_Pelanggan),
    Tanggal_Masuk DATE NOT NULL,
    Tanggal_Selesai DATE NOT NULL,
    Diterima_Oleh VARCHAR(255) NOT NULL
);


CREATE TABLE Detail_Transaksi (
    ID_Detail SERIAL PRIMARY KEY,
    ID_Transaksi INT REFERENCES Transaksi(ID_Transaksi),
    ID_Layanan INT REFERENCES Jenis_Layanan(ID_Layanan),
    Jumlah INT NOT NULL,
    Total_Harga NUMERIC(10, 2) NOT NULL
);

