<template :key="update">
        <div class="pistas d-flex flex-wrap p-3 justify-content-center m-3">
            <div class="card p-3 card-reser" v-for="court in this.courts" :key="court.id">
                <img src="http://manzasport.com/wp-content/uploads/2018/04/Pista-de-padel-Paquito-Navarro-by-Manzasport-1.jpeg"/>
                <h1>Sector: {{ court.sector }}</h1>
                <p class="reser-descrip">Precio hora: {{ court.price_h }}</p>
            </div>
        </div> 
</template>
<script>

import { onMounted } from '@vue/runtime-core';
import { useGetAllCourts, useGetSportCourts,useGetDateReservation } from '../composables/courts/useGetAllCourts';

export default({
    data(){
        return{
            courts:[],
            update: 0
        }
    },
    async setup() {

        console.log('refresh');
        var sl = localStorage.getItem('slug')
        var dat = localStorage.getItem('date')
        var hin = localStorage.getItem('hini')
        var hfi = localStorage.getItem('hfin')
        var arr = []
        if (sl && dat && hin && hfi){
            arr.push(dat)
            arr.push(hin)
            arr.push(hfi)
            arr.push(sl)
            console.log('entra');
            const { courts, count } = await useGetDateReservation(arr);
            
            localStorage.removeItem('date')
            localStorage.removeItem('hini')
            localStorage.removeItem('hfin')
            // this.courts = courts;
            this.update += 1;
            return { courts, count };
        }else if (sl){
            const { courts, count } = await useGetSportCourts(sl);
            console.log(courts.value)
            // this.courts = courts;
            return { courts, count };
        } else {
            const { courts, count } = await useGetAllCourts();
            console.log(courts.value)
            // this.courts = courts;
            return { courts, count };
        }


    },
})
</script>
