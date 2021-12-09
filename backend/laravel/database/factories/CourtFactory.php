<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;

class CourtFactory extends Factory
{
    /**
     * Define the model's default state.
     *
     * @return array
     */
    public function definition()
    {
        return [
            "id_sport"=>$this->IdSport(),
            "sector"=>$this->Sector(),
            "price_h"=>$this->Price(),
            "img"=>$this->imgCourt()
        ];
    }
    private function IdSport(){
        $idSport = random_int(1,7);
        $_SESSION['IdSport']=$idSport;
        return $idSport;
    }

    private function Sector(){
        $randNum = random_int(1,5);
        $_SESSION['Sector'] = $randNum;
        return $randNum;
    }
    private function Price(){
        $price = 0;
        switch($_SESSION['IdSport']){
            case 1:
                $price = 8;
                break;
            case 2:
                $price = 15;
                break;
            case 3:
                $price = 20;
                break;
            case 4:
                $price = 25;
                break;
            case 5:
                $price = 25;
                break;
            case 6:
                $price = 50;
                break;
            case 7:
                $price = 6;
                break;
        }
        return $price;
    }
    private function imgCourt(){
        $img=array(
            "https://ucjcsportsclub.es/wp-content/uploads/2015/03/padel3-1024x780.jpg",
            "https://www.bbva.com/wp-content/uploads/2017/08/bbva-balon-futbol-2017-08-11.jpg",
            "https://www.fundeu.es/wp-content/uploads/2021/01/voleibol-webefesptwelve224491.jpg",
            "https://mundoentrenamiento.com/wp-content/uploads/2017/05/doble-penalti-en-futbol-sala-800x500.jpg",
            "https://www.basquetplus.com/sites/default/files/main/articles/redes.jpg",
            "https://raquetados.com/wp-content/uploads/2019/09/Frontenis.jpg",
            "https://static2.abc.es/media/bienestar/2019/08/02/tenis-abecedario-kgNF--620x349@abc.jpg"
            );
        $image = "";
        switch($_SESSION['IdSport']){
            case 1:
                $image = $img[6];
                break;
            case 2:
                $image = $img[5];
                break;
            case 3:
                $image = $img[4];
                break;
            case 4:
                $image = $img[3];
                break;
            case 5:
                $image = $img[2];
                break;
            case 6:
                $image = $img[1];
                break;
            case 7:
                $image = $img[0];
                break;
        }
        return $image;
    }
}
