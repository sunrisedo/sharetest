package main

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}
