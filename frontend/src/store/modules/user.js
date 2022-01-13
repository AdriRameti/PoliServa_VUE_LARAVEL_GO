import Constant from "../../Constant";
import UserServices from "../../services/UserServices";
import { useToast } from "vue-toastification";

export const user = {
    namespaced: true,

    state: {
        g2faverified: false,
    },
    mutations: {
        [Constant.LOGIN_USER]: (state, payload) => {

            state.user = payload;
            state.token = payload.token;
            state.google2fa_secret = payload.google2fa_secret;

            localStorage.setItem('token', payload.token);
        },
        [Constant.REGISTER_USER]: (state, payload) => {
            state.user = payload;
            state.token = payload.token;
            state.google2fa_secret = payload.google2fa_secret;

            localStorage.setItem('token', payload.token);
        },
        [Constant.ENABLE2FA]: (state, payload) => {
            state.google2fa_secret = 'secret';
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
            state.user = payload;
        },
        [Constant.VALIDATE_REGISTER]: (state, payload) => {
            state.code = payload
            localStorage.setItem('verifyCode',state.code);
        },
        [Constant.DELETE_USER]: (state, payload) => {
            state.code = payload
            localStorage.setItem('verifyCode',state.code);
        },
        [Constant.DESTROY_USER]: (state, payload) => {
            localStorage.setItem('delete',1);
        },
        [Constant.LOGOUT]: (state, payload) => {
            localStorage.removeItem('token');
            state.user = undefined;
        },
        [Constant.GET_ALL_USERS]: (state, payload) => {
            state.allUsers = payload;
        },
        [Constant.CREATE_USER]: (state, payload) => {
            state.allUsers.push(payload);
        },
        [Constant.UPDATE_USER_ADMIN]: (state, payload) => {
            var newData = state.allUsers.map(data => data.uuid == payload.uuid ? payload : data );
            state.allUsers = newData;
        },
        [Constant.DELETE_USER_ADMIN]: (state, payload) => {
            var newData = state.allUsers.filter((data) => { return data.uuid != payload.uuid })
            state.allUsers = newData;
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
                        
                    } else if (payload.from == 'delete') {
                        
                        UserServices.destroyUser().then(data =>{
                            if(data){
                                store.commit(Constant.LOGOUT, data);
                                
                                toastr.success("You're acount has been deleted successfully", {
                                    timeout: 1500
                                });
        
                                setTimeout(() => {
                                    window.location.reload();
                                }, 1800);
                            }
                        })
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
                store.commit(Constant.UPDATE_USER, data.data)
            });

        },
        [Constant.VALIDATE_REGISTER]: (store, payload) => {
            
            UserServices.validaRegister(payload).then(data =>{
                console.log(data.data);
                store.commit(Constant.VALIDATE_REGISTER, data.data);
            });
        },
        [Constant.REGISTER_USER]: (store, payload) => {
            
            var toastr = useToast();
            UserServices.register(payload).then(data =>{
                store.commit(Constant.REGISTER_USER, data.data.data);

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
        [Constant.DELETE_USER]: (store, payload) =>{
            UserServices.deleteUser(payload).then(data =>{
                store.commit(Constant.DELETE_USER, data.data);
            })
        },
        [Constant.DESTROY_USER]: (store, payload) =>{
            UserServices.destroyUser().then(data =>{
                if(data){
                    store.commit(Constant.DESTROY_USER, data);
                }
            })
        },
        [Constant.LOGOUT]: (store, payload) => {
            store.commit(Constant.LOGOUT);
        },
        [Constant.GET_ALL_USERS]: (store, payload) => {
            UserServices.getAllUsers().then(data => {
                store.commit(Constant.GET_ALL_USERS, data.data);
            });
        },
        [Constant.CREATE_USER]: (store, payload) => {

            var toastr = useToast();

            UserServices.createUser(payload).then(data => {

                if (data.status != 200) {
                    toastr.error("Something failed trying to process the request.", {
                        timeout: 1500
                    });
                } else {
                    if (data.data == 'already used') {
                        toastr.error("Mail already used.", {
                            timeout: 1500
                        });
                    } else if (data.data == "can't bind") {
                        toastr.error("Something failed trying to process the request.", {
                            timeout: 1500
                        });
                    } else {
                        store.commit(Constant.CREATE_USER, data.data);
                        toastr.success("Used created successfully.", {
                            timeout: 1500
                        });
                    }
                }

                
            });
        },
        [Constant.UPDATE_USER_ADMIN]: (store, payload) => {
            var toastr = useToast();

            UserServices.updateUserAdmin(payload).then(data => {

                if (data.status != 200) {
                    toastr.error("Something failed trying to process the request.", {
                        timeout: 1500
                    });
                } else {
                    if (data.data == 'already used') {
                        toastr.error("Mail already used.", {
                            timeout: 1500
                        });
                    } else if (data.data == "can't bind") {
                        toastr.error("Something failed trying to process the request.", {
                            timeout: 1500
                        });
                    } else {
                        store.commit(Constant.UPDATE_USER_ADMIN, data.data);
                        toastr.success("Used updated successfully.", {
                            timeout: 1500
                        });
                    }
                }  

            });

        },
        [Constant.DELETE_USER_ADMIN]: (store, payload) => {

            var toastr = useToast();

            UserServices.deleteUserAdmin(payload).then(data => {
                if (data.status == 200) {
                    if (data.data == "Se ha eliminado con éxito") {
                        store.commit(Constant.DELETE_USER_ADMIN, payload);
                        toastr.success("User deleted successfully.", {
                            timeout: 1500
                        });
                    } else {
                        toastr.error("Something failed trying to process the request.", {
                            timeout: 1500
                        });
                    }
                } else {
                    toastr.error("Something failed trying to process the request.", {
                        timeout: 1500
                    });
                }
                
            });
        }
    },
    getters: {
        getUser(state) {
            return state.user;
        }
    }
}