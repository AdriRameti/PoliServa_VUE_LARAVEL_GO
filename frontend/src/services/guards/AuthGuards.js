import store from '@/store';

export default {

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