<template>
    <div>
        <Login v-show="login"/>
        <Register v-show="register"/>
        <p class="text-center" v-show="register">{{ $t("HAVE_ACCOUNT") }} <a v-on:click="changeLogin()">{{ $t("SIGN_IN") }}</a></p>
        <p class="text-center" v-show="login">{{ $t("NO_ACCOUNT") }} <a v-on:click="changeRegister()">{{ $t("REGISTER") }}</a></p>
    </div>
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