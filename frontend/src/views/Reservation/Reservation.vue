<template>
    <div>
        <div class="input-group">
            <Datepicker class="input-group-text" v-model="picked" :on-change="onChange(picked)" />
            <select class="form-select">
            <option value="" selected disabled hidden>Inicio</option>
            <option>08:00</option>
            <option>09:00</option>
            <option>10:00</option>
            <option>11:00</option>
            <option>12:00</option>
            <option>13:00</option>
            <option>14:00</option>
            <option>15:00</option>
            <option>16:00</option>
            <option>17:00</option>
            <option>18:00</option>
            <option>19:00</option>
            <option>20:00</option>
            <option>21:00</option>
        </select>
        <select class="form-select">
            <option value="" selected disabled hidden>Fin</option>
            <option>09:00</option>
            <option>10:00</option>
            <option>11:00</option>
            <option>12:00</option>
            <option>13:00</option>
            <option>14:00</option>
            <option>15:00</option>
            <option>16:00</option>
            <option>17:00</option>
            <option>18:00</option>
            <option>19:00</option>
            <option>20:00</option>
            <option>21:00</option>
            <option>22:00</option>
        </select>
        </div>
        
    </div>
</template>
<script>
import Datepicker from 'vue3-datepicker/dist/vue3-datepicker.esm'
export default({
    components: {
        Datepicker
    },
    data: ()=> {
        return{
            picked: new Date()
        };
       
    },
    methods: {
        onChange(event){
            console.log(event)
        }
    }
})
</script>