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
                    
                    <button class="btn btn-primary btn-lg btn-block" v-on:click="check2fa()" :disabled="otp.length != 6">Send OTP</button>

                    </div>
                </div>
                </div>
            </div>
        </div>
  </section>
</template>
<script>

import { useStore } from 'vuex';
import Constant from "../../Constant";

export default({


    data() {
        return {
            otp: ''
        }
    },
    methods: {
        check2fa() {
            this.store.dispatch("user/" + Constant.CHECK2FA, {'otp': this.otp, 'from': 'otp'});
        }
    },
    setup() {
        const store = useStore();

        return { store }
    },
})
</script>

<style lang="scss">
</style>