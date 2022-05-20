import store from '@/store';
import UserServices from "../../services/UserServices";

export default {

    checkIfToken(to, from, next) {
        if (localStorage.getItem('token') && !store.state.user.user) {
            UserServices.getUser().then(data => {
                store.state.user.user = data.data.data;
            })
        }

        next();
    },

    authGuardUser(to, from, next) {

        if (localStorage.getItem('token') && store.state.user.user) {
            next();
        } else {
            next("/");
        }
        
    },

    authGuardAdmin(to, from, next) {

        if (localStorage.getItem('token') && store.state.user.user && store.state.user.user.role == 'admin') {
            next();
        } else {
            next("/");
        }

    }
}