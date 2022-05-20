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
    reservations.value = data;
    return { reservations }

}
export const useGetUserCourtReservation = async () => {
    let reservations = ref([]);
    let res = await fetch(secret.GO_URL + "reservations/usercourtreservation");
    let data = await res.json();
    reservations.value = data;
    return { reservations }

}

export const useGetSessionId = async()=>{
    var total_price = localStorage.getItem("total_price");
    let token = localStorage.getItem('token');
    let res = await fetch(secret.LARAVEL_URL + "reservation/usergetsessionid",{
        method:'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify({total_price:parseInt(total_price)})
    });
    let data = await res.json();
    return data.id;
}
