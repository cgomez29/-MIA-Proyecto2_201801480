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