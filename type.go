package jasmine

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RencanaStudi struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_matkul string             `bson:"nama_matkul,omitempty" json:"nama_matkul,omitempty"`
	Kode_matkul int                `bson:"kode_matkul,omitempty" json:"kode_matkul,omitempty"`
	Status      string             `bson:"status,omitempty" json:"status,omitempty"`
	Sks         int                `bson:"sks,omitempty" json:"sks,omitempty"`
	Kelas       string             `bson:"kelas,omitempty" json:"kelas,omitempty"`
}
type Nilai struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_matkul string             `bson:"nama_matkul,omitempty" json:"nama_matkul,omitempty"`
	Kode_matkul string             `bson:"kode_matkul,omitempty" json:"kode_matkul,omitempty"`
	Sks         string             `bson:"sks,omitempty" json:"sks,omitempty"`
	Grade       string             `bson:"grade,omitempty" json:"grade,omitempty"`
}

type MataKuliah struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_matkul string             `bson:"nama_matkul,omitempty" json:"nama_matkul,omitempty"`
	Kode_matkul string             `bson:"kode_matkul,omitempty" json:"kode_matkul,omitempty"`
	Nama_dosen  string             `bson:"nama_dosen,omitempty" json:"nama_dosen,omitempty"`
	Sks         string             `bson:"sks,omitempty" json:"sks,omitempty"`
	Gambar      string             `bson:"gambar,omitempty" json:"gambar,omitempty"`
}
type Transkrip struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_matkul string             `bson:"nama_matkul,omitempty" json:"nama_matkul,omitempty"`
	Kode_matkul string             `bson:"kode_matkul,omitempty" json:"kode_matkul,omitempty"`
	Sks         string             `bson:"sks,omitempty" json:"sks,omitempty"`
	Grade       string             `bson:"grade,omitempty" json:"grade,omitempty"`
}
type Users struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama          string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Npm           string             `bson:"npm,omitempty" json:"npm,omitempty"`
	Program       string             `bson:"program,omitempty" json:"program,omitempty"`
	Program_Studi string             `bson:"program_studi,omitempty" json:"program_studi,omitempty"`
	Kelas         string             `bson:"kelas,omitempty" json:"kelas,omitempty"`
	Dosen_Wali    string             `bson:"dosen_wali,omitempty" json:"dosen_wali,omitempty"`
}
