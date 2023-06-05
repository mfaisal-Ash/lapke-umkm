package url

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sidiq200/lapke-umkm/config"
	"github.com/sidiq200/lapke-umkm/controller/handler"
)

func SetuplapRoutes(router fiber.Router) {

	Mongo := config.DBMongo("lapke-umkm")

	UMK := handler.UMKHandler{Mongo}

	router.Get("/penjualan/:namaproduk", UMK.GetDataPenjualan)
	router.Get("/pemasukan/:namapemasukan", UMK.GetDataPemasukan)
	router.Get("/pengeluaran/:namapengeluaran", UMK.GetDataPengeluaran)
	router.Post("/inspenjualan", UMK.InsertDataPenjualan)
	router.Post("/inspengeluaran", UMK.InsPengeluaran)
	router.Get("/getpengeluaran", UMK.GetAllPengeluaran)
	router.Get("/getpenjualan", UMK.GetAllPenjualan)
	router.Get("/getpemasukan", UMK.GetAllPemasukan)
	router.Get("/getlaporan", UMK.KalkulasiLaporan)
	

}
