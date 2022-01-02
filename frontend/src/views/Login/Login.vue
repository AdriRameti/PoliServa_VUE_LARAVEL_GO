<template>
    <Login v-show="login"/>
    <Register v-show="register"/>
    <p v-show="register">Ya tienes cuenta? <a v-on:click="changeLogin()">Login</a></p>
    <p v-show="login">No tienes cuenta? <a v-on:click="changeRegister()">Registrarte</a></p>
</template>
<script>
//import { defineComponent } from '@vue/composition-api'

import Login from '@/components/Login.vue';
import Register from '@/components/Register.vue';

export default({

    name: 'LoginView',
    components: {
        Login,
        Register
    },

    data() {
        return {
            login: true,
            register: false
        }
    },
    methods: {
        changeRegister() {
            this.register = true;
            this.login = false;
        },
        changeLogin() {
            this.login = true;
            this.register = false;
        }
    },
    setup() {
        
    },
})
</script>

<style lang="scss">
    a {
        cursor: pointer;
    }
</style>