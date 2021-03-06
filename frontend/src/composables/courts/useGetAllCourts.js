import { ref, watch, computed } from "vue";
// import Api from "../../services/Api"
import { secret } from "../../secret";
import { useToast } from "vue-toastification";

export const useGetAllCourts = async () => {
    let courts = ref([]);

    let count = computed(() => courts.value.length);

    let res = await fetch(secret.GO_URL + "courts/");
    let data = await res.json();
    
    courts.value = data;

    return { courts, count }                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    
}
export const useGetSportCourts = async (slug) => {

    let courts = ref([]);

    let count = computed(() => courts.value.length);

    let res = await fetch(secret.GO_URL + "courts/sport?slug="+slug);
    // let res = Api(`http://localhost:3000/api/`).get(`courts`);
    let data = await res.json();
    
    courts.value = data;
    
    return { courts, count }                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    
}
export const useGetSport = async (idSport) =>{
    let courts = ref([]);

    let count = computed(() => courts.value.length);

    let res = await fetch(secret.GO_URL + "sports/sport?slug="+slug);
    let data = await res.json();
    
    courts.value = data;
    
    return { courts, count }  
}
export const useGetCarouselCourts = async () => {
    
    let courts = ref([]);

    let count = computed(() => courts.value.length);

    let res = await fetch(secret.GO_URL + "courts/carousel");
    // let res = Api(`http://localhost:3000/api/`).get(`courts`);
    let data = await res.json();
    
    courts.value = data;

    return { courts, count }                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    
}
export const useGetDateReservation = async (infoDate) => {
    var date = infoDate[0]
    var hini = infoDate[1]
    var hfin = infoDate[2]
    var slug = infoDate[3]
    
    let courts = ref([]);

    let count = computed(() => courts.value.length);

    let res = await fetch(secret.GO_URL + "courts/datereservation?date="+date+"&hini="+hini+"&hfin="+hfin+"&slug="+slug);
    let data = await res.json();
    courts.value = data;

    return { courts, count }                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    
}
export const useCreateReservation = async () => {
    
    let courts = ref([]);

    let count = computed(() => courts.value.length);

    let res = await fetch(secret.GO_URL + "reservations/");
    let data = await res.json();
    
    courts.value = data;

    return { courts, count }                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    
}
export const useInsertReservation = async (Rdata) => {
    let reservation = ref([]);
    let token = localStorage.getItem('token')
    let res = await fetch(secret.LARAVEL_URL + "reservation/insertReservations",{
        method:'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify(Rdata)
    });
    let data = await res.json();

    return {data}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   
}

export const useDeleteCourt = async (id) => {

    let toastr = useToast();

    let res = await fetch(secret.GO_URL + "courts/?id=" + id, {
        method: 'DELETE'
    });

    if (res.status == 200) {

        toastr.success("Deleted successfully.", {
            timeout: 1500
        });

        return 'success'
    } else {
        toastr.error("Something failed trying to process the request.", {
            timeout: 1500
        });

        return 'err'
    }
}

export const useCreateCourt = async (info) => {

    let court = ref();

    let toastr = useToast();

    let res = await fetch(secret.GO_URL + "courts/?id_sport="+ info.id_sport + "&img=" + info.img + "&price_h=" + info.price_h + "&sector=" + info.sector , {
        method: 'POST',
    });

    if (res.status == 201) {

        let data = await res.json();

        court.value = data;

        toastr.success("Created successfully.", {
            timeout: 1500
        });

        return { court }

    } else {
        toastr.error("Something failed trying to process the request.", {
            timeout: 1500
        });

        return 'err'
    }
}

export const useUpdateCourt = async (info) => {
    
    let court = ref();

    let toastr = useToast();

    let res = await fetch(secret.GO_URL + "courts/?id=" + info.id + "&id_sport=" + info.id_sport + "&img=" + info.img + "&price_h=" + info.price_h + "&sector=" + info.sector , {
        method: 'PUT',
    });

    if (res.status == 200) {

        let data = await res.json();

        court.value = data;

        toastr.success("Updated successfully.", {
            timeout: 1500
        });

        return { court }

    } else {
        toastr.error("Something failed trying to process the request.", {
            timeout: 1500
        });

        return 'err'
    }
}