# Mini e-wallet using GO

- buat table dengan nama `ewallet` pada DBMS MySQL

- edit file .env untuk koneksi database dengan nilai variabel disesuaikan dengan konfigurasi MySQL dan juga PORT untuk aplikasi:

`USER=root `<br/>
`PASSWORD=mysql `<br/>
`HOST=localhost `<br/>
`DATABASE=ewallet `<br/>
`PORT=8080 `<br/>

- pastikan version dari golang minimal v1.11

- untuk menjalankan aplikasi, buka terminal dan arahkan ke direktori di mana file `main.go` berada, dan jalankan `go run main.go`

ENDPOINTS | METHOD | NAME |KETERANGAN
----------|--------|------|------
`host:port/register`|POST|REGISTER|input didapatkan dari FORM dengan data yang dibutuhkan adalah `username`, `email`, `password`
`host:port/login`|POST|LOGIN|input didapatkan dari FORM dengan data yang dibutuhkan adalah `username`, `password`
`host:port/logout`|GET|LOGOUT|
`host:port/wallet/create`|POST|CREATE NEW WALLET OF LOGGED IN USER|
`host:port/wallet/all`|GET|LIST OF ALL WALLET OF LOGGED IN USER|
`host:port/wallet/delete?walletid=1`|DELETE|DELETE WALLET BY ID|`walletid`=id dari wallet
`host:port/wallet/addbalance?walletid=1&type=debit&code=008&newbalance=2000`|PUT|TOPUP WALLET|`walletid`=id dari wallet<br/>`type`=jenis transaksi(debit/credit)<br/>`code`=kode bank (ex. 008)<br/>`newbalance`=jumlah topup<br/>
`host:port/wallet/transfer?fromwallet=1&towallet=2&type=debit&balance=2000`|POST|TRANSFER BALANCE|`fromwallet`=id dari wallet yang mentransfer<br/>`towallet=id dari wallet yang menerima transfer`<br/>`type`=jenis transaksi(debit/credit)`balance`=jumlah transfer
`host:port/bank/all`|GET|LIST OF ALL BANKS|



