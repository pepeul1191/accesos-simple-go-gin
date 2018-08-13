<!DOCTYPE html>
<html>
<head>
    <title>{{.title}}</title>
    <meta charset="UTF-8">
    <link rel="shortcut icon" href="{{.constants.BASE_URL}}favicon.ico">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    {{loadCSS .csss | raw}}
    <script type="text/javascript">
      var BASE_URL = "{{.constants.BASE_URL}}";
      var STATICS_URL  = "{{.constants.STATIC_URL}}";
      var MODULOS_JSON = JSON.parse('[{"url" : "contenidos/", "nombre" : "Contenidos"}, {"url" : "ubicaciones/", "nombre" : "Ubicaciones"}]');
      var ITEMS_JSON = JSON.parse('[{"subtitulo":"","items":[{"item":"Ubicaciones del Perú","url":"ubicaciones/#/ubicacion"},{"item":"Autocompletar","url":"ubicaciones/#/autocompletar"}]}]');
      var DATA = JSON.parse('{"mensaje":false,"titulo_pagina":"Gesti\u00f3n de Ubicaciones","modulo":"Ubicaciones"}');
      var CSRF = "PKBcauXg6sTXz7Ddlty0nejVgoUodXL89KNxcrfwkEme0Huqtj6jjt4fP7v2uF4L";
      var CSRF_KEY = 'csrf_val';
    </script>
</head>
<body>
  <label id="defaultTargetMensajes"></label>
  <!-- Inicio modal -->
  <button type="button" class="btn btn-primary btn-lg oculto" data-toggle="modal" data-target="#modal-container" id="btnModal">Launch demo modal</button>
  <div class="modal fade" id="modal-container" tabindex="-1" role="dialog" aria-labelledby="exampleModalLabel" aria-hidden="true">
    <div class="modal-dialog" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="exampleModalLabel">Modal title</h5>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">&times;</span>
          </button>
        </div>
        <div class="modal-body">
          ...
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
          <button type="button" class="btn btn-primary">Save changes</button>
        </div>
      </div>
    </div>
  </div>
  <!-- Fin modal -->
  <div id="app"></div>
  <!-- Handlebars Templates -->
  <script id="template" type="text/x-handlebars-template">
    {{ "{{> header}}" | raw }}
    {{ "{{> breadcrumb}}" | raw}}
    {{ "{{> contenido}}" | raw}}
    {{ "{{> footer}}" | raw}}
  </script>
  <script id="header-template" type="text/x-handlebars-template">
    <header>
      <a href="{{"{{BASE_URL}}"|raw}}">Inicio</a>
      <a href="{{"{{BASE_URL}}"|raw}}ayuda">Ayuda</a>
      <a href="{{"{{BASE_URL}}"|raw}}login/ver" class="pull-right">Pepe Valdivia</a>
      <a href="{{"{{BASE_URL}}"|raw}}login/cerrar" class="pull-right">Cerrar Sesión</a>
    </header>
  </nav>
  <!-- Fin Header -->
  </script>
  <script id="breadcrumb-template" type="text/x-handlebars-template">
    <nav>
      <h1><i class="fa fa-pencil h1" aria-hidden="true"></i>Gestor de Contenidos - COA</h1>
      {{"{{{menuModulos}}}" | raw }}
    </nav>
  </script>
  <script id="contenido-template" type="text/x-handlebars-template">
    <div id="body-app" class="row">
      <aside class="col-md-2">
        {{"{{{menuSubModulos}}}" | raw }}
      </aside>
      <section class="col-md-10" id="workspace">
        <!-- Inicio Yield-->
        {{ "{{> yield}}" | raw}}
        <!-- Fin Yield-->
      </section>
    </div>
  </script>
  <script id="footer-template" type="text/x-handlebars-template">
    <footer>
      <p>Powered by: <a href="http://softweb.pe/">Software Web Perú</a> © 2011-2018 </p>
    </footer>
  </script>