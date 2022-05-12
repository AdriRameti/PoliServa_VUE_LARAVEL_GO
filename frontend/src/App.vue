<template>
  <div id="nav">
    <Nav @relv="reloadView" v-show="$route.name != 'Otp'"/>
    <router-view :key="reload"/>
    <Cookies class="cookies"/>
    <Footer/>
  </div>
</template>
<script>
import Nav from './components/Nav.vue'
import Footer from './components/Footer.vue'
import Cookies from './components/Cookies.vue'
export default ({
  components:{Nav, Footer, Cookies},
  data() {
    return {
      reload: 0,
    }
  },
  methods: { 
    reloadView() {
      this.reload += 1;
    }
  },
});
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  //text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
  position: absolute;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}

.cookies {
  position: relative;
}

@import "~ag-grid-community/dist/styles/ag-grid.css";
@import "~ag-grid-community/dist/styles/ag-theme-alpine.css";

</style>
