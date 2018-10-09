require 'httparty'
require 'json'

Given("Generar petición HTTP {string} con headers") do |url|
  @url = CONSTANTS[:BASE_URL] + url
  @headers = {
    CONSTANTS[:CSRF][:key] => CONSTANTS[:CSRF][:value],
    'Content-Type' => 'application/x-www-form-urlencoded',
    'charset' => 'utf-8'
  }
end

When("Ejecutar petición HTTP") do
  @response = HTTParty.post(
    @url,
    headers: @headers,
    body: 'data=' + @data.to_json,
  )
end

Then("Se debe obtener un status code success {int}") do |status_code|
  if SHOW_RESPONSE
    puts @response.code
    puts @response.body
  end
  expect(@response.code).to be == status_code
end

Then("Se debe obtener el id mongo generado") do
  if @response.code != 200
    fail('Error, se obtuvo un response code 404')
  else
    rpta = JSON.parse(@response.body)
    mongo_id = rpta['mensaje'][1]
    expect(mongo_id.length).to be == 24
  end
end

Then("No se debe obtener el id generado") do
  if @response.code != 200
    fail('Error, se obtuvo un response code 404')
  else
    rpta = JSON.parse(@response.body)
    nuevos = rpta['mensaje'].length
    expect(nuevos).to be == 1
  end
end
