import { ref, reactive, renderSlot } from "vue";
import { secret } from "../../secret";
import { useToast } from "vue-toastification";

export const useGetAllSports = async () => {
    
    let sports = reactive([]);

    let res = await fetch(secret.GO_URL + "sports/");
    let data = await res.json();
    
    sports.value = data;

    return { sports }

}

export const useDeleteSport = async (slug) => {

    let sports = reactive([]);
    let toastr = useToast();

    let res = await fetch(secret.GO_URL + "sports/?slug=" + slug, {
        method: 'DELETE'
    });

    if (res.status == 200) {

        let data = await res.json();

        sports.value = data;

        toastr.success("Deleted successfully.", {
            timeout: 1500
        });

        return { sports }
    } else {
        toastr.error("Something failed trying to process the request.", {
            timeout: 1500
        });

        return 'err'
    } 

}

export const useCreateSport = async (info) => {

    let sport = ref();

    let toastr = useToast();

    let res = await fetch(secret.GO_URL + "sports/?name="+ info.name + "&img=" + info.img, {
        method: 'POST',
    });

    if (res.status == 200) {

        let data = await res.json();

        sport.value = data;

        toastr.success("Created successfully.", {
            timeout: 1500
        });

        return { sport }

    } else {
        toastr.error("Something failed trying to process the request.", {
            timeout: 1500
        });

        return 'err'
    }
}

export const useUpdateSport = async (info) => {
    
    let sport = ref();

    let toastr = useToast();

    let res = await fetch(secret.GO_URL + "sports/?slug=" + info.slug + "&name="+ info.name + "&img=" + info.img, {
        method: 'PUT',
    });

    if (res.status == 200) {

        let data = await res.json();

        sport.value = data;

        toastr.success("Updated successfully.", {
            timeout: 1500
        });

        return { sport }

    } else {
        toastr.error("Something failed trying to process the request.", {
            timeout: 1500
        });

        return 'err'
    }
}