<template>
    <div class="w-100">

    <div class="container-fluid">
    <div class="row flex-nowrap">
        <div class="background-lateral-menu col-auto col-md-3 col-xl-2 px-sm-2 px-0">
            <div class="d-flex flex-column align-items-center align-items-sm-start px-3 pt-2 text-white min-vh-100">
                <ul class="nav nav-pills flex-column mb-sm-auto mb-0 align-items-center align-items-sm-start" id="menu">
                    <li>
                        <a class="nav-link px-0 align-middle text-white" v-on:click="showDashboard()">
                            <i class="fab fa-accessible-icon"></i> <span class="ms-1 d-none d-sm-inline">{{ $t("STATS") }}</span> 
                        </a>
                    </li>
                    <li>
                        <a class="nav-link px-0 align-middle text-white" v-on:click="showSports()">
                            <i class="fs-4 bi-table"></i> <span class="ms-1 d-none d-sm-inline">{{ $t("SPORTS") }}</span></a>
                    </li>
                    <li>
                        <a class="nav-link px-0 align-middle text-white" v-on:click="showCourts()">
                            <i class="fs-4 bi-bootstrap"></i> <span class="ms-1 d-none d-sm-inline">{{ $t("COURTS") }}</span>
                        </a>
                    </li>
                    <li>
                        <a class="nav-link px-0 align-middle text-white" v-on:click="showUsers()">
                            <i class="fs-4 bi-bootstrap"></i> <span class="ms-1 d-none d-sm-inline">{{ $t("USERS") }}</span>
                        </a>
                    </li>
                </ul>
                <hr>
            </div>
        </div>
        <div class="col py-3">
            <div class="w-100 h-100" v-show="dashboard">
                <Suspense>
                    <AdminStats></AdminStats>
                </Suspense>  
            </div>  
            <div class="w-100 h-100" v-show="sports">
                <Suspense>
                    <ListSportsAdmin></ListSportsAdmin>
                </Suspense>  
            </div>
            <div class="w-100 h-100" v-show="courts">
                <Suspense>
                    <ListCourtsAdmin></ListCourtsAdmin>
                </Suspense>  
            </div>
            <div class="w-100 h-100" v-show="users">
                <Suspense>
                    <ListUsersAdmin></ListUsersAdmin>
                </Suspense>  
            </div>    
        </div>
    </div>
</div>
        
    </div>
</template>
<script>
//import { defineComponent } from '@vue/composition-api'
import AdminStats from '@/components/AdminStats.vue'
import ListSportsAdmin from '@/components/ListSportsAdmin.vue'
import ListCourtsAdmin from '@/components/ListCourtsAdmin.vue'
import ListUsersAdmin from '@/components/ListUsersAdmin.vue'

export default ({
    components: {
        AdminStats,
        ListSportsAdmin,
        ListCourtsAdmin,
        ListUsersAdmin
    },
  data() {
    return {
        dashboard: true,
        sports: false,
        courts: false,
        users: false
    }   
  },
  methods: {
    showDashboard() {
        this.dashboard = true;
        this.sports = false;
        this.courts = false;
        this.users = false;
    },
    showSports() {
        this.dashboard = false;
        this.sports = true;
        this.courts = false;
        this.users = false;
    },
    showCourts() {
        this.dashboard = false;
        this.sports = false;
        this.courts = true;
        this.users = false;
    },
    showUsers() {
        this.dashboard = false;
        this.sports = false;
        this.courts = false;
        this.users = true;
    }
  },
  setup() {

   
  },
  
})
</script>

<style>
    .background-lateral-menu{
        background-color: #1B4332;
    }
</style>
<style lang="scss">
    a {
        cursor: pointer;
    }
</style>