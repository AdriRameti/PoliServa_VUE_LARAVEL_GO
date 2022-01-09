import store from '@/store';

export default {

    authGuardUser(to, from, next) {

        if (localStorage.getItem('token') && store.state.user.user) {
            next();
        } else {
            next("/");
        }
        
    }
}