<?php
 
namespace App\Traits;
 
use Illuminate\Http\Request;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Response;
use App\Repositories\UserRepository;

 
trait Google2FATrait {

    public static function en2fa() {

        $google2fa = app('pragmarx.google2fa');

        $secret = $google2fa->generateSecretKey();

        $userRepository = new UserRepository();

        $user = $userRepository->getUser();

        $user['google2fa_secret'] = encrypt($secret);

        $user->save();

        $QR_Image = self::generateQR($user);

        $user2 = $userRepository->getUser();

        $user2['QR'] = $QR_Image;

        return $user2;

    }

    public static function dis2fa() {
        $userRepository = new UserRepository();

        $user = $userRepository->getUser();

        $user['google2fa_secret'] = null;

        $user->save();

        return $user;
    }

    public static function generateQR($user) {

        $google2fa = app('pragmarx.google2fa');
    
        $QR_Image = $google2fa->getQRCodeInline(
            'Poliserva',
            $user['mail'],
            $user['google2fa_secret']
        );

        return $QR_Image;
    }

    public static function checkOTP($digits) {

        $google2fa = app('pragmarx.google2fa');

        $userRepository = new UserRepository();

        $user = $userRepository->getUser();

        $secret = $user['google2fa_secret'];
    
        if ($google2fa->verifyGoogle2FA($secret, $digits) == 1) {
            return 'verified';
        } else {    
            return 'not verified';
        }
    }

}