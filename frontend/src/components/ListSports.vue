<template>
    <div class="pistas d-flex flex-wrap p-3 justify-content-center m-3">

        <div class="card p-3 card-reser" v-for="sport in sportsData" :key="sport.id">
            <img class="w-100" :src="sport.img"/>
            <h1 class="text-center uppercase fs-3">{{ sport.name }}</h1>
            <p class="reser-descrip">Lorem Ipsum is simply dummy text of the printing and typesetting industry.</p>
            <button class="btn btn-dark text-white" v-on:click="setSport(sport.slug)">Mirar pistas</button>
        </div>
    </div>
</template>
<script>

import { useGetAllSports } from '../composables/sports/useSports';
import { reactive } from 'vue';

export default({

    methods:{
        setSport(slug) {
            localStorage.setItem('slug', slug);
            window.location.href="/#/reservation";
        }
    },
    async setup() {

        const { sports } = await useGetAllSports();

        var sportsData = reactive(sports.value);

        return { sportsData };
    },
})
</script>

<style lang="scss">

</style>