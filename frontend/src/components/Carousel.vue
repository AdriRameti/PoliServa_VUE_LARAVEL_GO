<template>
    <div id="carouselExampleControls" class="carousel slide" data-bs-ride="carousel">
        <div class="carousel-inner">

          <div class="carousel-item" v-for="(court, index) in courts" :key="court.id" :class="{ active: index == 0 }">
            <img :src="court.img" class="d-block w-100" alt="...">
          </div>

        </div>
        <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleControls" data-bs-slide="prev">
          <span class="carousel-control-prev-icon" aria-hidden="true"></span>
          <span class="visually-hidden">Previous</span>
        </button>
        <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleControls" data-bs-slide="next">
          <span class="carousel-control-next-icon" aria-hidden="true"></span>
          <span class="visually-hidden">Next</span>
        </button>
    </div>
</template>
<script>

import { useGetAllCourts, useGetSportCourts, useGetCarouselCourts } from '../composables/courts/useGetAllCourts';

export default({
    async setup() {
        const { courts } = await useGetCarouselCourts();

        return { courts };
    },
})
</script>

<style lang="scss">
.carousel{
    height: 500px;
}
.carousel {
    img{
        height: 500px;
    }
}
</style>