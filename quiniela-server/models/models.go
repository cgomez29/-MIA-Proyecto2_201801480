package models

/* USUARIO */

type USUARIO struct {
	IdUsuario uint `json:"idUsuario"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Tier string `json:"tier"`
	FechaNacimiento string `json:"date"`
	Email string `json:"email"`
	Photo string `json:"img"`
	IdRol uint `json:"idRol"`
}

/* END USUARIO */

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
/* END DEPORTES*/