Feature: Change Activation Key
  Change Activation Key

  Scenario: Change Activation user
    Given Generar petición HTTP "key/activation/reset_by_user_id" con headers
    Given Crear POST data key-activation-reset sin errores
    When Ejecutar petición HTTP Form Data
    Then Se debe obtener un status code success 200