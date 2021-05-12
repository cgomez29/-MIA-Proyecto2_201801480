package controller

import (
	"../database"
	"../models"
	"fmt"
	"github.com/gofiber/fiber"
	"strconv"
	"strings"
	"time"
)

var allSeason []models.SEASON
/* MEMBERSHIP */
const gold int = 900
const silver int = 450
const bronze int = 150
/* WEIGHING */
const b_first float64 = 0.60
const b_second float64 = 0.30
const b_third float64 = 0.1
/* TIER - WEIGHING*/
const a_gold int = 3
const a_silver int = 2
const a_bronze int = 1
/* PREMIO TOTAL DISPONIBLE A REPARTIR 80% */
const premio_total float64 = 0.8

func BulkLoad2(c *fiber.Ctx) error  {
	var data models.Temporal
	var user models.USUARIO
	var temporada models.Temporada
	var jornada models.Jornada
	var deporte models.DEPORTE
	var evento models.EVENTO

	if err:= c.BodyParser(&data); err != nil {
		return err
	}
	var query string
	//INSERTANDO USUARIO
	query = fmt.Sprintf("CALL sp_insert_usuario_bl('%s','%s','%s','%s')",
		data.Username, data.Password, data.Nombre, data.Apellido)
	_,err := database.ExecuteQuery(query)

	if err != nil  {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Usuario repetido",
		})
	}

	// Recuperando id de usuario ingresado
	query = fmt.Sprintf("SELECT idUsuario FROM USUARIO WHERE username = '%s'", data.Username)

	rows,err := database.ExecuteQuery(query)

	if err != nil  {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Error running query",
			"error": err,
		})
	}

	for rows.Next() {
		rows.Scan(&user.IdUsuario)
	}

	for _, i:= range data.Resultados {
		//INSERTANDO TEMPORADAS
		//formando la fecha
		dateTemporada := fmt.Sprintf("%s%s%s%s",i.Temporada[0:4],"/",i.Temporada[6:],"/01")
		query = fmt.Sprintf("CALL sp_insert_temporada_bl('%s','%s')", i.Temporada, dateTemporada)
		database.ExecuteQuery(query)

		//ASIGNANDO  tier para temporada actual
		//Obteniendo id temporada

		query = fmt.Sprintf("SELECT idTemporada FROM TEMPORADA WHERE nombre = '%s'", i.Temporada )

		rows,err := database.ExecuteQuery(query)
		if err != nil  {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "Error running query",
				"error": err,
			})
		}

		for rows.Next() {
			rows.Scan(&temporada.IdTemporada)
		}

		var idMembresia int
		if i.Tier == "gold" {
			idMembresia = 1
		} else if i.Tier == "silver" {
			idMembresia = 2
		} else {
			idMembresia = 3
		}

		query = fmt.Sprintf("CALL sp_insert_detalle_usuario_bl(%s,%s,%s)",
			strconv.Itoa(int(temporada.IdTemporada)), strconv.Itoa(idMembresia), strconv.Itoa(int(user.IdUsuario)))
		database.ExecuteQuery(query)

		for _, j:= range i.Jornadas {
			//CREANDO JORNADAS
			semana := fmt.Sprintf("%s",i.Temporada[6:])
			query = fmt.Sprintf("CALL sp_insert_jornada_bl('%s','%s',%s,%s)",
				j.Jornada, dateTemporada, semana, strconv.Itoa(int(temporada.IdTemporada)))
			database.ExecuteQuery(query)

			//Obteniendo id Jornada actual
			query = fmt.Sprintf("SELECT idJornada FROM JORNADA WHERE name = '%s' AND idTemporada = %s",
				j.Jornada, strconv.Itoa(int(temporada.IdTemporada)))

			rows,err := database.ExecuteQuery(query)
			if err != nil  {
				c.Status(fiber.StatusInternalServerError)
				return c.JSON(fiber.Map{
					"message": "Error running query",
					"error": err,
				})
			}

			for rows.Next() {
				rows.Scan(&jornada.IdJornada)
			}

			for _, k:= range j.Predicciones {
				// INSERTANDO DEPORTES
				query = fmt.Sprintf("CALL sp_insert_deporte_bl('%s')", k.Deporte)
				database.ExecuteQuery(query)

				//Obteniendo id Deporte
				query = fmt.Sprintf("SELECT idDeporte FROM DEPORTE WHERE nombre = '%s'",
						k.Deporte)

				rows,err = database.ExecuteQuery(query)
				if err != nil  {
					c.Status(fiber.StatusInternalServerError)
					return c.JSON(fiber.Map{
						"message": "Error running query",
						"error": err,
					})
				}

				for rows.Next() {
					rows.Scan(&deporte.IdDeporte)
				}

				//CREANDO EVENTOS
				query = fmt.Sprintf("CALL sp_insert_evento_bl('%s','%s','%s',%s,%s)",
						k.Fecha, k.Local, k.Visitante, strconv.Itoa(jornada.IdJornada), strconv.Itoa(int(deporte.IdDeporte)))
				database.ExecuteQuery(query)

				//OBTENIENDO id de EVENTO
				query = fmt.Sprintf("SELECT idEvento FROM EVENTO WHERE idJornada = %s AND idDeporte = %s AND fecha_hora = TO_TIMESTAMP('%s', 'YYYY-MM-DD HH24:MI:SS.FF')",
					strconv.Itoa(jornada.IdJornada), strconv.Itoa(int(deporte.IdDeporte)), k.Fecha)

				rows,err = database.ExecuteQuery(query)
				if err != nil  {
					c.Status(fiber.StatusInternalServerError)
					return c.JSON(fiber.Map{
						"message": "Error running query",
						"error": err,
					})
				}

				for rows.Next() {
					rows.Scan(&evento.IdEvento)
				}

				//CREANDO RESULTADOS
				query = fmt.Sprintf("CALL sp_insert_resultado_bl(%s,%s,%s)",
					strconv.Itoa(int(k.Resultado.Visitante)), strconv.Itoa(int(k.Resultado.Local)),
						strconv.Itoa(evento.IdEvento))
				database.ExecuteQuery(query)

				//CREANDO PREDICCIONES
				query = fmt.Sprintf("CALL sp_insert_prediccion_bl(%s,%s,%s,%s)",
					strconv.Itoa(int(k.Prediccion.Local)), strconv.Itoa(int(k.Prediccion.Visitante)),
					strconv.Itoa(evento.IdEvento), strconv.Itoa(int(user.IdUsuario)))
				database.ExecuteQuery(query)

			}
		}
	}

	/*

	//*/

	return c.JSON(fiber.Map{
		"message": "Successful",
	})
}

func Bulk(data models.Temporal)  {
	var user models.USUARIO
	var temporada models.Temporada
	var jornada models.Jornada
	var deporte models.DEPORTE
	var evento models.EVENTO
	var idMembresia int
	var score int = 0
	var p10 int = 0
	var p5 int = 0
	var p3 int = 0
	var p0 int = 0

	var query string
	//INSERTANDO USUARIO
	query = fmt.Sprintf("CALL sp_insert_usuario_bl('%s','%s','%s','%s')",
		data.Username, data.Password, data.Nombre, data.Apellido)
	database.Execute(query)

	// Recuperando id de usuario ingresado
	query = fmt.Sprintf("SELECT idUsuario FROM USUARIO WHERE username = '%s'", data.Username)

	rows,_ := database.ExecuteQuery(query)

	for rows.Next() {
		rows.Scan(&user.IdUsuario)
	}

	for _, i:= range data.Resultados {
		//INSERTANDO TEMPORADAS
		//formando la fecha
		dateTemporada := fmt.Sprintf("%s%s%s%s",i.Temporada[0:4],"/",i.Temporada[6:],"/01")
		query = fmt.Sprintf("CALL sp_insert_temporada_bl('%s','%s')", i.Temporada, dateTemporada)
		database.Execute(query)

		//ASIGNANDO  tier para temporada actual
		//Obteniendo id temporada

		query = fmt.Sprintf("SELECT idTemporada FROM TEMPORADA WHERE nombre = '%s'", i.Temporada )

		rows,_ := database.ExecuteQuery(query)
		for rows.Next() {
			rows.Scan(&temporada.IdTemporada)
		}

		//GUARDANDO EL ID DE LAS TEMPORADAS INGRESADAS
		var season models.SEASON
		season.IdSeason = int(temporada.IdTemporada)
		allSeason = append(allSeason, season)

		// SELECCION MEMBRESIA ACTUAL
		if i.Tier == "gold" {
			idMembresia = 3
		} else if i.Tier == "silver" {
			idMembresia = 2
		} else {
			idMembresia = 1
		}
		query = fmt.Sprintf("CALL sp_insert_detalle_usuario_bl(%s,%s,%s)",
			strconv.Itoa(int(temporada.IdTemporada)), strconv.Itoa(idMembresia), strconv.Itoa(int(user.IdUsuario)))
		database.Execute(query)

		for _, j:= range i.Jornadas {
			//CREANDO JORNADAS
			semana := fmt.Sprintf("%s",j.Jornada[1:])
			query = fmt.Sprintf("CALL sp_insert_jornada_bl('%s','%s',%s,%s)",
				j.Jornada, strings.ReplaceAll(strings.Split(dateTemporada, " ")[0], "/","-"), semana, strconv.Itoa(int(temporada.IdTemporada)))
			database.Execute(query)

			//Obteniendo id Jornada actual
			query = fmt.Sprintf("SELECT idJornada FROM JORNADA WHERE name = '%s' AND idTemporada = %s",
				j.Jornada, strconv.Itoa(int(temporada.IdTemporada)))

			rows,_ := database.ExecuteQuery(query)

			for rows.Next() {
				rows.Scan(&jornada.IdJornada)
			}

			for _, k:= range j.Predicciones {
				// INSERTANDO DEPORTES
				query = fmt.Sprintf("CALL sp_insert_deporte_bl('%s')", k.Deporte)
				database.Execute(query)

				//Obteniendo id Deporte
				query = fmt.Sprintf("SELECT idDeporte FROM DEPORTE WHERE nombre = '%s'",
					k.Deporte)

				rows,_ = database.ExecuteQuery(query)

				for rows.Next() {
					rows.Scan(&deporte.IdDeporte)
				}

				//CREANDO EVENTOS
				query = fmt.Sprintf("CALL sp_insert_evento_bl('%s','%s','%s',%s,%s)",
					strings.ReplaceAll(k.Fecha, "/","-"), k.Local, k.Visitante, strconv.Itoa(jornada.IdJornada), strconv.Itoa(int(deporte.IdDeporte)))
				database.Execute(query)

				//OBTENIENDO id de EVENTO
				query = fmt.Sprintf("SELECT idEvento FROM EVENTO WHERE idJornada = %s AND idDeporte = %s AND fecha_hora = TO_TIMESTAMP('%s', 'DD-MM-YYYY HH24:MI')",
					strconv.Itoa(jornada.IdJornada), strconv.Itoa(int(deporte.IdDeporte)), strings.ReplaceAll(k.Fecha, "/","-") )

				rows,_ = database.ExecuteQuery(query)

				for rows.Next() {
					rows.Scan(&evento.IdEvento)
				}

				//CREANDO RESULTADOS
				query = fmt.Sprintf("CALL sp_insert_resultado_bl(%s,%s,%s)",
					strconv.Itoa(int(k.Resultado.Visitante)), strconv.Itoa(int(k.Resultado.Local)),
					strconv.Itoa(evento.IdEvento))
				database.Execute(query)

				//CREANDO PREDICCIONES
				query = fmt.Sprintf("CALL sp_insert_prediccion_bl(%s,%s,%s,%s)",
					strconv.Itoa(int(k.Prediccion.Local)), strconv.Itoa(int(k.Prediccion.Visitante)),
					strconv.Itoa(evento.IdEvento), strconv.Itoa(int(user.IdUsuario)))
				database.Execute(query)

				// CALCULANDO PUNTOS
				// Puntaje total, P10,P5,P3,P0

				score = 0
				p10 = 0
				p5 = 0
				p3 = 0
				p0 = 0

				if k.Resultado.Local == k.Prediccion.Local &&
					k.Resultado.Visitante == k.Prediccion.Visitante { // P10
					p10 = 1
					score = 10
				} else if k.Resultado.Local > k.Resultado.Visitante { //Ganador: Local
					if k.Prediccion.Local > k.Prediccion.Visitante {
						if (k.Resultado.Local - k.Prediccion.Local) == 2 {
							p5 = 1
							score = 5
						}  else {
							p3 = 1
							score = 3
						}
					} else {
						p0 = 1
						score = 0
					}
				} else if k.Resultado.Local < k.Resultado.Visitante { //Ganador: Visitante
					if k.Prediccion.Local < k.Prediccion.Visitante {
						if (k.Resultado.Local - k.Prediccion.Local) == 2 {
							p5 = 1
							score = 5
						}  else {
							p3 = 1
							score = 3
						}
					} else {
						p0 = 1
						score = 0
					}
				}
				// ACTUALIZANDO EL DETALLE USUARIO
				query = fmt.Sprintf("CALL sp_update_detalleu_score_bl(%s,%s,%s,%s,%s,%s,%s)",
					strconv.Itoa(score), strconv.Itoa(p10),strconv.Itoa(p5),strconv.Itoa(p3),strconv.Itoa(p0),
					strconv.Itoa(int(temporada.IdTemporada)),strconv.Itoa(int(user.IdUsuario)))
				database.Execute(query)

			}
		}
	}
}

func BulkLoad(c *fiber.Ctx) error {
	var data models.ArrayTemporal

	if err:= c.BodyParser(&data); err != nil {
		return err
	}

	for _, value := range data{
		Bulk(value)
		//fmt.Println(value.Username)
	}
	//CALCULANDO RECOMPENSAS POR TEMPORADAS
	Rewards()

	return c.JSON(fiber.Map{
		"message": "Successful",
	})
}

func Rewards() {
	/**
	*	g: Golden
	*	s: Silver
	*	b: Bronze
	**/

	var reward models.RewardBulkLoad
	var totalTier models.TotalTemporada
	var query string
	var multiplicador float64
	var total int
	var premio float64
	//Obteniendo el total de ganancias


	for _, i:= range allSeason {
		total = 0 // se reinicia por cada temporada
		query = fmt.Sprintf("SELECT COUNT(idMembresia), idMembresia FROM DETALLE_USUARIO WHERE idTemporada = %s GROUP BY idMembresia", strconv.Itoa(i.IdSeason))
		rows,_ := database.ExecuteQuery(query)
		for rows.Next() {
			rows.Scan(&totalTier.Cantidad, &reward.IdMembresia)
			if reward.IdMembresia == 3 { //GOLD
				total += totalTier.Cantidad * gold
			} else if reward.IdMembresia == 2 { //SILVER
				total += totalTier.Cantidad * silver
			} else { //BRONZE
				total += totalTier.Cantidad * bronze
			}
		}

		// Obteniendo los tres usuarios ganadores
		query = fmt.Sprintf("SELECT score, idUsuario, idMembresia FROM DETALLE_USUARIO WHERE idTemporada = %s ORDER BY score DESC FETCH FIRST 3 ROWS ONLY", strconv.Itoa(i.IdSeason))
		rows,_ = database.ExecuteQuery(query)
		var user models.Multiplicador
		var winners []models.Multiplicador

		for rows.Next() {
			rows.Scan(&reward.Score, &reward.IdUsuario, &reward.IdMembresia)
			// CALCULANDO EL MULTIPLICADOR DE TIER
			if reward.IdMembresia == 3 { //GOLD
				multiplicador = float64(1 + (a_gold / (a_gold + a_silver + a_bronze)))
			} else if reward.IdMembresia == 2 { //SILVER
				multiplicador = float64(1 + (a_silver / (a_gold + a_silver + a_bronze)))
			} else { //BRONZE
				multiplicador = float64(1 + (a_bronze / (a_gold + a_silver + a_bronze)))
			}
			// GUARDANDO LOS GANADORES CON SU MULTIPLICADOR
			user.IdUsuario = reward.IdUsuario
			user.Score = reward.Score
			user.IdMembresia = reward.IdMembresia
			user.IdTemporada= i.IdSeason
			user.Mtier = multiplicador
			winners =  append(winners, user)
		}
		// REALIZANDO LA SUMATORIA DE MULTIPLICADORES DE TIER
		var key = 0
		var sumatoria float64
		for _, j:= range winners {
			if key == 0 { //GOLD
				sumatoria += b_first*j.Mtier
				key++
			} else if key == 1 { //SILVER
				sumatoria += b_second*j.Mtier
				key++
			} else { //BRONZE
				sumatoria += b_third*j.Mtier
			}
		}
		// CALCULANDO LOS PREMIOS
		var tierString string
		key = 0
		for _, j:= range winners {
			fmt.Println("KEYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY",key)
			if j.IdMembresia == 3 { //GOLD
				tierString = "Gold"
			} else if j.IdMembresia == 2 { //SILVER
				tierString = "Silver"
			} else { //BRONZE
				tierString = "Bronze"
			}

			if key == 0 { //Primer lugar
				premio = float64(total) * premio_total * b_first * (1 + j.Mtier - sumatoria)
				key++
			} else if key == 1 { //Segundo Lugar
				premio = float64(total) * premio_total * b_second * (1 + j.Mtier - sumatoria)
				key++
			} else { //Tercer lugar
				premio = float64(total) * premio_total * b_third * (1 + j.Mtier - sumatoria)
			}

			// INSERTANDO LAS RECOMPENZAS DE LOS GANADORES
			query = fmt.Sprintf("CALL sp_insert_recompensa_bl(%s,%s,'%s',%s,%s)",
				strconv.Itoa(j.Score), strconv.FormatFloat(premio,'E', -1, 32),
				tierString, strconv.Itoa(j.IdUsuario), strconv.Itoa(j.IdTemporada))
			database.Execute(query)
		}
	}
}

//TemporadaActual crea la temporada si no existe
func TemporadaActual()  {
	t := time.Now()
	month := t.Month() // type time.Month
	year := t.Year() // type time.Month
	fecha := fmt.Sprintf("1-%s-%s",strconv.Itoa(int(month)),strconv.Itoa(int(year)))
	fecha_inicio := fmt.Sprintf("%s-%s-01",strconv.Itoa(int(year)),"05")
	name := fmt.Sprintf("%s-Q%s",strconv.Itoa(int(year)),strconv.Itoa(int(month)))

	query := fmt.Sprintf("CALL sp_insert_temporada_actual('%s','%s','%s')",
		name, fecha, fecha_inicio)
	database.Execute(query)

	var idTemporada int
	var idUsuario int
	var tier string
	//Obteniendo el id de la temporada actual
	query = fmt.Sprintf("SELECT idTemporada FROM TEMPORADA WHERE nombre = '%s' ", name)

	rows,err := database.ExecuteQuery(query)

	if err != nil  {
		return
	}

	for rows.Next() {
		rows.Scan(&idTemporada)
	}
	//Obteniendo usuario y su tier
	query = "SELECT idUsuario, tier FROM USUARIO WHERE tier != '-' AND idRol = 2"

	rows,err = database.ExecuteQuery(query)

	if err != nil  {
		return
	}
	for rows.Next() {
		rows.Scan(&idUsuario, &tier)
		query = fmt.Sprintf("CALL sp_insert_detalle_usuario_bl(%s,%s,%s)",
			strconv.Itoa(idTemporada), tier, strconv.Itoa(idUsuario))
		database.Execute(query)
	}
	defer rows.Close()
}

func JornadaDetalle (c *fiber.Ctx) error  {
	t := time.Now()
	day := t.Day() // type time.Month
	month := t.Month() // type time.Month
	year := t.Year() // type time.Month
	name := fmt.Sprintf("%s-Q%s",strconv.Itoa(int(year)),strconv.Itoa(int(month)))

	var idTemporada string
	var participantes string
	var fecha_fin string
	var fecha_inicio string
	var jornada string
	var idJornada string
	var semanaJornada string

	query := fmt.Sprintf("SELECT COUNT(DETALLE_USUARIO.id_detalle_usuario), TEMPORADA.idTemporada, TEMPORADA.fechainicio, TEMPORADA.fechafin " +
		"FROM DETALLE_USUARIO INNER JOIN TEMPORADA ON DETALLE_USUARIO.idTemporada = TEMPORADA.idTemporada " +
		"WHERE TEMPORADA.nombre = '%s' GROUP BY TEMPORADA.idTemporada, TEMPORADA.fechainicio, TEMPORADA.fechafin", name)

	rows,err := database.ExecuteQuery(query)

	if err != nil {
		return nil
	}

	for rows.Next() {
		rows.Scan(&participantes, &idTemporada, &fecha_inicio, &fecha_fin)
	}

	//Verificando jornada actual si no se crea
	if day < 8 {
		jornada = "J1"
		semanaJornada = "1"
	} else if day < 15 {
		jornada = "J2"
		semanaJornada = "2"
	} else if day < 22 {
		jornada = "J3"
		semanaJornada = "3"
	} else {
		jornada = "J4"
		semanaJornada = "4"
	}

	query = fmt.Sprintf("CALL sp_insert_jornada_bl('%s','%s',%s,%s)",
		jornada, fecha_inicio[0:10], semanaJornada, idTemporada)
	database.Execute(query)

	query = fmt.Sprintf("SELECT idJornada FROM JORNADA WHERE name = '%s' AND  idTemporada =%s", jornada, idTemporada)
	rows,err = database.ExecuteQuery(query)

	if err != nil {
		return nil
	}

	for rows.Next() {
		rows.Scan(&idJornada)
	}

	return c.JSON(fiber.Map{
		"temporada": name,
		"participantes": participantes,
		"fecha": fecha_fin,
		"jornada": jornada,
		"idJornada": idJornada,
	})
}

func HomeDetalle(c *fiber.Ctx) error  {
	t := time.Now()
	month := t.Month() // type time.Month
	year := t.Year() // type time.Month
	name := fmt.Sprintf("%s-Q%s",strconv.Itoa(int(year)),strconv.Itoa(int(month)))

	query := fmt.Sprintf("SELECT COUNT(idMembresia), idMembresia " +
		"FROM DETALLE_USUARIO INNER JOIN TEMPORADA ON DETALLE_USUARIO.idTemporada = TEMPORADA.idTemporada " +
		"WHERE TEMPORADA.nombre = '%s' GROUP BY idMembresia", name)
	rows,err := database.ExecuteQuery(query)
	if err != nil {
		return nil
	}

	var total = 0 // se reinicia por cada temporada
	var totalTier models.TotalTemporada
	var reward models.RewardBulkLoad

	var count_gold = 0 	//cantidad de participates con tier gold
	var count_silver = 0 //cantidad de participates con tier silver
	var count_bronze = 0 //cantidad de participates con tier bronze

	for rows.Next() {
		rows.Scan(&totalTier.Cantidad, &reward.IdMembresia)
		if reward.IdMembresia == 3 { //GOLD
			total += totalTier.Cantidad * gold
			count_gold = totalTier.Cantidad
		} else if reward.IdMembresia == 2 { //SILVER
			total += totalTier.Cantidad * silver
			count_silver = totalTier.Cantidad
		} else { //BRONZE
			total += totalTier.Cantidad * bronze
			count_bronze = totalTier.Cantidad
		}
	}

	return c.JSON(fiber.Map{
		"capital": total,
		"gold": count_gold,
		"silver": count_silver,
		"bronze": count_bronze,
	})
}
