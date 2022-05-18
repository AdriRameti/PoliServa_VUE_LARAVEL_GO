<template>
  <section class="container mt-5">
   <!--Contact heading-->
   <div class="row">
      <!--Grid column-->
      <div class="col-sm-12 mb-4 col-md-5">
         <!--Form with header-->
         <div class="card rounded-0">
            <div class="card-header p-0">
               <div class="header-contact text-center py-2">
                  <h3 class="text-white"><i class="fa fa-envelope"></i> {{ $t("CONTACT_WITH_US") }}:</h3>
                  <p class="m-0">{{ $t("COMMUNICATION") }}</p>
               </div>
            </div>
            <div class="card-body p-3">
               
                  <div class="form-group">
                  <label> {{ $t("NAME") }} </label>
                  <div class="input-group">
                     <input value="" type="text" name="data[name]" class="form-control" id="inlineFormInputGroupUsername" :placeholder="$t('ENTER_YOUR_NAME')">
                  </div>
				</div>
                  <div class="form-group">
                     <label>{{ $t("EMAIL") }}</label>
                     <div class="input-group mb-2 mb-sm-0">
                        <input type="email" value="" name="data[email]" class="form-control" id="inlineFormInputGroupUsername" :placeholder="$t('ENTER_THE_EMAIL')">
                     </div>
                  </div>
                  <div class="form-group">
                     <label>{{ $t("SUBJECT") }}</label>
                     <div class="input-group mb-2 mb-sm-0">
                        <input type="text" name="data[subject]" class="form-control" id="inlineFormInputGroupUsername" :placeholder="$t('ENTER_THE_SUBJECT')">
                     </div>
                  </div>
                  <div class="form-group">
                     <label>{{ $t("MESSAGE") }}</label>
                     <div class="input-group mb-2 mb-sm-0">
                        <input type="text" class="form-control" name="mesg" :placeholder="$t('ENTER_THE_MESSAGE')">
                     </div>
                  </div>
                  <br/>
                  <div class="text-center">
                     <input type="submit" name="submit" :value="$t('SEND')" class="btn btn-block rounded-0 py-2 text-white">
                  </div>
             
			     </div>
				  
            </div>
         </div>
      <!--Grid column-->
	  
      <!--Grid column-->
      <div class="col-sm-12 col-md-7">
         <!--Google map-->
         <div class="mb-4">
            <iframe src="https://www.google.com/maps/d/u/0/embed?mid=1gHLFZwOKEl7B7kEmS9Trtukj-mqi6NVY&ehbc=2E312F" width="100%" height="450" frameborder="0" style="border:0;" allowfullscreen="" aria-hidden="false" tabindex="0"></iframe>
         </div>
         <!--Buttons-->
         <div class="row text-center">
            <div class="col-md-4">
               <a class="btn-info px-3 py-2 rounded text-white mb-2 d-inline-block"><i class="fa fa-map-marker"></i></a>
               <p> Plaza del ayuntamiento </p>
            </div>
            <div class="col-md-4">
               <a class="btn-info px-3 py-2 rounded text-white mb-2 d-inline-block"><i class="fa fa-phone"></i></a>
               <p>+34 - 654 65 46 54</p>
            </div>
            <div class="col-md-4">
               <a class="btn-info px-3 py-2 rounded text-white mb-2 d-inline-block"><i class="fa fa-envelope"></i></a>
               <p> poliserva@info.com</p>
            </div>
         </div>
      </div>
      <!--Grid column-->
	    </div>


</section>
</template>

<script>
export default {

}
</script>

<style>
.header-contact {
   background-color: #1B4332 !important;
}
.btn{
   background-color: #1B4332 !important;
}
.btn-info{
   background-color: #1B4332 !important;
}
</style>