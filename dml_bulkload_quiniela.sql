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
    INSERT INTO USUARIO(username, password, name, surname, tier, fecha_registro, idRol)
    VALUES (v_username, v_password, v_name, v_surname,'-',SYSTIMESTAMP,2);
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
    VALUES (v_name, TO_DATE(v_fecha_inicio_fin, 'yyyy/mm/dd'), LAST_DAY(TO_DATE(v_fecha_inicio_fin, 'yyyy/mm/dd')), 0);
    COMMIT;
END;

/* DETALLE USUARIO */

CREATE OR REPLACE PROCEDURE sp_insert_detalle_usuario_bl ( 
    v_idTemporada IN NUMBER,
    v_idMembresia IN NUMBER,
    v_idUsuario IN NUMBER
    )
AS
BEGIN
    INSERT INTO DETALLE_USUARIO(score, incremento, pos_anterior, p_10, p_5, p_3, p_0, idTemporada, idMembresia, idUsuario)
    VALUES (0,0,0,0,0,0,0,v_idTemporada, v_idMembresia, v_idUsuario);
    COMMIT;
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
BEGIN

    IF (v_semana = 1 ) THEN    
        INSERT INTO JORNADA(name, fecha_hora_inicio, fecha_hora_fin, idFase, idTemporada)
        VALUES (v_name, TO_TIMESTAMP(v_fecha_inicio, 'YYYY-MM-DD HH24:MI:SS.FF'), (TO_TIMESTAMP(v_fecha_inicio, 'YYYY-MM-DD HH24:MI:SS.FF')+7),
                1, v_temporada); -- VERIFICAR EN QUE FASE ES QUE ENTRA LA JORNADA CARGADA
    END IF;
    
    IF (v_semana = 2 ) THEN    
        INSERT INTO JORNADA(name, fecha_hora_inicio, fecha_hora_fin, idFase, idTemporada)
        VALUES (v_name, (TO_TIMESTAMP(v_fecha_inicio, 'YYYY-MM-DD HH24:MI:SS.FF')+7), (TO_TIMESTAMP(v_fecha_inicio, 'YYYY-MM-DD HH24:MI:SS.FF')+14),
                1, v_temporada); -- VERIFICAR EN QUE FASE ES QUE ENTRA LA JORNADA CARGADA
    END IF;
 
    IF (v_semana = 3 ) THEN    
        INSERT INTO JORNADA(name, fecha_hora_inicio, fecha_hora_fin, idFase, idTemporada)
        VALUES (v_name, (TO_TIMESTAMP(v_fecha_inicio, 'YYYY-MM-DD HH24:MI:SS.FF')+14), (TO_TIMESTAMP(v_fecha_inicio, 'YYYY-MM-DD HH24:MI:SS.FF')+21),
                1, v_temporada); -- VERIFICAR EN QUE FASE ES QUE ENTRA LA JORNADA CARGADA
    END IF;
    
    IF (v_semana = 4 ) THEN    
        INSERT INTO JORNADA(name, fecha_hora_inicio, fecha_hora_fin, idFase, idTemporada)
        VALUES (v_name, (TO_TIMESTAMP(v_fecha_inicio, 'YYYY-MM-DD HH24:MI:SS.FF')+21),LAST_DAY(TO_TIMESTAMP(v_fecha_inicio, 'YYYY-MM-DD HH24:MI:SS.FF')) ,
                1, v_temporada); -- VERIFICAR EN QUE FASE ES QUE ENTRA LA JORNADA CARGADA
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
BEGIN
    INSERT INTO EVENTO(fecha_hora, estado, local, visitante, idJornada, idDeporte)
    VALUES (TO_TIMESTAMP(v_fechahora, 'YYYY-MM-DD HH24:MI:SS.FF'), 0, v_local, v_visitante,
            v_idJornada, v_idDeporte);
    COMMIT;
END;

/*RESULTADO*/
CREATE OR REPLACE PROCEDURE sp_insert_resultado_bl ( 
    v_visitante IN NUMBER,
    v_local IN NUMBER,
    v_idEvento IN NUMBER
    )
AS
BEGIN
    INSERT INTO RESULTADO(visitante, local, idEvento) 
    VALUES (v_visitante, v_local, v_idEvento);
    COMMIT;
END;

/*PREDICCION*/
CREATE OR REPLACE PROCEDURE sp_insert_prediccion_bl ( 
    v_local IN NUMBER,
    v_visitante IN NUMBER,
    v_idEvento IN NUMBER,
    v_idUsuario IN NUMBER
    )
AS
BEGIN
    INSERT INTO PREDICCION(local, visitante, idEvento, idUsuario) 
    VALUES (v_local, v_visitante,v_idEvento, v_idUsuario);
    COMMIT;
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


/*EXEC sp_insert_jornada_bl ('J1', '2018/3/1', 1, 1);
EXEC sp_insert_jornada_bl ('J2', '2018/3/1', 2, 1);
EXEC sp_insert_jornada_bl ('J3', '2018/3/1', 3, 1);
EXEC sp_insert_jornada_bl ('J4', '2018/3/1', 4, 1);*/



SELECT * FROM USUARIO;
SELECT * FROM TEMPORADA;
SELECT * FROM DETALLE_USUARIO WHERE idTemporada = 3;
SELECT * FROM JORNADA;
SELECT * FROM DEPORTE;
SELECT * FROM EVENTO;
SELECT * FROM RESULTADO;
SELECT * FROM PREDICCION;
SELECT * FROM MEMBRESIA;
SELECT * FROM RECOMPENSA WHERE idTemporada = 1;
SELECT * FROM RECOMPENSA;
SELECT * FROM USUARIO;

UPDATE USUARIO set photo = 'scorpio.jpg' , fecha_nacimiento = sysdate where idUsuario = 2;

EXEC sp_insert_evento_bl('5/5/5','efe','ef',1,1);

COMMIT;

SELECT score, idUsuario, idMembresia FROM DETALLE_USUARIO 
WHERE idTemporada = 1
ORDER BY score DESC FETCH FIRST 3 ROWS ONLY;

SELECT * FROM DETALLE_USUARIO ;

SELECT COUNT(idMembresia), idMembresia FROM DETALLE_USUARIO WHERE idTemporada = 1 GROUP BY idMembresia;


COMMIT;
-- 1 8 15  22



