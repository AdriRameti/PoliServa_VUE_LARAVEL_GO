<template>
<div>
  <section>
    <div class="container py-5 h-100">
      <div class="row d-flex justify-content-center align-items-center h-100">
        <div class="col-12 col-md-8 col-lg-6 col-xl-5">
          <div class="card shadow-2-strong" style="border-radius: 1rem;">
            <div class="card-body p-4 text-center">

              <h3 class="mb-3">Sign in</h3>

              <div class="form-outline mb-4">
                <input type="email" id="typeEmailX-2" v-model="email" class="form-control form-control-lg" placeholder="Email"/>
              </div>

              <div class="form-outline mb-4">
                <input type="password" id="typePasswordX-2" v-model="password" class="form-control form-control-lg" placeholder="Password"/>
              </div>

              <!-- Checkbox -->
              
              <button class="btn btn-primary btn-lg btn-block" v-on:click="login()">Login</button>

            </div>
          </div>
        </div>
      </div>
    </div>
    <div id="firebaseui-auth-container"></div>
  </section>
</div>
</template>

<script>

import { firebaseConfig } from '../firebaseConfig';
import firebase from 'firebase/compat/app';
import { useStore } from 'vuex';
import { useToast } from "vue-toastification";
import Constant from "../Constant";

firebase.initializeApp(firebaseConfig);
import * as firebaseui from 'firebaseui';
import 'firebaseui/dist/firebaseui.css';
import { onBeforeMount, onMounted } from '@vue/runtime-core';
export default({

  data() {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    validateForm() {
      var regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

      if (this.email.length == 0 && this.password.length == 0) {
        return "empty";
      } else if (this.email.length == 0) {
        return "email empty";
      } else if (this.password.length == 0) {
        return "password empty";
      }

      if (regex.test(this.email) == false) { 
        return "email not valid";
      }

    },
    login() {

      var val = this.validateForm();

      if (val == "empty") {

        this.toastr.error("Email and password mustn't be empty", {
          timeout: 1500
        });
        return

      } else if (val == "email empty") {

        this.toastr.error("Email mustn't be empty", {
          timeout: 1500
        });
        return

      } else if (val == "password empty") {

        this.toastr.error("Password mustn't be empty", {
          timeout: 1500
        });
        return

      } else if (val == "email not valid") {

        this.toastr.error("Email must be like example@example.example", {
          timeout: 1500
        });
        return

      }

      this.store.dispatch("user/" + Constant.LOGIN_USER, {mail: this.email, pass: this.password})

    }
  },
   setup() {
    const toastr = useToast();
    const store = useStore();
    const uiConfig = {
      signInFlow: 'popup',
      signinSuccessUrl:'http://localhost:4200/#/',
      signInOptions:[ 
        firebase.auth.GoogleAuthProvider.PROVIDER_ID,
        firebase.auth.GithubAuthProvider.PROVIDER_ID
      ],
      callbacks:{
        signInSuccessWithAuthResult: function(authResult,redirectUrl){
          const PayloadPrefix= authResult.additionalUserInfo.profile;
          const PrivatePayloadPrefix = authResult.user
          const payloadUser = {
            fullName:PayloadPrefix.name,
            id: PayloadPrefix.id,
            img: PayloadPrefix.picture,
            is_blocked: false,
            mail:PayloadPrefix.email,
            name: PayloadPrefix.given_name,
            role:'client',
            surnames:PayloadPrefix.family_name,
            token:PrivatePayloadPrefix.refreshToken,
            uuid:PrivatePayloadPrefix.uid,
            social:1
          };
          store.dispatch('user/'+Constant.SOCIAL_LOGIN_USER,payloadUser);
          return true;
        }
      }
    }
    
      let ui = firebaseui.auth.AuthUI.getInstance();
      if (!ui) {
          ui = new firebaseui.auth.AuthUI(firebase.auth());
      }
      onMounted(()=>{
        ui.start('#firebaseui-auth-container' ,uiConfig);
      })
      


    return { toastr, store }
  },
})
</script>
