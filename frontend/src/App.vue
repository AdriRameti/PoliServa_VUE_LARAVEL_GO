<template>
  <div id="nav">
    <Nav @relv="reloadView" v-show="$route.name != 'Otp'"/>
    <router-view :key="reload"/>
    <Footer/>
  </div>
</template>
<script>
import Nav from './components/Nav.vue'
import Footer from './components/Footer.vue'
export default ({
  components:{Nav,Footer},
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

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}

@import "~ag-grid-community/dist/styles/ag-grid.css";
@import "~ag-grid-community/dist/styles/ag-theme-alpine.css";

</style>
