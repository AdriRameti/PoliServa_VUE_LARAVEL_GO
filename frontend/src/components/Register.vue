<template>
  <section>
    <div class="container py-5 h-100">
      <div class="row d-flex justify-content-center align-items-center h-100" v-if="showView==true">
        <div class="col-12 col-md-8 col-lg-6 col-xl-5">
          <div class="card shadow-2-strong" style="border-radius: 1rem;">
            <div class="card-body p-4 text-center">

              <h3 class="mb-3">Register</h3>

              <div class="form-outline mb-4">
                <input type="text" id="typeNameX-2" v-model="name" class="form-control form-control-lg" placeholder="Name"/>
              </div>

              <div class="form-outline mb-4">
                <input type="text" id="typeSurnamesX-2" v-model="surnames" class="form-control form-control-lg" placeholder="Surnames"/>
              </div>

              <div class="form-outline mb-4">
                <input type="email" id="typeEmailX-2" v-model="email" class="form-control form-control-lg" placeholder="Email"/>
              </div>

              <div class="form-outline mb-4">
                <input type="password" id="typePasswordX-2" v-model="password"  class="form-control form-control-lg" placeholder="Password"/>
              </div>

              <div class="progress" v-show="this.password.length>0">
                <div class="progress-bar" id="security"></div>
              </div>
              <br v-show="this.password.length>0"/>
              <div class="form-outline mb-4">
                <input type="password" id="typeRepeatPasswordX-2" v-model="repeatPassword" class="form-control form-control-lg" placeholder="Repeat Password"/>
              </div>           
              <!-- Checkbox -->
              
              <button class="btn btn-primary btn-lg btn-block" @click="register()">Register</button>

            </div>
          </div>
        </div>
      </div>
 <div class="row d-flex justify-content-center align-items-center h-100" v-if="showView==false">
        <div class="col-12 col-md-8 col-lg-6 col-xl-5">
          <div class="card shadow-2-strong" style="border-radius: 1rem;">
            <div class="card-body p-4 text-center">

              <h3 class="mb-3">Introduce the verify code</h3>

              <div class="form-outline mb-4">
                <input type="text" id="typeCodeX-2" v-model="code" class="form-control form-control-lg" placeholder="Insert code"/>
              </div>      
              <!-- Checkbox -->

              <button class="btn btn-primary btn-lg btn-block" @click="send()">Send</button>

            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import { useToast } from "vue-toastification";
import { useStore } from 'vuex';
import Constant from "../Constant";
export default ({

  data(){
    return{
      progress:"",
      name:"",
      surnames:"",
      email:"",
      password:"",
      repeatPassword:"",
      showView: true,
      code:''
    }
  },
  watch:{
    password(pass){
      if(pass.length<4){
        var element = document.getElementById("security");
        element.classList.remove("w-50")
        element.classList.remove("bg-danger")
        element.classList.remove("w-75")
        element.classList.remove("bg-warning")
        element.classList.remove("w-100")
        element.classList.remove("bg-success")
        element.classList.add("w-25");
        element.classList.add("bg-danger")
      }else if(pass.length>4 && pass.length<6){
        var element = document.getElementById("security");
        element.classList.remove("w-25")
        element.classList.remove("bg-danger")
        element.classList.remove("w-75")
        element.classList.remove("bg-warning")
        element.classList.remove("w-100")
        element.classList.remove("bg-success")
        element.classList.add("w-50");
        element.classList.add("bg-danger")
      }else if(pass.length>6 && pass.length<9){
        var element = document.getElementById("security");
        element.classList.remove("w-25")
        element.classList.remove("bg-danger")
        element.classList.remove("w-50")
        element.classList.remove("bg-danger")
        element.classList.remove("w-100")
        element.classList.remove("bg-success")
        element.classList.add("w-75");
        element.classList.add("bg-warning")
      }else if(pass.length>9){
        var element = document.getElementById("security");
        element.classList.remove("w-25")
        element.classList.remove("bg-danger")
        element.classList.remove("w-50")
        element.classList.remove("bg-danger")
        element.classList.remove("w-75")
        element.classList.remove("bg-warning")
        element.classList.add("w-100");
        element.classList.add("bg-success")
      }
    }
  },
  methods:{
    printSecurityPassword(){
      console.log('hola')
      if(this.password.length<4){
      var element = document.getElementById("security");
      element.classList.add("w-33");
      element.classList.add("bg-error")
      }
    },
    send(){
      var obj = {
        'name':this.name,
        'surnames':this.surnames,
        'mail':this.email,
        'password':this.password
      }
        if(this.code!= localStorage.getItem('verifyCode')){
          this.toastr.error("The current code is different to the code that we send", {
          timeout: 2500
        })
        }else{
          this.store.dispatch("user/" + Constant.REGISTER_USER, {info: obj})

        }
        

    },
    validateForm() {
      var regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

      if (this.name.length == 0 && this.surnames.length == 0 && this.mail.length == 0
         && this.password.length == 0 && this.repeatPassword.length == 0) {
        return "empty";
      }else if (this.name.length == 0) {
        return "name empty";
      }else if (this.surnames.length == 0) {
        return "surnames empty";
      } else if (this.email.length == 0) {
        return "email empty";
      } else if (this.password.length == 0) {
        return "password empty";
      }else if(this.repeatPassword.length == 0){
        return "repeatPassword"
      }
      if (regex.test(this.email) == false) { 
        return "email not valid";
      }

    },
    register(){

      var val = this.validateForm();

      if (val == "empty") {

        this.toastr.error("The register form mustn't be empty", {
          timeout: 2500
        });
        return

      } else if (val == "name empty") {

        this.toastr.error("Name mustn't be empty", {
          timeout: 2500
        });
        return

      } else if (val == "surnames empty") {

        this.toastr.error("Surnames mustn't be empty", {
          timeout: 2500
        });
        return

      } else if (val == "email empty") {

        this.toastr.error("Email mustn't be empty", {
          timeout: 2500
        });
        return

      } else if (val == "password empty") {

        this.toastr.error("Password mustn't be empty", {
          timeout: 2500
        });
        return

      } else if (val == "repeatPassword empty") {

        this.toastr.error("Repeat Password mustn't be empty", {
          timeout: 2500
        });
        return

      } else if (val == "email not valid") {

        this.toastr.error("Email must be like example@example.example", {
          timeout: 2500
        });
        return

      }
      if(this.password != this.repeatPassword){
        this.toastr.error("The passwords must be equals", {
          timeout: 2500
        });
        return
      }else{
        this.toastr.warning("Check your email and enter the verification code that we have sent you", {
          timeout: 5500
        });  
        this.showView = false;
        this.store.dispatch("user/" + Constant.VALIDATE_REGISTER, {mail: this.email})
      }

    }
  },
  setup() {
    const toastr = useToast();
    const store = useStore();
    return{toastr,store}
  },
})
</script>
