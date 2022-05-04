<template>
    <footer>
        <div class="container">
            <div class="row">
            <div class="col-md-4 footer-column">
                <ul class="nav flex-column">
                    <li class="nav-item">
                        <span class="footer-title">Reservas</span>
                    </li>
                </ul>
            </div>
            <div class="col-md-4 footer-column">
                <ul class="nav flex-column">
                    <li class="nav-item">
                        <span class="footer-title">Compañía</span>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#">Sobre nosotros</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#">Nuestra Historia</a>
                    </li>
                </ul>
            </div>
            <div class="col-md-4 footer-column">
                <ul class="nav flex-column">
                    <li class="nav-item">
                        <span class="footer-title">Contacto & Soporte</span>
                    </li>
                    <li class="nav-item">
                        <span class="nav-link"><i class="fas fa-phone"></i>+34 969 69 69 69</span>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#"><i class="fas fa-comments"></i>Live chat</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#"><i class="fas fa-envelope"></i>Contact us</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#"><i class="fas fa-star"></i>Give feedback</a>
                    </li>
                </ul>
            </div>
            </div>

            <div class="text-center"><i class="fas fa-ellipsis-h"></i></div>
            
            <div class="row text-center">
            <div class="col-md-4 box">
                <span class="copyright quick-links">Copyright &copy; Poliserva
                </span>
            </div>
            <div class="col-md-4 box">
                <ul class="list-inline social-buttons">
                <li class="list-inline-item">
                    <a href="https://twitter.com/" target="_blank">
                    <i class="fab fa-twitter"></i>
                </a>
                </li>
                <li class="list-inline-item">
                    <a href="https://facebook.com/" target="_blank">
                    <i class="fab fa-facebook-f"></i>
                </a>
                </li>
                <li class="list-inline-item">
                    <a href="https://www.linkedin.com/" target="_blank">
                    <i class="fab fa-linkedin-in"></i>
                </a>
                </li>
                </ul>
            </div>
            <div class="col-md-4 box">
                <ul class="list-inline quick-links">
                <li class="list-inline-item">
                    <a href="#">Privacy Policy</a>
                </li>
                <li class="list-inline-item">
                    <a href="#">Terms of Use</a>
                </li>
                </ul>
            </div>
            </div>
        </div>
    </footer>
</template>
<style>
a {
  color: #6c757d;
}

a:hover {
  color: #EB5E28;
  text-decoration: none;
}

::-moz-selection {
  background: #EB5E28;
  text-shadow: none;
}

::selection {
  background: #EB5E28;
  text-shadow: none;
}

footer {
  padding: 2rem 0;
  background-color: #212529;
}

.footer-column:not(:first-child) {
  padding-top: 2rem;
}

@media (min-width: 768px) {
  .footer-column:not(:first-child) {
    padding-top: 0rem;
  }
}

.footer-column {
  text-align: center;
}

.footer-column .nav-item .nav-link {
  padding: 0.1rem 0;
}

.footer-column .nav-item span.nav-link {
  color: #6c757d;
}

.footer-column .nav-item span.footer-title {
  font-size: 14px;
  font-weight: 700;
  color: #fff;
  text-transform: uppercase;
}

.footer-column .nav-item .fas {
  margin-right: 0.5rem;
}

.footer-column ul {
  display: inline-block;
}

@media (min-width: 768px) {
  .footer-column ul {
    text-align: left;
  }
}

ul.social-buttons {
  margin-bottom: 0;
}

ul.social-buttons li a:active,
ul.social-buttons li a:focus,
ul.social-buttons li a:hover {
  background-color: #EB5E28;
}

ul.social-buttons li a {
  font-size: 20px;
  line-height: 40px;
  display: block;
  width: 40px;
  height: 40px;
  -webkit-transition: all 0.3s;
  transition: all 0.3s;
  color: #fff;
  border-radius: 100%;
  outline: 0;
  background-color: #1a1d20;
}

footer .quick-links {
  font-size: 90%;
  line-height: 40px;
  margin-bottom: 0;
  text-transform: none;
  font-family: Montserrat, "Helvetica Neue", Helvetica, Arial, sans-serif;
}

.copyright {
  color: white;
}

.fa-ellipsis-h {
  color: white;
  padding: 2rem 0;
}
</style>
<script>
export default {
    setup() {
        
    },
}
</script>