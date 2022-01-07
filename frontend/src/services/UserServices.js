import Api from "./Api"
import { secret } from "../secret"

export default {
    login(data) {
        console.log(data);
        return Api(`${secret.LARAVEL_URL}`).post(`users/login/`, data);
    },
    enable2fa() {
        return Api(`${secret.LARAVEL_URL}`).post(`users/enable2fa/`);
    },
    disable2fa() {
        return Api(`${secret.LARAVEL_URL}`).post(`users/disable2fa/`);
    },
    check2fa(data) {
        return Api(`${secret.LARAVEL_URL}`).post(`users/check2fa/`, {'one_time_password': data})
    }
}