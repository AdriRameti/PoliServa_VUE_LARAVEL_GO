<template>
  <div class="container">
    <div class="row profile">
		<div class="col-md-3">
			<div class="profile-sidebar">
				<!-- SIDEBAR USERPIC -->
				<div class="profile-userpic">
					<img :src="store.state.user.user.img" class="img-responsive" alt="">
				</div>
				<!-- END SIDEBAR USERPIC -->
				<!-- SIDEBAR USER TITLE -->
				<div class="profile-usertitle">
					<div class="profile-usertitle-name">
						{{ store.state.user.user.fullName }}
					</div>
					<div class="profile-usertitle-job">
						
					</div>
				</div>
				<!-- END SIDEBAR USER TITLE -->
				<!-- SIDEBAR BUTTONS -->
				<div class="profile-userbuttons">
					<button type="button" class="btn btn-success btn-sm" data-bs-target="#modal2fa" data-bs-toggle="modal" v-show="!store.state.user.google2fa_secret" @click="enable2fa()">Enable 2fa</button>
          <button type="button" class="btn btn-danger btn-sm" v-show="store.state.user.google2fa_secret" @click="disable2fa()">Dissable 2fa</button>

					<button type="button" class="btn btn-danger btn-sm">Delete user</button>
				</div>
				<!-- END SIDEBAR BUTTONS -->
				<!-- SIDEBAR MENU -->
				<div class="profile-usermenu">
					<ul class="nav">
						<li class="active">
							<a @click="showOverview()">
							<i class="glyphicon glyphicon-home"></i>
							Overview </a>
						</li>
						<li>
							<a @click="showAccSettings()">
							<i class="glyphicon glyphicon-user"></i>
							Account Settings </a>
						</li>
					</ul>
				</div>
				<!-- END MENU -->
			   
           <div class="portlet light bordered">
            <!-- STAT -->
            <div class="row list-separated profile-stat">
              <div class="col-md-4 col-sm-4 col-xs-6">
                  <div class="uppercase profile-stat-title"> 51 </div>
                  <div class="uppercase profile-stat-text"> Reservations </div>
              </div>
            </div>
            <!-- END STAT -->
          </div>                   
                                          
			</div>
		</div>
		<div class="col-md-9">
          <div class="profile-content">
            <p v-show="overview">Hola estadisticas</p>

            <section v-show="accSettngs">
              <div class="container h-100">
                  <div class="col-12 col-md-8 col-lg-6 col-xl-5">
                      <div class="card-body p-4">
                        <h3 class="mb-3">Account Settings</h3>

                        <div class="form-outline mb-4">
                          <label for="typeNameX-2" class="form-label float-left">Name</label>
                          <input type="email" id="typeNameX-2" class="form-control form-control-lg" v-model="updateNameValue" :placeholder="user.name"/>
                        </div>

                        <div class="form-outline mb-4">
                          <label for="typeSurnamesX-2" class="form-label float-left">Surnames</label>
                          <input type="email" id="typeSurnamesX-2" class="form-control form-control-lg" v-model="updateSurnamesValue" :placeholder="user.surnames"/>
                        </div>

                        <div class="form-outline mb-4">
                          <label for="typeEmailX-2" class="form-label float-left">Email</label>
                          <input type="email" id="typeEmailX-2" class="form-control form-control-lg" v-model="updateMailValue" :placeholder="user.mail"/>
                        </div>

                        <div class="form-outline mb-4">
                          <label for="typePasswordX-2" class="form-label float-left">Password</label>
                          <input type="password" id="typePasswordX-2" v-model="updatePassValue" class="form-control form-control-lg" placeholder="Password"/>
                        </div>

                        <div class="form-outline mb-4">
                          <label for="typeProfileX-2" class="form-label float-left">Profile image</label>
                          <input type="file" id="typeProfileX-2" class="form-control form-control-lg" @change="previewFile"/>
                        </div>
                        <div class="profile-userpic mt-2 mb-3" v-show="profileImgPreview">
                          <p>Profile photo preview</p>
                          <img :src="profileImgURL" class="img-responsive" alt="">
                        </div>

                        <button class="btn btn-primary btn-lg btn-block" v-on:click="updateUser()">Update</button>
                      </div>
                </div>
              </div>
            </section>


          </div>
		</div>
	</div>


  <div id="modal2fa" class="modal fade" data-bs-backdrop="static" data-bs-keyboard="false">
    <div class="modal-dialog modal-dialog-centered">
        <div class="modal-content">
            <div class="modal-header">
                <h4 class="modal-title">Two-Factor Authentication</h4>
                <button type="button" class="close btn btn-dark" data-bs-dismiss="modal" @click="disable2fa()">
                    <span aria-hidden="true">X</span>
                </button>
            </div> 
            <div class="modal-body">
              <h4>Two-factor authentication increases the security of your Ledgy account.</h4>
              <p>All you need is a compatible app on your smartphone, for example:</p>
              <ul>
                <li>Google Authentificator</li>
                <li>Duo</li>
                <li>Authy</li>
              </ul>
              <div v-html="store.state.user.user.QR"></div>
              <p>Scan this image with your app. You will see a 6-digit code on your screen. Enter the code below to verify your phone and complete the setup.</p>
              <div class="form-outline mb-4">
                <input type="text" v-model="otp" class="form-control form-control-lg" placeholder="One Time Password"/>
              </div>
              <button class="btn btn-primary btn-lg btn-block" v-on:click="check2fa()" v-show="!store.state.user.g2faverified" :disabled="otp.length != 6">Verify Two-Factor Authentication</button>
              <button class="btn btn-primary btn-lg btn-block" v-show="store.state.user.g2faverified" data-bs-dismiss="modal">Close</button>
            </div>
        </div>
    </div>
  </div>
</div>
</template>

<script>

import { useStore } from 'vuex';
import Constant from "../Constant";

export default {
  setup() {
    const store = useStore();

    console.log(window.location.origin)

    return { store }
  },
  data() {
    return {
      user: this.store.getters["user/getUser"] ? this.store.getters["user/getUser"] : {},
      otp: '',
      overview: true,
      accSettngs: false,
      updateNameValue: '',
      updateSurnamesValue: '',
      updateMailValue: '',
      updatePassValue: '',
      profileImgURL: '',
      profileImgPreview: false
    }
  },
  methods: {
    enable2fa() {
      this.store.dispatch("user/" + Constant.ENABLE2FA);
    },
    disable2fa() {
      this.store.dispatch("user/" + Constant.DISABLE2FA);
    },
    check2fa() {
      this.store.dispatch("user/" + Constant.CHECK2FA, {'otp': this.otp});
    },
    showOverview() {
      this.overview = true;
      this.accSettngs = false;
    },
    showAccSettings() {
      this.overview = false;
      this.accSettngs = true;
    },
    previewFile(event) {
      this.profileImg = event.target.files[0];

      if (this.profileImg.type.split('/')[0] == 'image') {
        this.profileImgURL = URL.createObjectURL(this.profileImg);
        this.profileImgPreview = true;
      } else {
        this.profileImgPreview = false;
      }
    },
    updateUser() {

      var userUpdate = new FormData()


      if (this.updateNameValue) {
        userUpdate.append('name', this.updateNameValue);
      }

      if (this.updateSurnamesValue) {
        userUpdate.append('surnames', this.updateSurnamesValue);
      }

      if (this.updateMailValue) {
        userUpdate.append('mail', this.updateMailValue);
      }

      if (this.updatePassValue) {
        userUpdate.append('pass', this.updatePassValue);
      }

      if (this.profileImg) { 
        userUpdate.append('img', this.profileImg);
      }

      userUpdate.append('_method', 'PUT');

      this.store.dispatch("user/" + Constant.UPDATE_USER, userUpdate);

    }
  }
}
</script>

<style>
body {
  background: #F1F3FA;
}

/* Profile container */
.profile {
  margin: 20px 0;
}

/* Profile sidebar */
.profile-sidebar {
  padding: 20px 0 10px 0;
  background: #fff;
}

.profile-userpic img {
  float: none;
  margin: 0 auto;
  width: 50%;
  height: 50%;
  -webkit-border-radius: 50% !important;
  -moz-border-radius: 50% !important;
  border-radius: 50% !important;
}

.profile-usertitle {
  text-align: center;
  margin-top: 20px;
}

.profile-usertitle-name {
  color: #5a7391;
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 7px;
}

.profile-usertitle-job {
  text-transform: uppercase;
  color: #5b9bd1;
  font-size: 12px;
  font-weight: 600;
  margin-bottom: 15px;
}

.profile-userbuttons {
  text-align: center;
  margin-top: 10px;
}

.profile-userbuttons .btn {
  text-transform: uppercase;
  font-size: 11px;
  font-weight: 600;
  padding: 6px 15px;
  margin-right: 5px;
}

.profile-userbuttons .btn:last-child {
  margin-right: 0px;
}
    
.profile-usermenu {
  margin-top: 30px;
}

.profile-usermenu ul li {
  border-bottom: 1px solid #f0f4f7;
}

.profile-usermenu ul li:last-child {
  border-bottom: none;
}

.profile-usermenu ul li a {
  color: #93a3b5;
  font-size: 14px;
  font-weight: 400;
}

.profile-usermenu ul li a i {
  margin-right: 8px;
  font-size: 14px;
}

.profile-usermenu ul li a:hover {
  background-color: #fafcfd;
  color: #5b9bd1;
}

.profile-usermenu ul li.active {
  border-bottom: none;
}

.profile-usermenu ul li.active a {
  color: #5b9bd1;
  background-color: #f6f9fb;
  border-left: 2px solid #5b9bd1;
  margin-left: -2px;
}

/* Profile Content */
.profile-content {
  padding: 20px;
  background: #fff;
  min-height: 460px;
}

a, button, code, div, img, input, label, li, p, pre, select, span, svg, table, td, textarea, th, ul {
    -webkit-border-radius: 0!important;
    -moz-border-radius: 0!important;
    border-radius: 0!important;
}
.dashboard-stat, .portlet {
    -webkit-border-radius: 4px;
    -moz-border-radius: 4px;
    -ms-border-radius: 4px;
    -o-border-radius: 4px;
}
.portlet {
    margin-top: 0;
    margin-bottom: 25px;
    padding: 0;
    border-radius: 4px;
}
.portlet.bordered {
    border-left: 2px solid #e6e9ec!important;
}
.portlet.light {
    padding: 12px 20px 15px;
    background-color: #fff;
}
.portlet.light.bordered {
    border: 1px solid #e7ecf1!important;
}
.list-separated {
    margin-top: 10px;
    margin-bottom: 15px;
}
.profile-stat {
    padding-bottom: 20px;
    border-bottom: 1px solid #f0f4f7;
}
.profile-stat-title {
    color: #7f90a4;
    font-size: 25px;
    text-align: center;
}
.uppercase {
    text-transform: uppercase!important;
}

.profile-stat-text {
    color: #5b9bd1;
    font-size: 10px;
    font-weight: 600;
    text-align: center;
}
.profile-desc-title {
    color: #7f90a4;
    font-size: 17px;
    font-weight: 600;
}
.profile-desc-text {
    color: #7e8c9e;
    font-size: 14px;
}
.margin-top-20 {
    margin-top: 20px!important;
}
[class*=" fa-"]:not(.fa-stack), [class*=" glyphicon-"], [class*=" icon-"], [class^=fa-]:not(.fa-stack), [class^=glyphicon-], [class^=icon-] {
    display: inline-block;
    line-height: 14px;
    -webkit-font-smoothing: antialiased;
}
.profile-desc-link i {
    width: 22px;
    font-size: 19px;
    color: #abb6c4;
    margin-right: 5px;
}
</style>