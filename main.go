package main

// Define a new data type "Triangle"
type perahuMotor struct {
	mesin  bool
	ukuran string
}

// Define a new data type "Square"
type perahuLayar struct {
	mesin  bool
	ukuran string
}

// Define a new data type "Rectangle"
type kapalPesiar struct {
	mesin  bool
	ukuran string
}

func (pm perahuMotor) Mesin() bool {
	return pm.mesin
}

func (pl perahuLayar) Mesin() bool {
	return pl.mesin
}

func (kp kapalPesiar) Mesin() bool {
	return kp.mesin
}

func (pm perahuMotor) Ukuran() string {
	return pm.ukuran
}

func (pl perahuLayar) Ukuran() string {
	return pl.ukuran
}

func (kp kapalPesiar) Ukuran() string {
	return kp.ukuran
}

// Define an interface as achieve abstraction
type Kapal interface {
	Mesin()
	Ukuran()
}

func main() {
	// Declare and assign values to varaibles
	pm := perahuMotor{
		mesin:  true,
		ukuran: "Kecil",
	}
	pl := perahuLayar{
		mesin:  false,
		ukuran: "Kecil",
	}
	kp := kapalPesiar{
		mesin:  true,
		ukuran: "Besar",
	}

}
