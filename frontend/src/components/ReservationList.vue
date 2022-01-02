<template>
        <div class="pistas d-flex flex-wrap p-3 justify-content-center m-3">
            <div class="card p-3 card-reser" v-for="court in courts" :key="court.id">
                <img src="http://manzasport.com/wp-content/uploads/2018/04/Pista-de-padel-Paquito-Navarro-by-Manzasport-1.jpeg"/>
                <h1>{{ court.id }}</h1>
                <p class="reser-descrip">Lorem Ipsum is simply dummy text of the printing and typesetting industry.</p>
            </div>
        </div> 
</template>
<script>

import { useGetAllCourts, useGetSportCourts,useGetDateReservation } from '../composables/courts/useGetAllCourts';

export default({
    async setup() {
        var sl = localStorage.getItem('slug')
        var dat = localStorage.getItem('date')
        var hin = localStorage.getItem('hini')
        var hfi = localStorage.getItem('hfin')
        var arr = []
        if (sl && dat && hin && hfi){
            arr.push(dat)
            arr.push(hin)
            arr.push(hfi)
            arr.push('tenis-5209')
            
            const { courts, count } = await useGetDateReservation(arr);
            
            return { courts, count };
        }else{
            const { courts, count } = await useGetSportCourts('tenis-5209');

            return { courts, count };
        }

    },
})
</script>
