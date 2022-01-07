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
        <select class="form-select hfin" @change="handleChange">
            <option value="" selected disabled hidden>Fin</option>
            <option data-time="09:00" id="9" value="9">09:00</option>
            <option data-time="10:00" id="10" value="10">10:00</option>
            <option data-time="11:00" id="11" value="11">11:00</option>
            <option data-time="12:00" id="12" value="12">12:00</option>
            <option data-time="13:00" id="13" value="13">13:00</option>
            <option data-time="14:00" id="14" value="14">14:00</option>
            <option data-time="15:00" id="15" value="15">15:00</option>
            <option data-time="16:00" id="16" value="16">16:00</option>
            <option data-time="17:00" id="17" value="17">17:00</option>
            <option data-time="18:00" id="18" value="18">18:00</option>
            <option data-time="19:00" id="19" value="19">19:00</option>
            <option data-time="20:00" id="20" value="20">20:00</option>
            <option data-time="21:00" id="21" value="21">21:00</option>
            <option data-time="22:00" id="22" value="22">22:00</option>
        </select>
        <button class="btn btn-dark text-white" v-on:click="search()" :disabled="apply">Aplicar</button>   
        </div>
        <Suspense>
            <reservation-list :key="updateList" 
            :dateSearch="{
                date: this.date,
                hini: this.hini,
                hfin: this.hfin,
                isActivated: this.isActivated
                }"
            />
        </Suspense>
    </div>
</template>
<style>
    .hfin>:disabled{
        color:rgba(160, 7, 7, 0.712);
    }
    select:focus > option:checked { 
        background: rgb(32, 117, 53) !important;
        color:white;
    }
</style>
<script>
import { useToast } from "vue-toastification";
import Datepicker from 'vue3-datepicker/dist/vue3-datepicker.esm'
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
            updateList: 0,
            isActivated: true,
            apply: true
        };
       
    },
    methods: {

        handleChange(e){
            
            this.hfin = e.target.options[e.target.options.selectedIndex].dataset.time
            if(this.hini && this.hfin){
                this.apply = false
            }
        },

        hIni(e){

            this.hini = e.target.options[e.target.options.selectedIndex].text

            this.disabledHour()

            if(this.hini && this.hfin){
                this.apply = false
            }
        },

        search(){

            this.isActivated = true

            const toastr = useToast();

            var day = this.picked.getDate()

            var month = this.picked.getMonth() +1

            var year = this.picked.getFullYear()

            this.date = day+'/'+month+'/'+year

            if(!this.date && !this.hini && !this.hfin){

                    toastr.error("Insert your date reservation", {
                        timeout: 1500
                    });

            }else{

                var verifyHour = this.verifyTime()

                if(verifyHour == false){

                    toastr.error("These hours have already passed", {
                        timeout: 2500
                    });
                }else{

                    localStorage.setItem('date',this.date)

                    localStorage.setItem('hini',this.hini)

                    localStorage.setItem('hfin',this.hfin)

                    this.isActivated = false
                }
            }

            this.updateList += 1;
        },
        disabledHour(){
            var hiniTime = this.hini.split(":")

            var hiniNumber = parseInt(hiniTime[0])

            for(var i = 9;i<23;i++){
                var id = document.getElementById(i).value
                if(id <= hiniNumber){
                    var element = document.getElementById(i)
              
                    element.setAttribute('disabled','')
                }
            }
        },
        verifyTime(){

            var currentTime = new Date();

            var currentHour = currentTime.getHours();

            var currentDay = currentTime.getDate();

            var currentMonth = currentTime.getMonth()+1;

            var currentYear = currentTime.getFullYear();

            var day = this.date.split("/")

            var dayNumber = parseInt(day[0]);

            var month = this.date.split("/")

            var monthNumber = parseInt(month[1]);

            var year = this.date.split("/")

            var yearNumber = parseInt(year[2]);

            var hiniTime = this.hini.split(":")

            var hiniNumber = parseInt(hiniTime[0])

            var hfinTime = this.hfin.split(":")

            var hfinNumber = parseInt(hfinTime[0])

            var currentDate = new Date(currentYear,currentMonth,currentDay,currentHour,0,0)

            var reservationDateHini = new Date(yearNumber,monthNumber,dayNumber,hiniNumber,0,0)

            var reservationDateHfin = new Date(yearNumber,monthNumber,dayNumber,hfinNumber,0,0)

            if(reservationDateHini.valueOf()<= currentDate.valueOf()&&reservationDateHfin.valueOf()<= currentDate.valueOf()){

                return false

            }else{

                return true
            }
        }
    }
})
</script>
