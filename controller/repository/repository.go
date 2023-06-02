package repository

import (
	"context"
	"fmt"
	acc "github.com/leekchan/accounting"
	"github.com/sidiq200/lapke-umkm/controller/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"math/big"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertPenjualan(db *mongo.Database, id int, namaproduk string, jumlahpenjualan int, tanggaldatamasuk string, cabang string) (InsertedID interface{}, err error) {
	var m model.Penjualan
	m.ID = id
	m.NamaProduk = namaproduk
	m.JumlahPenjualan = jumlahpenjualan
	m.TanggalDataMasuk = tanggaldatamasuk
	m.Cabang = cabang
	return InsertOneDoc(db, "penjualan", m), err
}

func GetPenjualanByNamaProduk(namaproduk string, db *mongo.Database) (data model.Penjualan, err error) {
	user := db.Collection("penjualan")
	filter := bson.M{"namaproduk": namaproduk}
	err = user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("GetDataPenjualanFormNamaProduk: %v\n", err)
	}
	return data, err
}
func InsertPengeluaran(db *mongo.Database, id int, nama string, jumlah int, tanggal string, cabang string) (InsertedID interface{}, err error) {
	var m model.Pengeluaran
	m.ID = id
	m.Namapengeluaran = nama
	m.Jumlah = jumlah
	m.Tanggal = tanggal
	m.Cabang = cabang
	return InsertOneDoc(db, "pengeluaran", m), err
}
func GetPengeluaranByNama(nama string, db *mongo.Database) (data model.Pengeluaran, err error) {
	user := db.Collection("pengeluaran")
	filter := bson.M{"namapengeluaran": nama}
	err = user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("GetDataPengeluaranFormNama: %v\n", err)
	}
	return data, err
}

func GetAllPengeluaran(cabang string, db *mongo.Database) (data []model.Pengeluaran, err error) {
	user := db.Collection("pengeluaran")
	filter := bson.D{{}}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetAllPengeluaran: %v\n", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetAllPenjualan(cabang string, db *mongo.Database) (data []model.Penjualan, err error) {
	user := db.Collection("penjualan")
	filter := bson.D{{}}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetAllPenjualan: %v\n", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func InsertRekap(db *mongo.Database, Pengeluaran []model.Pengeluaran, penjualan []model.Penjualan, jumlahbersih, jumlahkotor int) (InsertedID interface{}, err error) {
	var m model.Recap
	m.Pengeluaran = Pengeluaran
	m.Penjualan = penjualan
	m.JumlahBersih = jumlahbersih
	m.JumlahKotor = jumlahkotor
	return InsertOneDoc(db, "rekap", m), err
}

func FormatRupiah(amount float64) string {
	lc := acc.LocaleInfo["IDR"]
	ac := acc.Accounting{
		Symbol:    lc.ComSymbol,
		Precision: 2,
		Thousand:  lc.ThouSep,
		Decimal:   lc.DecSep,
	}
	return ac.FormatMoneyBigFloat(big.NewFloat(amount))
}

func InsertPemasukan(db *mongo.Database, id int, namapemasukan string, jumlahpemasukan int, tanggaldatamasuk string, cabang string) (InsertedID interface{}, err error) {
	var m model.Pemasukan
	m.ID = id
	m.NamaPemasukan = namapemasukan
	m.JumlahPemasukan = jumlahpemasukan
	m.TanggalDataMasuk = tanggaldatamasuk
	m.Cabang = cabang
	return InsertOneDoc(db, "pemasukan", m), err
}

func GetAllPemasukan(cabang string, db *mongo.Database) (data []model.Pemasukan, err error) {
	user := db.Collection("pemasukan")
	filter := bson.D{{}}
	cursor, err := user.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("GetAllPemasukan: %v\n", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func HitungTotalKeuangan(pemasukan []model.Pemasukan, pengeluaran []model.Pengeluaran) (total model.Total) {
	totalPemasukan := 0
	totalPengeluaran := 0

	// Menghitung total pemasukan
	for _, p := range pemasukan {
		totalPemasukan += p.JumlahPemasukan
	}

	// Menghitung total pengeluaran
	for _, peng := range pengeluaran {
		totalPengeluaran += peng.Jumlah
	}

	// Menghitung jumlah bersih
	total.JumlahBersih = totalPemasukan - totalPengeluaran
	total.TotalPemasukan = totalPemasukan
	total.TotalPengeluaran = totalPengeluaran

	return total
}
