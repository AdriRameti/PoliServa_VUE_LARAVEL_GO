import { ref } from "vue";
import { secret } from "../../secret";

export const useGetAllSports = async () => {
    
    let sports = ref([]);

    let res = await fetch(secret.GO_URL + "sports/");
    let data = await res.json();
    
    sports.value = data;

    return { sports }

}