<template>
    <div>
        <h3 class="text-center w-100">{{ $t("TOTAL_RESERVATIONS_BY_SPORT") }}</h3>
        <column-chart :data="arrSport"></column-chart>
        <br/>
        <h3 class="text-center">{{ $t("TOTAL_RESERVATIONS_BY_COURT") }}</h3>
        <column-chart :data="arrCourt"></column-chart>
    </div>
</template>

<script>
import { useStore } from 'vuex';
import { useGetUserSportReservation,useGetUserCourtReservation } from '../composables/reservations/useReservations';
export default ({
    async setup() {
        const store = useStore();
        var id = store.state.user.user.id;
        var reserArr =[id,'admin']
        var reservationSport = await useGetUserSportReservation(reserArr);
        var ReservationSportValue = reservationSport.reservations.value
        var reservationCourt = await useGetUserCourtReservation();
        var ReservationCourtValue = reservationCourt.reservations.value
        var arrSport = []
        var arrCourt = []
        if(ReservationSportValue.length == 0){
          var sport = 'No Sports'
          var count = 0
          var sportArr = [sport,count]
          arrSport.push(sportArr)
        }else{
          for(var i=0;i<ReservationSportValue.length;i++){
            var sport = ReservationSportValue[i].name.toUpperCase() 
            var count = ReservationSportValue[i].numreser
            var sportArr = [sport,count]
            arrSport.push(sportArr)
          }
        }
        if(ReservationCourtValue.length == 0){
          var court = 'No Sports'
          var count = 0
          var courtArr = [court,count]
          arrCourt.push(courtArr)
        }else{
          for(var i=0;i<ReservationCourtValue.length;i++){
            var court = "Court nº "+ReservationCourtValue[i].id 
            var count = ReservationCourtValue[i].numreser
            var courtArr = [court,count]
            arrCourt.push(courtArr)
          }
        }
        return {arrSport,arrCourt}
    },
})
</script>
