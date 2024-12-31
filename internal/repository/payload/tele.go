package payload

type SimpanDataEmployee struct {
	Userid    int64
	Username  string
	FirstName string
	LastName  string
	Chat_id   int64
}

type ListTiketAktif struct {
	NoTiket string
	Nowa    string
}

type GrabTiketAktif struct {
	NoTiket string
	Nowa    string
	Judul   string
	Isi     string
}

type TeleInsertChat struct {
	NoTiket string
	Dari    string
	Pesan   string
	Attch   string
	Kepada  string
}
