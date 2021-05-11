package models

/* USUARIO */

type USUARIO struct {
	IdUsuario uint `json:"idUsuario"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Tier string `json:"tier"`
	Fecha string `json:"fecha"`
	Email string `json:"email"`
	Photo string `json:"file"`
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
	Imagen string `json:"file"`
	Color string `json:"color"`
}

type ArrayDeporte []DEPORTE
/* END DEPORTES*/

/* PREDICCIONES Usuario */
type PREDICCION_USUARIO struct {
	IdPrediccion int `json:"idPrediccion"`
	Local string `json:"local"`
	Visitante string `json:"visitante"`
	IdEvento string `json:"idEvento"`
	IdUsuario int `json:"idUsuario"`
}
