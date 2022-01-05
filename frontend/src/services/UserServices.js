import Api from "./Api"
import { secret } from "../secret"

export default {
    login(data) {
        console.log(data);
        return Api(`${secret.LARAVEL_URL}`).post(`users/login/`, data);
    }
}