<template>
    <section>
        <div class="container py-5 h-100">
            <div class="row d-flex justify-content-center align-items-center h-100">
                <div class="col-12 col-md-8 col-lg-6 col-xl-5">
                <div class="card shadow-2-strong" style="border-radius: 1rem;">
                    <div class="card-body p-4 text-center">

                    <h3 class="mb-3">Validate user</h3>

                    <div class="form-outline mb-4">
                        <input type="text" v-model="otp" class="form-control form-control-lg" placeholder="One Time Password"/>
                    </div>
                    
                    <button class="btn btn-primary btn-lg btn-block" v-on:click="validate()">Send OTP</button>

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
            otp: ''
        }
    },
    methods: {
        validate() {

            const toastr = useToast();
            
            if (this.otp.length == 0) {
                toastr.error("One Time Password mustn't be empty", {
                    timeout: 1500
                });
                return
            }

            console.log(this.otp)
        }
    },
    setup() {
        
    },
})
</script>

<style lang="scss">
</style>