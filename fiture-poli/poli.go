package fiture-poli

type Poli struct{
	Id_Poli int `gorm:primaryKey json:"id_poli"`
	Nama_Poli string `json:"nama_poli"`
}