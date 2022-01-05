<template>
  <div id="nav">
    <Nav @relv="reloadView" :key="nav"/>
    <router-view :key="reload"/>
  </div>
</template>
<script>
import Nav from './components/Nav.vue'

export default ({
  components:{Nav},
  watch:{
    $route (to, from) {
      this.reloadNav()
    }
  },
  data() {
    return {
      reload: 0,
      nav: 0
    }
  },
  methods: { 
    reloadView() {
      this.reload += 1;
    },
    reloadNav() {
      this.nav += 1;
    }
  },
});
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
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
</style>
