package models

type ResultadoV struct{
	Visitante uint `json:"visitante"`
	Local uint `json:"local"`
}

type PreddicionV struct{
	Visitante uint `json:"visitante"`
	Local uint `json:"local"`
}

type PrediccionesV struct {
	Deporte string `json:"deporte"`
	Fecha string `json:"fecha"`
	Visitante string `json:"visitante"`
	Local string `json:"local"`
	Prediccion PreddicionV `json:"prediccion"`
	Resultado ResultadoV `json:"resultado"`
}

type ArrayPredicciones []PrediccionesV

type  JornadasV struct {
	Jornada string `json:"jornada"`
	Predicciones ArrayPredicciones `json:"predicciones"`
}

type ArrayJornadas []JornadasV

type ResultadosV struct {
	Temporada string `json:"temporada"`
	Tier string `json:"tier"`
	Jornadas ArrayJornadas `json:"jornadas"`
}

type ArrayResultados []ResultadosV

type Temporal struct {
	Nombre    string `json:"nombre"`
	Apellido    string `json:"apellido"`
	Password    string `json:"password"`
	Username    string `json:"username"`
	Resultados ArrayResultados `json:"resultados"`
}

type ArrayTemporal  map[string]Temporal

/* TEMPORADA */

type  Temporada struct {
	IdTemporada uint `json:"idTemporada"`
	Nombre string `json:"nombre"`
	FechaInicio string `json:"fechainicio"`
	FechaFin string `json:"fechafin"`
	Estado uint `json:"estado"`
}

/* JORNADA */
type  Jornada struct {
	IdJornada int `json:"idJornada"`
	Nombre string `json:"name"`
	FechaInicio string `json:"fechainicio"`
	FechaFin string `json:"fechafin"`
	IdFase int `json:"idFase"`
	IdTemporada uint `json:"idTemporada"`
}

/* EVENTO */

type EVENTO struct {
	IdEvento int `json:"idEvento"`
	FechaHora string `json:"fechahora"`
	Estado int `json:"estado"`
	Local string `json:"local"`
	Visitante string `json:"visitante"`
	IdJornada int `json:"idJornada"`
	IdDeporte int `json:"idDeporte"`
}

/* PREDICCIONES */
type PREDICCION struct {
	IdPrediccion int `json:"idPrediccion"`
	Local int `json:"local"`
	Visitante int `json:"visitante"`
	IdEvento int `json:"idEvento"`
	IdUsuario int `json:"idUsuario"`
}

/* RESULTADOS */
type RESULTADO struct {
	IdResultado int `json:"idResultado"`
	Local int `json:"local"`
	Visitante int `json:"visitante"`
	IdEvento int `json:"idEvento"`
}
