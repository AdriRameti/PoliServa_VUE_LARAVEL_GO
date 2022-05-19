<?php

namespace App\Repositories;

use App\Models\User;
use Illuminate\Support\Facades\Hash;
use App\Traits\UtilsTrait;
use Illuminate\Support\Facades\Session;

class UserRepository{

    use UtilsTrait;

    public function login($data) {

        $user = User::where('mail', $data['mail'])->first();

        if ($user) {
            if (Hash::check($data['pass'], $user['pass']) == 1) {

                // if ($user['google2fa_secret']) {

                //     $user['is_blocked'] = true;
                //     $user->save();
                // }

                return $user['uuid'];
            } else {
                return "password don't match";
            }
        } else {
            return "user not found" ;
        }

    }

    public function register($data)
    {

        if($data['social']==1){
            $data = [
                'name'=>$data['name'],
                'uuid'=>self::generteUUID(),
                'surnames'=>$data['surnames'],
                'mail'=>$data['mail'],
                'pass'=>'',
                'img'=>$data['img'],
                'google2fa_secret'=> null,
                'role' => $data['role'],
                'is_blocked' => $data['is_blocked']
            ];
            $user = User::where('mail', $data['mail'])->first();
            if($user){
                $user = $user;
            }else{
                $user = User::create($data);
            }
            
        }else{
            $data = [
                'name'=>$data['name'],
                'uuid'=>self::generteUUID(),
                'surnames'=>$data['surnames'],
                'mail'=>$data['mail'],
                'pass'=>$data['pass'],
                'img'=>"https://upload.wikimedia.org/wikipedia/commons/thumb/7/7c/User_font_awesome.svg/2048px-User_font_awesome.svg.png",
                'google2fa_secret'=> null,
                'role' => 'client',
                'is_blocked' => false
            ];
            $user = User::create($data);
        }

        return $user;
    }

    public function getUser() {
        $uuid = self::getUuid();

        $user = User::where('uuid', $uuid)->first();

        return $user;
    }
    public function check2fa(){
        $uuid = self::getUuid();

        $user = User::where('uuid', $uuid)->whereNotNull('google2fa_secret')->first();
        
        return $user;
    }
    public function logOut(){
        Session::forget('token');
    }

}