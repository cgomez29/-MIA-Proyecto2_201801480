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

/*EXEC sp_insert_jornada_bl ('J1', '2018/3/1', 1, 1);
EXEC sp_insert_jornada_bl ('J2', '2018/3/1', 2, 1);
EXEC sp_insert_jornada_bl ('J3', '2018/3/1', 3, 1);
EXEC sp_insert_jornada_bl ('J4', '2018/3/1', 4, 1);*/



SELECT * FROM USUARIO;
SELECT * FROM TEMPORADA;
SELECT * FROM DETALLE_USUARIO;
SELECT * FROM JORNADA;
SELECT * FROM DEPORTE;

SELECT idUsuario FROM USUARIO WHERE username = 'jpook0@army.mil';

EXEC sp_insert_detalle_usuario_bl(1,1,1);
SELECT idTemporada FROM TEMPORADA WHERE nombre = '2020-Q4';


COMMIT;
-- 1 8 15  22



