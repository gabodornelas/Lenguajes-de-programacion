package main
import f "fmt"
import m "math" // importa Log2 y Floor
func fi(n float64) float64{// funcion de fibonacci
	if(n<2){
		return n
	}else{
		return fi(n-1)+fi(n-2)
	}
}
func main() {
    var n float64
    _, err := f.Scan(&n)
    if err != nil {
        f.Println("Error", err)
        return
    }// los numeros de narayana fueron simplificados
    f.Println(fi(m.Floor(m.Log2(n*n*(n*n-1)/12))+1))
}
