<template>
  <section class="container mt-5">
   <!--Contact heading-->
   <div class="row">
      <!--Grid column-->
      <div class="col-sm-12 mb-4 col-md-5 centrado">
         <!--Form with header-->
         <div class="card rounded-0">
            <div class="card-header p-0">
               <div class="header-contact text-white text-center py-2">
                  <h3><i class="fa fa-envelope"></i> {{ $t("REPORT_ISSUE") }}:</h3>
               </div>
            </div>
            <div class="card-body p-3">
                  <div class="form-group">
                     <label>{{ $t("SUBJECT") }}</label>
                     <div class="input-group mb-2 mb-sm-0">
                        <input type="text" name="data[subject]" class="form-control" id="inlineFormInputGroupUsername" :placeholder="$t('ENTER_THE_SUBJECT')" v-model="subject">
                     </div>
                  </div>
                  <div class="form-group">
                     <label>{{ $t("MESSAGE") }}</label>
                     <div class="input-group mb-2 mb-sm-0">
                        <input type="text" class="form-control" name="mesg" :placeholder="$t('ENTER_THE_ISSUE')" v-model="issue">
                     </div>
                  </div>
                  <br/>
                  <div class="text-center">
                     <button type="button" class="btn btn-block rounded-0 py-2 text-white" @click="send_issue()">{{ $t("SEND") }}</button>
                  </div>
             
			     </div>
				  
            </div>
        </div>

	</div>



</section>
</template>

<script>
import { useStore } from 'vuex';
import Constant from "../Constant";

export default {
   methods: {
      send_issue() {
         this.store.dispatch("user/" + Constant.SEND_ISSUE, {subject: this.subject, issue: this.issue});
      }
   },
   data() {
      return {
         subject: "",
         issue: ""
      }
   },
   setup() {
    const store = useStore();

    return { store }
   }
}
</script>

<style>

section .row .centrado {
    margin: 0 auto;
}

.btn{
   background-color: #1B4332 !important;
}
.btn-info{
   background-color: #1B4332 !important;
}
</style>