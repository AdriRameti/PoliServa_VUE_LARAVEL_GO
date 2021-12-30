import { ref, watch, computed } from "vue";
// import Api from "../../services/Api"
import { secret } from "../../secret";

export const useGetAllCourts = async () => {
    let courts = ref([]);

    let count = computed(() => courts.value.length);

    let res = await fetch(secret.GO_URL + "courts/");
    // let res = Api(`http://localhost:3000/api/`).get(`courts`);
    let data = await res.json();
    
    courts.value = data;

    console.log(courts.value, count.value);

    return { courts, count }                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    
}
export const useGetSportCourts = async (slug) => {
    console.log(slug)
    let courts = ref([]);

    let count = computed(() => courts.value.length);

    let res = await fetch(secret.GO_URL + "courts/sport?slug="+slug);
    // let res = Api(`http://localhost:3000/api/`).get(`courts`);
    let data = await res.json();
    
    courts.value = data;
    
    return { courts, count }                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    
}
export const useGetCarouselCourts = async () => {
    
    let courts = ref([]);

    let res = await fetch(secret.GO_URL + "courts/carousel");
    // let res = Api(`http://localhost:3000/api/`).get(`courts`);
    let data = await res.json();
    
    courts.value = data;

    return { courts }

}