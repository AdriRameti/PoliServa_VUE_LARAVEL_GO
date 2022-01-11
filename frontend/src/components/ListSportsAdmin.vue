<template>
    
    <h1>Deportes</h1>

    <div class="mb-2 d-flex justify-content-end">
        <button @click="getSelectedRows()" class="btn btn-success m-1">Create</button>
        <button @click="getSelectedRows()" class="btn btn-primary m-1">Update</button>
        <button @click="getSelectedRows('delete')" class="btn btn-danger m-1">Delete</button>
    </div>

    <ag-grid-vue id="myGrid" style="width: 100%;"
        class="ag-theme-alpine"
        :columnDefs="columnDefs"
        :rowData="rowData"
        rowSelection="single"
        domLayout='autoHeight'
        @gridSizeChanged="gridSizeChanged" :key="grid">
    </ag-grid-vue>

</template>

<style>

</style>
<script>

import { useToast } from "vue-toastification";
import { useGetAllSports, useDeleteSport } from '../composables/sports/useSports';

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
            grid: 0
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
            }
        },
        gridSizeChanged(params) {

            this.gridApi.value = params.api;
            this.columnApi.value = params.columnApi;

            this.gridApi.value.sizeColumnsToFit();

        }
    },

})
</script>
