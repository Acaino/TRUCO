package main

import(
		"fmt"
		"math/rand"
)


func creacionMazo(){
	var palos[4] string = [4]string{"oro"; "copa"; "espada"; "basto"}
	var numeros[10] string = [10]string{"1"; "2", "3", "4", "5", "6", "7", "10", "11", "12"}
	var cartasTotales[40] string
	var k int = 0

	for i:= 0; i<len(palos); i++{
		for j:=0; j<len(numero); j++{
			cartasTotales[k] = palos[i]+" de "+numeros[j]
			k++
		}
	}
	return cartasTotales
}

func mezclar(){

	var indiceRand1, indiceRand2 int
	var contenidoRand1, contenidoRand2 string
	cartasMezcladas = mazo()
	for i:= 0; i<100; i++{
		indiceRand1 = rand.Intn(len(cartasTotales))
		contenidoRand1 = cartasTotales[indiceRand1] 
		indiceRand2 = rand.Intn(len(cartasTotales))
		contenidoRand2 = cartasTotales[indiceRand2]
		cartasTotales[indiceRand1] = contenidoRand2
		cartasTotales[indiceRand2] = contenidoRand1
	}
	
}
func menu(){ //tengo que hacer que esta función sea variable en función (ja) de los estadíos del juego
	


}
func main(){
	var respuesta byte
	creacionMazo()
	mezclar()
	var cartasPlayerA[3]string = []
	var cartasPlayerB[3]string
	

	fmt.Printf("Bienvenido al Truco 2021\n\n")
	fmt.Printf("¿Desea comenzar? s/n", &comenzar)

	for 
}