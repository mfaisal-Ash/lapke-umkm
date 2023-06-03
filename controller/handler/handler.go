package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidiq200/lapke-umkm/config"
	"github.com/sidiq200/lapke-umkm/controller/helper/json"
	"github.com/sidiq200/lapke-umkm/controller/model"
	"github.com/sidiq200/lapke-umkm/controller/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type UMKHandler struct {
	Mongo *mongo.Database
}

// GetDataPengeluaran godoc
// @Summary Mengambil data Pengeluaran (single).
// @Description get data pengeluaran.
// @Tags Lapke-umkm
// @Accept application/json
// @Produce json
// @Param namapengeluaran path string true "Masukan namapengeluaran"
// @Success 200 {object} model.Pengeluaran{}
// @Router /lapuak/pengeluaran/{namapengeluaran} [get]
func (db *UMKHandler) GetDataPengeluaran(c *fiber.Ctx) (err error) {
	namapenjualan := c.Params("namapengeluaran")
	getdata, err := repository.GetPengeluaranByNama(namapenjualan, config.DBMongo("lapke-umkm"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Data Tidak ada")
	}
	return json.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Data Pengeluaran berhasil diambil",
		Data:    getdata,
	}.WriteToBody(c)
}

// GetDataPenjualan godoc
// @Summary Mengambil data Penjualan (single).
// @Description get data Penjualan.
// @Tags Lapke-umkm
// @Accept application/json
// @Produce json
// @Param NamaProduk path string true "Masukan namaproduk"
// @Success 200 {object} model.Penjualan{}
// @Router /lapuak/penjualan/{NamaProduk} [get]
func (db *UMKHandler) GetDataPenjualan(c *fiber.Ctx) (err error) {
	namaproduk := c.Params("NamaProduk")
	getdata, err := repository.GetPenjualanByNamaProduk(namaproduk, config.DBMongo("lapke-umkm"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Data tidak Ditemukan")
	}
	return json.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Data Penjualan berhasil diambil",
		Data:    getdata,
	}.WriteToBody(c)
}

// InsertDataPenjualan godoc
// @Summary insert data penjualan.
// @Description get data penjualan.
// @Tags Lapke-umkm
// @Accept application/json
// @Param request body model.Penjualan true "Payload Body [RAW]"
// @Produce json
// @Success 200 {object} model.Penjualan
// @Router /lapuak/inspenjualan [post]
func (db *UMKHandler) InsertDataPenjualan(c *fiber.Ctx) (err error) {
	database := config.DBMongo("lapke-umkm")
	var penjualan model.Penjualan
	if err := c.BodyParser(&penjualan); err != nil {
		return err
	}
	Inserted, err := repository.InsertPenjualan(database,
		penjualan.ID,
		penjualan.NamaProduk,
		penjualan.JumlahPenjualan,
		penjualan.TanggalDataMasuk,
		penjualan.Cabang,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return json.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Data Berhasil Disimpan",
		Data:    Inserted,
	}.WriteToBody(c)
}

// InsPengeluaran godoc
// @Summary insert data Pengeluaran.
// @Description get data Pengeluaran.
// @Tags Lapke-umkm
// @Accept application/json
// @Param request body model.Pengeluaran true "Payload Body [RAW]"
// @Produce json
// @Success 200 {object} model.Pengeluaran
// @Router /lapuak/inspengeluaran [post]
func (db *UMKHandler) InsPengeluaran(c *fiber.Ctx) (err error) {
	database := config.DBMongo("lapke-umkm")
	var pengeluaran model.Pengeluaran
	if err := c.BodyParser(&pengeluaran); err != nil {
		return err
	}
	Inserted, err := repository.InsertPengeluaran(database,
		pengeluaran.ID,
		pengeluaran.Namapengeluaran,
		pengeluaran.Jumlah,
		pengeluaran.Tanggal,
		pengeluaran.Cabang,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	return json.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Data Berhasil Disimpan",
		Data:    Inserted,
	}.WriteToBody(c)
}

// KalkulasiLaporan godoc
// @Summary Kalkulasi Jumlah Laporan Keuangan.
// @Description Get Data Jumlah.
// @Tags Lapke-umkm
// @Accept application/json
// @Produce json
// @Success 200 {object} model.Recap
// @Router /lapuak/getlaporan [get]
func (db *UMKHandler) KalkulasiLaporan(c *fiber.Ctx) (err error) {
	cabang := "pekanbaru"
	getdatapengeluaran, err := repository.GetAllPengeluaran(cabang, config.DBMongo("lapke-umkm"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Tidak Ada Data Pengeluaran")
	}
	getdatapenjualan, err := repository.GetAllPenjualan(cabang, config.DBMongo("lapke-umkm"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Tidak Ada Data Penjualan")
	}
	jmlpenjualan := 0
	for i, _ := range getdatapenjualan {
		jmlpenjualan += getdatapenjualan[i].JumlahPenjualan
	}
	jmlpengeluaran := 0
	for z, _ := range getdatapengeluaran {
		jmlpengeluaran += getdatapengeluaran[z].Jumlah
	}
	jumlahpenjualan := float64(jmlpenjualan)
	jumlahpengeluaran := float64(jmlpengeluaran)
	jumlahakhir := float64(jmlpenjualan - jmlpengeluaran)

	jmlpengeluaranrp := repository.FormatRupiah(jumlahpengeluaran)
	jmlpemasukanrp := repository.FormatRupiah(jumlahpemasukan)
	jmlakhirrupiah := repository.FormatRupiah(jumlahakhir)
	jmlakhirrupiah := repository.FormatRupiah(jumlahakhir)
	jmlpenjualanrp := repository.FormatRupiah(jumlahpenjualan)

	data := model.RecapResponse{
		Penjualan:         getdatapenjualan,
		Pemasukan:       getdatapemasukan,
		Pengeluaran:       getdatapengeluaran,
		JumlahKotor:       jmlpenjualanrup,
		JumlahPengeluaran: jmlpengeluaranrup,
		JumlahBersih:      jmlakhirrupiah,
	}

	_, err = repository.InsertRekap(config.DBMongo("lapke-umkm"),
		getdatapengeluaran,
		getdatapenjualan,
		jmlpenjualan,
		int(jumlahakhir),
	)

	return json.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Data Rekap Berhasil Disimpan!",
		Data:    data,
	}.WriteToBody(c)
}

// GetAllPengeluaran godoc
// @Summary Get data Pengeluaran.
// @Description get data Pengeluaran.
// @Tags Lapke-umkm
// @Accept application/json
// @Produce json
// @Success 200 {object} model.Pengeluaran
// @Router /lapuak/getpengeluaran [get]
func (db *UMKHandler) GetAllPengeluaran(c *fiber.Ctx) (err error) {
	cabang := "Surabaya"
	getdata, err := repository.GetAllPengeluaran(cabang, config.DBMongo("lapke-umkm"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Data tidak ada")
	}
	return json.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Data Berhasil diambil",
		Data:    getdata,
	}.WriteToBody(c)
}

// GetAllPenjualan godoc
// @Summary Get data Penjualan.
// @Description get data Penjualan.
// @Tags Lapke-umkm
// @Accept application/json
// @Produce json
// @Success 200 {object} model.Penjualan
// @Router /lapuak/getpenjualan [get]
func (db *UMKHandler) GetAllPenjualan(c *fiber.Ctx) (err error) {
	cabang := "surabaya"
	getdata, err := repository.GetAllPenjualan(cabang, config.DBMongo("lapke-umkm"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Data tidak ada")
	}
	return json.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Data Berhasil diambil",
		Data:    getdata,
	}.WriteToBody(c)
}
