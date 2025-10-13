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

func fi(n float64) float64{
	if(n<2){
		return n
	}else{
		return fi(n-1)+fi(n-2)
	}
}

func main() {
    var numero float64

    f.Print("Número: ")
    _, err := f.Scan(&numero)
    if err != nil {
        f.Println("Error", err)
        return
    }

	fakm1 := fa(numero-2)
	fak := fakm1*(numero-1)
	fan := fak*numero*(numero+1)

    f.Println("Wadefoc:", fi(m.Floor(m.Log2( ((fan*fan)/((numero+1)*fak*fakm1*12)) ))+1))
}
