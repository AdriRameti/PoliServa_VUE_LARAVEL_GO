<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <div class="container-fluid">
      <a class="navbar-brand text-white" href="#">Poliserva</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav ms-md-auto gap-2">
          <li class="nav-item rounded">
            <a class="nav-link active text-white" aria-current="page" href="#"><i class="bi bi-house-fill me-2"></i>Home</a>
          </li>
          <li class="nav-item rounded">
            <a class="nav-link text-white" v-on:click="clearSport()"><i class="bi bi-people-fill me-2"></i>Reservar</a>
          </li>

          <li class="nav-item rounded" v-show="!token">
              <a class="nav-link active text-white" aria-current="page" href="/#/login"><i class="bi bi-house-fill me-2"></i>Login</a>
          </li>
          
          <li class="nav-item dropdown rounded" v-show="token">
            <a class="nav-link dropdown-toggle text-white" href="#" id="navbarDropdown" role="button" data-bs-toggle="dropdown" aria-expanded="false"><i class="bi bi-person-fill me-2"></i>Profile</a>
            <ul class="dropdown-menu dropdown-menu-end" aria-labelledby="navbarDropdown">
              <li><a class="dropdown-item" href="#">Account</a></li>
              <li>
                <hr class="dropdown-divider">
              </li>
              <li><a class="dropdown-item" href="#">Logout</a></li>
            </ul>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<style>

.nav-item {
  cursor: pointer;
}

.navbar-nav .nav-item:hover {
  background-color: rgba(180, 190, 203, 0.4);
}

</style>
<script>

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
        window.location.href="/#/reservation";
      }
    }
  },
  setup() {
    
  },
})
</script>
