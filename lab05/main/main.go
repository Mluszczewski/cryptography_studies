package main

import (
	"bytes"
	"fmt"
)

func fmtBits(data []byte) []byte {
	var buf bytes.Buffer
	for _, b := range data {
		fmt.Fprintf(&buf, "%04d", b)

	}
	return buf.Bytes()
}

func main() {
	data1 := fmtBits([]byte("dc0b0f1d649a5c7754f0444139a6baf9"))
	data2 := fmtBits([]byte("ca0d5de4ba2609ac9160749fa8d305f9"))

	// data1 := fmtBits([]byte("b43400f6db40f9d31d554f12e33133dc014fade8"))
	// data2 := fmtBits([]byte("d34737c78a0d87b4302d46f4493336c042db5a20"))

	// data1 := fmtBits([]byte("f3182a6f57d7be1102907286e4403b74b7878fb922183390c0953df3"))
	// data2 := fmtBits([]byte("9c3771fb34a5c633ebef897ae23044fbbdd2b5fff251a95f302ff586"))

	// data1 := fmtBits([]byte("5e13c59b0ecafd797bb5bf20ccc3c3cbb9dc4cafcf82c065fe7c5edd861f12b4"))
	// data2 := fmtBits([]byte("31502f0e7397a5d98b925477b3f7dd61258f049e0e057b9737cb798525b46b37"))

	// data1 := fmtBits([]byte("789a6dba8f360c4c6ad586823580ca068655fa7da2641d5a41c4a5ddb63a18b88a4fed1fce929958f6aa2fc44ea09526"))
	// data2 := fmtBits([]byte("895059b58b4b83e8ff516b819024661785a2261777385338fbcc17789905b5a3a96c75dcd2d254ec6d43ff60dc4e5e14"))

	// data1 := fmtBits([]byte("5d9598addb3c13d203e0473e540bc59b52242e2f607755c89c2d63d2504e4439758fea576b09a2323fbecb925a26b5954f078925f9eac0a43a4321dcdca2f6df"))
	// data2 := fmtBits([]byte("127577010f8f54e47d63dbd6939ec0d1f3dddd83fde25885a8b1e63b251b60a8a7b759ae99457bee08bff791006eca4c069edee09b1c58c6f09998fc8f55100a"))

	var counter int
	var i int

	for i, _ = range data1 {
		i++
	}
	for b, _ := range data1 {
		dupa := data1[b] - data2[b]
		if dupa != 0 {
			counter++
		}
	}
	fmt.Println(i)
	fmt.Println(counter)
	percentage := float32((float32(counter) / float32(i)) * 100)
	fmt.Println("Precentage: ", percentage, "%")
	fmt.Println("Liczba rozniacych sie bitow:", counter, "z", i, "procentowo:", percentage, "%")
}
