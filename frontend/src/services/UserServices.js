import Api from "./Api"
import { secret } from "../secret"

export default {
    login(data) {
        return Api(`${secret.LARAVEL_URL}`).post(`users/login/`, data);
    },
    enable2fa() {
        return Api(`${secret.LARAVEL_URL}`).post(`users/enable2fa/`);
    },
    disable2fa() {
        return Api(`${secret.LARAVEL_URL}`).post(`users/disable2fa/`);
    },
    check2fa(data) {
        return Api(`${secret.LARAVEL_URL}`).post(`users/check2fa/`, {'one_time_password': data});
    },
    updateUser(data) {
        return Api(`${secret.LARAVEL_URL}`).post(`users/update`, data);
    },
    validaRegister(mail){  
        return Api(`${secret.LARAVEL_URL}`).post(`users/sendMailRegister/`,{'mail':mail.mail});
    },
    register(info){ 
        return Api(`${secret.LARAVEL_URL}`).post(`users/mailRegister/`,{'info':info});
    },
    deleteUser(data){
        return Api(`${secret.LARAVEL_URL}`).post(`users/deleteUser`,{'one_time_password':data})
    },
    destroyUser(){
        return Api(`${secret.LARAVEL_URL}`).delete(`users/destroy`)
    }
}