<template>
    <footer>
        <div class="container">
            <div class="row">
            <div class="col-md-4 footer-column">
                <ul class="nav flex-column">
                    <li class="nav-item">
                        <span class="footer-title">{{ $t("RESERVATIONS") }}</span>
                    </li>
                </ul>
            </div>
            <div class="col-md-4 footer-column">
                <ul class="nav flex-column">
                    <li class="nav-item">
                        <span class="footer-title">{{ $t("COMPANY") }}</span>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link color-ent" href="#">{{ $t("ABOUT_US") }}</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link color-ent" href="#">{{ $t("OUR_HISTORY") }}</a>
                    </li>
                </ul>
            </div>
            <div class="col-md-4 footer-column">
                <ul class="nav flex-column">
                    <li class="nav-item">
                        <span class="footer-title">{{ $t("CONTACT_AND_SUPPORT") }}</span>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link color-ent" href="/#/contact"><i class="fas fa-envelope"></i>{{ $t("CONTACT_US") }}</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link color-ent" href="#"><i class="fas fa-star"></i>{{ $t("GIVE_FEEDBACK") }}</a>
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
                    <a class="color-ent" href="#">{{ $t("PRIVACY_POLICY") }}</a>
                </li>
                <li class="list-inline-item">
                    <a class="color-ent" href="#">{{ $t("TERMS_OF_USE") }}</a>
                </li>
                </ul>
            </div>
            </div>
        </div>
    </footer>
</template>
<style>
.color-ent {
  color: #40916C !important;
}

.color-ent:hover {
  color: #74C69D;
  text-decoration: none;
}

::-moz-selection {
  background: #40916C;
  text-shadow: none;
}

::selection {
  background: #40916C;
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

.footer-column .nav-item {
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
  background-color: #40916C;
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