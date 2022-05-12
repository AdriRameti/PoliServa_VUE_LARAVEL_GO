<template>
  <nav class="navbar navbar-expand-lg navbar-dark">
    <div class="container-fluid">
      <a class="navbar-brand text-white" href="#">Poliserva</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav ms-md-auto gap-2">
          <li class="nav-item rounded">
            <a class="nav-link active text-white" aria-current="page" href="#"><i class="bi bi-house-fill me-2"></i>{{ $t("HOME") }}</a>
          </li>
          <li class="nav-item rounded">
            <a class="nav-link text-white" v-on:click="clearSport()"><i class="bi bi-people-fill me-2"></i>{{ $t("RESERVE") }}</a>
          </li>

          <li class="nav-item rounded">
            <a class="nav-link text-white" v-show="store.state.user.user && store.state.user.user.role == 'admin'" href="/#/dashboard"><i class="bi bi-people-fill me-2"></i>{{ $t("DASHBOARD") }}</a>
          </li>

          <li class="nav-item rounded" v-show="!store.state.user.user">
              <a class="nav-link active text-white" aria-current="page" href="/#/login"><i class="bi bi-house-fill me-2"></i>{{ $t("LOGIN") }}</a>
          </li>

          <li class="nav-item dropdown rounded" v-show="store.state.user.user">
            <a class="nav-link dropdown-toggle text-white" href="#" id="navbarDropdown" role="button" data-bs-toggle="dropdown" aria-expanded="false"><i class="bi bi-person-fill me-2"></i>{{ $t("PROFILE") }}</a>
            <ul class="dropdown-menu dropdown-menu-end" aria-labelledby="navbarDropdown">
              <li><a class="dropdown-item" href="/#/profile">{{ $t("ACCOUNT") }}</a></li>
              <li>
                <hr class="dropdown-divider">
              </li>
              <li><a class="dropdown-item" v-on:click="logout()">{{ $t("LOGOUT") }}</a></li>
            </ul>
          </li>
          <li class="nav-item dropdown rounded">
            <select class="form-select" v-model="$i18n.locale">
              <option value="en">English</option>
              <option value="es">Español</option>
              <option value="val">Valencià</option>
            </select>
          </li>

        </ul>
      </div>
    </div>
  </nav>
</template>

<style>

.navbar{
  background-color: #1B4332;
}
.nav-item {
  cursor: pointer;
}

.navbar-nav .nav-item:hover {
  background-color: rgba(180, 190, 203, 0.4);
}

</style>
<script>

import { useToast } from "vue-toastification";
import { useStore } from 'vuex';
import Constant from "../Constant";

export default({
  data() {
    return {
      token: localStorage.getItem('token') ? true : false
    }
  },
  methods:{
    clearSport(slug) {

      localStorage.removeItem('slug');

      if (window.location.hash == '#/reservation') {
        this.$emit('relv');
      } else {
        this.$router.push({name: 'Reservation'})
      }
    },
    logout() {

      this.toastr.success("Logout success", {
          timeout: 1500
      });

      this.$router.push({name: 'Home'});

      this.store.dispatch("user/" + Constant.LOGOUT);

    },
    changeLang(lang) {
      this.store.dispatch(Constant.CHANGE_LANG, lang);
    }
  },
  setup() {
    const toastr = useToast();
    const store = useStore();

    return { toastr, store }
  }
})
</script>
