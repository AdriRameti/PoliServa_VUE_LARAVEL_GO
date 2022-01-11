import { ref } from "vue";
import { secret } from "../../secret";

export const useGetUserReservation = async (arr) => {

    let reservations = ref([]);
    let id = arr[0];
    let fIni = arr[1];
    let fFin = arr[2];
    let res = await fetch(secret.GO_URL + "reservations/userreservation?id="+id+"&fIni="+fIni+"&fFin="+fFin);
     let data = await res.json();
    
    reservations.value = data;
    return { reservations }

}

export const useGetUserSportReservation = async (id) => {
    let reservations = ref([]);
    let res = await fetch(secret.GO_URL + "reservations/usersportreservation?id="+id);
     let data = await res.json();
    console.log(data);
    reservations.value = data;
    return { reservations }

}