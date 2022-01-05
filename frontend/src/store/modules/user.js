import Constant from "../../Constant";
import UserServices from "../../services/UserServices";
import { useToast } from "vue-toastification";

export const user = {
    namespaced: true,

    state: {
        
    },
    mutations: {
        [Constant.LOGIN_USER]: (state, payload) => {

            state.user = payload
            state.token = payload.token

            localStorage.setItem('token', payload.token)
        },
        [Constant.REGISTER_USER]: (state, payload) => {
            
        },
        [Constant.VALIDATE_OTP]: (state, payload) => {
            
        },
        [Constant.VALIDATE_REGISTER]: (state, payload) => {
            
        }
    },
    actions: {
        [Constant.LOGIN_USER]: (store, payload) => {

            var toastr = useToast();

            UserServices.login(payload).then((data) => {

                
                if (data.data.message == "password don't match") {
                    toastr.error("Password don't match", {
                        timeout: 1500
                    });
                    return

                } else if (data.data.message == "user not found") {
                    toastr.error("User not found", {
                        timeout: 1500
                    });
                    return
                }

                store.commit(Constant.LOGIN_USER, data.data.data);

                toastr.success("Login success", {
                    timeout: 1500
                });

                // $router.push({name: 'Profile'})
                window.location.href = "/#/";

            });
            

        },
        [Constant.REGISTER_USER]: (store, payload) => {
            
        },
        [Constant.VALIDATE_OTP]: (store, payload) => {
            
        },
        [Constant.VALIDATE_REGISTER]: (store, payload) => {
            
        }
    },
    getters: {
        getUser(state) {

            console.log('getter', state)
            return state.user;
        }
    }
}