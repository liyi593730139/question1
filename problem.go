package main

import (
	"math/big"
	"gaby"
)

type Part struct {
	Num big.Int
}

var mapa = make(map[string]Part)

func update(id string, changeFunc func(*Part) (bool)) {
	pn, ok := mapa[id] // 取出
	if !ok {
		return
	}
	ok = changeFunc(&pn) // 修改
	if !ok {
		return
	}
	mapa[id] = pn // 塞回去
}

func main() {

	mapa["1"] = Part{
		Num: *big.NewInt(1111),
	}

	update("1", func(part *Part) bool {
		part.Num = *big.NewInt(22222)
		return true
	})

	println(gaby.ToJson(mapa))
}
