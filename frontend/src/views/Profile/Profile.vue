<template>
    <Suspense>
        <Profile/>
    </Suspense>
</template>
<script>
//import { defineComponent } from '@vue/composition-api'

import Profile from '@/components/Profile.vue'

export default({

    name: 'ProfileView',
    components: {
        Profile
    },

})
</script>