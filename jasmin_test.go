package jasmine

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "dhs",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertRencanastudi(t *testing.T) {
	nama_matkul := "Kewirausahaan"
	kode_matkul := "TI43304"
	status := "Diterima"
	sks := "3"
	kelas := "2A"
	hasil := InsertRencanastudi(nama_matkul, kode_matkul, status, sks, kelas)
	fmt.Println(hasil)
}

func TestInsertNilaidhs(t *testing.T) {
	nama_matkul := "Network Programming"
	kode_matkul := "TI42253"
	sks := "3"
	grade := "A"
	hasil := InsertNilai(nama_matkul, kode_matkul, sks, grade)
	fmt.Println(hasil)
}
func TestInsertMataKuliah(t *testing.T) {
	nama_matkul := "Network Programming"
	kode_matkul := "TI42253"
	nama_dosen := "M. Yusril Helmi"
	sks := "3"
	gambar := "images/network.png"
	hasil := InsertMataKuliah(nama_matkul, kode_matkul, nama_dosen, sks, gambar)
	fmt.Println(hasil)
}
func TestInsertTranskrip(t *testing.T) {
	nama_matkul := "Algoritma dan Struktur Data I"
	kode_matkul := "TI41061"
	sks := "3"
	grade := "C"
	hasil := InsertTranskrp(nama_matkul, kode_matkul, sks, grade)
	fmt.Println(hasil)
}
func TestInsertUserdhs(t *testing.T) {
	nama := "Jasmine Mutiara Bintang"
	npm := "1214012"
	program := "Reguler (REG)"
	program_studi := "D4 Teknik Informatika"
	tahun_akademik := "2021"
	kelas := "2A"
	dosen_wali := "Muhammad Yusril Helmi Setywan, S.Kom,. M.Kom"
	hasil := InsertUserdhs(nama, npm, program, program_studi, tahun_akademik, kelas, dosen_wali)
	fmt.Println(hasil)
}
