<template>
    <div>
        <div class="text-center">
            <h1>Users</h1>
        </div>

        <div class="mb-2 d-flex justify-content-end">
            <button v-on:click="showCreate()" class="btn btn-success m-1">Create</button>
            <button v-on:click="getSelectedRows('update')" class="btn btn-primary m-1">Update</button>
            <button v-on:click="getSelectedRows('delete')" class="btn btn-danger m-1">Delete</button>
        </div>
        
        <div class="text-center">
            <p v-if="create">Create</p>
            <p v-if="update">Update</p>
        </div>
        <div class="d-flex flex-row justify-content-between mb-3" v-if="create">
            <input id="cMail" v-model="cMail" class="form-control form-control-sm" type="text" placeholder="Mail">
            <input id="cPass" v-model="cPass" class="form-control form-control-sm" type="password" placeholder="Password">
            <input id="cName" v-model="cName" class="form-control form-control-sm" type="text" placeholder="Name">
            <input id="cSurnames" v-model="cSurnames" class="form-control form-control-sm w-75" type="text" placeholder="Surnames">
            <input id="cImg" v-model="cImg" class="form-control form-control-sm w-75" type="text" placeholder="Image">
            <input id="cRole" v-model="cRole" class="form-control form-control-sm w-75" type="text" placeholder="Role">
            <input id="cBlocked" v-model="cBlocked" class="form-control form-control-sm w-75" type="text" placeholder="Blocked">
            <button v-on:click="createCourt()" class="btn btn-primary m-1">Send</button>
        </div>

        <div class="d-flex d-inline-block" v-if="update">
            <input id="uMail" v-model="uMail" class="form-control form-control-sm" type="text" :placeholder="uMail">
            <input id="uPass" v-model="uPass" class="form-control form-control-sm" type="password" placeholder="Password">
            <input id="uName" v-model="uName" class="form-control form-control-sm" type="text" :placeholder="uName">
            <input id="uSurnames" v-model="uSurnames" class="form-control form-control-sm w-75" type="text" :placeholder="uSurnames">
            <input id="uImg" v-model="uImg" class="form-control form-control-sm w-75" type="text" :placeholder="uImg">
            <input id="uRole" v-model="uRole" class="form-control form-control-sm w-75" type="text" :placeholder="uRole">
            <input id="uBlocked" v-model="uBlocked" class="form-control form-control-sm w-75" type="text" :placeholder="uBlocked">
            <button v-on:click="updateCourt()" class="btn btn-primary m-1">Send</button>
        </div>

        <ag-grid-vue id="myGrid" style="width: 100%;"
            class="ag-theme-alpine"
            :columnDefs="columnDefs"
            :rowData="store.state.user.allUsers"
            rowSelection="single"
            domLayout='autoHeight'
            @gridSizeChanged="gridSizeChanged" :key="grid">
        </ag-grid-vue>
    </div>

</template>

<style>

</style>
<script>

import { useToast } from "vue-toastification";
import { useStore } from "vuex";
import Constant from '../Constant';

import { AgGridVue } from "ag-grid-vue3";
import { ref } from "vue";

export default({
    components: {
        AgGridVue,
    },
    async setup() {
        const toastr = useToast();
        const store = useStore();

        store.dispatch("user/" + Constant.GET_ALL_USERS);

        return { store, toastr }
    },
    data() {
        return {
            columnDefs: [
                { headerName: "ID", field: "id", sortable: true, checkboxSelection: true },
                { headerName: "UUID", field: "uuid", sortable: true, resizable: true },
                { headerName: "Mail", field: "mail", sortable: true, filter: true, resizable: true },
                { headerName: "Name", field: "name", sortable: true, filter: true, resizable: true },
                { headerName: "Surnames", field: "surnames", sortable: true, filter: true, resizable: true },
                { headerName: "Image", field: "img", resizable: true },
                { headerName: "Role", field: "role", sortable: true, filter: true },
                { headerName: "Blocked", field: "is_blocked", sortable: true, filter: true },
            ],
            gridApi: ref({}),
            columnApi: ref({}),
            grid: 0,
            create: false,
            update: false,
            cMail: '',
            cPass: '',
            cName: '',
            cSurnames: '',
            cImg: '',
            cRole: '',
            cBlocked: '',
            uUUID: '',
            uMail: '',
            uPass: '',
            uName: '',
            uSurnames: '',
            uImg: '',
            uRole: '',
            uBlocked: '',
        }
    },
    methods:{
        async getSelectedRows(op) {
            const selectedNodes = this.gridApi.value.getSelectedNodes();
            const selectedData = selectedNodes.map( node => node.data );
            const uuid = selectedData.map( node => `${node.uuid}`);

            if (op == "delete") {

                if (selectedNodes.length == 1) {
                    this.store.dispatch("user/" + Constant.DELETE_USER_ADMIN, {'uuid': uuid})
                } else {
                    this.toastr.error("You must select one.", {
                        timeout: 1500
                    });
                }

            } else if (op == "update") {

                if (this.update && this.uMail == selectedData[0].mail) {
                    this.update = false;
                    this.create = false;
                } else {
                    if (selectedNodes.length == 1) {

                        const data = selectedData[0];

                        this.uUUID = uuid;
                        this.uMail = data.mail;
                        this.uName = data.name;
                        this.uSurnames = data.surnames;
                        this.uImg = data.img;
                        this.uRole = data.role;
                        this.uBlocked = data.is_blocked.toString();

                        this.showUpdate();
                    
                    } else {
                        this.toastr.error("You must select one.", {
                            timeout: 1500
                        });
                    }
                }

            }
        },
        gridSizeChanged(params) {

            this.gridApi.value = params.api;
            this.columnApi.value = params.columnApi;

            this.gridApi.value.sizeColumnsToFit();

        },
        showCreate() {

            if (this.create) {
                this.create = false;
                this.update = false;
                this.cMail = '',
                this.cPass = '',
                this.cName = '',
                this.cSurnames = '',
                this.cImg = '',
                this.cRole = '',
                this.cBlocked = ''
            } else {
                this.create = true;
                this.update = false;
            }
            
        },
        showUpdate() {
       
            this.create = false;
            this.update = true;
            
        },
        async createCourt() {
            if (this.cMail.length != 0 && this.cPass.length != 0 && this.cName.length != 0 && this.cSurnames.length != 0 && this.cImg.length != 0 && this.cRole.length != 0 && this.cBlocked.length != 0) {

                var regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
                var blocked = this.checkIfBlocked();

                if (regex.test(this.cMail) != false && (blocked != 'err')) { 
                    this.store.dispatch("user/" + Constant.CREATE_USER, {'mail': this.cMail, 'pass':this.cPass, 'name': this.cName, 'surnames': this.cSurnames, 'img': this.cImg, 'role': this.cRole, 'is_blocked': blocked});
                } else if (regex.test(this.uMail) != false) {
                    this.toastr.error("Email must be like example@example.example", {
                        timeout: 1500
                    }); 
                }


            } else {
                this.toastr.error("All fields must be filled.", {
                    timeout: 1500
                });
            }
        },
        async updateCourt() {
            if (this.uMail.length != 0 && this.uPass.length != 0 && this.uName.length != 0 && this.uSurnames.length != 0 && this.uImg.length != 0 && this.uRole.length != 0 && this.uBlocked.length != 0) {

                var regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
                var blocked = this.checkIfBlocked();

                if (regex.test(this.uMail) != false && (blocked != 'err')) {
                    this.store.dispatch("user/" + Constant.UPDATE_USER_ADMIN, {'uuid': this.uUUID, 'mail': this.uMail, 'pass':this.uPass, 'name': this.uName, 'surnames': this.uSurnames, 'img': this.uImg, 'role': this.uRole, 'is_blocked': blocked});
                } else if (regex.test(this.uMail) != false) {
                    this.toastr.error("Email must be like example@example.example", {
                        timeout: 1500
                    }); 
                }

            } else {
                this.toastr.error("All fields must be filled.", {
                    timeout: 1500
                }); 
            }
        },
        checkIfBlocked() {
            if (this.create) {
                if (this.cBlocked == 'true') {
                    return true
                } else if (this.cBlocked == 'false') {
                    return false
                } else {
                    this.toastr.error("Only true or false are valid in Blocked field.", {
                        timeout: 1500
                    });
                    return 'err'
                }
            } else if (this.update) {
                if (this.uBlocked == 'true') {
                    return true
                } else if (this.uBlocked == 'false') {
                    return false
                } else {
                    this.toastr.error("Only true or false are valid in Blocked field.", {
                        timeout: 1500
                    });
                    return 'err'
                }
            }
        }
    },

})
</script>
