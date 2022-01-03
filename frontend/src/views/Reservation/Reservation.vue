<template>
    <div>
        <div class="input-group">
            <Datepicker class="input-group-text" v-model="picked" />
            <select class="form-select" v-on:change="hIni($event)">
            <option value="" selected disabled hidden>Inicio</option>
            <option data-time="08:00">08:00</option>
            <option data-time="09:00">09:00</option>
            <option data-time="10:00">10:00</option>
            <option data-time="11:00">11:00</option>
            <option data-time="12:00">12:00</option>
            <option data-time="13:00">13:00</option>
            <option data-time="14:00">14:00</option>
            <option data-time="15:00">15:00</option>
            <option data-time="16:00">16:00</option>
            <option data-time="17:00">17:00</option>
            <option data-time="18:00">18:00</option>
            <option data-time="19:00">19:00</option>
            <option data-time="20:00">20:00</option>
            <option data-time="21:00">21:00</option>
        </select>
        <select class="form-select" @change="handleChange">
            <option value="" selected disabled hidden>Fin</option>
            <option data-time="09:00">09:00</option>
            <option data-time="10:00">10:00</option>
            <option data-time="11:00">11:00</option>
            <option data-time="12:00">12:00</option>
            <option data-time="13:00">13:00</option>
            <option data-time="14:00">14:00</option>
            <option data-time="15:00">15:00</option>
            <option data-time="16:00">16:00</option>
            <option data-time="17:00">17:00</option>
            <option data-time="18:00">18:00</option>
            <option data-time="19:00">19:00</option>
            <option data-time="20:00">20:00</option>
            <option data-time="21:00">21:00</option>
            <option data-time="22:00">22:00</option>
        </select>
        <a class="btn btn-dark text-white" v-on:click="search()">Aplicar</a>   
        </div>
        <Suspense>
            <reservation-list :key="updateList"/>
        </Suspense>
    </div>
</template>
<script>
import Datepicker from 'vue3-datepicker/dist/vue3-datepicker.esm'
import { Store, useStore } from "vuex"
import ReservationList from '../../components/ReservationList.vue'
export default({
    components: {
        Datepicker,
        ReservationList
    },
    data: ()=> {
        return{
            picked: new Date(),
            date: '',
            hini: '',
            hfin:'',
            filter: [],
            updateList: 0
        };
       
    },
    methods: {
        handleChange(e){
            this.hfin = e.target.options[e.target.options.selectedIndex].dataset.time
        },
        hIni(e){
            this.hini = e.target.options[e.target.options.selectedIndex].text
        },
        search(){
            var day = this.picked.getDate()
            var month = this.picked.getMonth() +1
            var year = this.picked.getFullYear()
            this.date = day+'/'+month+'/'+year

            if(!this.date && !this.hini && !this.hfin){
                console.log('Variables vacias');
            }else{
                localStorage.setItem('date',this.date)
                localStorage.setItem('hini',this.hini)
                localStorage.setItem('hfin',this.hfin)
            }
            this.updateList += 1;
            console.log(this.date,this.hini,this.hfin)
        }
    },
    setup(){
    }
})
</script>