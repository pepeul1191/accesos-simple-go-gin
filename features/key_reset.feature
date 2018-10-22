Feature: Validate Reset Key
  Validate Reset Key

  Scenario: Validate user
    Given Generar petición HTTP "key/reset" con headers
    Given Crear POST data key-reset sin errores
    When Ejecutar petición HTTP Form Data
    Then Se debe obtener un status code success 200