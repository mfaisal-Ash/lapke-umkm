package model

type ReturnData struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

type Pemasukan struct {
	ID               int    `json:"id" bson:"id"`
	NamaPemasukan    string `json:"namapemasukan" bson:"namapemasukan"`
	JumlahPemasukan  int    `json:"jumlahpemasukan" bson:"jumlahpemasukan"`
	TanggalDataMasukPem string `json:"tanggaldatamasukPemasuk" bson:"tanggaldatamasukPemasuk"`
	Cabang           string `json:"cabang" bson:"cabang"`
}

type Penjualan struct {
	ID               int    `json:"id" bson:"id"`
	NamaProduk       string `json:"namaproduk" bson:"namaproduk"`
	JumlahPenjualan  int    `json:"jumlahpenjualan" bson:"jumlahpenjualan"`
	TanggalDataMasuk string `json:"tanggaldatamasuk" bson:"tanggaldatamasuk"`
	Cabang           string `json:"cabang" bson:"cabang"`
}

type Pengeluaran struct {
	ID              int    `json:"id" bson:"id"`
	Namapengeluaran string `json:"namapengeluaran" bson:"namapengeluaran"`
	Jumlah          int    `json:"jumlah" bson:"jumlah"`
	Tanggal         string `json:"tanggal" bson:"tanggal"`
	Cabang          string `json:"cabang" bson:"cabang"`
}

type User struct {
	ID       int    `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Recap struct {
	Penjualan    []Penjualan   `json:"penjualan" bson:"penjualan"`
	Pemasukan  []Pemasukan `json:"pemasukan" bson:"pemasukan"`
	Pengeluaran  []Pengeluaran `json:"pengeluaran" bson:"pengeluaran"`
	JumlahKotor  int           `json:"jumlahkotor" bson:"jumlahkotor"`
	JumlahBersih int           `json:"jumlahbersih" bson:"jumlahbersih"`
}

type RecapResponse struct {
	Penjualan         []Penjualan `json:"penjualan" bson:"penjualan"`
	Pengeluaran       []Pengeluaran `json:"pengeluaran" bson:"pengeluaran"`
	Pemasukan       []Pemasukan `json:"pemasukan" bson:"pemasukan"`
	JumlahKotor       int         `json:"jumlahkotor" bson:"jumlahkotor"`
	JumlahPemasukan int         `json:"jumlahPemasukan" bson:"jumlahPemasukan"`
	JumlahPengeluaran int         `json:"jumlahPengeluaran" bson:"jumlahPengeluaran"`
	JumlahBersih      int         `json:"jumlahbersih" bson:"jumlahbersih"`
}

type Total struct {
	TotalPemasukan  int `json:"totalpemasukan"`
	TotalPengeluaran int `json:"totalpengeluaran"`
	JumlahBersih    int `json:"jumlahbersih"`
}

type Keuangan struct {
	Pemasukan   []Pemasukan   `json:"pemasukan" bson:"pemasukan"`
	Pengeluaran []Pengeluaran `json:"pengeluaran" bson:"pengeluaran"`
	Total       Total         `json:"total" bson:"total"`
}
