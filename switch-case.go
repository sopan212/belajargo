package main

import "fmt"

func main() {
	var point = 4
	switch point {
	case 8:
		fmt.Println("keren")
	case 7:
		fmt.Println("bagus")
	case 5:
		fmt.Println("cukup")
	default:
		fmt.Println("kamu harus berusaha lebih giat!")
	}
	//switch dengan banyak kondisi
	var target = 9
	switch target {
	case 10:
		fmt.Println("kamu keren bisa mencapai target sempurna")
	case 6, 7, 8:
		fmt.Println("kamu hebat!")
	default:
		fmt.Println("not bad!")
	}

	//kurung kurawal pada case dan default
	var targetBaru = 4
	switch targetBaru {
	case 10:
		fmt.Println("kamu keren bisa mencapai target sempurna")
	case 5, 6, 7, 8, 9:
		fmt.Println("kamu hebat!")
	default:
		{
			fmt.Println("not bad!")
			fmt.Println("you can do better!")
		}

	}
	fmt.Println("===========================================")
	//swictcase dengan gaya if else
	var targetBaruLagi = 6
	switch {
	case targetBaruLagi == 10:
		fmt.Println("kamu keren bisa mencapai target sempurna")
	case targetBaruLagi >= 5:
		fmt.Println("keren kamu bisa achive target!")
	case targetBaruLagi < 5:
		fmt.Println("kamu harus berusaha lebih giat lagi!")
	default:
		{
			fmt.Println("not bad!")
			fmt.Println("you can do better!")
		}

	}

	//keyword falltrough di switch case
	fmt.Println("<===========================================================>")
	var targetBaruLagiNih = 8
	switch {
	case targetBaruLagiNih == 10:
		fmt.Println("kamu keren bisa mencapai target sempurna")
	case targetBaruLagiNih >= 5:
		fmt.Println("keren kamu bisa achive target!")
		fallthrough
	case targetBaruLagiNih < 5:
		fmt.Println("kamu harus berusaha lebih giat lagi!")
	default:
		{
			fmt.Println("not bad!")
			fmt.Println("you can do better!")
		}

	}

	//pengkondisian bersarang
	fmt.Println("============================")

	var nilaiAkhir = 0

	if nilaiAkhir > 70 {
		switch nilaiAkhir {
		case 90:
			fmt.Println("perfect!!")
		default:
			fmt.Println("nice try")
		}
	} else {
		if nilaiAkhir == 50 {
			fmt.Println("not bad")
		} else if nilaiAkhir == 30 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it!")
			if nilaiAkhir == 0 {
				fmt.Println("try harder")
			}

		}

	}

}
