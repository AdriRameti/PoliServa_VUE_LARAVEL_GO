<template>
    <div id="about-main">
        <div class="jumbotron">
            <div class="jumbotron-inner">
                <div class="top-box">
                    <div class="content-box">
                        <h1>
                            Sobre Poliserva
                        </h1>
                        <p>
                            Poliserva es una plataforma de experiencia para desarrolladores de API web.  <br /> Nuestra misión es hacer que los desarrolladores sean productivos a través de la generación automática de código.
                        </p>
                    </div>
                </div>
            </div>
            <div class="img-layer-container">
                <div class="team-image" id="team-image">
                    <img src="https://apimatic.io/img/theme/aboutUs/images-1.png" class="img-about" />
                </div>

                <div class="circles-container">
                    <div class="img-1">
                        <img src="https://apimatic.io/img/theme/aboutUs/Circles-1-1.svg" />                        
                    </div>
                    <div class="img-2">
                        <img src="https://apimatic.io/img/theme/aboutUs/Circles-2-1.svg" />                        
                    </div>                    
                </div>           
            </div>
        </div>
        <div class="story-container">
            <div class="need-for-dx-container">
                <h3 class="text-center">
                    Necesidad de Poliserva
                </h3>
                <p>
                    Las interfaces de programación de aplicaciones (API) han conquistado el mundo y ahora son el estándar de facto de la comunicación de software. Casi todos los productos de software hoy en día consumen API. El modelo de negocio de numerosas empresas de todo el mundo se basa en el consumo de sus API. Los proveedores de API, por lo tanto, se esfuerzan por aumentar las tasas de adopción de API gastando millones de dólares cada año para mejorar la experiencia del desarrollador. Esto generalmente se hace al proporcionar kits de desarrollo de software (SDK) y documentación de API al desarrollador que consume sus API. Sin embargo, desarrollar SDK y escribir documentación son tareas arduas, monótonas y propensas a errores. Es un proceso lento y cuesta mucho tiempo y dinero.
                </p>
                <div class="img-container">
                    <img src="https://apimatic.io/img/theme/aboutUs/dxFlow.svg" alt="apimatic developer experience process" class="img-responsive"/>
                </div>
            </div>
            <div class="container-divider"></div>
            <div class="our-tech-container">
                <h3 class="text-center">
                    Nuestras Tecnologías
                </h3>
                <p>
                    Nuestro motor de generación de código permite a los proveedores de API generar SDK para sus API en cuestión de minutos y por una fracción del costo. Proporcionamos herramientas como nuestro editor de API y el transformador de API para ayudar aún más a los proveedores de API a minimizar el tiempo necesario para enviar SDK de excelente calidad a los desarrolladores que utilizan sus API. Nuestro motor de generación de código también es capaz de generar documentación sucinta y sin errores para API y SDK, ambos. La documentación de los SDK incluye capturas de pantalla dinámicas que detallan las instrucciones de uso adaptadas a la API específica del proveedor y también fragmentos de código que muestran ejemplos de uso. Como guinda del pastel, proporcionamos portales DX bellamente diseñados para encapsular esta documentación.
                </p>
                <div class="img-container">
                    <img src="https://apimatic.io/img/theme/cgaasIcons/cgaasProcess.gif" alt="apimatic code generation engine" class="img-responsive"/>
                </div>
            </div>
            <div class="container-divider"></div>
            <div class="origin-story-container">
                <h3 class="text-center">
                    Nuestra historia
                </h3>
                <p>
                    Mientras realizaban un trabajo de investigación para sus trabajos finales del instituto IES L'Estació en 2022, nuestros fundadores se encontraron con uno de los puntos débiles de las reservas de pistas online: el responsivo de la página y la dificultat de interacción del usuario. Se dieron cuenta de que los administradores de las páginas web de reservas no cuidaban el trato con el cliente haciendo que estos añadiera calificaciones negativas a la página.
                </p>
                <p>
                    Estaba tan claro que, aunque los administradores sabían que había malas valoraciones, no iban a mejorar la aplicación que nos pusimos manos a la obra para empezar el desarrollo de la aplicación y poder presentarla como futuro proyecto al ayuntamiento de Ontinyent.
                </p>
            </div>
            <div class="container-divider"></div>
            <div class="today-container">
                <h3 class="text-center">
                    Avances
                </h3>
                <p>
                   Poliserva ha recorrido un corto camino desde su creación hace . Pero esto no impide que el proyecto crezca con nuevos y fuertes valores los cuales son:
                </p>
                <ul>
                    <li>Cuidar el medio ambiente, cuidar nuestro futuro</li>
                    <li>Transparencia con el cliente</li>
                    <li>Crecer como persona y con el equipo</li>
                    <li>Tener visión de futuro</li>

                </ul>
            </div>
            <div class="container-divider"></div>
        </div>
    </div>
</template>
<style>
@import url(https://fonts.googleapis.com/css?family=Open+Sans:400,300,700,400italic);
.product-navbar {
  height: 40px;
  width: 100%;
  background: #090b0f;
  border-bottom: 1px solid #090b0f;
  font-family: "Open Sans", sans-serif;
  font-size: 14px;
  line-height: 1.42857143;
  color: #333;
}
.product-navbar .product-navbar-centered {
  max-width: 90vw;
  height: 20px;
  margin: 0 5vw;
}
.product-navbar .product-navbar-centered ul {
  margin: 0;
  padding: 0;
}
.product-navbar .product-navbar-centered ul li {
  float: left;
  list-style: none;
  margin-right: 30px;
  padding-top: 9px;
}
.product-navbar .product-navbar-centered ul li a {
  color: #6f7782;
  font-size: 12px;
  font-weight: 700;
  text-decoration: none;
  text-transform: uppercase;
  text-shadow: none !important;
}

.hero {
  background: #171b23;
  padding-left: 0;
  padding-right: 0;
}
.hero .container {
  margin: 0 5vw;
  padding: 0 0vw;
  max-width: 90vw;
}
.hero .navbar-brand {
  color: #fff;
}
.hero .navbar-brand img {
  margin-right: 7px;
  margin-top: -7px;
}

#about-main {
  font-family: "Open Sans", sans-serif;
  margin: 0 auto;
}
#about-main .row {
  margin-left: 0;
  margin-right: 0;
}
#about-main .jumbotron {
  width: 100%;
  height: 530px;
  background: #31705448;
  position: relative;
  margin: 0;
  z-index: 3;
  border-radius: 0px;
}
#about-main .jumbotron .img-layer-container {
  position: relative;
  max-width: 1100px;
  margin: 0 auto;
  display: flex;
  justify-content: center;
}
#about-main .jumbotron .img-layer-container img {
  max-width: 95vw;
}
#about-main .jumbotron .img-layer-container .team-image {
  position: absolute;
}
#about-main .jumbotron .img-layer-container .team-image img {
  position: relative;
  top:-52px;
  z-index: 3;
  max-height: 90% !important;
}


#about-main .jumbotron .img-layer-container .circles-container {
  position: absolute;
  max-width: 1100px;
  margin: 0 auto;
  display: flex;
  justify-content: center;
}
#about-main .jumbotron .img-layer-container .circles-container .img-1 {
  position: absolute;
  z-index: 2;
}
#about-main .jumbotron .img-layer-container .circles-container .img-2 {
  position: absolute;
  z-index: 5;
}
#about-main .jumbotron .jumbotron-inner {
  max-width: 1350px;
  margin: 0 auto;
}
#about-main .jumbotron .jumbotron-inner a {
  text-decoration: none;
}
#about-main .jumbotron .jumbotron-inner .top-box {
  max-width: 950px;
  margin: 0 auto;
}
#about-main .jumbotron .jumbotron-inner .top-box .content-box {
  margin: 0 50px 35px;
}
#about-main .jumbotron .jumbotron-inner .top-box .content-box h1 {
  text-align: center;
  color: #212529;
  font-size: 40px;
  font-weight: bold;
  margin-bottom: 20px;
  padding-top: 20px;
}
#about-main .jumbotron .jumbotron-inner .top-box .content-box p {
  text-align: center;
  color: #212529af;
  font-size: 18px;
  font-weight: 500;
  margin: 0;
  margin-bottom: 15px;
}
@media (max-width: 1401px) {
  #about-main .jumbotron {
    height: 595px !important;
  }
  #about-main .jumbotron .jumbotron-inner {
    max-width: 900px;
  }
  #about-main .jumbotron .jumbotron-inner .top-box {
    max-width: 900px;
  }
  #about-main .jumbotron .jumbotron-inner .top-box .content-box {
    /* margin-top: 30px; */
    margin-bottom: 50px;
  }
  #about-main .jumbotron .jumbotron-inner .top-box .content-box h3 {
    font-size: 35px;
    margin: 5px 0px 20px 0px;
    color: #6f7782;
    font-weight: 300;
  }
}
@media (max-width: 1000px) {
  #about-main .jumbotron .img-layer-container {
    width: 100%;
  }
}
@media (max-width: 815px) {
  #about-main .jumbotron .jumbotron-inner {
    max-width: 600px;
  }
}
@media (max-width: 640px) {
  #about-main .jumbotron {
    height: 400px !important;
  }
  #about-main .jumbotron .jumbotron-inner {
    max-width: 550px;
    padding: 0 20px;
  }
  #about-main .jumbotron .img-1,
#about-main .jumbotron .img-2,
#about-main .jumbotron .team-image {
    display: none;
  }
}
@media (max-width: 420px) {
  #about-main .jumbotron .jumbotron-inner {
    max-width: 600px;
    padding: 0 20px;
  }
}
@media (max-width: 400px) {
  #about-main .jumbotron .jumbotron-inner .top-box .content-box {
    margin: 30px 0;
  }
  #about-main .jumbotron .jumbotron-inner .top-box .content-box h1 {
    font-size: 35px;
  }
  #about-main .jumbotron .jumbotron-inner .top-box .content-box p {
    font-size: 16px;
  }
}
#about-main .story-container {
  height: 100%;
  width: 100%;
  margin-top: 150px;
}
@media (max-width: 640px) {
  #about-main .story-container {
    margin-top: 80px;
  }
}
#about-main .story-container .container-divider {
  margin-top: 100px;
}
@media (max-width: 640px) {
  #about-main .story-container .container-divider {
    margin-top: 50px;
  }
}
#about-main .story-container .need-for-dx-container,
#about-main .story-container .our-tech-container,
#about-main .story-container .origin-story-container,
#about-main .story-container .today-container {
  max-width: 950px;
  margin: 0 auto;
  padding: 0 20px;
}
@media (max-width: 815px) {
  #about-main .story-container .need-for-dx-container,
#about-main .story-container .our-tech-container,
#about-main .story-container .origin-story-container,
#about-main .story-container .today-container {
    width: 100%;
  }
}
#about-main .story-container .need-for-dx-container .img-container,
#about-main .story-container .our-tech-container .img-container,
#about-main .story-container .origin-story-container .img-container,
#about-main .story-container .today-container .img-container {
  width: 100%;
  display: flex;
  justify-content: center;
}
#about-main .story-container .need-for-dx-container .img-container img,
#about-main .story-container .our-tech-container .img-container img,
#about-main .story-container .origin-story-container .img-container img,
#about-main .story-container .today-container .img-container img {
  margin: 50px auto;
}
#about-main .story-container .need-for-dx-container h3,
#about-main .story-container .our-tech-container h3,
#about-main .story-container .origin-story-container h3,
#about-main .story-container .today-container h3 {
  color: #225540;
}
#about-main .story-container .need-for-dx-container p,
#about-main .story-container .our-tech-container p,
#about-main .story-container .origin-story-container p,
#about-main .story-container .today-container p {
  font-size: 15px;
  margin-top: 20px;
  color: #676a6c;
}
@media (max-width: 640px) {
  #about-main .story-container .need-for-dx-container p,
#about-main .story-container .our-tech-container p,
#about-main .story-container .origin-story-container p,
#about-main .story-container .today-container p {
    font-size: 13px;
  }
}
#about-main .story-container .need-for-dx-container ul,
#about-main .story-container .our-tech-container ul,
#about-main .story-container .origin-story-container ul,
#about-main .story-container .today-container ul {
  margin-top: 30px;
  list-style-image: url("https://apimatic.io/img/theme/aboutUs/bullet.svg");
}
#about-main .story-container .need-for-dx-container ul li,
#about-main .story-container .our-tech-container ul li,
#about-main .story-container .origin-story-container ul li,
#about-main .story-container .today-container ul li {
  line-height: 25px;
  padding-left: 5px;
  color: #676a6c;
}
</style>