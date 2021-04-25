package models

type Prueba struct {
	Prueba_id int
	Nombre    string
}

type ArrayPrueba []Prueba

/* DEPORTES */
type DEPORTE struct {
	IdDeporte uint `json:"idDeporte"`
	Nombre string `json:"name"`
	Imagen string `json:"img"`
	Color string `json:"color"`
}

type ArrayDeporte []DEPORTE
