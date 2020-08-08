# Mini e-wallet using GO

- buat table dengan nama ewallet pada DBMS MySQL

- edit file .env untuk koneksi database dengan nilai variabel disesuaikan dengan konfigurasi MySQL dan juga PORT untuk aplikasi:

`USER=root `<br/>
`PASSWORD=mysql `<br/>
`HOST=localhost `<br/>
`DATABASE=ewallet `<br/>
`PORT=8080 `<br/>

- untuk menjalankan aplikasi, buka terminal dan arahkan ke direktori di mana file "main.go" berada, dan jalankan "go run main.go"


Daftar Endpoints :<br/>
`host:port/register` Method:POST REGISTER <br/>
`host:port/login` Method:POST LOGIN <br/>
`host:port/logout` Method:GET LOGOUT <br/>
`host:port/wallet/create` Method:POST CREATE WALLET <br/>
`host:port/wallet/all` METHOD:GET LIST OF ALL WALLET OF A CERTAIN USER <br/>
`host:port/wallet/delete?walletid=1` Method:DELETE DELETE WALLET BY ID <br/>
`host:port/wallet/addbalance?walletid=1&type=debit&code=008&newbalance=2000` Method:PUT TOPUP WALLET <br/>
`host:port/wallet/transfer?fromwallet=1&towallet=2&type=debit&balance=2000` Method:POST TRANSFER BALANCE <br/>
`host:port/bank/all` Method:GET LIST OF ALL BANKS <br/>



