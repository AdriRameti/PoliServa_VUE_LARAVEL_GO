<template>
  <div class="home">
    <Suspense>
      <Carousel/>
    </Suspense>

    <h1 class="mt-3 text-center">{{ $t("SPORTS") }}</h1>

    <Suspense>
      <ListSports/>
    </Suspense>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from '@/components/HelloWorld.vue'
import Carousel from '@/components/Carousel.vue'
import ListSports from '@/components/ListSports.vue'

export default {
  name: 'Home',
  components: {
    Carousel,
    ListSports
  }
}
</script>
