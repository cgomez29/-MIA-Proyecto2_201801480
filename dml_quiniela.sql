
/**
*   Stored Procedured required
**/


CREATE OR REPLACE PROCEDURE sp_auth ( 
    email IN VARCHAR2,
    password IN VARCHAR2
    )
AS
    auth SYS_REFCURSOR;
BEGIN
    IF auth IS NULL THEN
        OPEN auth FOR
        SELECT * FROM CLIENTE WHERE CLIENTE.email = email AND CLIENTE.password = password; 
        DBMS_SQL.RETURN_RESULT(auth);
        -- dbms_output.put_line(foo);
    END IF;
    
END;

CREATE OR REPLACE PROCEDURE sp_membership_payment (
    local IN NUMBER,    
    visitante IN NUMBER,    
    idEvento IN NUMBER,    
    idCliente IN NUMBER    
)
AS 
BEGIN
    INSERT INTO PREDICCION(local, visitante, idEvento, idCliente) 
    VALUES (local, visitante, idEvento, idCliente);
END;



EXECUTE sp_membership_payment(1,1,1,1);

EXECUTE sp_auth('efe','');



/**
*   PAG. 18
*   Deporte, Local, Visitante, Prediccion, Resultado, Puntos, Fecha
*/

SELECT DEPORTE.nombre FROM DETALLE_CLIENTE
INNER JOIN CLIENTE ON DETALLE_CLIENTE.idCliente = CLIENTE.idCliente
INNER JOIN PREDICCION ON CLIENTE.idCliente = PREDICCION.idCliente
INNER JOIN EVENTO ON PREDICCION.idEvento = EVENTO.idEvento
INNER JOIN DEPORTE ON EVENTO.idDeporte = DEPORTE.idDeporte
INNER JOIN DETALLE_EVENTO ON EVENTO.idEvento = DETALLE_EVENTO.idEvento ;
 

select * from admins;


INSERT INTO ADMINS(username, password, email) VALUES ('cgomez29','cris','cris@cris.com');

COMMIT;

CREATE TABLE PRUEBA (
    prueba_id INTEGER GENERATED ALWAYS AS IDENTITY(START WITH 1 INCREMENT BY 1) NOT NULL,
    nombre VARCHAR(45) NOT NULL,
    CONSTRAINT prueba_id PRIMARY KEY(prueba_id)
);



insert into Prueba(nombre) VALUES('cris');


/* DEPORTE */
CREATE OR REPLACE PROCEDURE sp_insert_deporte (
    nombre IN VARCHAR,    
    imagen IN VARCHAR,    
    color IN VARCHAR    
)
AS 
BEGIN
    INSERT INTO DEPORTE(nombre, imagen, color) 
    VALUES (nombre, imagen, color);
    COMMIT;
END;


CREATE OR REPLACE PROCEDURE sp_delete_deporte (
    id IN NUMBER    
)
AS 
BEGIN
    DELETE FROM DEPORTE WHERE idDeporte = id;
    COMMIT;
END;

CREATE OR REPLACE PROCEDURE sp_update_deporte (
    imagen IN VARCHAR,    
    color IN VARCHAR,
    idDeporte IN NUMBER
)
AS 
BEGIN
    UPDATE DEPORTE SET imagen = imagen, color = color WHERE idDeporte = idDeporte;
    COMMIT;
END;


commit;


/* END DEPORTE */

select * from DEPORTE;

EXECUT sp_delete_deporte(25);


/*DROP TABLE PREDICCION;
DROP TABLE DETALLE_EVENTO;
DROP TABLE EQUIPO;
DROP TABLE EVENTO;
DROP TABLE DETALLE_CLIENTE;
DROP TABLE JORNADA;
DROP TABLE RECOMPENSA;
DROP TABLE MENSAJE;
DROP TABLE CHAT;
DROP TABLE ADMINS;
DROP TABLE CLIENTE;
DROP TABLE TEMPORADA;
DROP TABLE MEMBRESIA;
DROP TABLE DEPORTE;
DROP TABLE FASE;
*/
