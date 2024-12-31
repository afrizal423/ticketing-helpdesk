# Ticketing Helpdesk
Repo ini hanya prototype untuk program ticketing helpdesk.
Program ini *hanya prototype* dan belum sempurna.

## Requirement
- Golang
- Oracle
- Redis

### Dependency
- [Whatsmeow](https://github.com/tulir/whatsmeow) untuk Whatsapp
- [go-telegram](https://github.com/go-telegram/bot) untuk Telegram
## How to install
- Clone repository ini menggunakan perintah berikut: `git clone https://github.com/afrizal423/ticketing-helpdesk`
- Copy file template yaml
    ```
    cp config.yaml.example config.yaml
    ```
- Isi beberapa konfigurasi pada file `config.yaml.example`
- Jalankan 
    ```
    go mod tidy
    ```
- Start server 
    ```
    go run cmd/main.go
    ```
- Scan qrcode untuk login WA