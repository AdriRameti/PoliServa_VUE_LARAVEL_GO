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

    console.log(courts.value, count.value);

    return { courts, count }                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    
}
export const useGetDateReservation = async (infoDate) => {
    console.log(infoDate)
    var date = infoDate[0]
    var hini = infoDate[1]
    var hfin = infoDate[2]
    var slug = infoDate[3]
    let courts = ref([]);

    let count = computed(() => courts.value.length);

    let res = await fetch(secret.GO_URL + "reservations/datereservation?date="+date+"&hini="+hini+"&hfin="+hfin+"&slug="+slug);
    // let res = Api(`http://localhost:3000/api/`).get(`courts`);
    let data = await res.json();
    
    courts.value = data;

    console.log(courts.value, count.value);

    return { courts, count }                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    
}