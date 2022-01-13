<template>
    <div>
        <div class="text-center">
            <h1>Courts</h1>
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
            <input id="cID_sport" v-model="cID_sport" class="form-control form-control-sm" type="text" placeholder="Id_sport">
            <input id="cImg" v-model="cImg" class="form-control form-control-sm" type="text" placeholder="Image">
            <input id="cPriceH" v-model="cPriceH" class="form-control form-control-sm w-75" type="text" placeholder="Precio/h">
            <input id="cSector" v-model="cSector" class="form-control form-control-sm w-75" type="text" placeholder="Sector">
            <button v-on:click="createCourt()" class="btn btn-primary m-1">Send</button>
        </div>

        <div class="d-flex d-inline-block" v-if="update">
            <input id="uID_sport" v-model="uID_sport" class="form-control form-control-sm" type="text" :placeholder="uID_sport">
            <input id="uImg" v-model="uImg" class="form-control form-control-sm" type="text" :placeholder="uImg">
            <input id="uPriceH" v-model="uPriceH" class="form-control form-control-sm w-75" type="text" :placeholder="uPriceH">
            <input id="uSector" v-model="uSector" class="form-control form-control-sm w-75" type="text" :placeholder="uSector">
            <button v-on:click="updateCourt()" class="btn btn-primary m-1">Send</button>
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
import { useGetAllCourts, useDeleteCourt, useCreateCourt, useUpdateCourt } from '../composables/courts/useGetAllCourts'

import { AgGridVue } from "ag-grid-vue3";
import { ref } from "vue";

export default({
    components: {
        AgGridVue,
    },
    async setup() {
        const toastr = useToast();

        const { courts, count } = await useGetAllCourts();

        const rowData = ref(courts.value);

        return { rowData, toastr }
    },
    data() {
        return {
            columnDefs: [
                { headerName: "ID", field: "id", sortable: true, filter: true, checkboxSelection: true },
                { headerName: "Sport ID", field: "id_sport", sortable: true },
                { headerName: "Image", field: "img", resizable: true },
                { headerName: "Price/h", field: "price_h", sortable: true },
                { headerName: "Sector", field: "sector", sortable: true}
            ],
            gridApi: ref({}),
            columnApi: ref({}),
            grid: 0,
            create: false,
            update: false,
            cID_sport: '',
            cImg: '',
            cPriceH: '',
            cSector: '',
            uID: '',
            uID_sport: '',
            uImg: '',
            uPriceH: '',
            uSector: '',
        }
    },
    methods:{
        async getSelectedRows(op) {
            const selectedNodes = this.gridApi.value.getSelectedNodes();
            const selectedData = selectedNodes.map( node => node.data );
            const sportSlug = selectedData.map( node => `${node.Sport.slug}`);

            if (op == "delete") {
                var del = await useDeleteCourt(selectedData[0].id);

                if (del != 'err') {

                    var newData = this.rowData.filter((data) => { return data.id != selectedData[0].id })
                    
                    this.rowData = newData;

                    this.grid += 1;
                }
            } else if (op == "update") {

                if (this.update && this.uID_sport == selectedData[0].id_sport) {
                    this.update = false;
                    this.create = false;
                } else {
                    if (selectedNodes.length == 1) {

                        const data = selectedData[0];

                        this.uID = data.id;
                        this.uID_sport = data.id_sport;
                        this.uImg = data.img;
                        this.uPriceH = data.price_h;
                        this.uSector = data.sector;

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
                this.cID_sport = '';
                this.cImg = '';
                this.cPriceH = '';
                this.cSector = '';
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
            if (this.cID_sport.length != 0 && this.cImg.length != 0 && this.cPriceH.length != 0 && this.cSector.length != 0) {

                var resC = await useCreateCourt({'id_sport': this.cID_sport, 'img': this.cImg, 'price_h': this.cPriceH, 'sector': this.cSector});

                if (resC != 'err') {

                    this.rowData.push(resC.court.value);

                    this.grid += 1;
                    this.cID_sport = '';
                    this.cImg = '';
                    this.cPriceH = '';
                    this.cSector = '';
                }

            } else {
                this.toastr.error("All fields must be filled.", {
                    timeout: 1500
                });
            }
        },
        async updateCourt() {
            if (this.uID_sport.length != 0 && this.uImg.length != 0 && this.uPriceH.length != 0 && this.uSector.length != 0) {

                var resU = await useUpdateCourt({'id': this.uID,'id_sport': this.uID_sport, 'img': this.uImg, 'price_h': this.uPriceH, 'sector': this.uSector});

                if (resU != 'err') {
                    
                    var newData = this.rowData.map(data => data.id == resU.court.value.id ? resU.court.value : data );
        
                    this.rowData = newData;
                
                    this.update = false;
                    this.uSlug = '';
                    this.uName = '';
                    this.uImg = '';
                    
                    this.grid += 1;

                }

            } else {
                this.toastr.error("All fields must be filled.", {
                    timeout: 1500
                }); 
            }
        }
    },

})
</script>
