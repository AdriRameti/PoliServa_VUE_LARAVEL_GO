<template>
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
  </section>
</template>

<script>

import { useToast } from "vue-toastification";

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

      const toastr = useToast();
      var val = this.validateForm();

      if (val == "empty") {

        toastr.error("Email and password mustn't be empty", {
          timeout: 1500
        });
        return

      } else if (val == "email empty") {

        toastr.error("Email mustn't be empty", {
          timeout: 1500
        });
        return

      } else if (val == "password empty") {

        toastr.error("Password mustn't be empty", {
          timeout: 1500
        });
        return

      } else if (val == "email not valid") {

        toastr.error("Email must be like example@example.example", {
          timeout: 1500
        });
        return

      }

      toastr.success("Login success", {
        timeout: 1500
      });


    }
  },
  setup() {
    
  },
})
</script>
