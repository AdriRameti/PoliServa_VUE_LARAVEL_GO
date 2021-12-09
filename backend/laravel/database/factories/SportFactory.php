<?php

namespace Database\Factories;

use Illuminate\Database\Eloquent\Factories\Factory;

class SportFactory extends Factory
{
    private $name = null;

    /**
     * Define the model's default state.
     *
     * @return array
     */
    public function definition()
    {
        return [
            "name"=>$this->nameSport(),
            "slug"=>$this->slugSport(),
            "img"=>$this->imgSport()
        ];
    }
    private function nameSport(){
        if ($this->name === null) {
            $this->name = array("padel","futbol","volei","futbol sala","basquet","frontenis","tenis");
        }

        $sport=array_pop($this->name);
        $_SESSION['NameSport']=$sport;    
        return $sport;
    }
    private function slugSport(){
        $nameSport = $_SESSION['NameSport'];
        $randNum = random_int(1000,9999);
        $slug= ''.$nameSport."-".$randNum.'';
        return $slug;
    }
    private function imgSport(){
        $img=array(
            "https://lh3.googleusercontent.com/proxy/o4Nfr0QZyq90Wms2V9krwC9r9zr6wycxPuh-D1RopIgn4x6Wauo_LHpRQR6RhnC_4qp7J1bapg3jd_vSGR_ieRxBNjr7xG32UTq1eGFVjU9Sx7idEhz4JiobK68mTu4EHbzplubk6D3LTB5e",
            "https://www.bbva.com/wp-content/uploads/2017/08/bbva-balon-futbol-2017-08-11.jpg",
            "https://www.fundeu.es/wp-content/uploads/2021/01/voleibol-webefesptwelve224491.jpg",
            "https://mundoentrenamiento.com/wp-content/uploads/2017/05/doble-penalti-en-futbol-sala-800x500.jpg",
            "https://www.basquetplus.com/sites/default/files/main/articles/redes.jpg",
            "https://raquetados.com/wp-content/uploads/2019/09/Frontenis.jpg",
            "https://static2.abc.es/media/bienestar/2019/08/02/tenis-abecedario-kgNF--620x349@abc.jpg"
            );
        $img_aleatoria=array_rand($img,1);
        $image = "";
        switch($_SESSION['NameSport']){
            case "padel":
                $image = $img[0];
                break;
            case "futbol":
                $image = $img[1];
                break;
            case "volei":
                $image= $img[2];
                break;
            case "futbol sala":
                $image = $img[3];
                break;
            case "basquet":
                $image = $img[4];
                break;
            case "frontenis":
                $image = $img[5];
                break;
            case "tenis":
                $image = $img[6];
                break;
        }
        return $image;
    }
}
