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

// GetDataPemasukan godoc
// @Summary Mengambil data pemasukan (single).
// @Description get data pemasukan.
// @Tags Lapke-umkm
// @Accept application/json
// @Produce json
// @Param namapemasukan path string true "Masukan namapemasukan"
// @Success 200 {object} model.Pemasukan{}
// @Router /lapuak/pemasukan/{namapemasukan} [get]
func (db *UMKHandler) GetDataPemasukan(c *fiber.Ctx) (err error) {
	namapenjualan := c.Params("namapemasukan")
	getdata, err := repository.GetPemasukanByNama(namapenjualan, config.DBMongo("lapke-umkm"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Data Tidak ada")
	}
	return json.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Data Pemasukan berhasil diambil",
		Data:    getdata,
	}.WriteToBody(c)
}

// GetDataPengeluaran godoc
// @Summary Mengambil data Pengeluaran (single).
// @Description get data pengeluaran.
// @Tags Lapke-umkm
// @Accept application/json
// @Produce json
// @Param namapengeluaranpath string true "Masukan namapengeluaran
// @Success 200 {object} model.Pengeluaran{}
// @Router /lapuak/pengeluaran/{namapengeluaran [get]
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

// InsertDataPemasukan godoc
// @Summary insert data pemasukan.
// @Description get data pemasukan.
// @Tags Lapke-umkm
// @Accept application/json
// @Param request body model.Pemasukan true "Payload Body [RAW]"
// @Produce json
// @Success 200 {object} model.Pemasukan
// @Router /lapuak/inspemasukan [post]
func (db *UMKHandler) InsertDataPemasukan(c *fiber.Ctx) (err error) {
	database := config.DBMongo("lapke-umkm")
	var pemasukan model.Pemasukan
	if err := c.BodyParser(&pemasukan); err != nil {
		return err
	}
	Inserted, err := repository.InsertPemasukan(database,
		pemasukan.ID,
		pemasukan.NamaPemasukan,
		pemasukan.JumlahPemasukan,
		pemasukan.TanggalDataMasukPem,
		pemasukan.Cabang,
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
	cabang := "surabaya"
	getdatapemasukan, err := repository.GetAllPemasukan(cabang, config.DBMongo("lapke-umkm"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Tidak Ada Data Pemasukan")
	}
	getdatapengeluaran, err := repository.GetAllPengeluaran(cabang, config.DBMongo("lapke-umkm"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Tidak Ada Data Pengeluaran")
	}
	getdatapenjualan, err := repository.GetAllPenjualan(cabang, config.DBMongo("lapke-umkm"))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Tidak Ada Data Penjualan")
	}

	// Menghitung total keuangan
	totalKeuangan := repository.HitungTotalKeuangan(getdatapemasukan,getdatapengeluaran)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Tidak Ada Data Total Keuangan")
	}

	jmlpenjualan := 0
	for i := range getdatapenjualan {
		jmlpenjualan += getdatapenjualan[i].JumlahPenjualan
	}

	jmlpemasukan := totalKeuangan.TotalPemasukan
	for i := range getdatapenjualan {
		jmlpemasukan += getdatapemasukan[i].JumlahPemasukan
	}
	
	jmlpengeluaran := totalKeuangan.TotalPengeluaran
	for i := range getdatapengeluaran {
		jmlpengeluaran += getdatapengeluaran[i].Jumlah
	}

	jumlahpenjualan := float64(jmlpenjualan)
	jumlahpemasukan := float64(jmlpemasukan)
	jumlahpengeluaran := float64(jmlpengeluaran)
	totalkeuangan := float64(jumlahpemasukan - jumlahpengeluaran - jumlahpenjualan)

	jmlpengeluaranrp := repository.FormatRupiah(jumlahpengeluaran)
	jmlpemasukanrp := repository.FormatRupiah(jumlahpemasukan)
	totalkeuanganrp := repository.FormatRupiah(totalkeuangan)
	jmlpenjualanrp := repository.FormatRupiah(jumlahpenjualan)

	data := model.RecapResponse{
		Penjualan:         getdatapenjualan,
		Pemasukan:         getdatapemasukan,
		Pengeluaran:       getdatapengeluaran,
		JumlahKotor:       jmlpenjualanrp,
		JumlahPemasukan:   jmlpemasukanrp,
		JumlahPengeluaran: jmlpengeluaranrp,
		JumlahBersih:      totalkeuanganrp,
	}

	_, err = repository.InsertRekap(config.DBMongo("lapke-umkm"),
		getdatapengeluaran,
		getdatapenjualan,
		jmlpenjualan,
		int(totalkeuangan),
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal menyimpan data rekap")
	}

	return json.ReturnData{
		Code:    200,
		Success: true,
		Status:  "Data Rekap Berhasil Disimpan!",
		Data:    data,
	}.WriteToBody(c)
}

// GetAllPemasukan godoc
// @Summary Get data Pemasukan.
// @Description get data Pemasukan.
// @Tags Lapke-umkm
// @Accept application/json
// @Produce json
// @Success 200 {object} model.Pemasukan
// @Router /lapumk/getpemasukan [get]
func (db *UMKHandler) GetAllPemasukan(c *fiber.Ctx) (err error) {
	cabang := "Surabaya"
	getdata, err := repository.GetAllPemasukan(cabang, config.DBMongo("lapke-umkm"))
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
// GetAllPemasukan godoc
// @Summary Get data Pemasukan.
// @Description get data Pemasukan.
// @Tags Lapke-umkm
// @Accept application/json
// @Produce json
// @Success 200 {object} model.Pengeluaran
// @Router /lapumk/getpemasukan [get]
func (db *UMKHandler) GetAllPengeluaran(c *fiber.Ctx) (err error) {
	cabang := "surabaya"
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
// @Router /lapumk/getpenjualan [get]
func (db *UMKHandler) GetAllPenjualan(c *fiber.Ctx) (err error) {
	cabang := "Surabaya"
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