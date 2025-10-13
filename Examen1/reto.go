package main

import f "fmt"
import m "math"

func fa(n float64) float64{
	if(n<2){
		return n
	}else{
		return n*fa(n-1)
	}
}

func na(n, k float64) float64{
	return ( fa(n)*fa(n) ) / ( n*fa(k)*fa(k-1)*fa(n-k)*fa(n-k+1) )
}

func fi(n float64) float64{
	if(n<2){
		return n
	}else{
		return fi(n-1)+fi(n-2)
	}
}

func w(n float64) float64{
	return fi(m.Floor(m.Log2(na(n+1,n-1)))+1)
}

func main() {
    var numero float64

    f.Print("Número: ")
    _, err := f.Scan(&numero)
    if err != nil {
        f.Println("Error", err)
        return
    }

    f.Println("Wadefoc:", w(numero))
}
