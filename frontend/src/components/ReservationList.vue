<template :key="update">
    <div>
        <div class="pistas d-flex flex-wrap p-3 justify-content-center m-3">
            <h1 v-if="this.courts == 0">No hay pistas disponibles en este horario</h1>
            <div class="card p-3 card-reser" v-for="court in this.courts" :key="court.id">
                <img :src="court.Sport.img"/>
                <h1> {{ court.Sport.name }}</h1>
                <p class="reser-descrip">Precio hora: {{ court.price_h }}</p>
                <p class="text-danger" v-show="dateSearch.isActivated == true">Introduzca la fecha de su reserva</p>
                <button id="confirm" class="btn btn-dark" data-bs-target="#modalReservation" data-bs-toggle="modal" :disabled="dateSearch.isActivated">Confirmar Reserva</button>
            </div>
        </div>
        <div class="modal fade" id="modalReservation">
            <div class="modal-dialog modal-dialog-centered">
                <div class="modal-content">
                    <div class="modal-header">
                        <h4 class="modal-title">Confirm your reservation</h4>
                        <button type="button" class="close btn btn-dark" data-bs-dismiss="modal">
                            <span aria-hidden="true">X</span>
                        </button>
                    </div> 
                    <div class="modal-body"> 
                        <p>
                            Fecha: {{ dateSearch.date }} <br/>
                            Hora Inicio: {{ dateSearch.hini }} <br/>
                            Hora Final: {{ dateSearch.hfin }} <br/>
                        </p>
                        <a class="btn btn-danger text-white m-5" data-bs-dismiss="modal">Cancelar</a>
                        <a class="btn btn-success text-white m-5" @click="reservar()">Confirmar</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>

import { onMounted } from '@vue/runtime-core';
import { useGetAllCourts, useGetSportCourts,useGetDateReservation, useCreateReservation } from '../composables/courts/useGetAllCourts';

export default ({
    data(){
        return{
            courts:[],
            date: "",
            hini: "",
            hfin: ""

        }
    },
    props:['dateSearch'],
    methods: {
        reservar(){
            // const reservation = useCreateReservation();
            
        }
    },
    async setup() {
    
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
            console.log(courts)
            localStorage.removeItem('date')
            localStorage.removeItem('hini')
            localStorage.removeItem('hfin')     
            localStorage.removeItem('slug') 
            return { courts, count }

        }else if (sl){

            const { courts, count } = await useGetSportCourts(sl);
            console.log(courts);
            return { courts, count };

        } else {

            const { courts, count } = await useGetAllCourts();
            console.log(courts);
            return { courts, count };
        }
    }
})

</script>
