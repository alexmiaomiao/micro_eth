# micro_eth
最简单的挖矿示例，使用比特币的共识算法pow（简化版）。没有实现地址，交易等（交易仅为一个字符串）。

go run main.go

lzt@SAMSUNG:~/Documents/mygo/src/github.com/alexmiaomiao/micro_eth$ go run main.go 
Mining the block containing "Miner lzt win 50 BTC||jm Send 1 BTC to ld"
000002525847027b43ab7d79ea1b7e5d131557b0406aec9f8e26739ae8c6edc1

Mining the block containing "Miner lzt win 50 BTC||ld Send 2 BTC to jm||ld Send 3 BTC to lzt"
00000c36ff9e8249742e2544cae2ea60497733329e31b02fe4102f9d999ec494

Number: 0
Parenthash: 0000000000000000000000000000000000000000000000000000000000000000
Hash: 0000000000000000000000000000000000000000000000000000000000000000
Data: Genesis Block

Number: 1
Parenthash: 0000000000000000000000000000000000000000000000000000000000000000
Hash: 000002525847027b43ab7d79ea1b7e5d131557b0406aec9f8e26739ae8c6edc1
Data: Miner lzt win 50 BTC||jm Send 1 BTC to ld

Number: 2
Parenthash: 000002525847027b43ab7d79ea1b7e5d131557b0406aec9f8e26739ae8c6edc1
Hash: 00000c36ff9e8249742e2544cae2ea60497733329e31b02fe4102f9d999ec494
Data: Miner lzt win 50 BTC||ld Send 2 BTC to jm||ld Send 3 BTC to lzt
