package jasmine

import (
	"fmt"
	"testing"
	
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertRencanaStudi(t *testing.T) {
	rencanastd : Rencanastudi{
		ID: primitive.NewObjectID(),
		id: "1",
		nama_matkul: "Kewirausahaan",
		kode_matkul: "TI43304",
		status: "Diterima",
		sks: "3",
		kelas: "2A",
	}
	insertedID := InsertRencanastudi("dhs", rencanastudi)
	if insertedID == nil {
		t.Errorf("Error inserting data into rencanastudi collection")
	}
}

func TestInsertNilai(t *testing.T) {
	nil := Nilai{
		ID: primitive.NewObjectID(),
		id: "1",
		nama_matkul: "Network Programming",
		kode_matkul: "TI42253",
		sks: "3",
		grade: "A",	
	}
	insertedID := InsertNilai("dhs", nilai)
	if insertedID == nil {
		t.Errorf("Error inserting data into nilai collection")
	}
}

func TestInsertMatakuliah(t *testing.T) {
	matkul := matakuliah{
		ID: primitive.NewObjectID(),
		id: "1",
		nama_matkul: "Network Programming",
		kode_matkul: "TI42253",
		nama_dosen: "M. Yusril Helmi",
		sks: "3",
		gambar: "images/network.png",	
	}
	insertedID := InsertMatakuliah("dhs", matakuliah)
	if insertedID == nil {
		t.Errorf("Error inserting data into matakuliah collection")
	}
}

func TestInsertTranskrip(t *testing.T) {
	Transkrip := transkrip{
		ID: primitive.NewObjectID(),
		id: "1",
		nama_matkul: "Algoritma dan Struktur Data I",
		kode_matkul: "TI41061",
		sks: "3",
		grade: "C",
	}
	insertedID := InsertTranskrip("dhs", transkrip)
	if insertedID == nil {
		t.Errorf("Error inserting data into Transkrip collection")
	}
}

func TestInsertUsers(t *testing.T) {
	usr := users{
		ID: primitive.NewObjectID(),
		id: "1",
		nama: "Jasmine Mutiara Bintang",
		npm: "1214012",
		program: "Reguler (REG)",
		program_studi: "D4 Teknik Informatika",
		tahun_akademik: "2021",	
		kelas: "2A",	
		dosen_wali: "Muhammad Yusril Helmi Setywan, S.Kom,. M.Kom",	
	}
	insertedID := InsertMatakuliah("dhs", users)
	if insertedID == nil {
		t.Errorf("Error inserting data into Users collection")
		
}