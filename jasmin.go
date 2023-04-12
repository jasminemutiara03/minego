package jasmine

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}
func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertRencanastudi(nama_matkul string, kode_matkul string, status string, sks string, kelas string) (InsertedID interface{}) {
	var rencanastudi RencanaStudi
	rencanastudi.Nama_matkul = nama_matkul
	rencanastudi.Kode_matkul = kode_matkul
	rencanastudi.Status = status
	rencanastudi.Sks = sks
	rencanastudi.Kelas = kelas

	return InsertOneDoc("Kewirausahaan", "rencanastudi", rencanastudi)
}
func InsertNilai(nama_matkul string, kode_matkul string, sks string, grade string) (InsertedID interface{}) {
	var nilai Nilaidhs
	nilai.Nama_matkul = nama_matkul
	nilai.Kode_matkul = kode_matkul
	nilai.Sks = sks
	nilai.Grade = grade

	return InsertOneDoc("Network Programming", "nilai", nilai)
}
func InsertMataKuliah(nama_matkul string, kode_matkul string, nama_dosen string, sks string, gambar string) (InsertedID interface{}) {
	var matakuliah MataKuliah
	matakuliah.Nama_matkul = nama_matkul
	matakuliah.Kode_matkul = kode_matkul
	matakuliah.Nama_dosen = nama_dosen
	matakuliah.Sks = sks
	matakuliah.Gambar = gambar

	return InsertOneDoc("Network Programming", "matakuliah", matakuliah)
}
func InsertTranskrp(nama_matkul string, kode_matkul string, sks string, grade string) (InsertedID interface{}) {
	var transkrip Transkrip
	transkrip.Nama_matkul = nama_matkul
	transkrip.Kode_matkul = kode_matkul
	transkrip.Sks = sks
	transkrip.Grade = grade

	return InsertOneDoc("Algoritma dan Struktur Data I", "transkrip", transkrip)
}
func InsertUserdhs(nama string, npm string, program string, program_studi string, tahun_akademik string, kelas string, dosen_wali string) (InsertedID interface{}) {
	var userdhs Userdhs
	userdhs.Nama = nama
	userdhs.Npm = npm
	userdhs.Program = program
	userdhs.Program_Studi = program_studi
	userdhs.Tahun_Akademik = tahun_akademik
	userdhs.Kelas = kelas
	userdhs.Dosen_Wali = dosen_wali

	return InsertOneDoc("Jasmine Mutiara Bintang", "userdhs", userdhs)
}
