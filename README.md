# Mini e-wallet using GO

- buat table dengan nama ewallet pada DBMS MySQL

- edit file .env untuk koneksi database dengan nilai variabel disesuaikan dengan konfigurasi MySQL dan juga PORT untuk aplikasi:

`USER=root `<br/>
`PASSWORD=mysql `<br/>
`HOST=localhost `<br/>
`DATABASE=ewallet `<br/>
`PORT=8080 `<br/>

- untuk menjalankan aplikasi, buka terminal dan arahkan ke direktori di mana file "main.go" berada, dan jalankan "go run main.go"

ENDPOINTS | METHOD | NAME
----------|--------|-----
`host:port/register`|POST|REGISTER
`host:port/login`|POST|LOGIN
`host:port/logout`|GET|LOGOUT
`host:port/wallet/create`|POST|CREATE WALLET
`host:port/wallet/all`|GET|LIST OF ALL WALLET OF A CERTAIN USER
`host:port/wallet/delete?walletid=1`|DELETE|DELETE WALLET BY ID
`host:port/wallet/addbalance?walletid=1&type=debit&code=008&newbalance=2000`|PUT|TOPUP WALLET
`host:port/wallet/transfer?fromwallet=1&towallet=2&type=debit&balance=2000`|POST|TRANSFER BALANCE
`host:port/bank/all`|GET|LIST OF ALL BANKS

- untuk endpoint REGISTER, input didapatkan dari FORM dengan data yang dibutuhkan adalah `username`, `email`, `password`
- untuk endpoint LOGIN, input didapatkan dari FORM dengan data yang dibutuhkan adalah `username`, `password`


