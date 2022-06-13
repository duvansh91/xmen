# xmen

REST API que valida si un humano es un mutante.

# Como instalarlo
Para ejecutar el api localmente se requiere tener instalado `Docker`y `Docker Compose`.
* Ubicado en la raíz del proyecto, Ejecutar el siguiente comando `docker-compose up`. Esto creará un contenedor para la API y un contenedor para la base de datos (mongoDB).

# Como ejecutarlo
Utilizando un cliente (Postman, insomnia, etc) puede probar las siguientes rutas:
* Validar si un humano es mutante:
  * Ruta: `/mutant/`
  * Tipo de petición: `POST`
  * Body:  
  
  ```json 
      {
        "dna": ["ATGCAA", "CAGTGC", "TTATGT", "AGAAGG", "TCCCTA", "TCACTG"]
      }
  ```
    El campo `dna` será una secuencia de cadenas de ADN representadas como una matriz, donde si se encuentra más de una secuencia de cuatro letras
    iguales, de forma oblicua, horizontal o vertical se considerará un mutante.
    
    Posibles códigos de respuesta:
    
    * `200` OK -> Cuando si es mutante.
    
    * `403` Forbidden -> Cuando no es mutante.
    
    
* Ver estadisticas de las validaciones que se han realizado:
  * Ruta: `/stats/`
  * Tipo de petición: `GET`

  Respuesta:
  
  ```json 
      {
          "stats": {
            "count_mutant_dna": 1,
            "count_human_dna": 1,
            "ratio": 0.5
          }
      }
  ```
