-- ROL
INSERT INTO ROL(name) VALUES ('admin');
INSERT INTO ROL(name) VALUES ('cliente');
 
-- ADMIN
INSERT INTO USUARIO(username, password, name, surname, tier, fecha_nacimiento, fecha_registro, email, photo, idRol)
VALUES ('cris','cris','Cristian','Gomez','GOLD',TO_TIMESTAMP('1998-10-29 00:00:00.000000000', 'YYYY-MM-DD HH24:MI:SS.FF'),
TO_TIMESTAMP('1998-10-29 00:00:00.000000000', 'YYYY-MM-DD HH24:MI:SS.FF'),'crisgomez029@gmail.com','localhost',1);
-- CLIENT
INSERT INTO USUARIO(username, password, name, surname, tier, fecha_nacimiento, fecha_registro, email, photo, idRol)
VALUES ('alex','alex','Alexander','Gomez','-',TO_TIMESTAMP('1998-10-29 00:00:00.000000000', 'YYYY-MM-DD HH24:MI:SS.FF'),
TO_TIMESTAMP('1998-10-29 00:00:00.000000000', 'YYYY-MM-DD HH24:MI:SS.FF'),'crisgomez029@gmail.com','localhost',2);

-- Tier 
INSERT INTO MEMBRESIA(nombre, precio) VALUES ('Gold', 900); 
INSERT INTO MEMBRESIA(nombre, precio) VALUES ('Silver', 450); 
INSERT INTO MEMBRESIA(nombre, precio) VALUES ('Bronze', 150); 

-- FASE 
INSERT INTO FASE(nombre) VALUES ('Activa'); 
INSERT INTO FASE(nombre) VALUES ('Calculo'); 
INSERT INTO FASE(nombre) VALUES ('Finalizada'); 

COMMIT;

SELECT * FROM MEMBRESIA;

/*
    INSERTS
*/

/* USUARIO */
CREATE OR REPLACE PROCEDURE sp_insert_usuario ( 
    v_username IN VARCHAR2,
    v_password IN VARCHAR2,
    v_name IN VARCHAR2,
    v_surname IN VARCHAR2,
    v_fecha_nacimiento IN VARCHAR2,
    v_photo IN VARCHAR2,
    v_email IN VARCHAR2
    )
AS
BEGIN
    INSERT INTO USUARIO(username, password, name, surname, tier,fecha_nacimiento, fecha_registro, email, photo, idRol)
    VALUES (v_username, v_password, v_name, v_surname,'-',
    (TO_TIMESTAMP(v_fecha_nacimiento, 'YYYY-MM-DD HH24:MI:SS.FF')),
    SYSTIMESTAMP,v_email,v_photo,2);
END;
       
CREATE OR REPLACE VIEW v_user AS
SELECT 
    idUsuario,username, name, surname, tier,fecha_nacimiento, email, photo, idRol
FROM 
    USUARIO;
    
SELECT * FROM v_user WHERE idUsuario = 2;

/**
*   Stored Procedured required
**/

CREATE OR REPLACE PROCEDURE sp_auth2 ( 
    v_username IN VARCHAR2,
    v_password IN VARCHAR2
    )
AS
BEGIN
    SELECT * FROM USUARIO WHERE username = v_username AND password = v_password;

END;


CREATE OR REPLACE PROCEDURE sp_auth ( 
    v_username IN VARCHAR2,
    v_password IN VARCHAR2
    )
AS
    rowU NUMBER(1) := 0;
BEGIN

        SELECT COUNT(username) INTO rowU FROM USUARIO WHERE USUARIO.username =
        v_username  AND USUARIO.password = v_password; 
        
        
        IF rowU = 0 THEN
            RAISE_APPLICATION_ERROR(-20111,'Login, Please check it');
        END IF;
        
        DBMS_OUTPUT.PUT_LINE('idUsuario' || (rowU));
        
END;

EXEC sp_auth('cris','cris');


SELECT * from USUARIO;

select standard_hash('password', 'MD5') from USUARIO;


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


/**
*   Triggers required
**/

CREATE OR REPLACE TRIGGER tr_auth 
BEFORE INSERT OR UPDATE ON USUARIO FOR EACH ROW

DECLARE
    v_username VARCHAR2(200);
BEGIN
    
    SELECT username INTO v_username FROM USUARIO WHERE username = :NEW.username FETCH FIRST 1 ROWS ONLY;

    /* Checking email */
    IF NOT REGEXP_LIKE (:NEW.email, '\s*\w+([-+.»]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*\s*') THEN 
        RAISE_APPLICATION_ERROR(-20111,'Invalid email Address, Please check it');
    END IF;
    
    IF LENGTH(:NEW.password) < 8 THEN
        RAISE_APPLICATION_ERROR(-20111,'Password 8, Please check it');
      --  RAISE_APPLICATION_ERROR(-20111,'Username already exists, Please check it');
    END IF;
    
    /* Unique username */
    IF  v_username = :NEW.username THEN
        RAISE_APPLICATION_ERROR(-20111,'Username, Please check it');
    END IF;
    
    /*
        > 8 characters
        at least one uppercase
        at least one lowercase
        at leaet one number        
    */
    
   
    
    /*IF NOT REGEXP_LIKE(:NEW.password, '^(?=\w*\d)(?=\w*[A-Z])(?=\w*[a-z])\S{8,}$') THEN
        RAISE_APPLICATION_ERROR(-20111,'Password, Please check it');
      --  RAISE_APPLICATION_ERROR(-20111,'Username already exists, Please check it');
    END IF;*/
    
EXCEPTION
    WHEN NO_DATA_FOUND THEN v_username:='';
END;

-- CLIENT
INSERT INTO USUARIO(username, password, name, surname, tier, fecha_nacimiento, fecha_registro, email, photo, idRol)
VALUES ('alexx2','alex416516541','Alexander','Gomez','-',TO_TIMESTAMP('1998-10-29 00:00:00.000000000', 'YYYY-MM-DD HH24:MI:SS.FF'),
SYSTIMESTAMP,'crisgomez029@gmail.com','localhost',2);


CREATE OR REPLACE TRIGGER tr_evento
BEFORE INSERT OR UPDATE ON EVENTO FOR EACH ROW
DECLARE
BEGIN
    IF :NEW.fecha_hora < sysdate THEN
        RAISE_APPLICATION_ERROR(-20111,'TIMESTAMP, Please check it');
    END IF;
    
END;





select LENGTH(password), password from USUARIO;

DELETE FROM CLIENTE WHERE idCliente = 4;
SELECT * from usuario;


commit;


/* END DEPORTE */

select * from DEPORTE;

EXECUT sp_delete_deporte(25);


DROP TABLE PREDICCION;
-- DROP TABLE DETALLE_EVENTO;
DROP TABLE RESULTADO;
-- DROP TABLE EQUIPO;
DROP TABLE EVENTO;
DROP TABLE DETALLE_USUARIO;
DROP TABLE JORNADA;
DROP TABLE RECOMPENSA;
DROP TABLE MENSAJE;
-- DROP TABLE ADMINS;
DROP TABLE USUARIO;
DROP TABLE TEMPORADA;
DROP TABLE MEMBRESIA;
DROP TABLE DEPORTE;
DROP TABLE FASE; 