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

export const useGetUserSportReservation = async (arr) => {
    let reservations = ref([]);
    let id = arr[0];
    let type = arr[1];
    let res = await fetch(secret.GO_URL + "reservations/usersportreservation?id="+id+"&type="+type);
     let data = await res.json();
    console.log(data);
    reservations.value = data;
    return { reservations }

}
export const useGetUserCourtReservation = async () => {
    let reservations = ref([]);
    let res = await fetch(secret.GO_URL + "reservations/usercourtreservation");
     let data = await res.json();
    console.log(data);
    reservations.value = data;
    return { reservations }

}
