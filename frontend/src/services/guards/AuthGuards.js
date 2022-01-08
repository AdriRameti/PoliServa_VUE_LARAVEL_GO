
export default {

    authGuardUser(to, from, next) {

        if (localStorage.getItem('token')) {
            next();
        } else {
            next("/");
        }
        
    }
}