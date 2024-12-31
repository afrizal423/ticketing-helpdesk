package payload

type SimpanDataClient struct {
	Nowa   string
	Nama   string
	Lokasi string
}

type DataClient struct {
	Nama   string
	Lokasi string
}

type SimpanTiketClient struct {
	NoTiket string
	Nowa    string
	Judul   string
	Isi     string
}

type ListTiketHeader struct {
	NoTiket string
	Status  string
}
