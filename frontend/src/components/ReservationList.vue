<template>
    <div>
        <ul>
            <li v-for="court in courts" :key="court.id">
                {{ court.id_sport }}
            </li>
        </ul>
    </div>
            <div class="pistas d-flex flex-wrap p-3 justify-content-center m-3">
            <div class="card p-3 card-reser">
                <img src="http://manzasport.com/wp-content/uploads/2018/04/Pista-de-padel-Paquito-Navarro-by-Manzasport-1.jpeg"/>
                <h1>Pista 1</h1>
                <p class="reser-descrip">Lorem Ipsum is simply dummy text of the printing and typesetting industry.</p>
            </div>
            <div class="card p-3 card-reser">
                <img src="http://manzasport.com/wp-content/uploads/2018/04/Pista-de-padel-Paquito-Navarro-by-Manzasport-1.jpeg"/>
                <h1>Pista 1</h1>
                <p class="reser-descrip">Lorem Ipsum is simply dummy text of the printing and typesetting industry.</p>
            </div>
            <div class="card p-3 card-reser">
                <img src="http://manzasport.com/wp-content/uploads/2018/04/Pista-de-padel-Paquito-Navarro-by-Manzasport-1.jpeg"/>
                <h1>Pista 1</h1>
                <p class="reser-descrip">Lorem Ipsum is simply dummy text of the printing and typesetting industry.</p>
            </div>
            <div class="card p-3 card-reser">
                <img src="http://manzasport.com/wp-content/uploads/2018/04/Pista-de-padel-Paquito-Navarro-by-Manzasport-1.jpeg"/>
                <h1>Pista 1</h1>
                <p class="reser-descrip">Lorem Ipsum is simply dummy text of the printing and typesetting industry.</p>
            </div>
            <div class="card p-3 card-reser">
                <img src="http://manzasport.com/wp-content/uploads/2018/04/Pista-de-padel-Paquito-Navarro-by-Manzasport-1.jpeg"/>
                <h1>Pista 1</h1>
                <p class="reser-descrip">Lorem Ipsum is simply dummy text of the printing and typesetting industry. </p>
            </div>
            <div class="card p-3 card-reser">
                <img src="http://manzasport.com/wp-content/uploads/2018/04/Pista-de-padel-Paquito-Navarro-by-Manzasport-1.jpeg"/>
                <h1>Pista 1</h1>
                <p class="reser-descrip">Lorem Ipsum is simply dummy text of the printing and typesetting industry. </p>
            </div>
        </div> 
</template>
<script>

import { useGetAllCourts } from '../composables/courts/useGetAllCourts';

export default({
    setup() {
        const { courts, count } = useGetAllCourts();

        console.log(courts);

        return { courts, count };
    },
})
</script>
