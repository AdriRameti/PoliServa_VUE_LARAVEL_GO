import store from '@/store';

export default {

    authGuardUser(to, from, next) {

        if (store.getters["user/getUser"]) {
            next();
        } else {
            next("/");
        }
        
    }
}