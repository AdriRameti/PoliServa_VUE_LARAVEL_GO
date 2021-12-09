<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;

class ReservationFactory extends Factory
{
    /**
     * Define the model"s default state.
     *
     * @return array
     */
    public function definition()
    {
        return [
            "id_user"=>1,
            "id_court"=>$this->idCourt(),
            "date"=>$this->date(),
            "hini"=>$this->hini(),
            "hfin"=>$this->hfin(),
            "total_price"=>$this->totalPrice()
        ];
    }
    private function idCourt(){
        $randNum = random_int(1,10);
        return $randNum;
    }
    private function date(){
        $day = random_int(1,31);
        $month = random_int(1,12);
        $year = random_int(2021,2050);
        $date = ''.$day.'/'.$month.'/'.$year.'';
        return $date;
    }
    private function hini(){
        $firstHour = random_int(0,2);
        if($firstHour==2){
            $secondHour = random_int(0,3);
        }else{
            $secondHour = random_int(0,9);
        }
        $string1 = ''.$firstHour.$secondHour.'';
        $firstMinute = random_int(0,5);
        $secondMinute = random_int(0,9);
        $string2 = ''.$firstMinute.$secondMinute.'';
        $hour = ''.$string1.':'.$string2.'';
        return $hour;
    }
    private function hfin(){
        $firstHour = random_int(0,2);
        if($firstHour==2){
            $secondHour = random_int(0,3);
        }else{
            $secondHour = random_int(0,9);
        }
        $string1 = ''.$firstHour.$secondHour.'';
        $firstMinute = random_int(0,5);
        $secondMinute = random_int(0,9);
        $string2 = ''.$firstMinute.$secondMinute.'';
        $hour = ''.$string1.':'.$string2.'';
        return $hour;
    }
    private function totalPrice(){
        $price = random_int(1,100);
        return $price;
    }
}
