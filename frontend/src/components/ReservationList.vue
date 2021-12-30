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

import { useGetAllCourts } from '../composables/courts/useGetAllCourts';

export default({
    async setup() {
        const { courts, count } = await useGetAllCourts();

        console.log('hola', courts.value);

        return { courts, count };
    },
})
</script>
