<template>
    <div class="pistas d-flex flex-wrap p-3 justify-content-center m-1">

        <div class="card p-3 card-reser" v-for="sport in sportsData" :key="sport.id">
            <img class="w-100 listSports" :src="sport.img"/>
            <h1 class="text-center uppercase fs-3">{{ sport.name }}</h1>
            <p class="reser-descrip">{{ $t("CARD_DESC") }}</p>
            <button class="btn-sport btn text-white" v-on:click="setSport(sport.slug)">{{ $t("CHECK_COURTS") }}</button>
        </div>
    </div>
</template>
<style>
    .btn-sport{
        background-color: #40916C;
    }
</style>
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

<style>
.listSports{
    height: 141.13px !important;
}
</style>