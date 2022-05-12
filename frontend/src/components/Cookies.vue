<template>
    <div id="modalCookies" data-bs-backdrop="static" class="modal fade" v-if="!localItem">
        <div class="modal-dialog modal-dialog-centered">
            <div class="modal-content">
                <div class="modal-body">
                    <div class="card">
                        <svg viewBox="0 0 512 512" width="100" title="cookie-bite">
                            <path d="M510.52 255.82c-69.97-.85-126.47-57.69-126.47-127.86-70.17 0-127-56.49-127.86-126.45-27.26-4.14-55.13.3-79.72 12.82l-69.13 35.22a132.221 132.221 0 0 0-57.79 57.81l-35.1 68.88a132.645 132.645 0 0 0-12.82 80.95l12.08 76.27a132.521 132.521 0 0 0 37.16 72.96l54.77 54.76a132.036 132.036 0 0 0 72.71 37.06l76.71 12.15c27.51 4.36 55.7-.11 80.53-12.76l69.13-35.21a132.273 132.273 0 0 0 57.79-57.81l35.1-68.88c12.56-24.64 17.01-52.58 12.91-79.91zM176 368c-17.67 0-32-14.33-32-32s14.33-32 32-32 32 14.33 32 32-14.33 32-32 32zm32-160c-17.67 0-32-14.33-32-32s14.33-32 32-32 32 14.33 32 32-14.33 32-32 32zm160 128c-17.67 0-32-14.33-32-32s14.33-32 32-32 32 14.33 32 32-14.33 32-32 32z" />
                        </svg>
                        <h3>Consentimiento de Cookies</h3>
                        <p>Utilizamos cookies en esta web, para asegurarse que la experiencia del usuario que visita nuestra web sea óptima.</p>
                        <div class="buttons__cta">
                            <button type="" data-bs-dismiss="modal" @click="acceptOrDenyCookies()">Lo entiendo</button>
                            <a href="https://policies.google.com/technologies/cookies?hl=es" target="_blank">Leer más</a>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <button id="modalCookiesButton" data-bs-target="#modalCookies" data-bs-toggle="modal" v-show="false"></button>
    </div>
</template>
<script>
import { onMounted } from "vue";
export default {
    methods: {
        acceptOrDenyCookies() {
            localStorage.setItem('cookies', " ");
        }
    },
    setup() {
       var localItem = localStorage.getItem('cookies');

       onMounted(() => {
            if (!localItem) {
                document.getElementById("modalCookiesButton").click();
            }
       }) 

        return { localItem };
    },
}
</script>
<style>
@import url("https://fonts.googleapis.com/css?family=Roboto:400,400i,700");
*{
  margin:0;
  padding:0;
  box-sizing: border-box;
}
.card{  
  width: 100%;
  height: 100%;
  padding: 20px 30px;
  max-width:100%;
  background: whitesmoke;
  text-align: center;
  border-radius: 14px;
  -webkit-box-shadow: 0px 4px 7px -3px #444; 
  box-shadow: 0px 4px 7px -3px #444;
  font-family: Roboto,sans-serif; 
}
.card svg{
  fill: #FFCC80;
  margin: 0 auto;
  margin-bottom: 15px;
}
.card h3 {
  color: #555;
  font-size: 1.5rem;
  margin-bottom: 10px;
}
.card p{
  color: #999;
  font-size:0.9rem;
  margin-bottom: 25px;
}
.buttons__cta{
  margin-bottom: 20px;
}
.buttons__cta button{
  color: white;
  cursor: pointer;
  font-weight: bold;
  padding: 8px 12px;
  margin-right: 20px;
  border-radius: 4px;  
  background: #40916C;  
  border: 1px solid #ccbb66;
}
.buttons__cta a {
  color: #FFCC80;
  text-decoration: none;
}

</style>