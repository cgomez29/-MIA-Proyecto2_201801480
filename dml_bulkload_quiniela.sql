/* bl: BULKLOAD*/

/* USUARIO */
CREATE OR REPLACE PROCEDURE sp_insert_usuario_bl ( 
    v_username IN VARCHAR2,
    v_password IN VARCHAR2,
    v_name IN VARCHAR2,
    v_surname IN VARCHAR2
    )
AS
BEGIN
    INSERT INTO USUARIO(username, password, name, surname, tier, fecha_nacimiento, fecha_registro, email, photo, idRol)
    VALUES (v_username, v_password, v_name, v_surname,'-',TO_DATE(SYSDATE, 'yyyy-mm-dd'),
    SYSTIMESTAMP,'x@x.com','x',2);
    COMMIT;
END;

/*DEPORTE*/
CREATE OR REPLACE PROCEDURE sp_insert_deporte_bl ( 
    v_name IN VARCHAR2
    )
AS
BEGIN
    INSERT INTO DEPORTE(nombre, imagen, color)
    VALUES (v_name, '-', '#9e9e9e');
    COMMIT;
END;

/*TEMPORADA*/
CREATE OR REPLACE PROCEDURE sp_insert_temporada_bl ( 
    v_name IN VARCHAR2,
    v_fecha_inicio_fin IN VARCHAR2
    )
AS
BEGIN
    INSERT INTO TEMPORADA(nombre, fechainicio, fechafin, estado)
    VALUES (v_name, TO_DATE(v_fecha_inicio_fin, 'yyyy-mm-dd'), LAST_DAY(TO_DATE(v_fecha_inicio_fin, 'yyyy-mm-dd')), 0);
    COMMIT;
END;

/* DETALLE USUARIO */

CREATE OR REPLACE PROCEDURE sp_insert_detalle_usuario_bl ( 
        v_idTemporada IN NUMBER,
        v_idMembresia IN NUMBER,
        v_idUsuario IN NUMBER
        )
    AS
        rowU NUMBER(1) := 0;
    BEGIN

        SELECT COUNT(id_detalle_usuario) INTO rowU  FROM DETALLE_USUARIO WHERE
        idTemporada = v_idTemporada AND idUsuario = v_idUsuario;
        
        IF rowU = 0 THEN
            INSERT INTO DETALLE_USUARIO(score, incremento, pos_anterior, p_10, p_5, p_3, p_0, idTemporada, idMembresia, idUsuario)
            VALUES (0,0,0,0,0,0,0,v_idTemporada, v_idMembresia, v_idUsuario);
            
            UPDATE USUARIO SET tier = v_idMembresia WHERE idUsuario = v_idUsuario;
            
            COMMIT;
        END IF;
END;

CREATE OR REPLACE PROCEDURE sp_update_detalleu_score_bl ( 
    v_score IN NUMBER,
    v_p10 IN NUMBER,
    v_p5 IN NUMBER,
    v_p3 IN NUMBER,
    v_p0 IN NUMBER,
    v_idTemporada IN NUMBER,
    v_idUsuario IN NUMBER
    )
AS
BEGIN
    UPDATE DETALLE_USUARIO SET score = (score + v_score), p_10 = (p_10 + v_p10),
         p_5 = (p_5 + v_p5), p_3 = (p_3 + v_p3), p_0 = (p_0 + v_p0)
    WHERE idUsuario = v_idUsuario AND idTemporada = v_idTemporada;
    COMMIT;
END;

/*JORNADA */
CREATE OR REPLACE PROCEDURE sp_insert_jornada_bl ( 
    v_name IN VARCHAR2,
    v_fecha_inicio IN VARCHAR2, -- Fecha inicio de la temporada
    v_semana IN NUMBER,
    v_temporada IN NUMBER
    )
AS
    rowU NUMBER(1) := 0;
BEGIN

    SELECT COUNT(idJornada) INTO rowU FROM JORNADA WHERE name = v_name AND 
    idTemporada = v_temporada;
    
    IF rowU = 0 THEN 
        IF (v_semana = 1 ) THEN    
            INSERT INTO JORNADA(name, fecha_hora_inicio, fecha_hora_fin, idFase, idTemporada)
            VALUES (v_name, TO_DATE(v_fecha_inicio, 'YYYY-MM-DD'), (TO_DATE(v_fecha_inicio, 'YYYY-MM-DD')+7),
                    1, v_temporada); -- VERIFICAR EN QUE FASE ES QUE ENTRA LA JORNADA CARGADA
        END IF;
        
        IF (v_semana = 2 ) THEN    
            INSERT INTO JORNADA(name, fecha_hora_inicio, fecha_hora_fin, idFase, idTemporada)
            VALUES (v_name, (TO_DATE(v_fecha_inicio, 'YYYY-MM-DD')+7), (TO_DATE(v_fecha_inicio, 'YYYY-MM-DD')+14),
                    1, v_temporada); -- VERIFICAR EN QUE FASE ES QUE ENTRA LA JORNADA CARGADA
        END IF;
     
        IF (v_semana = 3 ) THEN    
            INSERT INTO JORNADA(name, fecha_hora_inicio, fecha_hora_fin, idFase, idTemporada)
            VALUES (v_name, (TO_DATE(v_fecha_inicio, 'YYYY-MM-DD')+14), (TO_DATE(v_fecha_inicio, 'YYYY-MM-DD')+21),
                    1, v_temporada); -- VERIFICAR EN QUE FASE ES QUE ENTRA LA JORNADA CARGADA
        END IF;
        
        IF (v_semana = 4 ) THEN    
            INSERT INTO JORNADA(name, fecha_hora_inicio, fecha_hora_fin, idFase, idTemporada)
            VALUES (v_name, (TO_DATE(v_fecha_inicio, 'YYYY-MM-DD')+21),LAST_DAY(TO_DATE(v_fecha_inicio, 'YYYY-MM-DD')) ,
                    1, v_temporada); -- VERIFICAR EN QUE FASE ES QUE ENTRA LA JORNADA CARGADA
        END IF;
    END IF;
    -- COMMIT;
END;

/*EVENTO*/
CREATE OR REPLACE PROCEDURE sp_insert_evento_bl ( 
    v_fechahora IN VARCHAR2,
    v_local IN VARCHAR2,
    v_visitante IN VARCHAR2,
    v_idJornada IN NUMBER,
    v_idDeporte IN NUMBER
    )
AS
    rowU NUMBER(1) := 0;
BEGIN

    SELECT COUNT(idEvento) INTO rowU FROM EVENTO WHERE fecha_hora = TO_DATE(v_fechahora, 'DD-MM-YYYY HH24:MI') AND
    local = v_local AND visitante = v_visitante AND idJornada = v_idJornada AND idDeporte = v_idDeporte;
    
    IF rowU = 0 THEN 
        INSERT INTO EVENTO(fecha_hora, estado, local, visitante,color, idJornada, idDeporte)
        VALUES (TO_DATE(v_fechahora, 'DD-MM-YYYY HH24:MI'), 0, v_local, v_visitante,'#8d8f91',
                v_idJornada, v_idDeporte);
        COMMIT;
    END IF;
END;

/*RESULTADO*/
CREATE OR REPLACE PROCEDURE sp_insert_resultado_bl ( 
    v_visitante IN NUMBER,
    v_local IN NUMBER,
    v_idEvento IN NUMBER
    )
AS
    rowU NUMBER(1) := 0;
BEGIN

    SELECT COUNT(idResultado) INTO rowU FROM RESULTADO WHERE idEvento = v_idEvento;
    
    IF rowU = 0 THEN
        INSERT INTO RESULTADO(visitante, local, idEvento) 
        VALUES (v_visitante, v_local, v_idEvento);
        COMMIT;
    END IF;
END;

/*PREDICCION*/
CREATE OR REPLACE PROCEDURE sp_insert_prediccion_bl ( 
    v_local IN NUMBER,
    v_visitante IN NUMBER,
    v_idEvento IN NUMBER,
    v_idUsuario IN NUMBER
    )
AS
    rowU NUMBER(1) := 0;
BEGIN

    SELECT COUNT(idPrediccion) INTO rowU FROM PREDICCION WHERE idEvento = v_idEvento AND 
    idUsuario = v_idUsuario;
    
    IF rowU = 0 THEN
        INSERT INTO PREDICCION(local, visitante, idEvento, idUsuario) 
        VALUES (v_local, v_visitante,v_idEvento, v_idUsuario);
        COMMIT;
    END IF;
END;

/* RECOMPENSA */
CREATE OR REPLACE PROCEDURE sp_insert_recompensa_bl ( 
    v_score IN NUMBER,
    v_premio IN FLOAT,
    v_tier IN VARCHAR2,
    v_idUsuario IN NUMBER,
    v_idTemporada IN NUMBER
    )
AS
BEGIN
    INSERT INTO RECOMPENSA(score, premio, tier, ultimo, fecha, incremento, idUsuario, idTemporada) 
    VALUES (v_score, v_premio, v_tier, 0, SYSDATE,0, v_idUsuario, v_idTemporada);
    COMMIT;
END;


/*TEMPORADA ACTUAL*/
CREATE OR REPLACE PROCEDURE sp_insert_temporada_actual ( 
    v_name IN VARCHAR2,
    v_fecha IN VARCHAR2,
    v_fecha_inicio IN VARCHAR2
    )
AS
    rowU NUMBER(1) := 0;
BEGIN

    SELECT COUNT(idTemporada) INTO rowU  FROM TEMPORADA WHERE fechainicio = TO_DATE(v_fecha_inicio, 'yyyy-mm-dd');
    
    IF rowU = 0 THEN
        INSERT INTO TEMPORADA(nombre, fechainicio, fechafin, estado)
        VALUES (v_name, TO_DATE(v_fecha_inicio, 'yyyy-mm-dd'), LAST_DAY(TO_DATE(v_fecha_inicio, 'yyyy-mm-dd')), 1);
        COMMIT;
    END IF;
END;


-- alter SESSION set NLS_DATE_FORMAT = 'DD-MM-YYYY HH24:MI'

COMMIT;

DELETE FROM TEMPORADA WHERE idTemporada = 564;
SELECT * FROM USUARIO where idUsuario = 17;
SELECT * FROM TEMPORADA;
SELECT * FROM DETALLE_USUARIO where idTemporada = 565;
SELECT * FROM JORNADA WHERE name = 'J2' AND  idTemporada = 561;
SELECT * FROM DEPORTE;
SELECT * FROM EVENTO WHERE idEvento = 561;
SELECT * FROM RESULTADO;
SELECT * FROM PREDICCION WHERE idUsuario = 2;
SELECT * FROM MEMBRESIA;
SELECT * FROM RECOMPENSA WHERE idTemporada = 1;
SELECT * FROM RECOMPENSA;
SELECT * FROM USUARIO;
SELECT * FROM FASE;



SELECT COUNT(DETALLE_USUARIO.id_detalle_usuario), TEMPORADA.nombre, TEMPORADA.fechafin
FROM DETALLE_USUARIO 
INNER JOIN TEMPORADA ON DETALLE_USUARIO.idTemporada = TEMPORADA.idTemporada
WHERE TEMPORADA.nombre = '2021-Q5'
GROUP BY TEMPORADA.nombre;

SELECT COUNT(DETALLE_USUARIO.id_detalle_usuario), TEMPORADA.nombre FROM DETALLE_USUARIO INNER JOIN TEMPORADA ON DETALLE_USUARIO.idTemporada = TEMPORADA.idTemporada WHERE TEMPORADA.nombre = '2021-Q5' GROUP BY TEMPORADA.nombre;




