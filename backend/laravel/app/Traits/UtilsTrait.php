<?php
 
namespace App\Traits;
 
use Illuminate\Http\Request;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Response;
 
trait UtilsTrait {

	public static function generteUUID(){
        $lenght1 = random_bytes(4);
        $lenght2 = random_bytes(2);
        $lenght3 = random_bytes(2);
        $lenght4 = random_bytes(2);
        $lenght5 = random_bytes(6);
        $uuid = ''.bin2hex($lenght1).'-'.bin2hex($lenght2).'-'.bin2hex($lenght3).'-'.bin2hex($lenght4).'-'.bin2hex($lenght5).'';
        return $uuid;
    } 
}