import Constant from "../../Constant";
import UserServices from "../../services/UserServices";
import { useToast } from "vue-toastification";

export const user = {
    namespaced: true,

    state: {
        g2faverified: false
    },
    mutations: {
        [Constant.LOGIN_USER]: (state, payload) => {

            state.user = payload;
            state.token = payload.token;

            localStorage.setItem('token', payload.token);
        },
        [Constant.REGISTER_USER]: (state, payload) => {
            
        },
        [Constant.ENABLE2FA]: (state, payload) => {
            state.google2fa_secret = payload.google2fa;
            state.user = payload;
        },
        [Constant.DISABLE2FA]: (state, payload) => {
            state.google2fa_secret = '';
            state.user = payload;
        },
        [Constant.CHECK2FA]: (state, payload) => {
            if (payload) {
                state.g2faverified = true;
            }
        },
        [Constant.UPDATE_USER]: (state, payload) => {
            state.user.img = payload.img;
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

                if (data.data.data.google2fa_secret) {
                    window.location.href = "/#/otp";
                } else {
                    window.location.href = "/#/";
                    toastr.success("Login success", {
                        timeout: 1500
                    });
                }

            });
            

        },
        [Constant.ENABLE2FA]: (store, payload) => {
            UserServices.enable2fa().then(data => {
                store.commit(Constant.ENABLE2FA, data.data);
            });
        },
        [Constant.DISABLE2FA]: (store, payload) => {
            UserServices.disable2fa().then(data => {
                store.commit(Constant.DISABLE2FA, data.data);
            });
        },
        [Constant.CHECK2FA]: (store, payload) => {

            var toastr = useToast();

            UserServices.check2fa(payload.otp).then(data => {

                if (data.data == "not verified") {
                    toastr.error("Incorrect code", {
                        timeout: 1500
                    });
                    return
                } else if (data.data == "verified") {
                    
                    
                    if (payload.from == 'otp') {

                        toastr.success("Login success", {
                            timeout: 1500
                        });

                        window.location.href="/#/";
                    } else {

                        toastr.success("Two-Factor Authentication activated", {
                            timeout: 1500
                        });

                        store.commit(Constant.CHECK2FA, true);
                    }

                }
            })
        },
        [Constant.UPDATE_USER]: (store, payload) => {

            UserServices.updateUser(payload).then(data => {
                console.log(data)

                store.commit(Constant.UPDATE_USER, data.data)
            });

        },
        [Constant.VALIDATE_REGISTER]: (store, payload) => {
            
        }
    },
    getters: {
        getUser(state) {

            console.log('getter', state);
            return state.user;
        }
    }
}