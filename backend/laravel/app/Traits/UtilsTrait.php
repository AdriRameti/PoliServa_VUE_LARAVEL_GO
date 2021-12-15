<?php
 
namespace App\Traits;
 
use Illuminate\Http\Request;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Response;
use Firebase\JWT\JWT;
use Firebase\JWT\Key;
use Illuminate\Support\Facades\Session;

trait UtilsTrait {

	public static function generteUUID(){
        $lenght1 = random_bytes(4);
        $lenght2 = random_bytes(2);
        $lenght3 = random_bytes(2);
        $lenght4 = random_bytes(2);
        $lenght5 = random_bytes(6);
        $uuid = ''.bin2hex($lenght1).'-'.bin2hex($lenght2).'-'.bin2hex($lenght3).'-'.bin2hex($lenght4).'-'.bin2hex($lenght5).'';
        session(['uuid'=>$uuid]);
        Session::save();

        return $uuid;
    }
    public static function encode(){
        $key = env("JWT_SECRET");
        $uuid = session('uuid');
        $payload = array(
            "uuid"=>$uuid,
            "iat" => now()->timestamp,
            "iot" => now()->timestamp + 3600,
        );
        $jwt = JWT::encode($payload, $key, 'HS256'); 
        return $jwt;
    }
    public static function decode($token){
        $decoded = JWT::decode($token, new Key(env("JWT_SECRET"), 'HS256'));
        return $decoded;
    } 
    public static function dataMail($info){
        $emailClient =$info[0];
        $emailServer = 'poliserva@infopoliserva.com';
        $code = self::generateCode();
        session(['code'=>$code]);
        Session::save();
        $message = 'Para completar el registro debe introducir el codigo de verificación: '.$code;
        $type = $info[1];
        $html = '';
        $subject = '';
        $body = '';
        $return = '';
        switch ($type){
            case 'register':
                $subject = 'Tu alta en Poliserva';
                $body = 'Gracias por unirte a Poliserva.';
            break;
        }
        $html .= "<html>";
        $html .= "<body>";
            $html .= "Asunto:";
            $html .= "<br><br>";
	       $html .= "<h4>". $subject ."</h4>";
           $html .= "<br><br>";
           $html .= "Mensaje:";
           $html .= "<br><br>";
           $html .= $message;
           $html .= "<br><br>";
	       $html .= $body;
	       $html .= "<br><br>";
	       $html .= "<p>Sent by Poliserva</p>";
		$html .= "</body>";
		$html .= "</html>";

        try{
            $result = self::sendMail($emailServer,$emailClient,$subject,$html);
        }catch(Exception $e){
            return $e;
        }
        return $result;
    }
    public static function sendMail($mailServ,$mailClien,$subject,$html){
        $info_file = parse_ini_file( $_SERVER['DOCUMENT_ROOT'].'\info.ini');
        $config = array();
        $config['api_key']=$info_file['api_key'];
        $config['api_url']=$info_file['api_url'];

        $message = array();
        $message['from'] = $mailServ;
        $message['to'] = $mailClien;
        $message['h:Reply-To'] = $mailServ;
        $message['subject'] = $subject;
        $message['html'] = $html;

        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL, $config['api_url']);
        curl_setopt($ch, CURLOPT_HTTPAUTH, CURLAUTH_BASIC);
        curl_setopt($ch, CURLOPT_USERPWD, "api:{$config['api_key']}");
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
        curl_setopt($ch, CURLOPT_CONNECTTIMEOUT, 10);
        curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, 0);
        curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 0);
        curl_setopt($ch, CURLOPT_POST, true);
        curl_setopt($ch, CURLOPT_POSTFIELDS,$message);
        $result = curl_exec($ch);
        curl_close($ch);
        return $result;
    }
    private static function generateCode(){
        $cod = random_int(00000,99999);
        return $cod;
    }
}