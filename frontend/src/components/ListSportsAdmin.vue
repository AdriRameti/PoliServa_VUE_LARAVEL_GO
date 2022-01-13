<template>
    <div>
        <div class="text-center">
            <h1>Sports</h1>
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
            <input id="cName" v-model="cName" class="form-control form-control-sm w-75" type="text" placeholder="Name">
            
            <input id="cImg" v-model="cImg" class="form-control form-control-sm w-75" type="text" placeholder="Image">
            <button v-on:click="createSport()" class="btn btn-primary m-1">Send</button>
        </div>

        <div class="d-flex d-inline-block" v-if="update">
            <input id="uName" v-model="uName" class="form-control form-control-sm" type="text" :placeholder="uName">
            <input id="uImg" v-model="uImg" class="form-control form-control-sm" type="text" :placeholder="uImg">
            <button v-on:click="updateSport()" class="btn btn-primary m-1">Send</button>
        </div>

        <ag-grid-vue id="myGrid" style="width: 100%;"
            class="ag-theme-alpine"
            :columnDefs="columnDefs"
            :rowData="rowData"
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
import { useGetAllSports, useDeleteSport, useCreateSport, useUpdateSport } from '../composables/sports/useSports';

import { AgGridVue } from "ag-grid-vue3";
import { ref, reactive } from "vue";

export default({
    components: {
        AgGridVue,
    },
    async setup() {
        const toastr = useToast();

        const { sports } = await useGetAllSports();

        const rowData = reactive(sports.value);

        return { rowData, toastr }
    },
    data() {
        return {
            columnDefs: [
                { headerName: "ID", field: "id", sortable: true, filter: true, checkboxSelection: true },
                { headerName: "Slug", field: "slug", sortable: true, filter: true },
                { headerName: "Name", field: "name", sortable: true },
                { headerName: "Image", field: "img", resizable: true },
            ],
            gridApi: ref({}),
            columnApi: ref({}),
            grid: 0,
            create: false,
            update: false,
            cName: '',
            cImg: '',
            uName: '',
            uImg: ''
        }
    },
    methods:{
        async getSelectedRows(op) {
            const selectedNodes = this.gridApi.value.getSelectedNodes();
            const selectedData = selectedNodes.map( node => node.data );
            const slug = selectedData.map( node => `${node.slug}`);

            if (op == "delete") {
                var del = await useDeleteSport(slug[0]);

                if (del != 'err') {
                    this.rowData = del.sports.value;
                    this.grid += 1;
                }
            } else if (op == "update") {

                if (this.update && this.uName == selectedData[0].name) {
                    this.update = false;
                    this.create = false;
                } else {
                    if (selectedNodes.length == 1) {

                        const data = selectedData[0];

                        this.uName = data.name;
                        this.uImg = data.img;
                        this.uSlug = slug;

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
                this.cName = '';
                this.cImg = '';
            } else {
                this.create = true;
                this.update = false;
            }
            
        },
        showUpdate() {
       
            this.create = false;
            this.update = true;
            
        },
        async createSport() {
            if (this.cName.length != 0 && this.cImg.length != 0) {

                var resC = await useCreateSport({'name': this.cName, 'img': this.cImg});

                if (resC != 'err') {

                    this.rowData.push(resC.sport.value);
                    this.grid += 1;
                    this.cName = '';
                    this.cImg = '';
                }

            } else {
                this.toastr.error("Both fields must be filled.", {
                    timeout: 1500
                });
            }
        },
        async updateSport() {
            if (this.uName.length != 0 && this.uImg.length != 0 && this.uSlug.length != 0) {

                var resU = await useUpdateSport({'slug': this.uSlug, 'name': this.uName, 'img': this.uImg});

                if (resU != 'err') {
                    
                    var newData = this.rowData.map(data => data.id == resU.sport.value.id ? resU.sport.value : data );
        
                    this.rowData = newData;
                
                    this.uSlug = '';
                    this.uName = '';
                    this.uImg = '';
                    
                    this.grid += 1;

                }

            } else {
                this.toastr.error("Both fields must be filled.", {
                    timeout: 1500
                }); 
            }
        }
    },

})
</script>
