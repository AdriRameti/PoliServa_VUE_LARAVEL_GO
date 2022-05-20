<template :key="update">
    <div>
        <div class="pistas d-flex flex-wrap p-3 justify-content-center m-3">
            <h1 v-if="this.courts == 0">{{ $t("NO_COURTS") }}</h1>
            <div class="card p-3 card-reser" v-for="court in this.courts" :key="court.id">
                <img class="reser-img" :src="court.Sport.img"/>
                <h1 class="text-center"> {{ court.Sport.name }}</h1>
                <p class="reser-descrip">{{ $t("PRICE_PER_HOUR") }}: {{ court.price_h }}</p>
                <p class="text-danger" v-show="dateSearch.isActivated == true">{{ $t("ENTER_YOUR_RESERVATION_DATE") }}</p>
                <button id="confirm" class="btn-reservation btn" data-bs-target="#modalReservation" data-bs-toggle="modal" :disabled="dateSearch.isActivated" v-on:click="saveId(court.id,court.price_h)">{{ $t("CONFIRM_RESERVATION") }}</button>
            </div>
        </div>
        <div class="modal fade" id="modalReservation">
            <div class="modal-dialog modal-dialog-centered">
                <div class="modal-content">
                    <div class="modal-header">
                        <h4 class="modal-title">{{ $t("CONFIRM_RESERVATION") }}</h4>
                        <button type="button" class="btn-modal close btn" data-bs-dismiss="modal">
                            <span aria-hidden="true" class="closed-key">X</span>
                        </button>
                    </div> 
                    <div class="modal-body"> 
                        <p>
                            {{ $t("DATE") }}: {{ dateSearch.date }} <br/><br/>
                            {{ $t("INITIAL_HOUR") }}: {{ dateSearch.hini }} <br/><br/>
                            {{ $t("END_HOUR") }}: {{ dateSearch.hfin }}
                        </p>
                        <a class="btn btn-danger text-white m-5" data-bs-dismiss="modal">{{ $t("CANCEL") }}</a>
                        <a class="btn btn-success text-white m-5" v-on:click="reservar()">{{ $t("CONFIRM") }}</a>
                    </div>
                </div>
            </div>
        </div>
        <stripe-checkout v-if="sessionId!=null"
        ref="checkoutRef"
        mode="payment"
        :pk="publishableKey"
        :sessionId="sessionId"
        />
    </div>
</template>
<style>
    .btn-reservation{
        background-color: #40916C;
    }
    .modal-header{
        background-color: #40916C;
        color:white;
    }
    .btn-modal{
        background-color: #1B4332;
    }
    .closed-key{
        color:white;
    }
    .modal-body{
        text-align: center;
    }
</style>
<script>

import { useGetAllCourts, useGetSportCourts,useGetDateReservation, useCreateReservation, useInsertReservation } from '../composables/courts/useGetAllCourts';
import { useToast } from "vue-toastification";
import { useStore } from "vuex";
import { StripeCheckout } from '@vue-stripe/vue-stripe';
import {useGetSessionId} from '../composables/reservations/useReservations';

export default ({
    components:{
     StripeCheckout,
    },
    data(){
        return{
            courts:[],
            date: "",
            hini: "",
            hfin: "",
            publishableKey:'pk_test_51KxbOyDYuFjRSXysBDcEnM12LTgFfpggzLvChDmXb76r6L2JF5L3LdtMLddXO3zanIgE87i3tNdoBj7doimMiLQ600bchwv1Iz',
            sessionId:null
        }
    },
    props:['dateSearch'],
    methods: {
        async reservar(){

            const toastr = useToast();

            if(localStorage.getItem('token') && this.store.state.user.user){

                var totalPrices = this.calculateHour();

                var idCourt = parseInt(localStorage.getItem('idCourt'));
            
                var obj = {
                id_court:idCourt,
                date: "09/05/2022",
                hini: this.dateSearch.hini,
                hfin: this.dateSearch.hfin,
                total_price:totalPrices
                }
                localStorage.setItem("total_price", totalPrices);
                const reservation = await useInsertReservation(obj)
                
                if(reservation.data !=1){

                    toastr.error("Had an error processing your reservation", {
                        timeout: 2500
                    });

                }else{
                    await this.getSessionId();
                    if(this.sessionId!=null){
                        await this.submit();
                    }

                    document.getElementById('modalReservation').click();

                    // this.$router.push({name: 'Home'});
                    
                }

            }else{
                document.getElementById('modalReservation').click();

               this.$router.push({name: 'Login'});
            }
        },
        async getSessionId(){
            const sessionId = await useGetSessionId();
            this.sessionId = sessionId;
        },
        submit () {
            // You will be redirected to Stripe's secure checkout page
            this.$refs.checkoutRef.redirectToCheckout();
        },
        calculateHour(){

            var priceHour = localStorage.getItem('priceHour')

            var hiniTime = this.dateSearch.hini.split(":")

            var hiniNumber = parseInt(hiniTime)

            var hfinTime = this.dateSearch.hfin.split(":")

            var hfinNumber = parseInt(hfinTime)

            var totalHoras = hfinNumber - hiniNumber

            var totalPrice = (totalHoras * priceHour) * 100;

            return totalPrice;
        },
        saveId(id,price){

            localStorage.setItem('idCourt',id);

            localStorage.setItem('priceHour',price)

        }
    },
    async setup() {

        var store = useStore();
    
        var sl = localStorage.getItem('slug')

        var dat = localStorage.getItem('date')

        var hin = localStorage.getItem('hini')

        var hfi = localStorage.getItem('hfin')

        var arr = []
    
        if (dat && hin && hfi){

            arr.push(dat)

            arr.push(hin)

            arr.push(hfi)

            if (sl) {

                arr.push(sl)

            }

            const { courts, count } = await useGetDateReservation(arr);

            localStorage.removeItem('date')

            localStorage.removeItem('hini')

            localStorage.removeItem('hfin')  
               
            localStorage.removeItem('slug') 
            return { courts, count, store }

        }else if (sl){

            const { courts, count } = await useGetSportCourts(sl);

            return { courts, count, store };

        } else {

            const { courts, count } = await useGetAllCourts();

            return { courts, count, store };
        }
    }
})

</script>
